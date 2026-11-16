--- vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestxml_deserializeOpAssociateAlias) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -69,7 +66,6 @@ func (m *awsRestxml_deserializeOpAssociateAlias) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -138,10 +134,8 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -186,7 +180,6 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -329,10 +322,8 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -377,7 +368,6 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -520,10 +510,8 @@ func (m *awsRestxml_deserializeOpCopyDistribution) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -568,7 +556,6 @@ func (m *awsRestxml_deserializeOpCopyDistribution) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -873,10 +860,8 @@ func (m *awsRestxml_deserializeOpCreateAnycastIpList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -921,7 +906,6 @@ func (m *awsRestxml_deserializeOpCreateAnycastIpList) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1047,10 +1031,8 @@ func (m *awsRestxml_deserializeOpCreateCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1095,7 +1077,6 @@ func (m *awsRestxml_deserializeOpCreateCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1232,10 +1213,8 @@ func (m *awsRestxml_deserializeOpCreateCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1280,7 +1259,6 @@ func (m *awsRestxml_deserializeOpCreateCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1408,10 +1386,8 @@ func (m *awsRestxml_deserializeOpCreateConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1456,7 +1432,6 @@ func (m *awsRestxml_deserializeOpCreateConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1582,10 +1557,8 @@ func (m *awsRestxml_deserializeOpCreateContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1630,7 +1603,6 @@ func (m *awsRestxml_deserializeOpCreateContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1761,10 +1733,8 @@ func (m *awsRestxml_deserializeOpCreateDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1809,7 +1779,6 @@ func (m *awsRestxml_deserializeOpCreateDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2123,10 +2092,8 @@ func (m *awsRestxml_deserializeOpCreateDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2171,7 +2138,6 @@ func (m *awsRestxml_deserializeOpCreateDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2303,10 +2269,8 @@ func (m *awsRestxml_deserializeOpCreateDistributionWit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2351,7 +2315,6 @@ func (m *awsRestxml_deserializeOpCreateDistributionWit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2665,10 +2628,8 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2713,7 +2674,6 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2850,10 +2810,8 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2898,7 +2856,6 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3035,10 +2992,8 @@ func (m *awsRestxml_deserializeOpCreateFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3083,7 +3038,6 @@ func (m *awsRestxml_deserializeOpCreateFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3211,10 +3165,8 @@ func (m *awsRestxml_deserializeOpCreateInvalidation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3259,7 +3211,6 @@ func (m *awsRestxml_deserializeOpCreateInvalidation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3388,10 +3339,8 @@ func (m *awsRestxml_deserializeOpCreateInvalidationFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3436,7 +3385,6 @@ func (m *awsRestxml_deserializeOpCreateInvalidationFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3565,10 +3513,8 @@ func (m *awsRestxml_deserializeOpCreateKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3613,7 +3559,6 @@ func (m *awsRestxml_deserializeOpCreateKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3738,10 +3683,8 @@ func (m *awsRestxml_deserializeOpCreateKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3786,7 +3729,6 @@ func (m *awsRestxml_deserializeOpCreateKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3917,10 +3859,8 @@ func (m *awsRestxml_deserializeOpCreateMonitoringSubsc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3960,7 +3900,6 @@ func (m *awsRestxml_deserializeOpCreateMonitoringSubsc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4068,10 +4007,8 @@ func (m *awsRestxml_deserializeOpCreateOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4116,7 +4053,6 @@ func (m *awsRestxml_deserializeOpCreateOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4238,10 +4174,8 @@ func (m *awsRestxml_deserializeOpCreateOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4286,7 +4220,6 @@ func (m *awsRestxml_deserializeOpCreateOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4423,10 +4356,8 @@ func (m *awsRestxml_deserializeOpCreatePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4471,7 +4402,6 @@ func (m *awsRestxml_deserializeOpCreatePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4593,10 +4523,8 @@ func (m *awsRestxml_deserializeOpCreateRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4636,7 +4564,6 @@ func (m *awsRestxml_deserializeOpCreateRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4744,10 +4671,8 @@ func (m *awsRestxml_deserializeOpCreateResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4792,7 +4717,6 @@ func (m *awsRestxml_deserializeOpCreateResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4929,10 +4853,8 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4977,7 +4899,6 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5129,10 +5050,8 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5177,7 +5096,6 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5332,10 +5250,8 @@ func (m *awsRestxml_deserializeOpCreateVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5380,7 +5296,6 @@ func (m *awsRestxml_deserializeOpCreateVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5514,10 +5429,8 @@ func (m *awsRestxml_deserializeOpDeleteAnycastIpList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5535,7 +5448,6 @@ func (m *awsRestxml_deserializeOpDeleteAnycastIpList) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5613,10 +5525,8 @@ func (m *awsRestxml_deserializeOpDeleteCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5634,7 +5544,6 @@ func (m *awsRestxml_deserializeOpDeleteCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5706,10 +5615,8 @@ func (m *awsRestxml_deserializeOpDeleteCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5727,7 +5634,6 @@ func (m *awsRestxml_deserializeOpDeleteCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5796,10 +5702,8 @@ func (m *awsRestxml_deserializeOpDeleteConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5817,7 +5721,6 @@ func (m *awsRestxml_deserializeOpDeleteConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5889,10 +5792,8 @@ func (m *awsRestxml_deserializeOpDeleteContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5910,7 +5811,6 @@ func (m *awsRestxml_deserializeOpDeleteContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5982,10 +5882,8 @@ func (m *awsRestxml_deserializeOpDeleteDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6003,7 +5901,6 @@ func (m *awsRestxml_deserializeOpDeleteDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6075,10 +5972,8 @@ func (m *awsRestxml_deserializeOpDeleteDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6096,7 +5991,6 @@ func (m *awsRestxml_deserializeOpDeleteDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6165,10 +6059,8 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6186,7 +6078,6 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6255,10 +6146,8 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6276,7 +6165,6 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6345,10 +6233,8 @@ func (m *awsRestxml_deserializeOpDeleteFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6366,7 +6252,6 @@ func (m *awsRestxml_deserializeOpDeleteFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6435,10 +6320,8 @@ func (m *awsRestxml_deserializeOpDeleteKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6456,7 +6339,6 @@ func (m *awsRestxml_deserializeOpDeleteKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6522,10 +6404,8 @@ func (m *awsRestxml_deserializeOpDeleteKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6543,7 +6423,6 @@ func (m *awsRestxml_deserializeOpDeleteKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6615,10 +6494,8 @@ func (m *awsRestxml_deserializeOpDeleteMonitoringSubsc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6630,7 +6507,6 @@ func (m *awsRestxml_deserializeOpDeleteMonitoringSubsc
 	output := &DeleteMonitoringSubscriptionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6696,10 +6572,8 @@ func (m *awsRestxml_deserializeOpDeleteOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6717,7 +6591,6 @@ func (m *awsRestxml_deserializeOpDeleteOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6786,10 +6659,8 @@ func (m *awsRestxml_deserializeOpDeleteOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6807,7 +6678,6 @@ func (m *awsRestxml_deserializeOpDeleteOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6879,10 +6749,8 @@ func (m *awsRestxml_deserializeOpDeletePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6900,7 +6768,6 @@ func (m *awsRestxml_deserializeOpDeletePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6969,10 +6836,8 @@ func (m *awsRestxml_deserializeOpDeleteRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6990,7 +6855,6 @@ func (m *awsRestxml_deserializeOpDeleteRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7056,10 +6920,8 @@ func (m *awsRestxml_deserializeOpDeleteResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7077,7 +6939,6 @@ func (m *awsRestxml_deserializeOpDeleteResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7149,10 +7010,8 @@ func (m *awsRestxml_deserializeOpDeleteStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7170,7 +7029,6 @@ func (m *awsRestxml_deserializeOpDeleteStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7239,10 +7097,8 @@ func (m *awsRestxml_deserializeOpDeleteVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7287,7 +7143,6 @@ func (m *awsRestxml_deserializeOpDeleteVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7419,10 +7274,8 @@ func (m *awsRestxml_deserializeOpDescribeFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7467,7 +7320,6 @@ func (m *awsRestxml_deserializeOpDescribeFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7581,10 +7433,8 @@ func (m *awsRestxml_deserializeOpDescribeKeyValueStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7629,7 +7479,6 @@ func (m *awsRestxml_deserializeOpDescribeKeyValueStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7749,10 +7598,8 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7797,7 +7644,6 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7927,10 +7773,8 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7975,7 +7819,6 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8105,10 +7948,8 @@ func (m *awsRestxml_deserializeOpGetAnycastIpList) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8153,7 +7994,6 @@ func (m *awsRestxml_deserializeOpGetAnycastIpList) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8273,10 +8113,8 @@ func (m *awsRestxml_deserializeOpGetCachePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8321,7 +8159,6 @@ func (m *awsRestxml_deserializeOpGetCachePolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8435,10 +8272,8 @@ func (m *awsRestxml_deserializeOpGetCachePolicyConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8483,7 +8318,6 @@ func (m *awsRestxml_deserializeOpGetCachePolicyConfig)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8597,10 +8431,8 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8645,7 +8477,6 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8759,10 +8590,8 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8807,7 +8636,6 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8921,10 +8749,8 @@ func (m *awsRestxml_deserializeOpGetConnectionGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8969,7 +8795,6 @@ func (m *awsRestxml_deserializeOpGetConnectionGroup) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9083,10 +8908,8 @@ func (m *awsRestxml_deserializeOpGetConnectionGroupByR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9131,7 +8954,6 @@ func (m *awsRestxml_deserializeOpGetConnectionGroupByR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9245,10 +9067,8 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9293,7 +9113,6 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9407,10 +9226,8 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9455,7 +9272,6 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9569,10 +9385,8 @@ func (m *awsRestxml_deserializeOpGetDistribution) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9617,7 +9431,6 @@ func (m *awsRestxml_deserializeOpGetDistribution) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9731,10 +9544,8 @@ func (m *awsRestxml_deserializeOpGetDistributionConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9779,7 +9590,6 @@ func (m *awsRestxml_deserializeOpGetDistributionConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9893,10 +9703,8 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9941,7 +9749,6 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10055,10 +9862,8 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10103,7 +9908,6 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10217,10 +10021,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10265,7 +10067,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10379,10 +10180,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10427,7 +10226,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10541,10 +10339,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10589,7 +10385,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10703,10 +10498,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10751,7 +10544,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10865,10 +10657,8 @@ func (m *awsRestxml_deserializeOpGetFunction) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10890,7 +10680,6 @@ func (m *awsRestxml_deserializeOpGetFunction) HandleDe
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10988,10 +10777,8 @@ func (m *awsRestxml_deserializeOpGetInvalidation) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11031,7 +10818,6 @@ func (m *awsRestxml_deserializeOpGetInvalidation) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11136,10 +10922,8 @@ func (m *awsRestxml_deserializeOpGetInvalidationForDis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11179,7 +10963,6 @@ func (m *awsRestxml_deserializeOpGetInvalidationForDis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11284,10 +11067,8 @@ func (m *awsRestxml_deserializeOpGetKeyGroup) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11332,7 +11113,6 @@ func (m *awsRestxml_deserializeOpGetKeyGroup) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11443,10 +11223,8 @@ func (m *awsRestxml_deserializeOpGetKeyGroupConfig) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11491,7 +11269,6 @@ func (m *awsRestxml_deserializeOpGetKeyGroupConfig) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11602,10 +11379,8 @@ func (m *awsRestxml_deserializeOpGetManagedCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11645,7 +11420,6 @@ func (m *awsRestxml_deserializeOpGetManagedCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11747,10 +11521,8 @@ func (m *awsRestxml_deserializeOpGetMonitoringSubscrip
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11790,7 +11562,6 @@ func (m *awsRestxml_deserializeOpGetMonitoringSubscrip
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11898,10 +11669,8 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11946,7 +11715,6 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12060,10 +11828,8 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12108,7 +11874,6 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12222,10 +11987,8 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12270,7 +12033,6 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12384,10 +12146,8 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12432,7 +12192,6 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12546,10 +12305,8 @@ func (m *awsRestxml_deserializeOpGetPublicKey) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12594,7 +12351,6 @@ func (m *awsRestxml_deserializeOpGetPublicKey) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12708,10 +12464,8 @@ func (m *awsRestxml_deserializeOpGetPublicKeyConfig) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12756,7 +12510,6 @@ func (m *awsRestxml_deserializeOpGetPublicKeyConfig) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12870,10 +12623,8 @@ func (m *awsRestxml_deserializeOpGetRealtimeLogConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12913,7 +12664,6 @@ func (m *awsRestxml_deserializeOpGetRealtimeLogConfig)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13018,10 +12768,8 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13066,7 +12814,6 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13180,10 +12927,8 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13228,7 +12973,6 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13342,10 +13086,8 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13390,7 +13132,6 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13504,10 +13245,8 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13552,7 +13291,6 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13666,10 +13404,8 @@ func (m *awsRestxml_deserializeOpGetVpcOrigin) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13714,7 +13450,6 @@ func (m *awsRestxml_deserializeOpGetVpcOrigin) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13834,10 +13569,8 @@ func (m *awsRestxml_deserializeOpListAnycastIpLists) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13877,7 +13610,6 @@ func (m *awsRestxml_deserializeOpListAnycastIpLists) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13985,10 +13717,8 @@ func (m *awsRestxml_deserializeOpListCachePolicies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14028,7 +13758,6 @@ func (m *awsRestxml_deserializeOpListCachePolicies) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14133,10 +13862,8 @@ func (m *awsRestxml_deserializeOpListCloudFrontOriginA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14176,7 +13903,6 @@ func (m *awsRestxml_deserializeOpListCloudFrontOriginA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14275,10 +14001,8 @@ func (m *awsRestxml_deserializeOpListConflictingAliase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14318,7 +14042,6 @@ func (m *awsRestxml_deserializeOpListConflictingAliase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14420,10 +14143,8 @@ func (m *awsRestxml_deserializeOpListConnectionGroups)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14463,7 +14184,6 @@ func (m *awsRestxml_deserializeOpListConnectionGroups)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14581,10 +14301,8 @@ func (m *awsRestxml_deserializeOpListContinuousDeploym
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14624,7 +14342,6 @@ func (m *awsRestxml_deserializeOpListContinuousDeploym
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14729,10 +14446,8 @@ func (m *awsRestxml_deserializeOpListDistributions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14772,7 +14487,6 @@ func (m *awsRestxml_deserializeOpListDistributions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14871,10 +14585,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14914,7 +14626,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15022,10 +14733,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByCa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15065,7 +14774,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByCa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15170,10 +14878,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15213,7 +14919,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15315,10 +15020,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByKe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15358,7 +15061,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByKe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15460,10 +15162,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15503,7 +15203,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15608,10 +15307,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15651,7 +15348,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15750,10 +15446,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15793,7 +15487,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15898,10 +15591,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByVp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15941,7 +15632,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByVp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16049,10 +15739,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByWe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16092,7 +15780,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByWe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16194,10 +15881,8 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16237,7 +15922,6 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16355,10 +16039,8 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16398,7 +16080,6 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16516,10 +16197,8 @@ func (m *awsRestxml_deserializeOpListDomainConflicts) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16559,7 +16238,6 @@ func (m *awsRestxml_deserializeOpListDomainConflicts) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16677,10 +16355,8 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16720,7 +16396,6 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16819,10 +16494,8 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16862,7 +16535,6 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16961,10 +16633,8 @@ func (m *awsRestxml_deserializeOpListFunctions) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17004,7 +16674,6 @@ func (m *awsRestxml_deserializeOpListFunctions) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17106,10 +16775,8 @@ func (m *awsRestxml_deserializeOpListInvalidations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17149,7 +16816,6 @@ func (m *awsRestxml_deserializeOpListInvalidations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17254,10 +16920,8 @@ func (m *awsRestxml_deserializeOpListInvalidationsForD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17297,7 +16961,6 @@ func (m *awsRestxml_deserializeOpListInvalidationsForD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17402,10 +17065,8 @@ func (m *awsRestxml_deserializeOpListKeyGroups) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17445,7 +17106,6 @@ func (m *awsRestxml_deserializeOpListKeyGroups) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17544,10 +17204,8 @@ func (m *awsRestxml_deserializeOpListKeyValueStores) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17587,7 +17245,6 @@ func (m *awsRestxml_deserializeOpListKeyValueStores) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17692,10 +17349,8 @@ func (m *awsRestxml_deserializeOpListOriginAccessContr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17735,7 +17390,6 @@ func (m *awsRestxml_deserializeOpListOriginAccessContr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17834,10 +17488,8 @@ func (m *awsRestxml_deserializeOpListOriginRequestPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17877,7 +17529,6 @@ func (m *awsRestxml_deserializeOpListOriginRequestPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17982,10 +17633,8 @@ func (m *awsRestxml_deserializeOpListPublicKeys) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18025,7 +17674,6 @@ func (m *awsRestxml_deserializeOpListPublicKeys) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18124,10 +17772,8 @@ func (m *awsRestxml_deserializeOpListRealtimeLogConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18167,7 +17813,6 @@ func (m *awsRestxml_deserializeOpListRealtimeLogConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18272,10 +17917,8 @@ func (m *awsRestxml_deserializeOpListResponseHeadersPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18315,7 +17958,6 @@ func (m *awsRestxml_deserializeOpListResponseHeadersPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18420,10 +18062,8 @@ func (m *awsRestxml_deserializeOpListStreamingDistribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18463,7 +18103,6 @@ func (m *awsRestxml_deserializeOpListStreamingDistribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18562,10 +18201,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18605,7 +18242,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18713,10 +18349,8 @@ func (m *awsRestxml_deserializeOpListVpcOrigins) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18756,7 +18390,6 @@ func (m *awsRestxml_deserializeOpListVpcOrigins) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18864,10 +18497,8 @@ func (m *awsRestxml_deserializeOpPublishFunction) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18907,7 +18538,6 @@ func (m *awsRestxml_deserializeOpPublishFunction) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19018,10 +18648,8 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19039,7 +18667,6 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19105,10 +18732,8 @@ func (m *awsRestxml_deserializeOpTestFunction) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19148,7 +18773,6 @@ func (m *awsRestxml_deserializeOpTestFunction) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19259,10 +18883,8 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19280,7 +18902,6 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19346,10 +18967,8 @@ func (m *awsRestxml_deserializeOpUpdateCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19394,7 +19013,6 @@ func (m *awsRestxml_deserializeOpUpdateCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19535,10 +19153,8 @@ func (m *awsRestxml_deserializeOpUpdateCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19583,7 +19199,6 @@ func (m *awsRestxml_deserializeOpUpdateCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19715,10 +19330,8 @@ func (m *awsRestxml_deserializeOpUpdateConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19763,7 +19376,6 @@ func (m *awsRestxml_deserializeOpUpdateConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19895,10 +19507,8 @@ func (m *awsRestxml_deserializeOpUpdateContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19943,7 +19553,6 @@ func (m *awsRestxml_deserializeOpUpdateContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20072,10 +19681,8 @@ func (m *awsRestxml_deserializeOpUpdateDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20120,7 +19727,6 @@ func (m *awsRestxml_deserializeOpUpdateDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20429,10 +20035,8 @@ func (m *awsRestxml_deserializeOpUpdateDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20477,7 +20081,6 @@ func (m *awsRestxml_deserializeOpUpdateDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20612,10 +20215,8 @@ func (m *awsRestxml_deserializeOpUpdateDistributionWit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20660,7 +20261,6 @@ func (m *awsRestxml_deserializeOpUpdateDistributionWit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20954,10 +20554,8 @@ func (m *awsRestxml_deserializeOpUpdateDomainAssociati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21002,7 +20600,6 @@ func (m *awsRestxml_deserializeOpUpdateDomainAssociati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21148,10 +20745,8 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21196,7 +20791,6 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21337,10 +20931,8 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21385,7 +20977,6 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21529,10 +21120,8 @@ func (m *awsRestxml_deserializeOpUpdateFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21577,7 +21166,6 @@ func (m *awsRestxml_deserializeOpUpdateFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21703,10 +21291,8 @@ func (m *awsRestxml_deserializeOpUpdateKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21751,7 +21337,6 @@ func (m *awsRestxml_deserializeOpUpdateKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21877,10 +21462,8 @@ func (m *awsRestxml_deserializeOpUpdateKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21925,7 +21508,6 @@ func (m *awsRestxml_deserializeOpUpdateKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22051,10 +21633,8 @@ func (m *awsRestxml_deserializeOpUpdateOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22099,7 +21679,6 @@ func (m *awsRestxml_deserializeOpUpdateOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22228,10 +21807,8 @@ func (m *awsRestxml_deserializeOpUpdateOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22276,7 +21853,6 @@ func (m *awsRestxml_deserializeOpUpdateOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22417,10 +21993,8 @@ func (m *awsRestxml_deserializeOpUpdatePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22465,7 +22039,6 @@ func (m *awsRestxml_deserializeOpUpdatePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22594,10 +22167,8 @@ func (m *awsRestxml_deserializeOpUpdateRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22637,7 +22208,6 @@ func (m *awsRestxml_deserializeOpUpdateRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22742,10 +22312,8 @@ func (m *awsRestxml_deserializeOpUpdateResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22790,7 +22358,6 @@ func (m *awsRestxml_deserializeOpUpdateResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22931,10 +22498,8 @@ func (m *awsRestxml_deserializeOpUpdateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22979,7 +22544,6 @@ func (m *awsRestxml_deserializeOpUpdateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23129,10 +22693,8 @@ func (m *awsRestxml_deserializeOpUpdateVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23177,7 +22739,6 @@ func (m *awsRestxml_deserializeOpUpdateVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23318,10 +22879,8 @@ func (m *awsRestxml_deserializeOpVerifyDnsConfiguratio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23361,7 +22920,6 @@ func (m *awsRestxml_deserializeOpVerifyDnsConfiguratio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
