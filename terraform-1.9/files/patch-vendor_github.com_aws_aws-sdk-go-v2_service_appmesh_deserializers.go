--- vendor/github.com/aws/aws-sdk-go-v2/service/appmesh/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/appmesh/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateGatewayRoute)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateGatewayRoute)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpCreateMesh) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpCreateMesh) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -390,10 +383,8 @@ func (m *awsRestjson1_deserializeOpCreateRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -433,7 +424,6 @@ func (m *awsRestjson1_deserializeOpCreateRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -563,10 +553,8 @@ func (m *awsRestjson1_deserializeOpCreateVirtualGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -606,7 +594,6 @@ func (m *awsRestjson1_deserializeOpCreateVirtualGatewa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -736,10 +723,8 @@ func (m *awsRestjson1_deserializeOpCreateVirtualNode) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -779,7 +764,6 @@ func (m *awsRestjson1_deserializeOpCreateVirtualNode) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -909,10 +893,8 @@ func (m *awsRestjson1_deserializeOpCreateVirtualRouter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -952,7 +934,6 @@ func (m *awsRestjson1_deserializeOpCreateVirtualRouter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1082,10 +1063,8 @@ func (m *awsRestjson1_deserializeOpCreateVirtualServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1125,7 +1104,6 @@ func (m *awsRestjson1_deserializeOpCreateVirtualServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1255,10 +1233,8 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayRoute)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1298,7 +1274,6 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayRoute)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1425,10 +1400,8 @@ func (m *awsRestjson1_deserializeOpDeleteMesh) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1468,7 +1441,6 @@ func (m *awsRestjson1_deserializeOpDeleteMesh) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1595,10 +1567,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1638,7 +1608,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1765,10 +1734,8 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1808,7 +1775,6 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualGatewa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1935,10 +1901,8 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualNode) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1978,7 +1942,6 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualNode) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2105,10 +2068,8 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualRouter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2148,7 +2109,6 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualRouter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2275,10 +2235,8 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2318,7 +2276,6 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2445,10 +2402,8 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayRout
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2488,7 +2443,6 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayRout
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2612,10 +2566,8 @@ func (m *awsRestjson1_deserializeOpDescribeMesh) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2655,7 +2607,6 @@ func (m *awsRestjson1_deserializeOpDescribeMesh) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2779,10 +2730,8 @@ func (m *awsRestjson1_deserializeOpDescribeRoute) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2822,7 +2771,6 @@ func (m *awsRestjson1_deserializeOpDescribeRoute) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2946,10 +2894,8 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2989,7 +2935,6 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualGate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3113,10 +3058,8 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualNode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3156,7 +3099,6 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualNode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3280,10 +3222,8 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualRout
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3323,7 +3263,6 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualRout
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3447,10 +3386,8 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3490,7 +3427,6 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3614,10 +3550,8 @@ func (m *awsRestjson1_deserializeOpListGatewayRoutes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3657,7 +3591,6 @@ func (m *awsRestjson1_deserializeOpListGatewayRoutes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3790,10 +3723,8 @@ func (m *awsRestjson1_deserializeOpListMeshes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3833,7 +3764,6 @@ func (m *awsRestjson1_deserializeOpListMeshes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3966,10 +3896,8 @@ func (m *awsRestjson1_deserializeOpListRoutes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4009,7 +3937,6 @@ func (m *awsRestjson1_deserializeOpListRoutes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4142,10 +4069,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4185,7 +4110,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4318,10 +4242,8 @@ func (m *awsRestjson1_deserializeOpListVirtualGateways
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4361,7 +4283,6 @@ func (m *awsRestjson1_deserializeOpListVirtualGateways
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4494,10 +4415,8 @@ func (m *awsRestjson1_deserializeOpListVirtualNodes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4537,7 +4456,6 @@ func (m *awsRestjson1_deserializeOpListVirtualNodes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4670,10 +4588,8 @@ func (m *awsRestjson1_deserializeOpListVirtualRouters)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4713,7 +4629,6 @@ func (m *awsRestjson1_deserializeOpListVirtualRouters)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4846,10 +4761,8 @@ func (m *awsRestjson1_deserializeOpListVirtualServices
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4889,7 +4802,6 @@ func (m *awsRestjson1_deserializeOpListVirtualServices
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5022,10 +4934,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5037,7 +4947,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5128,10 +5037,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5143,7 +5050,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5231,10 +5137,8 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayRoute)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5274,7 +5178,6 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayRoute)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5404,10 +5307,8 @@ func (m *awsRestjson1_deserializeOpUpdateMesh) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5447,7 +5348,6 @@ func (m *awsRestjson1_deserializeOpUpdateMesh) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5574,10 +5474,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5617,7 +5515,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5747,10 +5644,8 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5790,7 +5685,6 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualGatewa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5920,10 +5814,8 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualNode) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5963,7 +5855,6 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualNode) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6093,10 +5984,8 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualRouter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6136,7 +6025,6 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualRouter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6266,10 +6154,8 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6309,7 +6195,6 @@ func (m *awsRestjson1_deserializeOpUpdateVirtualServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
