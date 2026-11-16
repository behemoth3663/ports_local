--- vendor/github.com/aws/aws-sdk-go-v2/service/eks/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/eks/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateAccessPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpAssociateAccessPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -223,10 +219,8 @@ func (m *awsRestjson1_deserializeOpAssociateEncryption
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -266,7 +260,6 @@ func (m *awsRestjson1_deserializeOpAssociateEncryption
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -393,10 +386,8 @@ func (m *awsRestjson1_deserializeOpAssociateIdentityPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -436,7 +427,6 @@ func (m *awsRestjson1_deserializeOpAssociateIdentityPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -568,10 +558,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessEntry) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -611,7 +599,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessEntry) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -735,10 +722,8 @@ func (m *awsRestjson1_deserializeOpCreateAddon) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -778,7 +763,6 @@ func (m *awsRestjson1_deserializeOpCreateAddon) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -902,10 +886,8 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -945,7 +927,6 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1072,10 +1053,8 @@ func (m *awsRestjson1_deserializeOpCreateEksAnywhereSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1115,7 +1094,6 @@ func (m *awsRestjson1_deserializeOpCreateEksAnywhereSu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1236,10 +1214,8 @@ func (m *awsRestjson1_deserializeOpCreateFargateProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1279,7 +1255,6 @@ func (m *awsRestjson1_deserializeOpCreateFargateProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1403,10 +1378,8 @@ func (m *awsRestjson1_deserializeOpCreateNodegroup) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1446,7 +1419,6 @@ func (m *awsRestjson1_deserializeOpCreateNodegroup) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1573,10 +1545,8 @@ func (m *awsRestjson1_deserializeOpCreatePodIdentityAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1616,7 +1586,6 @@ func (m *awsRestjson1_deserializeOpCreatePodIdentityAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1740,10 +1709,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessEntry) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1755,7 +1722,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessEntry) 
 	output := &DeleteAccessEntryOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1834,10 +1800,8 @@ func (m *awsRestjson1_deserializeOpDeleteAddon) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1877,7 +1841,6 @@ func (m *awsRestjson1_deserializeOpDeleteAddon) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1998,10 +1961,8 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2041,7 +2002,6 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2162,10 +2122,8 @@ func (m *awsRestjson1_deserializeOpDeleteEksAnywhereSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2205,7 +2163,6 @@ func (m *awsRestjson1_deserializeOpDeleteEksAnywhereSu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2323,10 +2280,8 @@ func (m *awsRestjson1_deserializeOpDeleteFargateProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2366,7 +2321,6 @@ func (m *awsRestjson1_deserializeOpDeleteFargateProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2484,10 +2438,8 @@ func (m *awsRestjson1_deserializeOpDeleteNodegroup) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2527,7 +2479,6 @@ func (m *awsRestjson1_deserializeOpDeleteNodegroup) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2651,10 +2602,8 @@ func (m *awsRestjson1_deserializeOpDeletePodIdentityAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2694,7 +2643,6 @@ func (m *awsRestjson1_deserializeOpDeletePodIdentityAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2812,10 +2760,8 @@ func (m *awsRestjson1_deserializeOpDeregisterCluster) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2855,7 +2801,6 @@ func (m *awsRestjson1_deserializeOpDeregisterCluster) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2979,10 +2924,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccessEntry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3022,7 +2965,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccessEntry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3137,10 +3079,8 @@ func (m *awsRestjson1_deserializeOpDescribeAddon) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3180,7 +3120,6 @@ func (m *awsRestjson1_deserializeOpDescribeAddon) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3301,10 +3240,8 @@ func (m *awsRestjson1_deserializeOpDescribeAddonConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3344,7 +3281,6 @@ func (m *awsRestjson1_deserializeOpDescribeAddonConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3486,10 +3422,8 @@ func (m *awsRestjson1_deserializeOpDescribeAddonVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3529,7 +3463,6 @@ func (m *awsRestjson1_deserializeOpDescribeAddonVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3653,10 +3586,8 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3696,7 +3627,6 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3814,10 +3744,8 @@ func (m *awsRestjson1_deserializeOpDescribeClusterVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3857,7 +3785,6 @@ func (m *awsRestjson1_deserializeOpDescribeClusterVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3981,10 +3908,8 @@ func (m *awsRestjson1_deserializeOpDescribeEksAnywhere
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4024,7 +3949,6 @@ func (m *awsRestjson1_deserializeOpDescribeEksAnywhere
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4142,10 +4066,8 @@ func (m *awsRestjson1_deserializeOpDescribeFargateProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4185,7 +4107,6 @@ func (m *awsRestjson1_deserializeOpDescribeFargateProf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4303,10 +4224,8 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4346,7 +4265,6 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4467,10 +4385,8 @@ func (m *awsRestjson1_deserializeOpDescribeInsight) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4510,7 +4426,6 @@ func (m *awsRestjson1_deserializeOpDescribeInsight) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4628,10 +4543,8 @@ func (m *awsRestjson1_deserializeOpDescribeNodegroup) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4671,7 +4584,6 @@ func (m *awsRestjson1_deserializeOpDescribeNodegroup) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4792,10 +4704,8 @@ func (m *awsRestjson1_deserializeOpDescribePodIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4835,7 +4745,6 @@ func (m *awsRestjson1_deserializeOpDescribePodIdentity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4953,10 +4862,8 @@ func (m *awsRestjson1_deserializeOpDescribeUpdate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4996,7 +4903,6 @@ func (m *awsRestjson1_deserializeOpDescribeUpdate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5114,10 +5020,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAccessP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5129,7 +5033,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAccessP
 	output := &DisassociateAccessPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5208,10 +5111,8 @@ func (m *awsRestjson1_deserializeOpDisassociateIdentit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5251,7 +5152,6 @@ func (m *awsRestjson1_deserializeOpDisassociateIdentit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5378,10 +5278,8 @@ func (m *awsRestjson1_deserializeOpListAccessEntries) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5421,7 +5319,6 @@ func (m *awsRestjson1_deserializeOpListAccessEntries) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5548,10 +5445,8 @@ func (m *awsRestjson1_deserializeOpListAccessPolicies)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5591,7 +5486,6 @@ func (m *awsRestjson1_deserializeOpListAccessPolicies)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5709,10 +5603,8 @@ func (m *awsRestjson1_deserializeOpListAddons) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5752,7 +5644,6 @@ func (m *awsRestjson1_deserializeOpListAddons) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5882,10 +5773,8 @@ func (m *awsRestjson1_deserializeOpListAssociatedAcces
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5925,7 +5814,6 @@ func (m *awsRestjson1_deserializeOpListAssociatedAcces
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6067,10 +5955,8 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6110,7 +5996,6 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6237,10 +6122,8 @@ func (m *awsRestjson1_deserializeOpListEksAnywhereSubs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6280,7 +6163,6 @@ func (m *awsRestjson1_deserializeOpListEksAnywhereSubs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6407,10 +6289,8 @@ func (m *awsRestjson1_deserializeOpListFargateProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6450,7 +6330,6 @@ func (m *awsRestjson1_deserializeOpListFargateProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6577,10 +6456,8 @@ func (m *awsRestjson1_deserializeOpListIdentityProvide
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6620,7 +6497,6 @@ func (m *awsRestjson1_deserializeOpListIdentityProvide
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6750,10 +6626,8 @@ func (m *awsRestjson1_deserializeOpListInsights) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6793,7 +6667,6 @@ func (m *awsRestjson1_deserializeOpListInsights) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6920,10 +6793,8 @@ func (m *awsRestjson1_deserializeOpListNodegroups) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6963,7 +6834,6 @@ func (m *awsRestjson1_deserializeOpListNodegroups) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7093,10 +6963,8 @@ func (m *awsRestjson1_deserializeOpListPodIdentityAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7136,7 +7004,6 @@ func (m *awsRestjson1_deserializeOpListPodIdentityAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7263,10 +7130,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7306,7 +7171,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7418,10 +7282,8 @@ func (m *awsRestjson1_deserializeOpListUpdates) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7461,7 +7323,6 @@ func (m *awsRestjson1_deserializeOpListUpdates) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7588,10 +7449,8 @@ func (m *awsRestjson1_deserializeOpRegisterCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7631,7 +7490,6 @@ func (m *awsRestjson1_deserializeOpRegisterCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7761,10 +7619,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7776,7 +7632,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7852,10 +7707,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7867,7 +7720,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7943,10 +7795,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccessEntry) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7986,7 +7836,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccessEntry) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8104,10 +7953,8 @@ func (m *awsRestjson1_deserializeOpUpdateAddon) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8147,7 +7994,6 @@ func (m *awsRestjson1_deserializeOpUpdateAddon) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8271,10 +8117,8 @@ func (m *awsRestjson1_deserializeOpUpdateClusterConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8314,7 +8158,6 @@ func (m *awsRestjson1_deserializeOpUpdateClusterConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8441,10 +8284,8 @@ func (m *awsRestjson1_deserializeOpUpdateClusterVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8484,7 +8325,6 @@ func (m *awsRestjson1_deserializeOpUpdateClusterVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8614,10 +8454,8 @@ func (m *awsRestjson1_deserializeOpUpdateEksAnywhereSu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8657,7 +8495,6 @@ func (m *awsRestjson1_deserializeOpUpdateEksAnywhereSu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8778,10 +8615,8 @@ func (m *awsRestjson1_deserializeOpUpdateNodegroupConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8821,7 +8656,6 @@ func (m *awsRestjson1_deserializeOpUpdateNodegroupConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8945,10 +8779,8 @@ func (m *awsRestjson1_deserializeOpUpdateNodegroupVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8988,7 +8820,6 @@ func (m *awsRestjson1_deserializeOpUpdateNodegroupVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9112,10 +8943,8 @@ func (m *awsRestjson1_deserializeOpUpdatePodIdentityAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9155,7 +8984,6 @@ func (m *awsRestjson1_deserializeOpUpdatePodIdentityAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
