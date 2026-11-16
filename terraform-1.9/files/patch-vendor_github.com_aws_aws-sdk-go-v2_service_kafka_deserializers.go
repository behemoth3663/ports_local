--- vendor/github.com/aws/aws-sdk-go-v2/service/kafka/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/kafka/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateScram
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateScram
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -226,10 +222,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateSc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -269,7 +263,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateSc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -405,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -448,7 +439,6 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -597,10 +587,8 @@ func (m *awsRestjson1_deserializeOpCreateClusterV2) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -640,7 +628,6 @@ func (m *awsRestjson1_deserializeOpCreateClusterV2) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -798,10 +785,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -841,7 +826,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1008,10 +992,8 @@ func (m *awsRestjson1_deserializeOpCreateReplicator) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1051,7 +1033,6 @@ func (m *awsRestjson1_deserializeOpCreateReplicator) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1203,10 +1184,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcConnection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1246,7 +1225,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcConnection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1429,10 +1407,8 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1472,7 +1448,6 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1603,10 +1578,8 @@ func (m *awsRestjson1_deserializeOpDeleteClusterPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1618,7 +1591,6 @@ func (m *awsRestjson1_deserializeOpDeleteClusterPolicy
 	output := &DeleteClusterPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1700,10 +1672,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1743,7 +1713,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1874,10 +1843,8 @@ func (m *awsRestjson1_deserializeOpDeleteReplicator) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1917,7 +1884,6 @@ func (m *awsRestjson1_deserializeOpDeleteReplicator) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2057,10 +2023,8 @@ func (m *awsRestjson1_deserializeOpDeleteVpcConnection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2100,7 +2064,6 @@ func (m *awsRestjson1_deserializeOpDeleteVpcConnection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2231,10 +2194,8 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2274,7 +2235,6 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2395,10 +2355,8 @@ func (m *awsRestjson1_deserializeOpDescribeClusterOper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2438,7 +2396,6 @@ func (m *awsRestjson1_deserializeOpDescribeClusterOper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2559,10 +2516,8 @@ func (m *awsRestjson1_deserializeOpDescribeClusterOper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2602,7 +2557,6 @@ func (m *awsRestjson1_deserializeOpDescribeClusterOper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2729,10 +2683,8 @@ func (m *awsRestjson1_deserializeOpDescribeClusterV2) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2772,7 +2724,6 @@ func (m *awsRestjson1_deserializeOpDescribeClusterV2) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2893,10 +2844,8 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2936,7 +2885,6 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3114,10 +3062,8 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3157,7 +3103,6 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3333,10 +3278,8 @@ func (m *awsRestjson1_deserializeOpDescribeReplicator)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3376,7 +3319,6 @@ func (m *awsRestjson1_deserializeOpDescribeReplicator)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3603,10 +3545,8 @@ func (m *awsRestjson1_deserializeOpDescribeVpcConnecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3646,7 +3586,6 @@ func (m *awsRestjson1_deserializeOpDescribeVpcConnecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3838,10 +3777,8 @@ func (m *awsRestjson1_deserializeOpGetBootstrapBrokers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3881,7 +3818,6 @@ func (m *awsRestjson1_deserializeOpGetBootstrapBrokers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4087,10 +4023,8 @@ func (m *awsRestjson1_deserializeOpGetClusterPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4130,7 +4064,6 @@ func (m *awsRestjson1_deserializeOpGetClusterPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4261,10 +4194,8 @@ func (m *awsRestjson1_deserializeOpGetCompatibleKafkaV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4304,7 +4235,6 @@ func (m *awsRestjson1_deserializeOpGetCompatibleKafkaV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4431,10 +4361,8 @@ func (m *awsRestjson1_deserializeOpListClientVpcConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4474,7 +4402,6 @@ func (m *awsRestjson1_deserializeOpListClientVpcConnec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4604,10 +4531,8 @@ func (m *awsRestjson1_deserializeOpListClusterOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4647,7 +4572,6 @@ func (m *awsRestjson1_deserializeOpListClusterOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4774,10 +4698,8 @@ func (m *awsRestjson1_deserializeOpListClusterOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4817,7 +4739,6 @@ func (m *awsRestjson1_deserializeOpListClusterOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4953,10 +4874,8 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4996,7 +4915,6 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5123,10 +5041,8 @@ func (m *awsRestjson1_deserializeOpListClustersV2) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5166,7 +5082,6 @@ func (m *awsRestjson1_deserializeOpListClustersV2) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5293,10 +5208,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5336,7 +5249,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5469,10 +5381,8 @@ func (m *awsRestjson1_deserializeOpListConfigurations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5512,7 +5422,6 @@ func (m *awsRestjson1_deserializeOpListConfigurations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5642,10 +5551,8 @@ func (m *awsRestjson1_deserializeOpListKafkaVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5685,7 +5592,6 @@ func (m *awsRestjson1_deserializeOpListKafkaVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5812,10 +5718,8 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5855,7 +5759,6 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5982,10 +5885,8 @@ func (m *awsRestjson1_deserializeOpListReplicators) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6025,7 +5926,6 @@ func (m *awsRestjson1_deserializeOpListReplicators) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6161,10 +6061,8 @@ func (m *awsRestjson1_deserializeOpListScramSecrets) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6204,7 +6102,6 @@ func (m *awsRestjson1_deserializeOpListScramSecrets) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6340,10 +6237,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6383,7 +6278,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6498,10 +6392,8 @@ func (m *awsRestjson1_deserializeOpListVpcConnections)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6541,7 +6433,6 @@ func (m *awsRestjson1_deserializeOpListVpcConnections)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6671,10 +6562,8 @@ func (m *awsRestjson1_deserializeOpPutClusterPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6714,7 +6603,6 @@ func (m *awsRestjson1_deserializeOpPutClusterPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6833,10 +6721,8 @@ func (m *awsRestjson1_deserializeOpRebootBroker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6876,7 +6762,6 @@ func (m *awsRestjson1_deserializeOpRebootBroker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7016,10 +6901,8 @@ func (m *awsRestjson1_deserializeOpRejectClientVpcConn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7031,7 +6914,6 @@ func (m *awsRestjson1_deserializeOpRejectClientVpcConn
 	output := &RejectClientVpcConnectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7116,10 +6998,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7137,7 +7017,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7216,10 +7095,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7237,7 +7114,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7316,10 +7192,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerCount) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7359,7 +7233,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerCount) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7493,10 +7366,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerStorage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7536,7 +7407,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerStorage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7670,10 +7540,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerType) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7713,7 +7581,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrokerType) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7853,10 +7720,8 @@ func (m *awsRestjson1_deserializeOpUpdateClusterConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7896,7 +7761,6 @@ func (m *awsRestjson1_deserializeOpUpdateClusterConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8033,10 +7897,8 @@ func (m *awsRestjson1_deserializeOpUpdateClusterKafkaV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8076,7 +7938,6 @@ func (m *awsRestjson1_deserializeOpUpdateClusterKafkaV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8216,10 +8077,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8259,7 +8118,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8392,10 +8250,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivity)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8435,7 +8291,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectivity)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8572,10 +8427,8 @@ func (m *awsRestjson1_deserializeOpUpdateMonitoring) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8615,7 +8468,6 @@ func (m *awsRestjson1_deserializeOpUpdateMonitoring) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8749,10 +8601,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8792,7 +8642,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8932,10 +8781,8 @@ func (m *awsRestjson1_deserializeOpUpdateSecurity) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8975,7 +8822,6 @@ func (m *awsRestjson1_deserializeOpUpdateSecurity) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9115,10 +8961,8 @@ func (m *awsRestjson1_deserializeOpUpdateStorage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9158,7 +9002,6 @@ func (m *awsRestjson1_deserializeOpUpdateStorage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
