--- vendor/github.com/aws/aws-sdk-go-v2/service/mediaconnect/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mediaconnect/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAddBridgeOutputs) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAddBridgeOutputs) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -225,10 +221,8 @@ func (m *awsRestjson1_deserializeOpAddBridgeSources) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -268,7 +262,6 @@ func (m *awsRestjson1_deserializeOpAddBridgeSources) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -404,10 +397,8 @@ func (m *awsRestjson1_deserializeOpAddFlowMediaStreams
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -447,7 +438,6 @@ func (m *awsRestjson1_deserializeOpAddFlowMediaStreams
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -580,10 +570,8 @@ func (m *awsRestjson1_deserializeOpAddFlowOutputs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -623,7 +611,6 @@ func (m *awsRestjson1_deserializeOpAddFlowOutputs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -759,10 +746,8 @@ func (m *awsRestjson1_deserializeOpAddFlowSources) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -802,7 +787,6 @@ func (m *awsRestjson1_deserializeOpAddFlowSources) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -935,10 +919,8 @@ func (m *awsRestjson1_deserializeOpAddFlowVpcInterface
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -978,7 +960,6 @@ func (m *awsRestjson1_deserializeOpAddFlowVpcInterface
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1111,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpCreateBridge) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1154,7 +1133,6 @@ func (m *awsRestjson1_deserializeOpCreateBridge) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1281,10 +1259,8 @@ func (m *awsRestjson1_deserializeOpCreateFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1324,7 +1300,6 @@ func (m *awsRestjson1_deserializeOpCreateFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1448,10 +1423,8 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1491,7 +1464,6 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1618,10 +1590,8 @@ func (m *awsRestjson1_deserializeOpDeleteBridge) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1661,7 +1631,6 @@ func (m *awsRestjson1_deserializeOpDeleteBridge) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1792,10 +1761,8 @@ func (m *awsRestjson1_deserializeOpDeleteFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1835,7 +1802,6 @@ func (m *awsRestjson1_deserializeOpDeleteFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1972,10 +1938,8 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2015,7 +1979,6 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2146,10 +2109,8 @@ func (m *awsRestjson1_deserializeOpDeregisterGatewayIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2189,7 +2150,6 @@ func (m *awsRestjson1_deserializeOpDeregisterGatewayIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2329,10 +2289,8 @@ func (m *awsRestjson1_deserializeOpDescribeBridge) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2372,7 +2330,6 @@ func (m *awsRestjson1_deserializeOpDescribeBridge) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2499,10 +2456,8 @@ func (m *awsRestjson1_deserializeOpDescribeFlow) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2542,7 +2497,6 @@ func (m *awsRestjson1_deserializeOpDescribeFlow) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2671,10 +2625,8 @@ func (m *awsRestjson1_deserializeOpDescribeFlowSourceM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2714,7 +2666,6 @@ func (m *awsRestjson1_deserializeOpDescribeFlowSourceM
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2865,10 +2816,8 @@ func (m *awsRestjson1_deserializeOpDescribeFlowSourceT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2908,7 +2857,6 @@ func (m *awsRestjson1_deserializeOpDescribeFlowSourceT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3032,10 +2980,8 @@ func (m *awsRestjson1_deserializeOpDescribeGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3075,7 +3021,6 @@ func (m *awsRestjson1_deserializeOpDescribeGateway) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3202,10 +3147,8 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3245,7 +3188,6 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3372,10 +3314,8 @@ func (m *awsRestjson1_deserializeOpDescribeOffering) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3415,7 +3355,6 @@ func (m *awsRestjson1_deserializeOpDescribeOffering) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3536,10 +3475,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3579,7 +3516,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3700,10 +3636,8 @@ func (m *awsRestjson1_deserializeOpGrantFlowEntitlemen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3743,7 +3677,6 @@ func (m *awsRestjson1_deserializeOpGrantFlowEntitlemen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3879,10 +3812,8 @@ func (m *awsRestjson1_deserializeOpListBridges) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3922,7 +3853,6 @@ func (m *awsRestjson1_deserializeOpListBridges) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4052,10 +3982,8 @@ func (m *awsRestjson1_deserializeOpListEntitlements) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4095,7 +4023,6 @@ func (m *awsRestjson1_deserializeOpListEntitlements) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4222,10 +4149,8 @@ func (m *awsRestjson1_deserializeOpListFlows) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4265,7 +4190,6 @@ func (m *awsRestjson1_deserializeOpListFlows) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4392,10 +4316,8 @@ func (m *awsRestjson1_deserializeOpListGatewayInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4435,7 +4357,6 @@ func (m *awsRestjson1_deserializeOpListGatewayInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4565,10 +4486,8 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4608,7 +4527,6 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4738,10 +4656,8 @@ func (m *awsRestjson1_deserializeOpListOfferings) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4781,7 +4697,6 @@ func (m *awsRestjson1_deserializeOpListOfferings) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4908,10 +4823,8 @@ func (m *awsRestjson1_deserializeOpListReservations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4951,7 +4864,6 @@ func (m *awsRestjson1_deserializeOpListReservations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5078,10 +4990,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5121,7 +5031,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5236,10 +5145,8 @@ func (m *awsRestjson1_deserializeOpPurchaseOffering) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5279,7 +5186,6 @@ func (m *awsRestjson1_deserializeOpPurchaseOffering) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5403,10 +5309,8 @@ func (m *awsRestjson1_deserializeOpRemoveBridgeOutput)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5446,7 +5350,6 @@ func (m *awsRestjson1_deserializeOpRemoveBridgeOutput)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5586,10 +5489,8 @@ func (m *awsRestjson1_deserializeOpRemoveBridgeSource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5629,7 +5530,6 @@ func (m *awsRestjson1_deserializeOpRemoveBridgeSource)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5769,10 +5669,8 @@ func (m *awsRestjson1_deserializeOpRemoveFlowMediaStre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5812,7 +5710,6 @@ func (m *awsRestjson1_deserializeOpRemoveFlowMediaStre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5949,10 +5846,8 @@ func (m *awsRestjson1_deserializeOpRemoveFlowOutput) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5992,7 +5887,6 @@ func (m *awsRestjson1_deserializeOpRemoveFlowOutput) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6129,10 +6023,8 @@ func (m *awsRestjson1_deserializeOpRemoveFlowSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6172,7 +6064,6 @@ func (m *awsRestjson1_deserializeOpRemoveFlowSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6309,10 +6200,8 @@ func (m *awsRestjson1_deserializeOpRemoveFlowVpcInterf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6352,7 +6241,6 @@ func (m *awsRestjson1_deserializeOpRemoveFlowVpcInterf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6494,10 +6382,8 @@ func (m *awsRestjson1_deserializeOpRevokeFlowEntitleme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6537,7 +6423,6 @@ func (m *awsRestjson1_deserializeOpRevokeFlowEntitleme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6674,10 +6559,8 @@ func (m *awsRestjson1_deserializeOpStartFlow) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6717,7 +6600,6 @@ func (m *awsRestjson1_deserializeOpStartFlow) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6854,10 +6736,8 @@ func (m *awsRestjson1_deserializeOpStopFlow) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6897,7 +6777,6 @@ func (m *awsRestjson1_deserializeOpStopFlow) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7034,10 +6913,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7055,7 +6932,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7134,10 +7010,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7155,7 +7029,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7234,10 +7107,8 @@ func (m *awsRestjson1_deserializeOpUpdateBridge) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7277,7 +7148,6 @@ func (m *awsRestjson1_deserializeOpUpdateBridge) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7404,10 +7274,8 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeOutput)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7447,7 +7315,6 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeOutput)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7583,10 +7450,8 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeSource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7626,7 +7491,6 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeSource)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7762,10 +7626,8 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeState) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7805,7 +7667,6 @@ func (m *awsRestjson1_deserializeOpUpdateBridgeState) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7945,10 +7806,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7988,7 +7847,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8112,10 +7970,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlowEntitleme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8155,7 +8011,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlowEntitleme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8288,10 +8143,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlowMediaStre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8331,7 +8184,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlowMediaStre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8464,10 +8316,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlowOutput) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8507,7 +8357,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlowOutput) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8640,10 +8489,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlowSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8683,7 +8530,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlowSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8816,10 +8662,8 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8859,7 +8703,6 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
