--- vendor/github.com/aws/aws-sdk-go-v2/service/panorama/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/panorama/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateApplicationIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateApplicationIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -211,10 +207,8 @@ func (m *awsRestjson1_deserializeOpCreateJobForDevices
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -254,7 +248,6 @@ func (m *awsRestjson1_deserializeOpCreateJobForDevices
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -375,10 +368,8 @@ func (m *awsRestjson1_deserializeOpCreateNodeFromTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -418,7 +409,6 @@ func (m *awsRestjson1_deserializeOpCreateNodeFromTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -540,10 +530,8 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -583,7 +571,6 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -719,10 +706,8 @@ func (m *awsRestjson1_deserializeOpCreatePackageImport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -762,7 +747,6 @@ func (m *awsRestjson1_deserializeOpCreatePackageImport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -884,10 +868,8 @@ func (m *awsRestjson1_deserializeOpDeleteDevice) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -927,7 +909,6 @@ func (m *awsRestjson1_deserializeOpDeleteDevice) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1052,10 +1033,8 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1067,7 +1046,6 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 	output := &DeletePackageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1152,10 +1130,8 @@ func (m *awsRestjson1_deserializeOpDeregisterPackageVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1167,7 +1143,6 @@ func (m *awsRestjson1_deserializeOpDeregisterPackageVe
 	output := &DeregisterPackageVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1252,10 +1227,8 @@ func (m *awsRestjson1_deserializeOpDescribeApplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1295,7 +1268,6 @@ func (m *awsRestjson1_deserializeOpDescribeApplication
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1552,10 +1524,8 @@ func (m *awsRestjson1_deserializeOpDescribeApplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1595,7 +1565,6 @@ func (m *awsRestjson1_deserializeOpDescribeApplication
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1782,10 +1751,8 @@ func (m *awsRestjson1_deserializeOpDescribeDevice) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1825,7 +1792,6 @@ func (m *awsRestjson1_deserializeOpDescribeDevice) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2112,10 +2078,8 @@ func (m *awsRestjson1_deserializeOpDescribeDeviceJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2155,7 +2119,6 @@ func (m *awsRestjson1_deserializeOpDescribeDeviceJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2359,10 +2322,8 @@ func (m *awsRestjson1_deserializeOpDescribeNode) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2402,7 +2363,6 @@ func (m *awsRestjson1_deserializeOpDescribeNode) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2654,10 +2614,8 @@ func (m *awsRestjson1_deserializeOpDescribeNodeFromTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2697,7 +2655,6 @@ func (m *awsRestjson1_deserializeOpDescribeNodeFromTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2924,10 +2881,8 @@ func (m *awsRestjson1_deserializeOpDescribePackage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2967,7 +2922,6 @@ func (m *awsRestjson1_deserializeOpDescribePackage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3146,10 +3100,8 @@ func (m *awsRestjson1_deserializeOpDescribePackageImpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3189,7 +3141,6 @@ func (m *awsRestjson1_deserializeOpDescribePackageImpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3399,10 +3350,8 @@ func (m *awsRestjson1_deserializeOpDescribePackageVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3442,7 +3391,6 @@ func (m *awsRestjson1_deserializeOpDescribePackageVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3655,10 +3603,8 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3698,7 +3644,6 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3819,10 +3764,8 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3862,7 +3805,6 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3983,10 +3925,8 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4026,7 +3966,6 @@ func (m *awsRestjson1_deserializeOpListApplicationInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4147,10 +4086,8 @@ func (m *awsRestjson1_deserializeOpListDevices) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4190,7 +4127,6 @@ func (m *awsRestjson1_deserializeOpListDevices) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4317,10 +4253,8 @@ func (m *awsRestjson1_deserializeOpListDevicesJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4360,7 +4294,6 @@ func (m *awsRestjson1_deserializeOpListDevicesJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4490,10 +4423,8 @@ func (m *awsRestjson1_deserializeOpListNodeFromTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4533,7 +4464,6 @@ func (m *awsRestjson1_deserializeOpListNodeFromTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4660,10 +4590,8 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4703,7 +4631,6 @@ func (m *awsRestjson1_deserializeOpListNodes) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4827,10 +4754,8 @@ func (m *awsRestjson1_deserializeOpListPackageImportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4870,7 +4795,6 @@ func (m *awsRestjson1_deserializeOpListPackageImportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4997,10 +4921,8 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5040,7 +4962,6 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5170,10 +5091,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5213,7 +5132,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5328,10 +5246,8 @@ func (m *awsRestjson1_deserializeOpProvisionDevice) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5371,7 +5287,6 @@ func (m *awsRestjson1_deserializeOpProvisionDevice) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5536,10 +5451,8 @@ func (m *awsRestjson1_deserializeOpRegisterPackageVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5551,7 +5464,6 @@ func (m *awsRestjson1_deserializeOpRegisterPackageVers
 	output := &RegisterPackageVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5633,10 +5545,8 @@ func (m *awsRestjson1_deserializeOpRemoveApplicationIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5648,7 +5558,6 @@ func (m *awsRestjson1_deserializeOpRemoveApplicationIn
 	output := &RemoveApplicationInstanceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5733,10 +5642,8 @@ func (m *awsRestjson1_deserializeOpSignalApplicationIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5776,7 +5683,6 @@ func (m *awsRestjson1_deserializeOpSignalApplicationIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5898,10 +5804,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5913,7 +5817,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5992,10 +5895,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6007,7 +5908,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6086,10 +5986,8 @@ func (m *awsRestjson1_deserializeOpUpdateDeviceMetadat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6129,7 +6027,6 @@ func (m *awsRestjson1_deserializeOpUpdateDeviceMetadat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
