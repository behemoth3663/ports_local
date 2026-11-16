--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateAppInstance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateAppInstance) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -222,10 +218,8 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceAd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -265,7 +259,6 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceAd
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -404,10 +397,8 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceBo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -447,7 +438,6 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceBo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -581,10 +571,8 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -624,7 +612,6 @@ func (m *awsRestjson1_deserializeOpCreateAppInstanceUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -758,10 +745,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -779,7 +764,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstance) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -870,10 +854,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceAd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -891,7 +873,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceAd
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -985,10 +966,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceBo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1006,7 +985,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceBo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1100,10 +1078,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1121,7 +1097,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppInstanceUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1215,10 +1190,8 @@ func (m *awsRestjson1_deserializeOpDeregisterAppInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1236,7 +1209,6 @@ func (m *awsRestjson1_deserializeOpDeregisterAppInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1324,10 +1296,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1367,7 +1337,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1491,10 +1460,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1534,7 +1501,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1658,10 +1624,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1701,7 +1665,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1828,10 +1791,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1871,7 +1832,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1995,10 +1955,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2038,7 +1996,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2162,10 +2119,8 @@ func (m *awsRestjson1_deserializeOpGetAppInstanceReten
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2205,7 +2160,6 @@ func (m *awsRestjson1_deserializeOpGetAppInstanceReten
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2345,10 +2299,8 @@ func (m *awsRestjson1_deserializeOpListAppInstanceAdmi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2388,7 +2340,6 @@ func (m *awsRestjson1_deserializeOpListAppInstanceAdmi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2533,10 +2484,8 @@ func (m *awsRestjson1_deserializeOpListAppInstanceBots
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2576,7 +2525,6 @@ func (m *awsRestjson1_deserializeOpListAppInstanceBots
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2721,10 +2669,8 @@ func (m *awsRestjson1_deserializeOpListAppInstances) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2764,7 +2710,6 @@ func (m *awsRestjson1_deserializeOpListAppInstances) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2897,10 +2842,8 @@ func (m *awsRestjson1_deserializeOpListAppInstanceUser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2940,7 +2883,6 @@ func (m *awsRestjson1_deserializeOpListAppInstanceUser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3073,10 +3015,8 @@ func (m *awsRestjson1_deserializeOpListAppInstanceUser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3116,7 +3056,6 @@ func (m *awsRestjson1_deserializeOpListAppInstanceUser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3258,10 +3197,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3301,7 +3238,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3425,10 +3361,8 @@ func (m *awsRestjson1_deserializeOpPutAppInstanceReten
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3468,7 +3402,6 @@ func (m *awsRestjson1_deserializeOpPutAppInstanceReten
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3608,10 +3541,8 @@ func (m *awsRestjson1_deserializeOpPutAppInstanceUserE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3651,7 +3582,6 @@ func (m *awsRestjson1_deserializeOpPutAppInstanceUserE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3787,10 +3717,8 @@ func (m *awsRestjson1_deserializeOpRegisterAppInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3830,7 +3758,6 @@ func (m *awsRestjson1_deserializeOpRegisterAppInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3973,10 +3900,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3994,7 +3919,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4085,10 +4009,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4106,7 +4028,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4194,10 +4115,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4237,7 +4156,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstance) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4368,10 +4286,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceBo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4411,7 +4327,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceBo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4545,10 +4460,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4588,7 +4501,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4722,10 +4634,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4765,7 +4675,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppInstanceUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
