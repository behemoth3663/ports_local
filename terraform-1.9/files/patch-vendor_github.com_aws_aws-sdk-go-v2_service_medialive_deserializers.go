--- vendor/github.com/aws/aws-sdk-go-v2/service/medialive/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/medialive/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAcceptInputDeviceTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -62,7 +59,6 @@ func (m *awsRestjson1_deserializeOpAcceptInputDeviceTr
 	output := &AcceptInputDeviceTransferOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -159,10 +155,8 @@ func (m *awsRestjson1_deserializeOpBatchDelete) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -202,7 +196,6 @@ func (m *awsRestjson1_deserializeOpBatchDelete) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -337,10 +330,8 @@ func (m *awsRestjson1_deserializeOpBatchStart) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -380,7 +371,6 @@ func (m *awsRestjson1_deserializeOpBatchStart) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -515,10 +505,8 @@ func (m *awsRestjson1_deserializeOpBatchStop) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -558,7 +546,6 @@ func (m *awsRestjson1_deserializeOpBatchStop) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -693,10 +680,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateSchedule
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -736,7 +721,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateSchedule
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -871,10 +855,8 @@ func (m *awsRestjson1_deserializeOpCancelInputDeviceTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -886,7 +868,6 @@ func (m *awsRestjson1_deserializeOpCancelInputDeviceTr
 	output := &CancelInputDeviceTransferOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -983,10 +964,8 @@ func (m *awsRestjson1_deserializeOpClaimDevice) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -998,7 +977,6 @@ func (m *awsRestjson1_deserializeOpClaimDevice) Handle
 	output := &ClaimDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1092,10 +1070,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1135,7 +1111,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1265,10 +1240,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelPlacem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1308,7 +1281,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelPlacem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1485,10 +1457,8 @@ func (m *awsRestjson1_deserializeOpCreateCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1528,7 +1498,6 @@ func (m *awsRestjson1_deserializeOpCreateCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1841,10 +1810,8 @@ func (m *awsRestjson1_deserializeOpCreateCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1884,7 +1851,6 @@ func (m *awsRestjson1_deserializeOpCreateCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2070,10 +2036,8 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2113,7 +2077,6 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2299,10 +2262,8 @@ func (m *awsRestjson1_deserializeOpCreateEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2342,7 +2303,6 @@ func (m *awsRestjson1_deserializeOpCreateEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2551,10 +2511,8 @@ func (m *awsRestjson1_deserializeOpCreateEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2594,7 +2552,6 @@ func (m *awsRestjson1_deserializeOpCreateEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2780,10 +2737,8 @@ func (m *awsRestjson1_deserializeOpCreateInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2823,7 +2778,6 @@ func (m *awsRestjson1_deserializeOpCreateInput) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2947,10 +2901,8 @@ func (m *awsRestjson1_deserializeOpCreateInputSecurity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2990,7 +2942,6 @@ func (m *awsRestjson1_deserializeOpCreateInputSecurity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3114,10 +3065,8 @@ func (m *awsRestjson1_deserializeOpCreateMultiplex) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3157,7 +3106,6 @@ func (m *awsRestjson1_deserializeOpCreateMultiplex) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3287,10 +3235,8 @@ func (m *awsRestjson1_deserializeOpCreateMultiplexProg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3330,7 +3276,6 @@ func (m *awsRestjson1_deserializeOpCreateMultiplexProg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3460,10 +3405,8 @@ func (m *awsRestjson1_deserializeOpCreateNetwork) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3503,7 +3446,6 @@ func (m *awsRestjson1_deserializeOpCreateNetwork) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3676,10 +3618,8 @@ func (m *awsRestjson1_deserializeOpCreateNode) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3719,7 +3659,6 @@ func (m *awsRestjson1_deserializeOpCreateNode) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3928,10 +3867,8 @@ func (m *awsRestjson1_deserializeOpCreateNodeRegistrat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3971,7 +3908,6 @@ func (m *awsRestjson1_deserializeOpCreateNodeRegistrat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4102,10 +4038,8 @@ func (m *awsRestjson1_deserializeOpCreatePartnerInput)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4145,7 +4079,6 @@ func (m *awsRestjson1_deserializeOpCreatePartnerInput)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4269,10 +4202,8 @@ func (m *awsRestjson1_deserializeOpCreateSdiSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4312,7 +4243,6 @@ func (m *awsRestjson1_deserializeOpCreateSdiSource) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4439,10 +4369,8 @@ func (m *awsRestjson1_deserializeOpCreateSignalMap) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4482,7 +4410,6 @@ func (m *awsRestjson1_deserializeOpCreateSignalMap) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4747,10 +4674,8 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4768,7 +4693,6 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4850,10 +4774,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4893,7 +4815,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5154,10 +5075,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPlacem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5197,7 +5116,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPlacem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5377,10 +5295,8 @@ func (m *awsRestjson1_deserializeOpDeleteCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5398,7 +5314,6 @@ func (m *awsRestjson1_deserializeOpDeleteCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5486,10 +5401,8 @@ func (m *awsRestjson1_deserializeOpDeleteCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5507,7 +5420,6 @@ func (m *awsRestjson1_deserializeOpDeleteCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5595,10 +5507,8 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5638,7 +5548,6 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5827,10 +5736,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5848,7 +5755,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5936,10 +5842,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5957,7 +5861,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6045,10 +5948,8 @@ func (m *awsRestjson1_deserializeOpDeleteInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6060,7 +5961,6 @@ func (m *awsRestjson1_deserializeOpDeleteInput) Handle
 	output := &DeleteInputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6154,10 +6054,8 @@ func (m *awsRestjson1_deserializeOpDeleteInputSecurity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6169,7 +6067,6 @@ func (m *awsRestjson1_deserializeOpDeleteInputSecurity
 	output := &DeleteInputSecurityGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6260,10 +6157,8 @@ func (m *awsRestjson1_deserializeOpDeleteMultiplex) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6303,7 +6198,6 @@ func (m *awsRestjson1_deserializeOpDeleteMultiplex) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6510,10 +6404,8 @@ func (m *awsRestjson1_deserializeOpDeleteMultiplexProg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6553,7 +6445,6 @@ func (m *awsRestjson1_deserializeOpDeleteMultiplexProg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6711,10 +6602,8 @@ func (m *awsRestjson1_deserializeOpDeleteNetwork) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6754,7 +6643,6 @@ func (m *awsRestjson1_deserializeOpDeleteNetwork) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6930,10 +6818,8 @@ func (m *awsRestjson1_deserializeOpDeleteNode) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6973,7 +6859,6 @@ func (m *awsRestjson1_deserializeOpDeleteNode) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7185,10 +7070,8 @@ func (m *awsRestjson1_deserializeOpDeleteReservation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7228,7 +7111,6 @@ func (m *awsRestjson1_deserializeOpDeleteReservation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7570,10 +7452,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchedule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7585,7 +7465,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchedule) Han
 	output := &DeleteScheduleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7676,10 +7555,8 @@ func (m *awsRestjson1_deserializeOpDeleteSdiSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7719,7 +7596,6 @@ func (m *awsRestjson1_deserializeOpDeleteSdiSource) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7849,10 +7725,8 @@ func (m *awsRestjson1_deserializeOpDeleteSignalMap) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7870,7 +7744,6 @@ func (m *awsRestjson1_deserializeOpDeleteSignalMap) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7958,10 +7831,8 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7979,7 +7850,6 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8061,10 +7931,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8104,7 +7972,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8228,10 +8095,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8271,7 +8136,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8529,10 +8393,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelPlac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8572,7 +8434,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelPlac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8749,10 +8610,8 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8792,7 +8651,6 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8978,10 +8836,8 @@ func (m *awsRestjson1_deserializeOpDescribeInput) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9021,7 +8877,6 @@ func (m *awsRestjson1_deserializeOpDescribeInput) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9284,10 +9139,8 @@ func (m *awsRestjson1_deserializeOpDescribeInputDevice
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9327,7 +9180,6 @@ func (m *awsRestjson1_deserializeOpDescribeInputDevice
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9573,10 +9425,8 @@ func (m *awsRestjson1_deserializeOpDescribeInputDevice
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9598,7 +9448,6 @@ func (m *awsRestjson1_deserializeOpDescribeInputDevice
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9733,10 +9582,8 @@ func (m *awsRestjson1_deserializeOpDescribeInputSecuri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9776,7 +9623,6 @@ func (m *awsRestjson1_deserializeOpDescribeInputSecuri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9940,10 +9786,8 @@ func (m *awsRestjson1_deserializeOpDescribeMultiplex) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9983,7 +9827,6 @@ func (m *awsRestjson1_deserializeOpDescribeMultiplex) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10187,10 +10030,8 @@ func (m *awsRestjson1_deserializeOpDescribeMultiplexPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10230,7 +10071,6 @@ func (m *awsRestjson1_deserializeOpDescribeMultiplexPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10385,10 +10225,8 @@ func (m *awsRestjson1_deserializeOpDescribeNetwork) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10428,7 +10266,6 @@ func (m *awsRestjson1_deserializeOpDescribeNetwork) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10601,10 +10438,8 @@ func (m *awsRestjson1_deserializeOpDescribeNode) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10644,7 +10479,6 @@ func (m *awsRestjson1_deserializeOpDescribeNode) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10853,10 +10687,8 @@ func (m *awsRestjson1_deserializeOpDescribeOffering) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10896,7 +10728,6 @@ func (m *awsRestjson1_deserializeOpDescribeOffering) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11167,10 +10998,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11210,7 +11039,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11549,10 +11377,8 @@ func (m *awsRestjson1_deserializeOpDescribeSchedule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11592,7 +11418,6 @@ func (m *awsRestjson1_deserializeOpDescribeSchedule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11728,10 +11553,8 @@ func (m *awsRestjson1_deserializeOpDescribeSdiSource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11771,7 +11594,6 @@ func (m *awsRestjson1_deserializeOpDescribeSdiSource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11898,10 +11720,8 @@ func (m *awsRestjson1_deserializeOpDescribeThumbnails)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11941,7 +11761,6 @@ func (m *awsRestjson1_deserializeOpDescribeThumbnails)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12071,10 +11890,8 @@ func (m *awsRestjson1_deserializeOpGetCloudWatchAlarmT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12114,7 +11931,6 @@ func (m *awsRestjson1_deserializeOpGetCloudWatchAlarmT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12424,10 +12240,8 @@ func (m *awsRestjson1_deserializeOpGetCloudWatchAlarmT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12467,7 +12281,6 @@ func (m *awsRestjson1_deserializeOpGetCloudWatchAlarmT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12650,10 +12463,8 @@ func (m *awsRestjson1_deserializeOpGetEventBridgeRuleT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12693,7 +12504,6 @@ func (m *awsRestjson1_deserializeOpGetEventBridgeRuleT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12899,10 +12709,8 @@ func (m *awsRestjson1_deserializeOpGetEventBridgeRuleT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12942,7 +12750,6 @@ func (m *awsRestjson1_deserializeOpGetEventBridgeRuleT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13125,10 +12932,8 @@ func (m *awsRestjson1_deserializeOpGetSignalMap) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13168,7 +12973,6 @@ func (m *awsRestjson1_deserializeOpGetSignalMap) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13430,10 +13234,8 @@ func (m *awsRestjson1_deserializeOpListChannelPlacemen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13473,7 +13275,6 @@ func (m *awsRestjson1_deserializeOpListChannelPlacemen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13606,10 +13407,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13649,7 +13448,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13782,10 +13580,8 @@ func (m *awsRestjson1_deserializeOpListCloudWatchAlarm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13825,7 +13621,6 @@ func (m *awsRestjson1_deserializeOpListCloudWatchAlarm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13955,10 +13750,8 @@ func (m *awsRestjson1_deserializeOpListCloudWatchAlarm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13998,7 +13791,6 @@ func (m *awsRestjson1_deserializeOpListCloudWatchAlarm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14128,10 +13920,8 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14171,7 +13961,6 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14304,10 +14093,8 @@ func (m *awsRestjson1_deserializeOpListEventBridgeRule
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14347,7 +14134,6 @@ func (m *awsRestjson1_deserializeOpListEventBridgeRule
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14477,10 +14263,8 @@ func (m *awsRestjson1_deserializeOpListEventBridgeRule
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14520,7 +14304,6 @@ func (m *awsRestjson1_deserializeOpListEventBridgeRule
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14650,10 +14433,8 @@ func (m *awsRestjson1_deserializeOpListInputDevices) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14693,7 +14474,6 @@ func (m *awsRestjson1_deserializeOpListInputDevices) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14826,10 +14606,8 @@ func (m *awsRestjson1_deserializeOpListInputDeviceTran
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14869,7 +14647,6 @@ func (m *awsRestjson1_deserializeOpListInputDeviceTran
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15005,10 +14782,8 @@ func (m *awsRestjson1_deserializeOpListInputs) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15048,7 +14823,6 @@ func (m *awsRestjson1_deserializeOpListInputs) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15181,10 +14955,8 @@ func (m *awsRestjson1_deserializeOpListInputSecurityGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15224,7 +14996,6 @@ func (m *awsRestjson1_deserializeOpListInputSecurityGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15357,10 +15128,8 @@ func (m *awsRestjson1_deserializeOpListMultiplexes) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15400,7 +15169,6 @@ func (m *awsRestjson1_deserializeOpListMultiplexes) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15533,10 +15301,8 @@ func (m *awsRestjson1_deserializeOpListMultiplexProgra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15576,7 +15342,6 @@ func (m *awsRestjson1_deserializeOpListMultiplexProgra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15712,10 +15477,8 @@ func (m *awsRestjson1_deserializeOpListNetworks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15755,7 +15518,6 @@ func (m *awsRestjson1_deserializeOpListNetworks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15888,10 +15650,8 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15931,7 +15691,6 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16064,10 +15823,8 @@ func (m *awsRestjson1_deserializeOpListOfferings) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16107,7 +15864,6 @@ func (m *awsRestjson1_deserializeOpListOfferings) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16240,10 +15996,8 @@ func (m *awsRestjson1_deserializeOpListReservations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16283,7 +16037,6 @@ func (m *awsRestjson1_deserializeOpListReservations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16416,10 +16169,8 @@ func (m *awsRestjson1_deserializeOpListSdiSources) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16459,7 +16210,6 @@ func (m *awsRestjson1_deserializeOpListSdiSources) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16592,10 +16342,8 @@ func (m *awsRestjson1_deserializeOpListSignalMaps) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16635,7 +16383,6 @@ func (m *awsRestjson1_deserializeOpListSignalMaps) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16765,10 +16512,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16808,7 +16553,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16926,10 +16670,8 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16969,7 +16711,6 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17099,10 +16840,8 @@ func (m *awsRestjson1_deserializeOpPurchaseOffering) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17142,7 +16881,6 @@ func (m *awsRestjson1_deserializeOpPurchaseOffering) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17272,10 +17010,8 @@ func (m *awsRestjson1_deserializeOpRebootInputDevice) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17287,7 +17023,6 @@ func (m *awsRestjson1_deserializeOpRebootInputDevice) 
 	output := &RebootInputDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17381,10 +17116,8 @@ func (m *awsRestjson1_deserializeOpRejectInputDeviceTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17396,7 +17129,6 @@ func (m *awsRestjson1_deserializeOpRejectInputDeviceTr
 	output := &RejectInputDeviceTransferOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17493,10 +17225,8 @@ func (m *awsRestjson1_deserializeOpRestartChannelPipel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17536,7 +17266,6 @@ func (m *awsRestjson1_deserializeOpRestartChannelPipel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17806,10 +17535,8 @@ func (m *awsRestjson1_deserializeOpStartChannel) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17849,7 +17576,6 @@ func (m *awsRestjson1_deserializeOpStartChannel) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18110,10 +17836,8 @@ func (m *awsRestjson1_deserializeOpStartDeleteMonitorD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18153,7 +17877,6 @@ func (m *awsRestjson1_deserializeOpStartDeleteMonitorD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18418,10 +18141,8 @@ func (m *awsRestjson1_deserializeOpStartInputDevice) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18433,7 +18154,6 @@ func (m *awsRestjson1_deserializeOpStartInputDevice) H
 	output := &StartInputDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18527,10 +18247,8 @@ func (m *awsRestjson1_deserializeOpStartInputDeviceMai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18542,7 +18260,6 @@ func (m *awsRestjson1_deserializeOpStartInputDeviceMai
 	output := &StartInputDeviceMaintenanceWindowOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18636,10 +18353,8 @@ func (m *awsRestjson1_deserializeOpStartMonitorDeploym
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18679,7 +18394,6 @@ func (m *awsRestjson1_deserializeOpStartMonitorDeploym
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18944,10 +18658,8 @@ func (m *awsRestjson1_deserializeOpStartMultiplex) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18987,7 +18699,6 @@ func (m *awsRestjson1_deserializeOpStartMultiplex) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19194,10 +18905,8 @@ func (m *awsRestjson1_deserializeOpStartUpdateSignalMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19237,7 +18946,6 @@ func (m *awsRestjson1_deserializeOpStartUpdateSignalMa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19502,10 +19210,8 @@ func (m *awsRestjson1_deserializeOpStopChannel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19545,7 +19251,6 @@ func (m *awsRestjson1_deserializeOpStopChannel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19806,10 +19511,8 @@ func (m *awsRestjson1_deserializeOpStopInputDevice) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19821,7 +19524,6 @@ func (m *awsRestjson1_deserializeOpStopInputDevice) Ha
 	output := &StopInputDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19915,10 +19617,8 @@ func (m *awsRestjson1_deserializeOpStopMultiplex) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19958,7 +19658,6 @@ func (m *awsRestjson1_deserializeOpStopMultiplex) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20165,10 +19864,8 @@ func (m *awsRestjson1_deserializeOpTransferInputDevice
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20180,7 +19877,6 @@ func (m *awsRestjson1_deserializeOpTransferInputDevice
 	output := &TransferInputDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20277,10 +19973,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20320,7 +20014,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20447,10 +20140,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20490,7 +20181,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20617,10 +20307,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelClass)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20660,7 +20348,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelClass)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20793,10 +20480,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelPlacem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20836,7 +20521,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelPlacem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21016,10 +20700,8 @@ func (m *awsRestjson1_deserializeOpUpdateCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21059,7 +20741,6 @@ func (m *awsRestjson1_deserializeOpUpdateCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21372,10 +21053,8 @@ func (m *awsRestjson1_deserializeOpUpdateCloudWatchAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21415,7 +21094,6 @@ func (m *awsRestjson1_deserializeOpUpdateCloudWatchAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21601,10 +21279,8 @@ func (m *awsRestjson1_deserializeOpUpdateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21644,7 +21320,6 @@ func (m *awsRestjson1_deserializeOpUpdateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21821,10 +21496,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21864,7 +21537,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22073,10 +21745,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventBridgeRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22116,7 +21786,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventBridgeRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22302,10 +21971,8 @@ func (m *awsRestjson1_deserializeOpUpdateInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22345,7 +22012,6 @@ func (m *awsRestjson1_deserializeOpUpdateInput) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22472,10 +22138,8 @@ func (m *awsRestjson1_deserializeOpUpdateInputDevice) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22515,7 +22179,6 @@ func (m *awsRestjson1_deserializeOpUpdateInputDevice) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22764,10 +22427,8 @@ func (m *awsRestjson1_deserializeOpUpdateInputSecurity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22807,7 +22468,6 @@ func (m *awsRestjson1_deserializeOpUpdateInputSecurity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22934,10 +22594,8 @@ func (m *awsRestjson1_deserializeOpUpdateMultiplex) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22977,7 +22635,6 @@ func (m *awsRestjson1_deserializeOpUpdateMultiplex) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23107,10 +22764,8 @@ func (m *awsRestjson1_deserializeOpUpdateMultiplexProg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23150,7 +22805,6 @@ func (m *awsRestjson1_deserializeOpUpdateMultiplexProg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23280,10 +22934,8 @@ func (m *awsRestjson1_deserializeOpUpdateNetwork) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23323,7 +22975,6 @@ func (m *awsRestjson1_deserializeOpUpdateNetwork) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23496,10 +23147,8 @@ func (m *awsRestjson1_deserializeOpUpdateNode) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23539,7 +23188,6 @@ func (m *awsRestjson1_deserializeOpUpdateNode) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23748,10 +23396,8 @@ func (m *awsRestjson1_deserializeOpUpdateNodeState) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23791,7 +23437,6 @@ func (m *awsRestjson1_deserializeOpUpdateNodeState) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24003,10 +23648,8 @@ func (m *awsRestjson1_deserializeOpUpdateReservation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24046,7 +23689,6 @@ func (m *awsRestjson1_deserializeOpUpdateReservation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24176,10 +23818,8 @@ func (m *awsRestjson1_deserializeOpUpdateSdiSource) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24219,7 +23859,6 @@ func (m *awsRestjson1_deserializeOpUpdateSdiSource) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
