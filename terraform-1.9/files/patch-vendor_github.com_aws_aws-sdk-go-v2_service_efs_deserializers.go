--- vendor/github.com/aws/aws-sdk-go-v2/service/efs/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/efs/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessPoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessPoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -289,10 +285,8 @@ func (m *awsRestjson1_deserializeOpCreateFileSystem) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -332,7 +326,6 @@ func (m *awsRestjson1_deserializeOpCreateFileSystem) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -640,10 +633,8 @@ func (m *awsRestjson1_deserializeOpCreateMountTarget) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -683,7 +674,6 @@ func (m *awsRestjson1_deserializeOpCreateMountTarget) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -922,10 +912,8 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -965,7 +953,6 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1165,10 +1152,8 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1186,7 +1171,6 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1265,10 +1249,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessPoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1286,7 +1268,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessPoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1365,10 +1346,8 @@ func (m *awsRestjson1_deserializeOpDeleteFileSystem) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1386,7 +1365,6 @@ func (m *awsRestjson1_deserializeOpDeleteFileSystem) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1468,10 +1446,8 @@ func (m *awsRestjson1_deserializeOpDeleteFileSystemPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1489,7 +1465,6 @@ func (m *awsRestjson1_deserializeOpDeleteFileSystemPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1571,10 +1546,8 @@ func (m *awsRestjson1_deserializeOpDeleteMountTarget) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1592,7 +1565,6 @@ func (m *awsRestjson1_deserializeOpDeleteMountTarget) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1674,10 +1646,8 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1695,7 +1665,6 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1777,10 +1746,8 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1798,7 +1765,6 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1877,10 +1843,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccessPoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1920,7 +1884,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccessPoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2047,10 +2010,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountPref
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2090,7 +2051,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountPref
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2208,10 +2168,8 @@ func (m *awsRestjson1_deserializeOpDescribeBackupPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2251,7 +2209,6 @@ func (m *awsRestjson1_deserializeOpDescribeBackupPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2372,10 +2329,8 @@ func (m *awsRestjson1_deserializeOpDescribeFileSystemP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2415,7 +2370,6 @@ func (m *awsRestjson1_deserializeOpDescribeFileSystemP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2546,10 +2500,8 @@ func (m *awsRestjson1_deserializeOpDescribeFileSystems
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2589,7 +2541,6 @@ func (m *awsRestjson1_deserializeOpDescribeFileSystems
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2722,10 +2673,8 @@ func (m *awsRestjson1_deserializeOpDescribeLifecycleCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2765,7 +2714,6 @@ func (m *awsRestjson1_deserializeOpDescribeLifecycleCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2880,10 +2828,8 @@ func (m *awsRestjson1_deserializeOpDescribeMountTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2923,7 +2869,6 @@ func (m *awsRestjson1_deserializeOpDescribeMountTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3062,10 +3007,8 @@ func (m *awsRestjson1_deserializeOpDescribeMountTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3105,7 +3048,6 @@ func (m *awsRestjson1_deserializeOpDescribeMountTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3223,10 +3165,8 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3266,7 +3206,6 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3396,10 +3335,8 @@ func (m *awsRestjson1_deserializeOpDescribeTags) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3439,7 +3376,6 @@ func (m *awsRestjson1_deserializeOpDescribeTags) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3572,10 +3508,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3615,7 +3549,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3742,10 +3675,8 @@ func (m *awsRestjson1_deserializeOpModifyMountTargetSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3763,7 +3694,6 @@ func (m *awsRestjson1_deserializeOpModifyMountTargetSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3851,10 +3781,8 @@ func (m *awsRestjson1_deserializeOpPutAccountPreferenc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3894,7 +3822,6 @@ func (m *awsRestjson1_deserializeOpPutAccountPreferenc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4006,10 +3933,8 @@ func (m *awsRestjson1_deserializeOpPutBackupPolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4049,7 +3974,6 @@ func (m *awsRestjson1_deserializeOpPutBackupPolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4170,10 +4094,8 @@ func (m *awsRestjson1_deserializeOpPutFileSystemPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4213,7 +4135,6 @@ func (m *awsRestjson1_deserializeOpPutFileSystemPolicy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4347,10 +4268,8 @@ func (m *awsRestjson1_deserializeOpPutLifecycleConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4390,7 +4309,6 @@ func (m *awsRestjson1_deserializeOpPutLifecycleConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4508,10 +4426,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4529,7 +4445,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4611,10 +4526,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4632,7 +4545,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4714,10 +4626,8 @@ func (m *awsRestjson1_deserializeOpUpdateFileSystem) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4757,7 +4667,6 @@ func (m *awsRestjson1_deserializeOpUpdateFileSystem) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5065,10 +4974,8 @@ func (m *awsRestjson1_deserializeOpUpdateFileSystemPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5108,7 +5015,6 @@ func (m *awsRestjson1_deserializeOpUpdateFileSystemPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
