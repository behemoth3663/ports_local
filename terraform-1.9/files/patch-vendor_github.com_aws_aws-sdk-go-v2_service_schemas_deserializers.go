--- vendor/github.com/aws/aws-sdk-go-v2/service/schemas/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/schemas/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateDiscoverer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateDiscoverer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -266,10 +262,8 @@ func (m *awsRestjson1_deserializeOpCreateRegistry) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -309,7 +303,6 @@ func (m *awsRestjson1_deserializeOpCreateRegistry) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -460,10 +453,8 @@ func (m *awsRestjson1_deserializeOpCreateSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -503,7 +494,6 @@ func (m *awsRestjson1_deserializeOpCreateSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -692,10 +682,8 @@ func (m *awsRestjson1_deserializeOpDeleteDiscoverer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -713,7 +701,6 @@ func (m *awsRestjson1_deserializeOpDeleteDiscoverer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -801,10 +788,8 @@ func (m *awsRestjson1_deserializeOpDeleteRegistry) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -822,7 +807,6 @@ func (m *awsRestjson1_deserializeOpDeleteRegistry) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -910,10 +894,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -931,7 +913,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1019,10 +1000,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1040,7 +1019,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1128,10 +1106,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchemaVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1149,7 +1125,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchemaVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1237,10 +1212,8 @@ func (m *awsRestjson1_deserializeOpDescribeCodeBinding
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1280,7 +1253,6 @@ func (m *awsRestjson1_deserializeOpDescribeCodeBinding
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1443,10 +1415,8 @@ func (m *awsRestjson1_deserializeOpDescribeDiscoverer)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1486,7 +1456,6 @@ func (m *awsRestjson1_deserializeOpDescribeDiscoverer)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1664,10 +1633,8 @@ func (m *awsRestjson1_deserializeOpDescribeRegistry) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1707,7 +1674,6 @@ func (m *awsRestjson1_deserializeOpDescribeRegistry) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1858,10 +1824,8 @@ func (m *awsRestjson1_deserializeOpDescribeSchema) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1901,7 +1865,6 @@ func (m *awsRestjson1_deserializeOpDescribeSchema) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2105,10 +2068,8 @@ func (m *awsRestjson1_deserializeOpExportSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2148,7 +2109,6 @@ func (m *awsRestjson1_deserializeOpExportSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2315,10 +2275,8 @@ func (m *awsRestjson1_deserializeOpGetCodeBindingSourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2335,7 +2293,6 @@ func (m *awsRestjson1_deserializeOpGetCodeBindingSourc
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2445,10 +2402,8 @@ func (m *awsRestjson1_deserializeOpGetDiscoveredSchema
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2488,7 +2443,6 @@ func (m *awsRestjson1_deserializeOpGetDiscoveredSchema
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2613,10 +2567,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2656,7 +2608,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2793,10 +2744,8 @@ func (m *awsRestjson1_deserializeOpListDiscoverers) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2836,7 +2785,6 @@ func (m *awsRestjson1_deserializeOpListDiscoverers) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2966,10 +2914,8 @@ func (m *awsRestjson1_deserializeOpListRegistries) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3009,7 +2955,6 @@ func (m *awsRestjson1_deserializeOpListRegistries) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3139,10 +3084,8 @@ func (m *awsRestjson1_deserializeOpListSchemas) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3182,7 +3125,6 @@ func (m *awsRestjson1_deserializeOpListSchemas) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3312,10 +3254,8 @@ func (m *awsRestjson1_deserializeOpListSchemaVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3355,7 +3295,6 @@ func (m *awsRestjson1_deserializeOpListSchemaVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3488,10 +3427,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3531,7 +3468,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3649,10 +3585,8 @@ func (m *awsRestjson1_deserializeOpPutCodeBinding) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3692,7 +3626,6 @@ func (m *awsRestjson1_deserializeOpPutCodeBinding) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3858,10 +3791,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3901,7 +3832,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4041,10 +3971,8 @@ func (m *awsRestjson1_deserializeOpSearchSchemas) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4084,7 +4012,6 @@ func (m *awsRestjson1_deserializeOpSearchSchemas) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4214,10 +4141,8 @@ func (m *awsRestjson1_deserializeOpStartDiscoverer) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4257,7 +4182,6 @@ func (m *awsRestjson1_deserializeOpStartDiscoverer) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4394,10 +4318,8 @@ func (m *awsRestjson1_deserializeOpStopDiscoverer) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4437,7 +4359,6 @@ func (m *awsRestjson1_deserializeOpStopDiscoverer) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4574,10 +4495,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4595,7 +4514,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4677,10 +4595,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4698,7 +4614,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4780,10 +4695,8 @@ func (m *awsRestjson1_deserializeOpUpdateDiscoverer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4823,7 +4736,6 @@ func (m *awsRestjson1_deserializeOpUpdateDiscoverer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5001,10 +4913,8 @@ func (m *awsRestjson1_deserializeOpUpdateRegistry) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5044,7 +4954,6 @@ func (m *awsRestjson1_deserializeOpUpdateRegistry) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5195,10 +5104,8 @@ func (m *awsRestjson1_deserializeOpUpdateSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5238,7 +5145,6 @@ func (m *awsRestjson1_deserializeOpUpdateSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
