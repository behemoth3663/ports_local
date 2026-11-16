--- vendor/github.com/aws/aws-sdk-go-v2/service/greengrassv2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/greengrassv2/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpAssociateServiceRol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -91,7 +88,6 @@ func (m *awsRestjson1_deserializeOpAssociateServiceRol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -207,10 +203,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateClien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -250,7 +244,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateClien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -371,10 +364,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateCl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -414,7 +405,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateCl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -535,10 +525,8 @@ func (m *awsRestjson1_deserializeOpCancelDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -578,7 +566,6 @@ func (m *awsRestjson1_deserializeOpCancelDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -706,10 +693,8 @@ func (m *awsRestjson1_deserializeOpCreateComponentVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -749,7 +734,6 @@ func (m *awsRestjson1_deserializeOpCreateComponentVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -919,10 +903,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -962,7 +944,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1111,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1132,7 +1111,6 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1220,10 +1198,8 @@ func (m *awsRestjson1_deserializeOpDeleteCoreDevice) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1241,7 +1217,6 @@ func (m *awsRestjson1_deserializeOpDeleteCoreDevice) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1329,10 +1304,8 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1350,7 +1323,6 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1438,10 +1410,8 @@ func (m *awsRestjson1_deserializeOpDescribeComponent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1481,7 +1451,6 @@ func (m *awsRestjson1_deserializeOpDescribeComponent) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1673,10 +1642,8 @@ func (m *awsRestjson1_deserializeOpDisassociateService
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1716,7 +1683,6 @@ func (m *awsRestjson1_deserializeOpDisassociateService
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1829,10 +1795,8 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1872,7 +1836,6 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2015,10 +1978,8 @@ func (m *awsRestjson1_deserializeOpGetComponentVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2058,7 +2019,6 @@ func (m *awsRestjson1_deserializeOpGetComponentVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2183,10 +2143,8 @@ func (m *awsRestjson1_deserializeOpGetConnectivityInfo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2226,7 +2184,6 @@ func (m *awsRestjson1_deserializeOpGetConnectivityInfo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2347,10 +2304,8 @@ func (m *awsRestjson1_deserializeOpGetCoreDevice) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2390,7 +2345,6 @@ func (m *awsRestjson1_deserializeOpGetCoreDevice) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2581,10 +2535,8 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2624,7 +2576,6 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2857,10 +2808,8 @@ func (m *awsRestjson1_deserializeOpGetServiceRoleForAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2900,7 +2849,6 @@ func (m *awsRestjson1_deserializeOpGetServiceRoleForAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3022,10 +2970,8 @@ func (m *awsRestjson1_deserializeOpListClientDevicesAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3065,7 +3011,6 @@ func (m *awsRestjson1_deserializeOpListClientDevicesAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3195,10 +3140,8 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3238,7 +3181,6 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3368,10 +3310,8 @@ func (m *awsRestjson1_deserializeOpListComponentVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3411,7 +3351,6 @@ func (m *awsRestjson1_deserializeOpListComponentVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3541,10 +3480,8 @@ func (m *awsRestjson1_deserializeOpListCoreDevices) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3584,7 +3521,6 @@ func (m *awsRestjson1_deserializeOpListCoreDevices) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3711,10 +3647,8 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3754,7 +3688,6 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3881,10 +3814,8 @@ func (m *awsRestjson1_deserializeOpListEffectiveDeploy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3924,7 +3855,6 @@ func (m *awsRestjson1_deserializeOpListEffectiveDeploy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3984,8 @@ func (m *awsRestjson1_deserializeOpListInstalledCompon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4097,7 +4025,6 @@ func (m *awsRestjson1_deserializeOpListInstalledCompon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4227,10 +4154,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4270,7 +4195,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4385,10 +4309,8 @@ func (m *awsRestjson1_deserializeOpResolveComponentCan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4428,7 +4350,6 @@ func (m *awsRestjson1_deserializeOpResolveComponentCan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4552,10 +4473,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4567,7 +4486,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4646,10 +4564,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4661,7 +4577,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4740,10 +4655,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4783,7 +4696,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
