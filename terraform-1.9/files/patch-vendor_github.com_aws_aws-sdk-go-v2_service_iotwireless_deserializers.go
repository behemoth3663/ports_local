--- vendor/github.com/aws/aws-sdk-go-v2/service/iotwireless/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotwireless/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateAwsAccount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAssociateAwsAccount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -221,10 +217,8 @@ func (m *awsRestjson1_deserializeOpAssociateMulticastG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -236,7 +230,6 @@ func (m *awsRestjson1_deserializeOpAssociateMulticastG
 	output := &AssociateMulticastGroupWithFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -324,10 +317,8 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -339,7 +330,6 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 	output := &AssociateWirelessDeviceWithFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -427,10 +417,8 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -442,7 +430,6 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 	output := &AssociateWirelessDeviceWithMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -530,10 +517,8 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -545,7 +530,6 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessDe
 	output := &AssociateWirelessDeviceWithThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -633,10 +617,8 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessGa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -676,7 +658,6 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessGa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -804,10 +785,8 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessGa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -819,7 +798,6 @@ func (m *awsRestjson1_deserializeOpAssociateWirelessGa
 	output := &AssociateWirelessGatewayWithThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -907,10 +885,8 @@ func (m *awsRestjson1_deserializeOpCancelMulticastGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -922,7 +898,6 @@ func (m *awsRestjson1_deserializeOpCancelMulticastGrou
 	output := &CancelMulticastGroupSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1010,10 +985,8 @@ func (m *awsRestjson1_deserializeOpCreateDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1053,7 +1026,6 @@ func (m *awsRestjson1_deserializeOpCreateDestination) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1190,10 +1162,8 @@ func (m *awsRestjson1_deserializeOpCreateDeviceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1233,7 +1203,6 @@ func (m *awsRestjson1_deserializeOpCreateDeviceProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1367,10 +1336,8 @@ func (m *awsRestjson1_deserializeOpCreateFuotaTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1410,7 +1377,6 @@ func (m *awsRestjson1_deserializeOpCreateFuotaTask) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1547,10 +1513,8 @@ func (m *awsRestjson1_deserializeOpCreateMulticastGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1590,7 +1554,6 @@ func (m *awsRestjson1_deserializeOpCreateMulticastGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1727,10 +1690,8 @@ func (m *awsRestjson1_deserializeOpCreateNetworkAnalyz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1770,7 +1731,6 @@ func (m *awsRestjson1_deserializeOpCreateNetworkAnalyz
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1907,10 +1867,8 @@ func (m *awsRestjson1_deserializeOpCreateServiceProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1950,7 +1908,6 @@ func (m *awsRestjson1_deserializeOpCreateServiceProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2084,10 +2041,8 @@ func (m *awsRestjson1_deserializeOpCreateWirelessDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2127,7 +2082,6 @@ func (m *awsRestjson1_deserializeOpCreateWirelessDevic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2264,10 +2218,8 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2307,7 +2259,6 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2441,10 +2392,8 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2484,7 +2433,6 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2621,10 +2569,8 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2664,7 +2610,6 @@ func (m *awsRestjson1_deserializeOpCreateWirelessGatew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2801,10 +2746,8 @@ func (m *awsRestjson1_deserializeOpDeleteDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2816,7 +2759,6 @@ func (m *awsRestjson1_deserializeOpDeleteDestination) 
 	output := &DeleteDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2904,10 +2846,8 @@ func (m *awsRestjson1_deserializeOpDeleteDeviceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2919,7 +2859,6 @@ func (m *awsRestjson1_deserializeOpDeleteDeviceProfile
 	output := &DeleteDeviceProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3007,10 +2946,8 @@ func (m *awsRestjson1_deserializeOpDeleteFuotaTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3022,7 +2959,6 @@ func (m *awsRestjson1_deserializeOpDeleteFuotaTask) Ha
 	output := &DeleteFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3107,10 +3043,8 @@ func (m *awsRestjson1_deserializeOpDeleteMulticastGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3122,7 +3056,6 @@ func (m *awsRestjson1_deserializeOpDeleteMulticastGrou
 	output := &DeleteMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3210,10 +3143,8 @@ func (m *awsRestjson1_deserializeOpDeleteNetworkAnalyz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3225,7 +3156,6 @@ func (m *awsRestjson1_deserializeOpDeleteNetworkAnalyz
 	output := &DeleteNetworkAnalyzerConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3313,10 +3243,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueuedMessage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3328,7 +3256,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueuedMessage
 	output := &DeleteQueuedMessagesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3413,10 +3340,8 @@ func (m *awsRestjson1_deserializeOpDeleteServiceProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3428,7 +3353,6 @@ func (m *awsRestjson1_deserializeOpDeleteServiceProfil
 	output := &DeleteServiceProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3516,10 +3440,8 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3531,7 +3453,6 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessDevic
 	output := &DeleteWirelessDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3616,10 +3537,8 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3631,7 +3550,6 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessDevic
 	output := &DeleteWirelessDeviceImportTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3719,10 +3637,8 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3734,7 +3650,6 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 	output := &DeleteWirelessGatewayOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3819,10 +3734,8 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3834,7 +3747,6 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 	output := &DeleteWirelessGatewayTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3919,10 +3831,8 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3934,7 +3844,6 @@ func (m *awsRestjson1_deserializeOpDeleteWirelessGatew
 	output := &DeleteWirelessGatewayTaskDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4019,10 +3928,8 @@ func (m *awsRestjson1_deserializeOpDeregisterWirelessD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4034,7 +3941,6 @@ func (m *awsRestjson1_deserializeOpDeregisterWirelessD
 	output := &DeregisterWirelessDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4116,10 +4022,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAwsAcco
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4131,7 +4035,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAwsAcco
 	output := &DisassociateAwsAccountFromPartnerAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4213,10 +4116,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMultica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4228,7 +4129,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMultica
 	output := &DisassociateMulticastGroupFromFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4313,10 +4213,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4328,7 +4226,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 	output := &DisassociateWirelessDeviceFromFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4416,10 +4313,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4431,7 +4326,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 	output := &DisassociateWirelessDeviceFromMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4516,10 +4410,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4531,7 +4423,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 	output := &DisassociateWirelessDeviceFromThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4619,10 +4510,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4634,7 +4523,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 	output := &DisassociateWirelessGatewayFromCertificateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4719,10 +4607,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4734,7 +4620,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWireles
 	output := &DisassociateWirelessGatewayFromThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4822,10 +4707,8 @@ func (m *awsRestjson1_deserializeOpGetDestination) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4865,7 +4748,6 @@ func (m *awsRestjson1_deserializeOpGetDestination) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5035,10 +4917,8 @@ func (m *awsRestjson1_deserializeOpGetDeviceProfile) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5078,7 +4958,6 @@ func (m *awsRestjson1_deserializeOpGetDeviceProfile) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5231,10 +5110,8 @@ func (m *awsRestjson1_deserializeOpGetEventConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5274,7 +5151,6 @@ func (m *awsRestjson1_deserializeOpGetEventConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5409,10 +5285,8 @@ func (m *awsRestjson1_deserializeOpGetFuotaTask) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5452,7 +5326,6 @@ func (m *awsRestjson1_deserializeOpGetFuotaTask) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5700,10 +5573,8 @@ func (m *awsRestjson1_deserializeOpGetLogLevelsByResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5743,7 +5614,6 @@ func (m *awsRestjson1_deserializeOpGetLogLevelsByResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5883,10 +5753,8 @@ func (m *awsRestjson1_deserializeOpGetMetricConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5926,7 +5794,6 @@ func (m *awsRestjson1_deserializeOpGetMetricConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6050,10 +5917,8 @@ func (m *awsRestjson1_deserializeOpGetMetrics) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6093,7 +5958,6 @@ func (m *awsRestjson1_deserializeOpGetMetrics) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6217,10 +6081,8 @@ func (m *awsRestjson1_deserializeOpGetMulticastGroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6260,7 +6122,6 @@ func (m *awsRestjson1_deserializeOpGetMulticastGroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6442,10 +6303,8 @@ func (m *awsRestjson1_deserializeOpGetMulticastGroupSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6485,7 +6344,6 @@ func (m *awsRestjson1_deserializeOpGetMulticastGroupSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6606,10 +6464,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkAnalyzerC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6649,7 +6505,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkAnalyzerC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6812,10 +6667,8 @@ func (m *awsRestjson1_deserializeOpGetPartnerAccount) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6855,7 +6708,6 @@ func (m *awsRestjson1_deserializeOpGetPartnerAccount) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6982,10 +6834,8 @@ func (m *awsRestjson1_deserializeOpGetPosition) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7025,7 +6875,6 @@ func (m *awsRestjson1_deserializeOpGetPosition) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7187,10 +7036,8 @@ func (m *awsRestjson1_deserializeOpGetPositionConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7230,7 +7077,6 @@ func (m *awsRestjson1_deserializeOpGetPositionConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7360,10 +7206,8 @@ func (m *awsRestjson1_deserializeOpGetPositionEstimate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7380,7 +7224,6 @@ func (m *awsRestjson1_deserializeOpGetPositionEstimate
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7487,10 +7330,8 @@ func (m *awsRestjson1_deserializeOpGetResourceEventCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7530,7 +7371,6 @@ func (m *awsRestjson1_deserializeOpGetResourceEventCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7671,10 +7511,8 @@ func (m *awsRestjson1_deserializeOpGetResourceLogLevel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7714,7 +7552,6 @@ func (m *awsRestjson1_deserializeOpGetResourceLogLevel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7839,10 +7676,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePosition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7859,7 +7694,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePosition
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7966,10 +7800,8 @@ func (m *awsRestjson1_deserializeOpGetServiceEndpoint)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8009,7 +7841,6 @@ func (m *awsRestjson1_deserializeOpGetServiceEndpoint)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8149,10 +7980,8 @@ func (m *awsRestjson1_deserializeOpGetServiceProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8192,7 +8021,6 @@ func (m *awsRestjson1_deserializeOpGetServiceProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8340,10 +8168,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessDevice) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8383,7 +8209,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessDevice) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8590,10 +8415,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessDeviceIm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8633,7 +8456,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessDeviceIm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8867,10 +8689,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessDeviceSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8910,7 +8730,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessDeviceSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9054,10 +8873,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGateway)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9097,7 +8914,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGateway)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9272,10 +9088,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9315,7 +9129,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9449,10 +9262,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9492,7 +9303,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9613,10 +9423,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9656,7 +9464,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9799,10 +9606,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9842,7 +9647,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10003,10 +9807,8 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10046,7 +9848,6 @@ func (m *awsRestjson1_deserializeOpGetWirelessGatewayT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10194,10 +9995,8 @@ func (m *awsRestjson1_deserializeOpListDestinations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10237,7 +10036,6 @@ func (m *awsRestjson1_deserializeOpListDestinations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10364,10 +10162,8 @@ func (m *awsRestjson1_deserializeOpListDeviceProfiles)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10407,7 +10203,6 @@ func (m *awsRestjson1_deserializeOpListDeviceProfiles)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10534,10 +10329,8 @@ func (m *awsRestjson1_deserializeOpListDevicesForWirel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10577,7 +10370,6 @@ func (m *awsRestjson1_deserializeOpListDevicesForWirel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10719,10 +10511,8 @@ func (m *awsRestjson1_deserializeOpListEventConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10762,7 +10552,6 @@ func (m *awsRestjson1_deserializeOpListEventConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10889,10 +10678,8 @@ func (m *awsRestjson1_deserializeOpListFuotaTasks) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10932,7 +10719,6 @@ func (m *awsRestjson1_deserializeOpListFuotaTasks) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11059,10 +10845,8 @@ func (m *awsRestjson1_deserializeOpListMulticastGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11102,7 +10886,6 @@ func (m *awsRestjson1_deserializeOpListMulticastGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11229,10 +11012,8 @@ func (m *awsRestjson1_deserializeOpListMulticastGroups
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11272,7 +11053,6 @@ func (m *awsRestjson1_deserializeOpListMulticastGroups
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11402,10 +11182,8 @@ func (m *awsRestjson1_deserializeOpListNetworkAnalyzer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11445,7 +11223,6 @@ func (m *awsRestjson1_deserializeOpListNetworkAnalyzer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11572,10 +11349,8 @@ func (m *awsRestjson1_deserializeOpListPartnerAccounts
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11615,7 +11390,6 @@ func (m *awsRestjson1_deserializeOpListPartnerAccounts
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11742,10 +11516,8 @@ func (m *awsRestjson1_deserializeOpListPositionConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11785,7 +11557,6 @@ func (m *awsRestjson1_deserializeOpListPositionConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11912,10 +11683,8 @@ func (m *awsRestjson1_deserializeOpListQueuedMessages)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11955,7 +11724,6 @@ func (m *awsRestjson1_deserializeOpListQueuedMessages)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12085,10 +11853,8 @@ func (m *awsRestjson1_deserializeOpListServiceProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12128,7 +11894,6 @@ func (m *awsRestjson1_deserializeOpListServiceProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12255,10 +12020,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12298,7 +12061,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12419,10 +12181,8 @@ func (m *awsRestjson1_deserializeOpListWirelessDeviceI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12462,7 +12222,6 @@ func (m *awsRestjson1_deserializeOpListWirelessDeviceI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12595,10 +12354,8 @@ func (m *awsRestjson1_deserializeOpListWirelessDevices
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12638,7 +12395,6 @@ func (m *awsRestjson1_deserializeOpListWirelessDevices
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12765,10 +12521,8 @@ func (m *awsRestjson1_deserializeOpListWirelessGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12808,7 +12562,6 @@ func (m *awsRestjson1_deserializeOpListWirelessGateway
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12935,10 +12688,8 @@ func (m *awsRestjson1_deserializeOpListWirelessGateway
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12978,7 +12729,6 @@ func (m *awsRestjson1_deserializeOpListWirelessGateway
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13105,10 +12855,8 @@ func (m *awsRestjson1_deserializeOpPutPositionConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13120,7 +12868,6 @@ func (m *awsRestjson1_deserializeOpPutPositionConfigur
 	output := &PutPositionConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13205,10 +12952,8 @@ func (m *awsRestjson1_deserializeOpPutResourceLogLevel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13220,7 +12965,6 @@ func (m *awsRestjson1_deserializeOpPutResourceLogLevel
 	output := &PutResourceLogLevelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13305,10 +13049,8 @@ func (m *awsRestjson1_deserializeOpResetAllResourceLog
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13320,7 +13062,6 @@ func (m *awsRestjson1_deserializeOpResetAllResourceLog
 	output := &ResetAllResourceLogLevelsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13405,10 +13146,8 @@ func (m *awsRestjson1_deserializeOpResetResourceLogLev
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13420,7 +13159,6 @@ func (m *awsRestjson1_deserializeOpResetResourceLogLev
 	output := &ResetResourceLogLevelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13505,10 +13243,8 @@ func (m *awsRestjson1_deserializeOpSendDataToMulticast
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13548,7 +13284,6 @@ func (m *awsRestjson1_deserializeOpSendDataToMulticast
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13676,10 +13411,8 @@ func (m *awsRestjson1_deserializeOpSendDataToWirelessD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13719,7 +13452,6 @@ func (m *awsRestjson1_deserializeOpSendDataToWirelessD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13841,10 +13573,8 @@ func (m *awsRestjson1_deserializeOpStartBulkAssociateW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13856,7 +13586,6 @@ func (m *awsRestjson1_deserializeOpStartBulkAssociateW
 	output := &StartBulkAssociateWirelessDeviceWithMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13941,10 +13670,8 @@ func (m *awsRestjson1_deserializeOpStartBulkDisassocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13956,7 +13683,6 @@ func (m *awsRestjson1_deserializeOpStartBulkDisassocia
 	output := &StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14041,10 +13767,8 @@ func (m *awsRestjson1_deserializeOpStartFuotaTask) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14056,7 +13780,6 @@ func (m *awsRestjson1_deserializeOpStartFuotaTask) Han
 	output := &StartFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14144,10 +13867,8 @@ func (m *awsRestjson1_deserializeOpStartMulticastGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14159,7 +13880,6 @@ func (m *awsRestjson1_deserializeOpStartMulticastGroup
 	output := &StartMulticastGroupSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14247,10 +13967,8 @@ func (m *awsRestjson1_deserializeOpStartSingleWireless
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14290,7 +14008,6 @@ func (m *awsRestjson1_deserializeOpStartSingleWireless
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14427,10 +14144,8 @@ func (m *awsRestjson1_deserializeOpStartWirelessDevice
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14470,7 +14185,6 @@ func (m *awsRestjson1_deserializeOpStartWirelessDevice
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14607,10 +14321,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14622,7 +14334,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14710,10 +14421,8 @@ func (m *awsRestjson1_deserializeOpTestWirelessDevice)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14753,7 +14462,6 @@ func (m *awsRestjson1_deserializeOpTestWirelessDevice)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14875,10 +14583,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14890,7 +14596,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14975,10 +14680,8 @@ func (m *awsRestjson1_deserializeOpUpdateDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14990,7 +14693,6 @@ func (m *awsRestjson1_deserializeOpUpdateDestination) 
 	output := &UpdateDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15075,10 +14777,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15090,7 +14790,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventConfigur
 	output := &UpdateEventConfigurationByResourceTypesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15172,10 +14871,8 @@ func (m *awsRestjson1_deserializeOpUpdateFuotaTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15187,7 +14884,6 @@ func (m *awsRestjson1_deserializeOpUpdateFuotaTask) Ha
 	output := &UpdateFuotaTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15275,10 +14971,8 @@ func (m *awsRestjson1_deserializeOpUpdateLogLevelsByRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15290,7 +14984,6 @@ func (m *awsRestjson1_deserializeOpUpdateLogLevelsByRe
 	output := &UpdateLogLevelsByResourceTypesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15378,10 +15071,8 @@ func (m *awsRestjson1_deserializeOpUpdateMetricConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15393,7 +15084,6 @@ func (m *awsRestjson1_deserializeOpUpdateMetricConfigu
 	output := &UpdateMetricConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15481,10 +15171,8 @@ func (m *awsRestjson1_deserializeOpUpdateMulticastGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15496,7 +15184,6 @@ func (m *awsRestjson1_deserializeOpUpdateMulticastGrou
 	output := &UpdateMulticastGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15584,10 +15271,8 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkAnalyz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15599,7 +15284,6 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkAnalyz
 	output := &UpdateNetworkAnalyzerConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15684,10 +15368,8 @@ func (m *awsRestjson1_deserializeOpUpdatePartnerAccoun
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15699,7 +15381,6 @@ func (m *awsRestjson1_deserializeOpUpdatePartnerAccoun
 	output := &UpdatePartnerAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15781,10 +15462,8 @@ func (m *awsRestjson1_deserializeOpUpdatePosition) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15796,7 +15475,6 @@ func (m *awsRestjson1_deserializeOpUpdatePosition) Han
 	output := &UpdatePositionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15881,10 +15559,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15896,7 +15572,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceEvent
 	output := &UpdateResourceEventConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15984,10 +15659,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourcePosit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15999,7 +15672,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourcePosit
 	output := &UpdateResourcePositionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16084,10 +15756,8 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16099,7 +15769,6 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessDevic
 	output := &UpdateWirelessDeviceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16184,10 +15853,8 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16199,7 +15866,6 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessDevic
 	output := &UpdateWirelessDeviceImportTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16287,10 +15953,8 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessGatew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16302,7 +15966,6 @@ func (m *awsRestjson1_deserializeOpUpdateWirelessGatew
 	output := &UpdateWirelessGatewayOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
