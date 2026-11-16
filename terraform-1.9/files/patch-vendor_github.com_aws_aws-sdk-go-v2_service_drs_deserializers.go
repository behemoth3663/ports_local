--- vendor/github.com/aws/aws-sdk-go-v2/service/drs/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/drs/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociateSourceNetw
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAssociateSourceNetw
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpCreateExtendedSourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -259,7 +253,6 @@ func (m *awsRestjson1_deserializeOpCreateExtendedSourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -386,10 +379,8 @@ func (m *awsRestjson1_deserializeOpCreateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -429,7 +420,6 @@ func (m *awsRestjson1_deserializeOpCreateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -553,10 +543,8 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -596,7 +584,6 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -856,10 +843,8 @@ func (m *awsRestjson1_deserializeOpCreateSourceNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -899,7 +884,6 @@ func (m *awsRestjson1_deserializeOpCreateSourceNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1030,10 +1014,8 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1045,7 +1027,6 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 	output := &DeleteJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1130,10 +1111,8 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchAction)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1145,7 +1124,6 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchAction)
 	output := &DeleteLaunchActionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1230,10 +1208,8 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1245,7 +1221,6 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchConfigu
 	output := &DeleteLaunchConfigurationTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1330,10 +1305,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1351,7 +1324,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1436,10 +1408,8 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1451,7 +1421,6 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 	output := &DeleteReplicationConfigurationTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1536,10 +1505,8 @@ func (m *awsRestjson1_deserializeOpDeleteSourceNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1551,7 +1518,6 @@ func (m *awsRestjson1_deserializeOpDeleteSourceNetwork
 	output := &DeleteSourceNetworkOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1636,10 +1602,8 @@ func (m *awsRestjson1_deserializeOpDeleteSourceServer)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1651,7 +1615,6 @@ func (m *awsRestjson1_deserializeOpDeleteSourceServer)
 	output := &DeleteSourceServerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1736,10 +1699,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobLogItems
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1779,7 +1740,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobLogItems
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1906,10 +1866,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1949,7 +1907,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2076,10 +2033,8 @@ func (m *awsRestjson1_deserializeOpDescribeLaunchConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2119,7 +2074,6 @@ func (m *awsRestjson1_deserializeOpDescribeLaunchConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2249,10 +2203,8 @@ func (m *awsRestjson1_deserializeOpDescribeRecoveryIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2292,7 +2244,6 @@ func (m *awsRestjson1_deserializeOpDescribeRecoveryIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2419,10 +2370,8 @@ func (m *awsRestjson1_deserializeOpDescribeRecoverySna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2462,7 +2411,6 @@ func (m *awsRestjson1_deserializeOpDescribeRecoverySna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2592,10 +2540,8 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2635,7 +2581,6 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2765,10 +2710,8 @@ func (m *awsRestjson1_deserializeOpDescribeSourceNetwo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2808,7 +2751,6 @@ func (m *awsRestjson1_deserializeOpDescribeSourceNetwo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2935,10 +2877,8 @@ func (m *awsRestjson1_deserializeOpDescribeSourceServe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2978,7 +2918,6 @@ func (m *awsRestjson1_deserializeOpDescribeSourceServe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3105,10 +3044,8 @@ func (m *awsRestjson1_deserializeOpDisconnectRecoveryI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3126,7 +3063,6 @@ func (m *awsRestjson1_deserializeOpDisconnectRecoveryI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3214,10 +3150,8 @@ func (m *awsRestjson1_deserializeOpDisconnectSourceSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3257,7 +3191,6 @@ func (m *awsRestjson1_deserializeOpDisconnectSourceSer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3475,10 +3408,8 @@ func (m *awsRestjson1_deserializeOpExportSourceNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3518,7 +3449,6 @@ func (m *awsRestjson1_deserializeOpExportSourceNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3646,10 +3576,8 @@ func (m *awsRestjson1_deserializeOpGetFailbackReplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3689,7 +3617,6 @@ func (m *awsRestjson1_deserializeOpGetFailbackReplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3842,10 +3769,8 @@ func (m *awsRestjson1_deserializeOpGetLaunchConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3885,7 +3810,6 @@ func (m *awsRestjson1_deserializeOpGetLaunchConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4080,10 +4004,8 @@ func (m *awsRestjson1_deserializeOpGetReplicationConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4123,7 +4045,6 @@ func (m *awsRestjson1_deserializeOpGetReplicationConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4380,10 +4301,8 @@ func (m *awsRestjson1_deserializeOpInitializeService) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4395,7 +4314,6 @@ func (m *awsRestjson1_deserializeOpInitializeService) 
 	output := &InitializeServiceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4477,10 +4395,8 @@ func (m *awsRestjson1_deserializeOpListExtensibleSourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4520,7 +4436,6 @@ func (m *awsRestjson1_deserializeOpListExtensibleSourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4650,10 +4565,8 @@ func (m *awsRestjson1_deserializeOpListLaunchActions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4693,7 +4606,6 @@ func (m *awsRestjson1_deserializeOpListLaunchActions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4823,10 +4735,8 @@ func (m *awsRestjson1_deserializeOpListStagingAccounts
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4866,7 +4776,6 @@ func (m *awsRestjson1_deserializeOpListStagingAccounts
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4996,10 +4905,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5039,7 +4946,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5160,10 +5066,8 @@ func (m *awsRestjson1_deserializeOpPutLaunchAction) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5203,7 +5107,6 @@ func (m *awsRestjson1_deserializeOpPutLaunchAction) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5430,10 +5333,8 @@ func (m *awsRestjson1_deserializeOpRetryDataReplicatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5473,7 +5374,6 @@ func (m *awsRestjson1_deserializeOpRetryDataReplicatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5691,10 +5591,8 @@ func (m *awsRestjson1_deserializeOpReverseReplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5734,7 +5632,6 @@ func (m *awsRestjson1_deserializeOpReverseReplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5865,10 +5762,8 @@ func (m *awsRestjson1_deserializeOpStartFailbackLaunch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5908,7 +5803,6 @@ func (m *awsRestjson1_deserializeOpStartFailbackLaunch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6032,10 +5926,8 @@ func (m *awsRestjson1_deserializeOpStartRecovery) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6075,7 +5967,6 @@ func (m *awsRestjson1_deserializeOpStartRecovery) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6196,10 +6087,8 @@ func (m *awsRestjson1_deserializeOpStartReplication) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6239,7 +6128,6 @@ func (m *awsRestjson1_deserializeOpStartReplication) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6360,10 +6248,8 @@ func (m *awsRestjson1_deserializeOpStartSourceNetworkR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6403,7 +6289,6 @@ func (m *awsRestjson1_deserializeOpStartSourceNetworkR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6527,10 +6412,8 @@ func (m *awsRestjson1_deserializeOpStartSourceNetworkR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6570,7 +6453,6 @@ func (m *awsRestjson1_deserializeOpStartSourceNetworkR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6691,10 +6573,8 @@ func (m *awsRestjson1_deserializeOpStopFailback) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6712,7 +6592,6 @@ func (m *awsRestjson1_deserializeOpStopFailback) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6794,10 +6673,8 @@ func (m *awsRestjson1_deserializeOpStopReplication) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6837,7 +6714,6 @@ func (m *awsRestjson1_deserializeOpStopReplication) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6958,10 +6834,8 @@ func (m *awsRestjson1_deserializeOpStopSourceNetworkRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7001,7 +6875,6 @@ func (m *awsRestjson1_deserializeOpStopSourceNetworkRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7125,10 +6998,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7146,7 +7017,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7231,10 +7101,8 @@ func (m *awsRestjson1_deserializeOpTerminateRecoveryIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7274,7 +7142,6 @@ func (m *awsRestjson1_deserializeOpTerminateRecoveryIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7395,10 +7262,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7416,7 +7281,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7501,10 +7365,8 @@ func (m *awsRestjson1_deserializeOpUpdateFailbackRepli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7522,7 +7384,6 @@ func (m *awsRestjson1_deserializeOpUpdateFailbackRepli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7607,10 +7468,8 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7650,7 +7509,6 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7851,10 +7709,8 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7894,7 +7750,6 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8018,10 +7873,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8061,7 +7914,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8324,10 +8176,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8367,7 +8217,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
