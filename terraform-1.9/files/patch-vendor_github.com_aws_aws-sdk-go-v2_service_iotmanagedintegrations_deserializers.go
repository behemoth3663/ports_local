--- vendor/github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotmanagedintegrations/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpCreateAccountAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -91,7 +88,6 @@ func (m *awsRestjson1_deserializeOpCreateAccountAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -249,10 +245,8 @@ func (m *awsRestjson1_deserializeOpCreateCloudConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -292,7 +286,6 @@ func (m *awsRestjson1_deserializeOpCreateCloudConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -417,10 +410,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -460,7 +451,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -585,10 +575,8 @@ func (m *awsRestjson1_deserializeOpCreateCredentialLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -628,7 +616,6 @@ func (m *awsRestjson1_deserializeOpCreateCredentialLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -784,10 +771,8 @@ func (m *awsRestjson1_deserializeOpCreateDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -827,7 +812,6 @@ func (m *awsRestjson1_deserializeOpCreateDestination) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -952,10 +936,8 @@ func (m *awsRestjson1_deserializeOpCreateEventLogConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -995,7 +977,6 @@ func (m *awsRestjson1_deserializeOpCreateEventLogConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1123,10 +1104,8 @@ func (m *awsRestjson1_deserializeOpCreateManagedThing)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1166,7 +1145,6 @@ func (m *awsRestjson1_deserializeOpCreateManagedThing)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1325,10 +1303,8 @@ func (m *awsRestjson1_deserializeOpCreateNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1368,7 +1344,6 @@ func (m *awsRestjson1_deserializeOpCreateNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1493,10 +1468,8 @@ func (m *awsRestjson1_deserializeOpCreateOtaTask) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1536,7 +1509,6 @@ func (m *awsRestjson1_deserializeOpCreateOtaTask) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1685,10 +1657,8 @@ func (m *awsRestjson1_deserializeOpCreateOtaTaskConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1728,7 +1698,6 @@ func (m *awsRestjson1_deserializeOpCreateOtaTaskConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1853,10 +1822,8 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1896,7 +1863,6 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2075,10 +2041,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2096,7 +2060,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2184,10 +2147,8 @@ func (m *awsRestjson1_deserializeOpDeleteCloudConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2205,7 +2166,6 @@ func (m *awsRestjson1_deserializeOpDeleteCloudConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2290,10 +2250,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2311,7 +2269,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2396,10 +2353,8 @@ func (m *awsRestjson1_deserializeOpDeleteCredentialLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2417,7 +2372,6 @@ func (m *awsRestjson1_deserializeOpDeleteCredentialLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2505,10 +2459,8 @@ func (m *awsRestjson1_deserializeOpDeleteDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2526,7 +2478,6 @@ func (m *awsRestjson1_deserializeOpDeleteDestination) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2611,10 +2562,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventLogConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2632,7 +2581,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventLogConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2717,10 +2665,8 @@ func (m *awsRestjson1_deserializeOpDeleteManagedThing)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2738,7 +2684,6 @@ func (m *awsRestjson1_deserializeOpDeleteManagedThing)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2832,10 +2777,8 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2853,7 +2796,6 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2938,10 +2880,8 @@ func (m *awsRestjson1_deserializeOpDeleteOtaTask) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2959,7 +2899,6 @@ func (m *awsRestjson1_deserializeOpDeleteOtaTask) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3047,10 +2986,8 @@ func (m *awsRestjson1_deserializeOpDeleteOtaTaskConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3068,7 +3005,6 @@ func (m *awsRestjson1_deserializeOpDeleteOtaTaskConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3153,10 +3089,8 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3174,7 +3108,6 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3265,10 +3198,8 @@ func (m *awsRestjson1_deserializeOpDeregisterAccountAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3286,7 +3217,6 @@ func (m *awsRestjson1_deserializeOpDeregisterAccountAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3371,10 +3301,8 @@ func (m *awsRestjson1_deserializeOpGetAccountAssociati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3414,7 +3342,6 @@ func (m *awsRestjson1_deserializeOpGetAccountAssociati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3610,10 +3537,8 @@ func (m *awsRestjson1_deserializeOpGetCloudConnector) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3653,7 +3578,6 @@ func (m *awsRestjson1_deserializeOpGetCloudConnector) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3819,10 +3743,8 @@ func (m *awsRestjson1_deserializeOpGetConnectorDestina
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3862,7 +3784,6 @@ func (m *awsRestjson1_deserializeOpGetConnectorDestina
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4042,10 +3963,8 @@ func (m *awsRestjson1_deserializeOpGetCredentialLocker
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4085,7 +4004,6 @@ func (m *awsRestjson1_deserializeOpGetCredentialLocker
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4252,10 +4170,8 @@ func (m *awsRestjson1_deserializeOpGetCustomEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4295,7 +4211,6 @@ func (m *awsRestjson1_deserializeOpGetCustomEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4426,10 +4341,8 @@ func (m *awsRestjson1_deserializeOpGetDefaultEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4469,7 +4382,6 @@ func (m *awsRestjson1_deserializeOpGetDefaultEncryptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4614,10 +4526,8 @@ func (m *awsRestjson1_deserializeOpGetDestination) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4657,7 +4567,6 @@ func (m *awsRestjson1_deserializeOpGetDestination) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4855,10 +4764,8 @@ func (m *awsRestjson1_deserializeOpGetDeviceDiscovery)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4898,7 +4805,6 @@ func (m *awsRestjson1_deserializeOpGetDeviceDiscovery)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5120,10 +5026,8 @@ func (m *awsRestjson1_deserializeOpGetEventLogConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5163,7 +5067,6 @@ func (m *awsRestjson1_deserializeOpGetEventLogConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5315,10 +5218,8 @@ func (m *awsRestjson1_deserializeOpGetHubConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5358,7 +5259,6 @@ func (m *awsRestjson1_deserializeOpGetHubConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5506,10 +5406,8 @@ func (m *awsRestjson1_deserializeOpGetManagedThing) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5549,7 +5447,6 @@ func (m *awsRestjson1_deserializeOpGetManagedThing) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5918,10 +5815,8 @@ func (m *awsRestjson1_deserializeOpGetManagedThingCapa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5961,7 +5856,6 @@ func (m *awsRestjson1_deserializeOpGetManagedThingCapa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6106,10 +6000,8 @@ func (m *awsRestjson1_deserializeOpGetManagedThingConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6149,7 +6041,6 @@ func (m *awsRestjson1_deserializeOpGetManagedThingConn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6314,10 +6205,8 @@ func (m *awsRestjson1_deserializeOpGetManagedThingMeta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6357,7 +6246,6 @@ func (m *awsRestjson1_deserializeOpGetManagedThingMeta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6493,10 +6381,8 @@ func (m *awsRestjson1_deserializeOpGetManagedThingStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6536,7 +6422,6 @@ func (m *awsRestjson1_deserializeOpGetManagedThingStat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6663,10 +6548,8 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6706,7 +6589,6 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6877,10 +6759,8 @@ func (m *awsRestjson1_deserializeOpGetOtaTask) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6920,7 +6800,6 @@ func (m *awsRestjson1_deserializeOpGetOtaTask) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7183,10 +7062,8 @@ func (m *awsRestjson1_deserializeOpGetOtaTaskConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7226,7 +7103,6 @@ func (m *awsRestjson1_deserializeOpGetOtaTaskConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7390,10 +7266,8 @@ func (m *awsRestjson1_deserializeOpGetProvisioningProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7433,7 +7307,6 @@ func (m *awsRestjson1_deserializeOpGetProvisioningProf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7605,10 +7478,8 @@ func (m *awsRestjson1_deserializeOpGetRuntimeLogConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7648,7 +7519,6 @@ func (m *awsRestjson1_deserializeOpGetRuntimeLogConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7778,10 +7648,8 @@ func (m *awsRestjson1_deserializeOpGetSchemaVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7821,7 +7689,6 @@ func (m *awsRestjson1_deserializeOpGetSchemaVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8002,10 +7869,8 @@ func (m *awsRestjson1_deserializeOpListAccountAssociat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8045,7 +7910,6 @@ func (m *awsRestjson1_deserializeOpListAccountAssociat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8175,10 +8039,8 @@ func (m *awsRestjson1_deserializeOpListCloudConnectors
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8218,7 +8080,6 @@ func (m *awsRestjson1_deserializeOpListCloudConnectors
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8345,10 +8206,8 @@ func (m *awsRestjson1_deserializeOpListConnectorDestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8388,7 +8247,6 @@ func (m *awsRestjson1_deserializeOpListConnectorDestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8515,10 +8373,8 @@ func (m *awsRestjson1_deserializeOpListCredentialLocke
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8558,7 +8414,6 @@ func (m *awsRestjson1_deserializeOpListCredentialLocke
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8688,10 +8543,8 @@ func (m *awsRestjson1_deserializeOpListDestinations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8731,7 +8584,6 @@ func (m *awsRestjson1_deserializeOpListDestinations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8858,10 +8710,8 @@ func (m *awsRestjson1_deserializeOpListDeviceDiscoveri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8901,7 +8751,6 @@ func (m *awsRestjson1_deserializeOpListDeviceDiscoveri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9034,10 +8883,8 @@ func (m *awsRestjson1_deserializeOpListDiscoveredDevic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9077,7 +8924,6 @@ func (m *awsRestjson1_deserializeOpListDiscoveredDevic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9213,10 +9059,8 @@ func (m *awsRestjson1_deserializeOpListEventLogConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9256,7 +9100,6 @@ func (m *awsRestjson1_deserializeOpListEventLogConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9383,10 +9226,8 @@ func (m *awsRestjson1_deserializeOpListManagedThingAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9426,7 +9267,6 @@ func (m *awsRestjson1_deserializeOpListManagedThingAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9553,10 +9393,8 @@ func (m *awsRestjson1_deserializeOpListManagedThings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9596,7 +9434,6 @@ func (m *awsRestjson1_deserializeOpListManagedThings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9729,10 +9566,8 @@ func (m *awsRestjson1_deserializeOpListManagedThingSch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9772,7 +9607,6 @@ func (m *awsRestjson1_deserializeOpListManagedThingSch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9908,10 +9742,8 @@ func (m *awsRestjson1_deserializeOpListNotificationCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9951,7 +9783,6 @@ func (m *awsRestjson1_deserializeOpListNotificationCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10078,10 +9909,8 @@ func (m *awsRestjson1_deserializeOpListOtaTaskConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10121,7 +9950,6 @@ func (m *awsRestjson1_deserializeOpListOtaTaskConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10248,10 +10076,8 @@ func (m *awsRestjson1_deserializeOpListOtaTaskExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10291,7 +10117,6 @@ func (m *awsRestjson1_deserializeOpListOtaTaskExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10421,10 +10246,8 @@ func (m *awsRestjson1_deserializeOpListOtaTasks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10464,7 +10287,6 @@ func (m *awsRestjson1_deserializeOpListOtaTasks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10594,10 +10416,8 @@ func (m *awsRestjson1_deserializeOpListProvisioningPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10637,7 +10457,6 @@ func (m *awsRestjson1_deserializeOpListProvisioningPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10770,10 +10589,8 @@ func (m *awsRestjson1_deserializeOpListSchemaVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10813,7 +10630,6 @@ func (m *awsRestjson1_deserializeOpListSchemaVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10946,10 +10762,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10989,7 +10803,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11107,10 +10920,8 @@ func (m *awsRestjson1_deserializeOpPutDefaultEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11150,7 +10961,6 @@ func (m *awsRestjson1_deserializeOpPutDefaultEncryptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11295,10 +11105,8 @@ func (m *awsRestjson1_deserializeOpPutHubConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11338,7 +11146,6 @@ func (m *awsRestjson1_deserializeOpPutHubConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11470,10 +11277,8 @@ func (m *awsRestjson1_deserializeOpPutRuntimeLogConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11491,7 +11296,6 @@ func (m *awsRestjson1_deserializeOpPutRuntimeLogConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11576,10 +11380,8 @@ func (m *awsRestjson1_deserializeOpRegisterAccountAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11619,7 +11421,6 @@ func (m *awsRestjson1_deserializeOpRegisterAccountAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11765,10 +11566,8 @@ func (m *awsRestjson1_deserializeOpRegisterCustomEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11808,7 +11607,6 @@ func (m *awsRestjson1_deserializeOpRegisterCustomEndpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11939,10 +11737,8 @@ func (m *awsRestjson1_deserializeOpResetRuntimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11960,7 +11756,6 @@ func (m *awsRestjson1_deserializeOpResetRuntimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12045,10 +11840,8 @@ func (m *awsRestjson1_deserializeOpSendConnectorEvent)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12088,7 +11881,6 @@ func (m *awsRestjson1_deserializeOpSendConnectorEvent)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12216,10 +12008,8 @@ func (m *awsRestjson1_deserializeOpSendManagedThingCom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12259,7 +12049,6 @@ func (m *awsRestjson1_deserializeOpSendManagedThingCom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12390,10 +12179,8 @@ func (m *awsRestjson1_deserializeOpStartAccountAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12433,7 +12220,6 @@ func (m *awsRestjson1_deserializeOpStartAccountAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12561,10 +12347,8 @@ func (m *awsRestjson1_deserializeOpStartDeviceDiscover
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12604,7 +12388,6 @@ func (m *awsRestjson1_deserializeOpStartDeviceDiscover
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12754,10 +12537,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12769,7 +12550,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12854,10 +12634,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12869,7 +12647,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12954,10 +12731,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12975,7 +12750,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13066,10 +12840,8 @@ func (m *awsRestjson1_deserializeOpUpdateCloudConnecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13087,7 +12859,6 @@ func (m *awsRestjson1_deserializeOpUpdateCloudConnecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13172,10 +12943,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13193,7 +12962,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13278,10 +13046,8 @@ func (m *awsRestjson1_deserializeOpUpdateDestination) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13299,7 +13065,6 @@ func (m *awsRestjson1_deserializeOpUpdateDestination) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13384,10 +13149,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventLogConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13405,7 +13168,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventLogConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13490,10 +13252,8 @@ func (m *awsRestjson1_deserializeOpUpdateManagedThing)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13511,7 +13271,6 @@ func (m *awsRestjson1_deserializeOpUpdateManagedThing)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13605,10 +13364,8 @@ func (m *awsRestjson1_deserializeOpUpdateNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13626,7 +13383,6 @@ func (m *awsRestjson1_deserializeOpUpdateNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13711,10 +13467,8 @@ func (m *awsRestjson1_deserializeOpUpdateOtaTask) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13732,7 +13486,6 @@ func (m *awsRestjson1_deserializeOpUpdateOtaTask) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
