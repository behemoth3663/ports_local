--- vendor/github.com/aws/aws-sdk-go-v2/service/groundstation/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/groundstation/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCancelContact) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -207,10 +203,8 @@ func (m *awsRestjson1_deserializeOpCreateConfig) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -250,7 +244,6 @@ func (m *awsRestjson1_deserializeOpCreateConfig) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -390,10 +383,8 @@ func (m *awsRestjson1_deserializeOpCreateDataflowEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -433,7 +424,6 @@ func (m *awsRestjson1_deserializeOpCreateDataflowEndpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +542,8 @@ func (m *awsRestjson1_deserializeOpCreateEphemeris) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -595,7 +583,6 @@ func (m *awsRestjson1_deserializeOpCreateEphemeris) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -714,10 +701,8 @@ func (m *awsRestjson1_deserializeOpCreateMissionProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -757,7 +742,6 @@ func (m *awsRestjson1_deserializeOpCreateMissionProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -876,10 +860,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfig) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -919,7 +901,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfig) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1056,10 +1037,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataflowEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1099,7 +1078,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataflowEndpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1218,10 +1196,8 @@ func (m *awsRestjson1_deserializeOpDeleteEphemeris) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1261,7 +1237,6 @@ func (m *awsRestjson1_deserializeOpDeleteEphemeris) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1380,10 +1355,8 @@ func (m *awsRestjson1_deserializeOpDeleteMissionProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1423,7 +1396,6 @@ func (m *awsRestjson1_deserializeOpDeleteMissionProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1542,10 +1514,8 @@ func (m *awsRestjson1_deserializeOpDescribeContact) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1585,7 +1555,6 @@ func (m *awsRestjson1_deserializeOpDescribeContact) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1869,10 +1838,8 @@ func (m *awsRestjson1_deserializeOpDescribeEphemeris) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1912,7 +1879,6 @@ func (m *awsRestjson1_deserializeOpDescribeEphemeris) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2115,10 +2081,8 @@ func (m *awsRestjson1_deserializeOpGetAgentConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2158,7 +2122,6 @@ func (m *awsRestjson1_deserializeOpGetAgentConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2286,10 +2249,8 @@ func (m *awsRestjson1_deserializeOpGetConfig) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2329,7 +2290,6 @@ func (m *awsRestjson1_deserializeOpGetConfig) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2485,10 +2445,8 @@ func (m *awsRestjson1_deserializeOpGetDataflowEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2528,7 +2486,6 @@ func (m *awsRestjson1_deserializeOpGetDataflowEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2692,10 +2649,8 @@ func (m *awsRestjson1_deserializeOpGetMinuteUsage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2735,7 +2690,6 @@ func (m *awsRestjson1_deserializeOpGetMinuteUsage) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2906,10 +2860,8 @@ func (m *awsRestjson1_deserializeOpGetMissionProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2949,7 +2901,6 @@ func (m *awsRestjson1_deserializeOpGetMissionProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3167,10 +3118,8 @@ func (m *awsRestjson1_deserializeOpGetSatellite) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3210,7 +3159,6 @@ func (m *awsRestjson1_deserializeOpGetSatellite) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3361,10 +3309,8 @@ func (m *awsRestjson1_deserializeOpListConfigs) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3404,7 +3350,6 @@ func (m *awsRestjson1_deserializeOpListConfigs) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3528,10 +3473,8 @@ func (m *awsRestjson1_deserializeOpListContacts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3571,7 +3514,6 @@ func (m *awsRestjson1_deserializeOpListContacts) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3695,10 +3637,8 @@ func (m *awsRestjson1_deserializeOpListDataflowEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3738,7 +3678,6 @@ func (m *awsRestjson1_deserializeOpListDataflowEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3862,10 +3801,8 @@ func (m *awsRestjson1_deserializeOpListEphemerides) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3905,7 +3842,6 @@ func (m *awsRestjson1_deserializeOpListEphemerides) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4029,10 +3965,8 @@ func (m *awsRestjson1_deserializeOpListGroundStations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4072,7 +4006,6 @@ func (m *awsRestjson1_deserializeOpListGroundStations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4196,10 +4129,8 @@ func (m *awsRestjson1_deserializeOpListMissionProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4239,7 +4170,6 @@ func (m *awsRestjson1_deserializeOpListMissionProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4363,10 +4293,8 @@ func (m *awsRestjson1_deserializeOpListSatellites) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4406,7 +4334,6 @@ func (m *awsRestjson1_deserializeOpListSatellites) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4530,10 +4457,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4573,7 +4498,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4688,10 +4612,8 @@ func (m *awsRestjson1_deserializeOpRegisterAgent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4731,7 +4653,6 @@ func (m *awsRestjson1_deserializeOpRegisterAgent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4850,10 +4771,8 @@ func (m *awsRestjson1_deserializeOpReserveContact) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4893,7 +4812,6 @@ func (m *awsRestjson1_deserializeOpReserveContact) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5012,10 +4930,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5027,7 +4943,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5106,10 +5021,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5121,7 +5034,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5200,10 +5112,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentStatus) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5243,7 +5153,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentStatus) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5362,10 +5271,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfig) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5405,7 +5312,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfig) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5542,10 +5448,8 @@ func (m *awsRestjson1_deserializeOpUpdateEphemeris) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5585,7 +5489,6 @@ func (m *awsRestjson1_deserializeOpUpdateEphemeris) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5704,10 +5607,8 @@ func (m *awsRestjson1_deserializeOpUpdateMissionProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5747,7 +5648,6 @@ func (m *awsRestjson1_deserializeOpUpdateMissionProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
