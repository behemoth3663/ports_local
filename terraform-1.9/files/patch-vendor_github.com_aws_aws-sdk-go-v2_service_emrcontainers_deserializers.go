--- vendor/github.com/aws/aws-sdk-go-v2/service/emrcontainers/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/emrcontainers/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCancelJobRun) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCancelJobRun) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpCreateJobTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpCreateJobTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -405,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateManagedEndpoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -448,7 +439,6 @@ func (m *awsRestjson1_deserializeOpCreateManagedEndpoi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -594,10 +584,8 @@ func (m *awsRestjson1_deserializeOpCreateSecurityConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -637,7 +625,6 @@ func (m *awsRestjson1_deserializeOpCreateSecurityConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -771,10 +758,8 @@ func (m *awsRestjson1_deserializeOpCreateVirtualCluste
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -814,7 +799,6 @@ func (m *awsRestjson1_deserializeOpCreateVirtualCluste
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -954,10 +938,8 @@ func (m *awsRestjson1_deserializeOpDeleteJobTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -997,7 +979,6 @@ func (m *awsRestjson1_deserializeOpDeleteJobTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1113,10 +1094,8 @@ func (m *awsRestjson1_deserializeOpDeleteManagedEndpoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1156,7 +1135,6 @@ func (m *awsRestjson1_deserializeOpDeleteManagedEndpoi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1281,10 +1259,8 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualCluste
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1324,7 +1300,6 @@ func (m *awsRestjson1_deserializeOpDeleteVirtualCluste
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1440,10 +1415,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobRun) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1483,7 +1456,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobRun) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1598,10 +1570,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1641,7 +1611,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1756,10 +1725,8 @@ func (m *awsRestjson1_deserializeOpDescribeManagedEndp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1799,7 +1766,6 @@ func (m *awsRestjson1_deserializeOpDescribeManagedEndp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1914,10 +1880,8 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1957,7 +1921,6 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2072,10 +2035,8 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualClus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2115,7 +2076,6 @@ func (m *awsRestjson1_deserializeOpDescribeVirtualClus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2230,10 +2190,8 @@ func (m *awsRestjson1_deserializeOpGetManagedEndpointS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2273,7 +2231,6 @@ func (m *awsRestjson1_deserializeOpGetManagedEndpointS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2413,10 +2370,8 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2456,7 +2411,6 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2577,10 +2531,8 @@ func (m *awsRestjson1_deserializeOpListJobTemplates) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2620,7 +2572,6 @@ func (m *awsRestjson1_deserializeOpListJobTemplates) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2741,10 +2692,8 @@ func (m *awsRestjson1_deserializeOpListManagedEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2784,7 +2733,6 @@ func (m *awsRestjson1_deserializeOpListManagedEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2905,10 +2853,8 @@ func (m *awsRestjson1_deserializeOpListSecurityConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2948,7 +2894,6 @@ func (m *awsRestjson1_deserializeOpListSecurityConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3069,10 +3014,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3112,7 +3055,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3227,10 +3169,8 @@ func (m *awsRestjson1_deserializeOpListVirtualClusters
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3270,7 +3210,6 @@ func (m *awsRestjson1_deserializeOpListVirtualClusters
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3391,10 +3330,8 @@ func (m *awsRestjson1_deserializeOpStartJobRun) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3434,7 +3371,6 @@ func (m *awsRestjson1_deserializeOpStartJobRun) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3580,10 +3516,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3595,7 +3529,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3674,10 +3607,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3689,7 +3620,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
