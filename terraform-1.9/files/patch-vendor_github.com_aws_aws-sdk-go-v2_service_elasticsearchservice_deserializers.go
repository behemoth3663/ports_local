--- vendor/github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAcceptInboundCrossC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAcceptInboundCrossC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -204,10 +200,8 @@ func (m *awsRestjson1_deserializeOpAddTags) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -225,7 +219,6 @@ func (m *awsRestjson1_deserializeOpAddTags) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -307,10 +300,8 @@ func (m *awsRestjson1_deserializeOpAssociatePackage) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -350,7 +341,6 @@ func (m *awsRestjson1_deserializeOpAssociatePackage) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -474,10 +464,8 @@ func (m *awsRestjson1_deserializeOpAuthorizeVpcEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -517,7 +505,6 @@ func (m *awsRestjson1_deserializeOpAuthorizeVpcEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -641,10 +628,8 @@ func (m *awsRestjson1_deserializeOpCancelDomainConfigC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -684,7 +669,6 @@ func (m *awsRestjson1_deserializeOpCancelDomainConfigC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -819,10 +803,8 @@ func (m *awsRestjson1_deserializeOpCancelElasticsearch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -862,7 +844,6 @@ func (m *awsRestjson1_deserializeOpCancelElasticsearch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -980,10 +961,8 @@ func (m *awsRestjson1_deserializeOpCreateElasticsearch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1023,7 +1002,6 @@ func (m *awsRestjson1_deserializeOpCreateElasticsearch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1150,10 +1128,8 @@ func (m *awsRestjson1_deserializeOpCreateOutboundCross
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1193,7 +1169,6 @@ func (m *awsRestjson1_deserializeOpCreateOutboundCross
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1339,10 +1314,8 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1382,7 +1355,6 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1509,10 +1481,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1552,7 +1522,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1676,10 +1645,8 @@ func (m *awsRestjson1_deserializeOpDeleteElasticsearch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1719,7 +1686,6 @@ func (m *awsRestjson1_deserializeOpDeleteElasticsearch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1837,10 +1803,8 @@ func (m *awsRestjson1_deserializeOpDeleteElasticsearch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1858,7 +1822,6 @@ func (m *awsRestjson1_deserializeOpDeleteElasticsearch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1937,10 +1900,8 @@ func (m *awsRestjson1_deserializeOpDeleteInboundCrossC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1980,7 +1941,6 @@ func (m *awsRestjson1_deserializeOpDeleteInboundCrossC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2092,10 +2052,8 @@ func (m *awsRestjson1_deserializeOpDeleteOutboundCross
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2135,7 +2093,6 @@ func (m *awsRestjson1_deserializeOpDeleteOutboundCross
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2247,10 +2204,8 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2290,7 +2245,6 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2414,10 +2368,8 @@ func (m *awsRestjson1_deserializeOpDeleteVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2457,7 +2409,6 @@ func (m *awsRestjson1_deserializeOpDeleteVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2575,10 +2526,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainAutoT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2618,7 +2567,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainAutoT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2745,10 +2693,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainChang
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2788,7 +2734,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainChang
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2906,10 +2851,8 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2949,7 +2892,6 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3067,10 +3009,8 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3110,7 +3050,6 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3228,10 +3167,8 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3271,7 +3208,6 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3386,10 +3322,8 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3429,7 +3363,6 @@ func (m *awsRestjson1_deserializeOpDescribeElasticsear
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3553,10 +3486,8 @@ func (m *awsRestjson1_deserializeOpDescribeInboundCros
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3596,7 +3527,6 @@ func (m *awsRestjson1_deserializeOpDescribeInboundCros
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3717,10 +3647,8 @@ func (m *awsRestjson1_deserializeOpDescribeOutboundCro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3760,7 +3688,6 @@ func (m *awsRestjson1_deserializeOpDescribeOutboundCro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3881,10 +3808,8 @@ func (m *awsRestjson1_deserializeOpDescribePackages) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3924,7 +3849,6 @@ func (m *awsRestjson1_deserializeOpDescribePackages) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3978,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservedEla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4097,7 +4019,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservedEla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4224,10 +4145,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservedEla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4267,7 +4186,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservedEla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4394,10 +4312,8 @@ func (m *awsRestjson1_deserializeOpDescribeVpcEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4437,7 +4353,6 @@ func (m *awsRestjson1_deserializeOpDescribeVpcEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4560,10 +4475,8 @@ func (m *awsRestjson1_deserializeOpDissociatePackage) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4603,7 +4516,6 @@ func (m *awsRestjson1_deserializeOpDissociatePackage) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4727,10 +4639,8 @@ func (m *awsRestjson1_deserializeOpGetCompatibleElasti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4770,7 +4680,6 @@ func (m *awsRestjson1_deserializeOpGetCompatibleElasti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4891,10 +4800,8 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionHi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4934,7 +4841,6 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionHi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5073,10 +4979,8 @@ func (m *awsRestjson1_deserializeOpGetUpgradeHistory) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5116,7 +5020,6 @@ func (m *awsRestjson1_deserializeOpGetUpgradeHistory) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5246,10 +5149,8 @@ func (m *awsRestjson1_deserializeOpGetUpgradeStatus) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5289,7 +5190,6 @@ func (m *awsRestjson1_deserializeOpGetUpgradeStatus) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5432,10 +5332,8 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5475,7 +5373,6 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5587,10 +5484,8 @@ func (m *awsRestjson1_deserializeOpListDomainsForPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5630,7 +5525,6 @@ func (m *awsRestjson1_deserializeOpListDomainsForPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5760,10 +5654,8 @@ func (m *awsRestjson1_deserializeOpListElasticsearchIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5803,7 +5695,6 @@ func (m *awsRestjson1_deserializeOpListElasticsearchIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5930,10 +5821,8 @@ func (m *awsRestjson1_deserializeOpListElasticsearchVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5973,7 +5862,6 @@ func (m *awsRestjson1_deserializeOpListElasticsearchVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6100,10 +5988,8 @@ func (m *awsRestjson1_deserializeOpListPackagesForDoma
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6143,7 +6029,6 @@ func (m *awsRestjson1_deserializeOpListPackagesForDoma
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6273,10 +6158,8 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6316,7 +6199,6 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6434,10 +6316,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6477,7 +6357,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6604,10 +6483,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpoints) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6647,7 +6524,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpoints) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6771,10 +6647,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointsFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6814,7 +6688,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointsFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6941,10 +6814,8 @@ func (m *awsRestjson1_deserializeOpPurchaseReservedEla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6984,7 +6855,6 @@ func (m *awsRestjson1_deserializeOpPurchaseReservedEla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7121,10 +6991,8 @@ func (m *awsRestjson1_deserializeOpRejectInboundCrossC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7164,7 +7032,6 @@ func (m *awsRestjson1_deserializeOpRejectInboundCrossC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7276,10 +7143,8 @@ func (m *awsRestjson1_deserializeOpRemoveTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7297,7 +7162,6 @@ func (m *awsRestjson1_deserializeOpRemoveTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7376,10 +7240,8 @@ func (m *awsRestjson1_deserializeOpRevokeVpcEndpointAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7391,7 +7253,6 @@ func (m *awsRestjson1_deserializeOpRevokeVpcEndpointAc
 	output := &RevokeVpcEndpointAccessOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7476,10 +7337,8 @@ func (m *awsRestjson1_deserializeOpStartElasticsearchS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7519,7 +7378,6 @@ func (m *awsRestjson1_deserializeOpStartElasticsearchS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7637,10 +7495,8 @@ func (m *awsRestjson1_deserializeOpUpdateElasticsearch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7680,7 +7536,6 @@ func (m *awsRestjson1_deserializeOpUpdateElasticsearch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7809,10 +7664,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7852,7 +7705,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7976,10 +7828,8 @@ func (m *awsRestjson1_deserializeOpUpdateVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8019,7 +7869,6 @@ func (m *awsRestjson1_deserializeOpUpdateVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8143,10 +7992,8 @@ func (m *awsRestjson1_deserializeOpUpgradeElasticsearc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8186,7 +8033,6 @@ func (m *awsRestjson1_deserializeOpUpgradeElasticsearc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
