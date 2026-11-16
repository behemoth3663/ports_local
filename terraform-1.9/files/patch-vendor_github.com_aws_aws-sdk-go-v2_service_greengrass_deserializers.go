--- vendor/github.com/aws/aws-sdk-go-v2/service/greengrass/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/greengrass/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateRoleToGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAssociateRoleToGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -204,10 +200,8 @@ func (m *awsRestjson1_deserializeOpAssociateServiceRol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -247,7 +241,6 @@ func (m *awsRestjson1_deserializeOpAssociateServiceRol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -363,10 +356,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -406,7 +397,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDefi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -573,10 +563,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -616,7 +604,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectorDefi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -756,10 +743,8 @@ func (m *awsRestjson1_deserializeOpCreateCoreDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -799,7 +784,6 @@ func (m *awsRestjson1_deserializeOpCreateCoreDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -966,10 +950,8 @@ func (m *awsRestjson1_deserializeOpCreateCoreDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1009,7 +991,6 @@ func (m *awsRestjson1_deserializeOpCreateCoreDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1149,10 +1130,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1192,7 +1171,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1314,10 +1292,8 @@ func (m *awsRestjson1_deserializeOpCreateDeviceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1357,7 +1333,6 @@ func (m *awsRestjson1_deserializeOpCreateDeviceDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1524,10 +1499,8 @@ func (m *awsRestjson1_deserializeOpCreateDeviceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1567,7 +1540,6 @@ func (m *awsRestjson1_deserializeOpCreateDeviceDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1707,10 +1679,8 @@ func (m *awsRestjson1_deserializeOpCreateFunctionDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1750,7 +1720,6 @@ func (m *awsRestjson1_deserializeOpCreateFunctionDefin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1917,10 +1886,8 @@ func (m *awsRestjson1_deserializeOpCreateFunctionDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1960,7 +1927,6 @@ func (m *awsRestjson1_deserializeOpCreateFunctionDefin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2100,10 +2066,8 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2143,7 +2107,6 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2310,10 +2273,8 @@ func (m *awsRestjson1_deserializeOpCreateGroupCertific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2353,7 +2314,6 @@ func (m *awsRestjson1_deserializeOpCreateGroupCertific
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2469,10 +2429,8 @@ func (m *awsRestjson1_deserializeOpCreateGroupVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2512,7 +2470,6 @@ func (m *awsRestjson1_deserializeOpCreateGroupVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2652,10 +2609,8 @@ func (m *awsRestjson1_deserializeOpCreateLoggerDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2695,7 +2650,6 @@ func (m *awsRestjson1_deserializeOpCreateLoggerDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2862,10 +2816,8 @@ func (m *awsRestjson1_deserializeOpCreateLoggerDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2905,7 +2857,6 @@ func (m *awsRestjson1_deserializeOpCreateLoggerDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3045,10 +2996,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3088,7 +3037,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceDefin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3255,10 +3203,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3298,7 +3244,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceDefin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3438,10 +3383,8 @@ func (m *awsRestjson1_deserializeOpCreateSoftwareUpdat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3481,7 +3424,6 @@ func (m *awsRestjson1_deserializeOpCreateSoftwareUpdat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3615,10 +3557,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3658,7 +3598,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3825,10 +3764,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3868,7 +3805,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4008,10 +3944,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4023,7 +3957,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorDefi
 	output := &DeleteConnectorDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4096,10 +4029,8 @@ func (m *awsRestjson1_deserializeOpDeleteCoreDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4111,7 +4042,6 @@ func (m *awsRestjson1_deserializeOpDeleteCoreDefinitio
 	output := &DeleteCoreDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4184,10 +4114,8 @@ func (m *awsRestjson1_deserializeOpDeleteDeviceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4199,7 +4127,6 @@ func (m *awsRestjson1_deserializeOpDeleteDeviceDefinit
 	output := &DeleteDeviceDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4272,10 +4199,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4287,7 +4212,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunctionDefin
 	output := &DeleteFunctionDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4360,10 +4284,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4375,7 +4297,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 	output := &DeleteGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4448,10 +4369,8 @@ func (m *awsRestjson1_deserializeOpDeleteLoggerDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4463,7 +4382,6 @@ func (m *awsRestjson1_deserializeOpDeleteLoggerDefinit
 	output := &DeleteLoggerDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4536,10 +4454,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4551,7 +4467,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceDefin
 	output := &DeleteResourceDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4624,10 +4539,8 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4639,7 +4552,6 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionD
 	output := &DeleteSubscriptionDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4712,10 +4624,8 @@ func (m *awsRestjson1_deserializeOpDisassociateRoleFro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4755,7 +4665,6 @@ func (m *awsRestjson1_deserializeOpDisassociateRoleFro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4871,10 +4780,8 @@ func (m *awsRestjson1_deserializeOpDisassociateService
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4914,7 +4821,6 @@ func (m *awsRestjson1_deserializeOpDisassociateService
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5027,10 +4933,8 @@ func (m *awsRestjson1_deserializeOpGetAssociatedRole) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5070,7 +4974,6 @@ func (m *awsRestjson1_deserializeOpGetAssociatedRole) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5195,10 +5098,8 @@ func (m *awsRestjson1_deserializeOpGetBulkDeploymentSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5238,7 +5139,6 @@ func (m *awsRestjson1_deserializeOpGetBulkDeploymentSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5384,10 +5284,8 @@ func (m *awsRestjson1_deserializeOpGetConnectivityInfo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5427,7 +5325,6 @@ func (m *awsRestjson1_deserializeOpGetConnectivityInfo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5548,10 +5445,8 @@ func (m *awsRestjson1_deserializeOpGetConnectorDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5591,7 +5486,6 @@ func (m *awsRestjson1_deserializeOpGetConnectorDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5763,10 +5657,8 @@ func (m *awsRestjson1_deserializeOpGetConnectorDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5806,7 +5698,6 @@ func (m *awsRestjson1_deserializeOpGetConnectorDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5960,10 +5851,8 @@ func (m *awsRestjson1_deserializeOpGetCoreDefinition) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6003,7 +5892,6 @@ func (m *awsRestjson1_deserializeOpGetCoreDefinition) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6175,10 +6063,8 @@ func (m *awsRestjson1_deserializeOpGetCoreDefinitionVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6218,7 +6104,6 @@ func (m *awsRestjson1_deserializeOpGetCoreDefinitionVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6372,10 +6257,8 @@ func (m *awsRestjson1_deserializeOpGetDeploymentStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6415,7 +6298,6 @@ func (m *awsRestjson1_deserializeOpGetDeploymentStatus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6560,10 +6442,8 @@ func (m *awsRestjson1_deserializeOpGetDeviceDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6603,7 +6483,6 @@ func (m *awsRestjson1_deserializeOpGetDeviceDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6775,10 +6654,8 @@ func (m *awsRestjson1_deserializeOpGetDeviceDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6818,7 +6695,6 @@ func (m *awsRestjson1_deserializeOpGetDeviceDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6972,10 +6848,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7015,7 +6889,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7187,10 +7060,8 @@ func (m *awsRestjson1_deserializeOpGetFunctionDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7230,7 +7101,6 @@ func (m *awsRestjson1_deserializeOpGetFunctionDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7384,10 +7254,8 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7427,7 +7295,6 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7599,10 +7466,8 @@ func (m *awsRestjson1_deserializeOpGetGroupCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7642,7 +7507,6 @@ func (m *awsRestjson1_deserializeOpGetGroupCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7776,10 +7640,8 @@ func (m *awsRestjson1_deserializeOpGetGroupCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7819,7 +7681,6 @@ func (m *awsRestjson1_deserializeOpGetGroupCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7953,10 +7814,8 @@ func (m *awsRestjson1_deserializeOpGetGroupVersion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7996,7 +7855,6 @@ func (m *awsRestjson1_deserializeOpGetGroupVersion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8141,10 +7999,8 @@ func (m *awsRestjson1_deserializeOpGetLoggerDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8184,7 +8040,6 @@ func (m *awsRestjson1_deserializeOpGetLoggerDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8356,10 +8211,8 @@ func (m *awsRestjson1_deserializeOpGetLoggerDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8399,7 +8252,6 @@ func (m *awsRestjson1_deserializeOpGetLoggerDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8544,10 +8396,8 @@ func (m *awsRestjson1_deserializeOpGetResourceDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8587,7 +8437,6 @@ func (m *awsRestjson1_deserializeOpGetResourceDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8759,10 +8608,8 @@ func (m *awsRestjson1_deserializeOpGetResourceDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8802,7 +8649,6 @@ func (m *awsRestjson1_deserializeOpGetResourceDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8947,10 +8793,8 @@ func (m *awsRestjson1_deserializeOpGetServiceRoleForAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8990,7 +8834,6 @@ func (m *awsRestjson1_deserializeOpGetServiceRoleForAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9112,10 +8955,8 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9155,7 +8996,6 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionDefi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9327,10 +9167,8 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9370,7 +9208,6 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionDefi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9524,10 +9361,8 @@ func (m *awsRestjson1_deserializeOpGetThingRuntimeConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9567,7 +9402,6 @@ func (m *awsRestjson1_deserializeOpGetThingRuntimeConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9679,10 +9513,8 @@ func (m *awsRestjson1_deserializeOpListBulkDeploymentD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9722,7 +9554,6 @@ func (m *awsRestjson1_deserializeOpListBulkDeploymentD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9840,10 +9671,8 @@ func (m *awsRestjson1_deserializeOpListBulkDeployments
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9883,7 +9712,6 @@ func (m *awsRestjson1_deserializeOpListBulkDeployments
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10001,10 +9829,8 @@ func (m *awsRestjson1_deserializeOpListConnectorDefini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10044,7 +9870,6 @@ func (m *awsRestjson1_deserializeOpListConnectorDefini
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10159,10 +9984,8 @@ func (m *awsRestjson1_deserializeOpListConnectorDefini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10202,7 +10025,6 @@ func (m *awsRestjson1_deserializeOpListConnectorDefini
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10320,10 +10142,8 @@ func (m *awsRestjson1_deserializeOpListCoreDefinitions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10363,7 +10183,6 @@ func (m *awsRestjson1_deserializeOpListCoreDefinitions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10478,10 +10297,8 @@ func (m *awsRestjson1_deserializeOpListCoreDefinitionV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10521,7 +10338,6 @@ func (m *awsRestjson1_deserializeOpListCoreDefinitionV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10639,10 +10455,8 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10682,7 +10496,6 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10800,10 +10613,8 @@ func (m *awsRestjson1_deserializeOpListDeviceDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10843,7 +10654,6 @@ func (m *awsRestjson1_deserializeOpListDeviceDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10958,10 +10768,8 @@ func (m *awsRestjson1_deserializeOpListDeviceDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11001,7 +10809,6 @@ func (m *awsRestjson1_deserializeOpListDeviceDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11119,10 +10926,8 @@ func (m *awsRestjson1_deserializeOpListFunctionDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11162,7 +10967,6 @@ func (m *awsRestjson1_deserializeOpListFunctionDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11277,10 +11081,8 @@ func (m *awsRestjson1_deserializeOpListFunctionDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11320,7 +11122,6 @@ func (m *awsRestjson1_deserializeOpListFunctionDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11438,10 +11239,8 @@ func (m *awsRestjson1_deserializeOpListGroupCertificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11481,7 +11280,6 @@ func (m *awsRestjson1_deserializeOpListGroupCertificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11593,10 +11391,8 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11636,7 +11432,6 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11751,10 +11546,8 @@ func (m *awsRestjson1_deserializeOpListGroupVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11794,7 +11587,6 @@ func (m *awsRestjson1_deserializeOpListGroupVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11912,10 +11704,8 @@ func (m *awsRestjson1_deserializeOpListLoggerDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11955,7 +11745,6 @@ func (m *awsRestjson1_deserializeOpListLoggerDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12070,10 +11859,8 @@ func (m *awsRestjson1_deserializeOpListLoggerDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12113,7 +11900,6 @@ func (m *awsRestjson1_deserializeOpListLoggerDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12231,10 +12017,8 @@ func (m *awsRestjson1_deserializeOpListResourceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12274,7 +12058,6 @@ func (m *awsRestjson1_deserializeOpListResourceDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12389,10 +12172,8 @@ func (m *awsRestjson1_deserializeOpListResourceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12432,7 +12213,6 @@ func (m *awsRestjson1_deserializeOpListResourceDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12550,10 +12330,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptionDef
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12593,7 +12371,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptionDef
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12708,10 +12485,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptionDef
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12751,7 +12526,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptionDef
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12869,10 +12643,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12912,7 +12684,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13021,10 +12792,8 @@ func (m *awsRestjson1_deserializeOpResetDeployments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13064,7 +12833,6 @@ func (m *awsRestjson1_deserializeOpResetDeployments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13186,10 +12954,8 @@ func (m *awsRestjson1_deserializeOpStartBulkDeployment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13229,7 +12995,6 @@ func (m *awsRestjson1_deserializeOpStartBulkDeployment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13351,10 +13116,8 @@ func (m *awsRestjson1_deserializeOpStopBulkDeployment)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13366,7 +13129,6 @@ func (m *awsRestjson1_deserializeOpStopBulkDeployment)
 	output := &StopBulkDeploymentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13439,10 +13201,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13460,7 +13220,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13533,10 +13292,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13554,7 +13311,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13627,10 +13383,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13670,7 +13424,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13795,10 +13548,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13810,7 +13561,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorDefi
 	output := &UpdateConnectorDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13883,10 +13633,8 @@ func (m *awsRestjson1_deserializeOpUpdateCoreDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13898,7 +13646,6 @@ func (m *awsRestjson1_deserializeOpUpdateCoreDefinitio
 	output := &UpdateCoreDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13971,10 +13718,8 @@ func (m *awsRestjson1_deserializeOpUpdateDeviceDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13986,7 +13731,6 @@ func (m *awsRestjson1_deserializeOpUpdateDeviceDefinit
 	output := &UpdateDeviceDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14059,10 +13803,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14074,7 +13816,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunctionDefin
 	output := &UpdateFunctionDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14147,10 +13888,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14162,7 +13901,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 	output := &UpdateGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14235,10 +13973,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroupCertific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14278,7 +14014,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroupCertific
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14412,10 +14147,8 @@ func (m *awsRestjson1_deserializeOpUpdateLoggerDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14427,7 +14160,6 @@ func (m *awsRestjson1_deserializeOpUpdateLoggerDefinit
 	output := &UpdateLoggerDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14500,10 +14232,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceDefin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14515,7 +14245,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceDefin
 	output := &UpdateResourceDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14588,10 +14317,8 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14603,7 +14330,6 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionD
 	output := &UpdateSubscriptionDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14676,10 +14402,8 @@ func (m *awsRestjson1_deserializeOpUpdateThingRuntimeC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14691,7 +14415,6 @@ func (m *awsRestjson1_deserializeOpUpdateThingRuntimeC
 	output := &UpdateThingRuntimeConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
