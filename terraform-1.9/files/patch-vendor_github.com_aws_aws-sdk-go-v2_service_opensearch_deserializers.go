--- vendor/github.com/aws/aws-sdk-go-v2/service/opensearch/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/opensearch/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAcceptInboundConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAcceptInboundConnec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -204,10 +200,8 @@ func (m *awsRestjson1_deserializeOpAddDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -247,7 +241,6 @@ func (m *awsRestjson1_deserializeOpAddDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -378,10 +371,8 @@ func (m *awsRestjson1_deserializeOpAddDirectQueryDataS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -421,7 +412,6 @@ func (m *awsRestjson1_deserializeOpAddDirectQueryDataS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -549,10 +539,8 @@ func (m *awsRestjson1_deserializeOpAddTags) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -570,7 +558,6 @@ func (m *awsRestjson1_deserializeOpAddTags) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -652,10 +639,8 @@ func (m *awsRestjson1_deserializeOpAssociatePackage) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -695,7 +680,6 @@ func (m *awsRestjson1_deserializeOpAssociatePackage) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -819,10 +803,8 @@ func (m *awsRestjson1_deserializeOpAssociatePackages) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -862,7 +844,6 @@ func (m *awsRestjson1_deserializeOpAssociatePackages) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -986,10 +967,8 @@ func (m *awsRestjson1_deserializeOpAuthorizeVpcEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1029,7 +1008,6 @@ func (m *awsRestjson1_deserializeOpAuthorizeVpcEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1153,10 +1131,8 @@ func (m *awsRestjson1_deserializeOpCancelDomainConfigC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1196,7 +1172,6 @@ func (m *awsRestjson1_deserializeOpCancelDomainConfigC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1331,10 +1306,8 @@ func (m *awsRestjson1_deserializeOpCancelServiceSoftwa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1374,7 +1347,6 @@ func (m *awsRestjson1_deserializeOpCancelServiceSoftwa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1492,10 +1464,8 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1535,7 +1505,6 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1717,10 +1686,8 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1760,7 +1727,6 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1887,10 +1853,8 @@ func (m *awsRestjson1_deserializeOpCreateOutboundConne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1930,7 +1894,6 @@ func (m *awsRestjson1_deserializeOpCreateOutboundConne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2090,10 +2053,8 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2133,7 +2094,6 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2260,10 +2220,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2303,7 +2261,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2427,10 +2384,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2442,7 +2397,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 	output := &DeleteApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2533,10 +2487,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2576,7 +2528,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2704,10 +2655,8 @@ func (m *awsRestjson1_deserializeOpDeleteDirectQueryDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2725,7 +2674,6 @@ func (m *awsRestjson1_deserializeOpDeleteDirectQueryDa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2810,10 +2758,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2853,7 +2799,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2971,10 +2916,8 @@ func (m *awsRestjson1_deserializeOpDeleteInboundConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3014,7 +2957,6 @@ func (m *awsRestjson1_deserializeOpDeleteInboundConnec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3126,10 +3068,8 @@ func (m *awsRestjson1_deserializeOpDeleteOutboundConne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3169,7 +3109,6 @@ func (m *awsRestjson1_deserializeOpDeleteOutboundConne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3281,10 +3220,8 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3324,7 +3261,6 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3448,10 +3384,8 @@ func (m *awsRestjson1_deserializeOpDeleteVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3491,7 +3425,6 @@ func (m *awsRestjson1_deserializeOpDeleteVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3609,10 +3542,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomain) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3652,7 +3583,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomain) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3770,10 +3700,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainAutoT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3813,7 +3741,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainAutoT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3940,10 +3867,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainChang
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3983,7 +3908,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainChang
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4101,10 +4025,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4144,7 +4066,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4262,10 +4183,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainHealt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4305,7 +4224,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainHealt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4534,10 +4452,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainNodes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4577,7 +4493,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainNodes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4701,10 +4616,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomains) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4744,7 +4657,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomains) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4859,10 +4771,8 @@ func (m *awsRestjson1_deserializeOpDescribeDryRunProgr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4902,7 +4812,6 @@ func (m *awsRestjson1_deserializeOpDescribeDryRunProgr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5033,10 +4942,8 @@ func (m *awsRestjson1_deserializeOpDescribeInboundConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5076,7 +4983,6 @@ func (m *awsRestjson1_deserializeOpDescribeInboundConn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5197,10 +5103,8 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5240,7 +5144,6 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5364,10 +5267,8 @@ func (m *awsRestjson1_deserializeOpDescribeOutboundCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5407,7 +5308,6 @@ func (m *awsRestjson1_deserializeOpDescribeOutboundCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5528,10 +5428,8 @@ func (m *awsRestjson1_deserializeOpDescribePackages) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5571,7 +5469,6 @@ func (m *awsRestjson1_deserializeOpDescribePackages) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5701,10 +5598,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservedIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5744,7 +5639,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservedIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5871,10 +5765,8 @@ func (m *awsRestjson1_deserializeOpDescribeReservedIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5914,7 +5806,6 @@ func (m *awsRestjson1_deserializeOpDescribeReservedIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6041,10 +5932,8 @@ func (m *awsRestjson1_deserializeOpDescribeVpcEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6084,7 +5973,6 @@ func (m *awsRestjson1_deserializeOpDescribeVpcEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6207,10 +6095,8 @@ func (m *awsRestjson1_deserializeOpDissociatePackage) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6250,7 +6136,6 @@ func (m *awsRestjson1_deserializeOpDissociatePackage) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6374,10 +6259,8 @@ func (m *awsRestjson1_deserializeOpDissociatePackages)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6417,7 +6300,6 @@ func (m *awsRestjson1_deserializeOpDissociatePackages)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6541,10 +6423,8 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6584,7 +6464,6 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6795,10 +6674,8 @@ func (m *awsRestjson1_deserializeOpGetCompatibleVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6838,7 +6715,6 @@ func (m *awsRestjson1_deserializeOpGetCompatibleVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6959,10 +6835,8 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7002,7 +6876,6 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7153,10 +7026,8 @@ func (m *awsRestjson1_deserializeOpGetDirectQueryDataS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7196,7 +7067,6 @@ func (m *awsRestjson1_deserializeOpGetDirectQueryDataS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7349,10 +7219,8 @@ func (m *awsRestjson1_deserializeOpGetDomainMaintenanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7392,7 +7260,6 @@ func (m *awsRestjson1_deserializeOpGetDomainMaintenanc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7576,10 +7443,8 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionHi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7619,7 +7484,6 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionHi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7758,10 +7622,8 @@ func (m *awsRestjson1_deserializeOpGetUpgradeHistory) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7801,7 +7663,6 @@ func (m *awsRestjson1_deserializeOpGetUpgradeHistory) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7931,10 +7792,8 @@ func (m *awsRestjson1_deserializeOpGetUpgradeStatus) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7974,7 +7833,6 @@ func (m *awsRestjson1_deserializeOpGetUpgradeStatus) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8117,10 +7975,8 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8160,7 +8016,6 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8293,10 +8148,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8336,7 +8189,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8460,10 +8312,8 @@ func (m *awsRestjson1_deserializeOpListDirectQueryData
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8503,7 +8353,6 @@ func (m *awsRestjson1_deserializeOpListDirectQueryData
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8633,10 +8482,8 @@ func (m *awsRestjson1_deserializeOpListDomainMaintenan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8676,7 +8523,6 @@ func (m *awsRestjson1_deserializeOpListDomainMaintenan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8806,10 +8652,8 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8849,7 +8693,6 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8961,10 +8804,8 @@ func (m *awsRestjson1_deserializeOpListDomainsForPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9004,7 +8845,6 @@ func (m *awsRestjson1_deserializeOpListDomainsForPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9134,10 +8974,8 @@ func (m *awsRestjson1_deserializeOpListInstanceTypeDet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9177,7 +9015,6 @@ func (m *awsRestjson1_deserializeOpListInstanceTypeDet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9304,10 +9141,8 @@ func (m *awsRestjson1_deserializeOpListPackagesForDoma
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9347,7 +9182,6 @@ func (m *awsRestjson1_deserializeOpListPackagesForDoma
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9477,10 +9311,8 @@ func (m *awsRestjson1_deserializeOpListScheduledAction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9520,7 +9352,6 @@ func (m *awsRestjson1_deserializeOpListScheduledAction
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9650,10 +9481,8 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9693,7 +9522,6 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9811,10 +9639,8 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9854,7 +9680,6 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9981,10 +9806,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10024,7 +9847,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10151,10 +9973,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpoints) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10194,7 +10014,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpoints) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10318,10 +10137,8 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointsFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10361,7 +10178,6 @@ func (m *awsRestjson1_deserializeOpListVpcEndpointsFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10488,10 +10304,8 @@ func (m *awsRestjson1_deserializeOpPurchaseReservedIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10531,7 +10345,6 @@ func (m *awsRestjson1_deserializeOpPurchaseReservedIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10668,10 +10481,8 @@ func (m *awsRestjson1_deserializeOpRejectInboundConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10711,7 +10522,6 @@ func (m *awsRestjson1_deserializeOpRejectInboundConnec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10823,10 +10633,8 @@ func (m *awsRestjson1_deserializeOpRemoveTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10844,7 +10652,6 @@ func (m *awsRestjson1_deserializeOpRemoveTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10923,10 +10730,8 @@ func (m *awsRestjson1_deserializeOpRevokeVpcEndpointAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10938,7 +10743,6 @@ func (m *awsRestjson1_deserializeOpRevokeVpcEndpointAc
 	output := &RevokeVpcEndpointAccessOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11023,10 +10827,8 @@ func (m *awsRestjson1_deserializeOpStartDomainMaintena
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11066,7 +10868,6 @@ func (m *awsRestjson1_deserializeOpStartDomainMaintena
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11191,10 +10992,8 @@ func (m *awsRestjson1_deserializeOpStartServiceSoftwar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11234,7 +11033,6 @@ func (m *awsRestjson1_deserializeOpStartServiceSoftwar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11352,10 +11150,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11395,7 +11191,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11591,10 +11386,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11634,7 +11427,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11762,10 +11554,8 @@ func (m *awsRestjson1_deserializeOpUpdateDirectQueryDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11805,7 +11595,6 @@ func (m *awsRestjson1_deserializeOpUpdateDirectQueryDa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11930,10 +11719,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11973,7 +11760,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainConfig)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12107,10 +11893,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12150,7 +11934,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12274,10 +12057,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageScope)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12317,7 +12098,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageScope)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12456,10 +12236,8 @@ func (m *awsRestjson1_deserializeOpUpdateScheduledActi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12499,7 +12277,6 @@ func (m *awsRestjson1_deserializeOpUpdateScheduledActi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12626,10 +12403,8 @@ func (m *awsRestjson1_deserializeOpUpdateVpcEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12669,7 +12444,6 @@ func (m *awsRestjson1_deserializeOpUpdateVpcEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12793,10 +12567,8 @@ func (m *awsRestjson1_deserializeOpUpgradeDomain) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12836,7 +12608,6 @@ func (m *awsRestjson1_deserializeOpUpgradeDomain) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
