--- vendor/github.com/aws/aws-sdk-go-v2/service/tnb/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/tnb/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpCancelSolNetworkOpe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -68,7 +65,6 @@ func (m *awsRestjson1_deserializeOpCancelSolNetworkOpe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -153,10 +149,8 @@ func (m *awsRestjson1_deserializeOpCreateSolFunctionPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -196,7 +190,6 @@ func (m *awsRestjson1_deserializeOpCreateSolFunctionPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -362,10 +355,8 @@ func (m *awsRestjson1_deserializeOpCreateSolNetworkIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -405,7 +396,6 @@ func (m *awsRestjson1_deserializeOpCreateSolNetworkIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -565,10 +555,8 @@ func (m *awsRestjson1_deserializeOpCreateSolNetworkPac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -608,7 +596,6 @@ func (m *awsRestjson1_deserializeOpCreateSolNetworkPac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -774,10 +761,8 @@ func (m *awsRestjson1_deserializeOpDeleteSolFunctionPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -795,7 +780,6 @@ func (m *awsRestjson1_deserializeOpDeleteSolFunctionPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -880,10 +864,8 @@ func (m *awsRestjson1_deserializeOpDeleteSolNetworkIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -901,7 +883,6 @@ func (m *awsRestjson1_deserializeOpDeleteSolNetworkIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -986,10 +967,8 @@ func (m *awsRestjson1_deserializeOpDeleteSolNetworkPac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1007,7 +986,6 @@ func (m *awsRestjson1_deserializeOpDeleteSolNetworkPac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1092,10 +1070,8 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1135,7 +1111,6 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1347,10 +1322,8 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1390,7 +1363,6 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1597,10 +1569,8 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1622,7 +1592,6 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1741,10 +1710,8 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1766,7 +1733,6 @@ func (m *awsRestjson1_deserializeOpGetSolFunctionPacka
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1885,10 +1851,8 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1928,7 +1892,6 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2122,10 +2085,8 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2165,7 +2126,6 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2355,10 +2315,8 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2398,7 +2356,6 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2601,10 +2558,8 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2626,7 +2581,6 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2745,10 +2699,8 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2770,7 +2722,6 @@ func (m *awsRestjson1_deserializeOpGetSolNetworkPackag
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2889,10 +2840,8 @@ func (m *awsRestjson1_deserializeOpInstantiateSolNetwo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2932,7 +2881,6 @@ func (m *awsRestjson1_deserializeOpInstantiateSolNetwo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3065,10 +3013,8 @@ func (m *awsRestjson1_deserializeOpListSolFunctionInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3108,7 +3054,6 @@ func (m *awsRestjson1_deserializeOpListSolFunctionInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3235,10 +3180,8 @@ func (m *awsRestjson1_deserializeOpListSolFunctionPack
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3278,7 +3221,6 @@ func (m *awsRestjson1_deserializeOpListSolFunctionPack
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3405,10 +3347,8 @@ func (m *awsRestjson1_deserializeOpListSolNetworkInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3448,7 +3388,6 @@ func (m *awsRestjson1_deserializeOpListSolNetworkInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3575,10 +3514,8 @@ func (m *awsRestjson1_deserializeOpListSolNetworkOpera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3618,7 +3555,6 @@ func (m *awsRestjson1_deserializeOpListSolNetworkOpera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3745,10 +3681,8 @@ func (m *awsRestjson1_deserializeOpListSolNetworkPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3788,7 +3722,6 @@ func (m *awsRestjson1_deserializeOpListSolNetworkPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3915,10 +3848,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3958,7 +3889,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4079,10 +4009,8 @@ func (m *awsRestjson1_deserializeOpPutSolFunctionPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4122,7 +4050,6 @@ func (m *awsRestjson1_deserializeOpPutSolFunctionPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4288,10 +4215,8 @@ func (m *awsRestjson1_deserializeOpPutSolNetworkPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4331,7 +4256,6 @@ func (m *awsRestjson1_deserializeOpPutSolNetworkPackag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4502,10 +4426,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4517,7 +4439,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4602,10 +4523,8 @@ func (m *awsRestjson1_deserializeOpTerminateSolNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4645,7 +4564,6 @@ func (m *awsRestjson1_deserializeOpTerminateSolNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4778,10 +4696,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4793,7 +4709,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4878,10 +4793,8 @@ func (m *awsRestjson1_deserializeOpUpdateSolFunctionPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4921,7 +4834,6 @@ func (m *awsRestjson1_deserializeOpUpdateSolFunctionPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5046,10 +4958,8 @@ func (m *awsRestjson1_deserializeOpUpdateSolNetworkIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5089,7 +4999,6 @@ func (m *awsRestjson1_deserializeOpUpdateSolNetworkIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5222,10 +5131,8 @@ func (m *awsRestjson1_deserializeOpUpdateSolNetworkPac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5265,7 +5172,6 @@ func (m *awsRestjson1_deserializeOpUpdateSolNetworkPac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5390,10 +5296,8 @@ func (m *awsRestjson1_deserializeOpValidateSolFunction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5433,7 +5337,6 @@ func (m *awsRestjson1_deserializeOpValidateSolFunction
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5599,10 +5502,8 @@ func (m *awsRestjson1_deserializeOpValidateSolNetworkP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5642,7 +5543,6 @@ func (m *awsRestjson1_deserializeOpValidateSolNetworkP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
