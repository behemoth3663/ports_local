--- vendor/github.com/aws/aws-sdk-go-v2/service/pinpointemail/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pinpointemail/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 	output := &CreateConfigurationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -163,7 +157,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 	output := &CreateConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -248,10 +241,8 @@ func (m *awsRestjson1_deserializeOpCreateDedicatedIpPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -263,7 +254,6 @@ func (m *awsRestjson1_deserializeOpCreateDedicatedIpPo
 	output := &CreateDedicatedIpPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -348,10 +338,8 @@ func (m *awsRestjson1_deserializeOpCreateDeliverabilit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -391,7 +379,6 @@ func (m *awsRestjson1_deserializeOpCreateDeliverabilit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -537,10 +524,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -580,7 +565,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -716,10 +700,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -731,7 +713,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 	output := &DeleteConfigurationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -813,10 +794,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -828,7 +807,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 	output := &DeleteConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -907,10 +885,8 @@ func (m *awsRestjson1_deserializeOpDeleteDedicatedIpPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -922,7 +898,6 @@ func (m *awsRestjson1_deserializeOpDeleteDedicatedIpPo
 	output := &DeleteDedicatedIpPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1004,10 +979,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1019,7 +992,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 	output := &DeleteEmailIdentityOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1101,10 +1073,8 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1144,7 +1114,6 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1292,10 +1261,8 @@ func (m *awsRestjson1_deserializeOpGetBlacklistReports
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1335,7 +1302,6 @@ func (m *awsRestjson1_deserializeOpGetBlacklistReports
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1450,10 +1416,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1493,7 +1457,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1637,10 +1600,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1680,7 +1641,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1795,10 +1755,8 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIp) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1838,7 +1796,6 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIp) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1953,10 +1910,8 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIps) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1996,7 +1951,6 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIps) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2120,10 +2074,8 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2163,7 +2115,6 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityDa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2317,10 +2268,8 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2360,7 +2309,6 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2499,10 +2447,8 @@ func (m *awsRestjson1_deserializeOpGetDomainDeliverabi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2542,7 +2488,6 @@ func (m *awsRestjson1_deserializeOpGetDomainDeliverabi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2657,10 +2602,8 @@ func (m *awsRestjson1_deserializeOpGetDomainStatistics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2700,7 +2643,6 @@ func (m *awsRestjson1_deserializeOpGetDomainStatistics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2820,10 +2762,8 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentity) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2863,7 +2803,6 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentity) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3015,10 +2954,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3058,7 +2995,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3179,10 +3115,8 @@ func (m *awsRestjson1_deserializeOpListDedicatedIpPool
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3222,7 +3156,6 @@ func (m *awsRestjson1_deserializeOpListDedicatedIpPool
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3343,10 +3276,8 @@ func (m *awsRestjson1_deserializeOpListDeliverabilityT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3386,7 +3317,6 @@ func (m *awsRestjson1_deserializeOpListDeliverabilityT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3510,10 +3440,8 @@ func (m *awsRestjson1_deserializeOpListDomainDeliverab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3553,7 +3481,6 @@ func (m *awsRestjson1_deserializeOpListDomainDeliverab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3677,10 +3604,8 @@ func (m *awsRestjson1_deserializeOpListEmailIdentities
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3720,7 +3645,6 @@ func (m *awsRestjson1_deserializeOpListEmailIdentities
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3841,10 +3765,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3884,7 +3806,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3999,10 +3920,8 @@ func (m *awsRestjson1_deserializeOpPutAccountDedicated
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4014,7 +3933,6 @@ func (m *awsRestjson1_deserializeOpPutAccountDedicated
 	output := &PutAccountDedicatedIpWarmupAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4090,10 +4008,8 @@ func (m *awsRestjson1_deserializeOpPutAccountSendingAt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4105,7 +4021,6 @@ func (m *awsRestjson1_deserializeOpPutAccountSendingAt
 	output := &PutAccountSendingAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4181,10 +4096,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4196,7 +4109,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetDeliveryOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4275,10 +4187,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4290,7 +4200,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetReputationOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4369,10 +4278,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4384,7 +4291,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetSendingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4463,10 +4369,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4478,7 +4382,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetTrackingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4557,10 +4460,8 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpInPoo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4572,7 +4473,6 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpInPoo
 	output := &PutDedicatedIpInPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4651,10 +4551,8 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpWarmu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4666,7 +4564,6 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpWarmu
 	output := &PutDedicatedIpWarmupAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4745,10 +4642,8 @@ func (m *awsRestjson1_deserializeOpPutDeliverabilityDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4760,7 +4655,6 @@ func (m *awsRestjson1_deserializeOpPutDeliverabilityDa
 	output := &PutDeliverabilityDashboardOptionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4845,10 +4739,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4860,7 +4752,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 	output := &PutEmailIdentityDkimAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4939,10 +4830,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityFee
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4954,7 +4843,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityFee
 	output := &PutEmailIdentityFeedbackAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5033,10 +4921,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityMai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5048,7 +4934,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityMai
 	output := &PutEmailIdentityMailFromAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5127,10 +5012,8 @@ func (m *awsRestjson1_deserializeOpSendEmail) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5170,7 +5053,6 @@ func (m *awsRestjson1_deserializeOpSendEmail) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5304,10 +5186,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5319,7 +5199,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5401,10 +5280,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5416,7 +5293,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5498,10 +5374,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5513,7 +5387,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 	output := &UpdateConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
