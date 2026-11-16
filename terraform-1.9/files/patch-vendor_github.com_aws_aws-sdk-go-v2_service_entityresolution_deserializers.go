--- vendor/github.com/aws/aws-sdk-go-v2/service/entityresolution/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/entityresolution/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAddPolicyStatement)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAddPolicyStatement)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -235,10 +231,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteUniqueId
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -278,7 +272,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteUniqueId
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -412,10 +405,8 @@ func (m *awsRestjson1_deserializeOpCreateIdMappingWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -455,7 +446,6 @@ func (m *awsRestjson1_deserializeOpCreateIdMappingWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -625,10 +615,8 @@ func (m *awsRestjson1_deserializeOpCreateIdNamespace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -668,7 +656,6 @@ func (m *awsRestjson1_deserializeOpCreateIdNamespace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -879,10 +866,8 @@ func (m *awsRestjson1_deserializeOpCreateMatchingWorkf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -922,7 +907,6 @@ func (m *awsRestjson1_deserializeOpCreateMatchingWorkf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1097,10 +1081,8 @@ func (m *awsRestjson1_deserializeOpCreateSchemaMapping
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1140,7 +1122,6 @@ func (m *awsRestjson1_deserializeOpCreateSchemaMapping
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1291,10 +1272,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdMappingWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1334,7 +1313,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdMappingWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1459,10 +1437,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdNamespace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1502,7 +1478,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdNamespace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1624,10 +1599,8 @@ func (m *awsRestjson1_deserializeOpDeleteMatchingWorkf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1667,7 +1640,6 @@ func (m *awsRestjson1_deserializeOpDeleteMatchingWorkf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1792,10 +1764,8 @@ func (m *awsRestjson1_deserializeOpDeletePolicyStateme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1835,7 +1805,6 @@ func (m *awsRestjson1_deserializeOpDeletePolicyStateme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1981,10 +1950,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchemaMapping
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2024,7 +1991,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchemaMapping
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2149,10 +2115,8 @@ func (m *awsRestjson1_deserializeOpGenerateMatchId) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2192,7 +2156,6 @@ func (m *awsRestjson1_deserializeOpGenerateMatchId) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2318,10 +2281,8 @@ func (m *awsRestjson1_deserializeOpGetIdMappingJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2361,7 +2322,6 @@ func (m *awsRestjson1_deserializeOpGetIdMappingJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2542,10 +2502,8 @@ func (m *awsRestjson1_deserializeOpGetIdMappingWorkflo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2585,7 +2543,6 @@ func (m *awsRestjson1_deserializeOpGetIdMappingWorkflo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2789,10 +2746,8 @@ func (m *awsRestjson1_deserializeOpGetIdNamespace) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2832,7 +2787,6 @@ func (m *awsRestjson1_deserializeOpGetIdNamespace) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3040,10 +2994,8 @@ func (m *awsRestjson1_deserializeOpGetMatchId) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3083,7 +3035,6 @@ func (m *awsRestjson1_deserializeOpGetMatchId) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3217,10 +3168,8 @@ func (m *awsRestjson1_deserializeOpGetMatchingJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3260,7 +3209,6 @@ func (m *awsRestjson1_deserializeOpGetMatchingJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3441,10 +3389,8 @@ func (m *awsRestjson1_deserializeOpGetMatchingWorkflow
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3484,7 +3430,6 @@ func (m *awsRestjson1_deserializeOpGetMatchingWorkflow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3693,10 +3638,8 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3736,7 +3679,6 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3879,10 +3821,8 @@ func (m *awsRestjson1_deserializeOpGetProviderService)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3922,7 +3862,6 @@ func (m *awsRestjson1_deserializeOpGetProviderService)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4127,10 +4066,8 @@ func (m *awsRestjson1_deserializeOpGetSchemaMapping) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4170,7 +4107,6 @@ func (m *awsRestjson1_deserializeOpGetSchemaMapping) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4364,10 +4300,8 @@ func (m *awsRestjson1_deserializeOpListIdMappingJobs) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4407,7 +4341,6 @@ func (m *awsRestjson1_deserializeOpListIdMappingJobs) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4537,10 +4470,8 @@ func (m *awsRestjson1_deserializeOpListIdMappingWorkfl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4580,7 +4511,6 @@ func (m *awsRestjson1_deserializeOpListIdMappingWorkfl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4707,10 +4637,8 @@ func (m *awsRestjson1_deserializeOpListIdNamespaces) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4750,7 +4678,6 @@ func (m *awsRestjson1_deserializeOpListIdNamespaces) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4877,10 +4804,8 @@ func (m *awsRestjson1_deserializeOpListMatchingJobs) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4920,7 +4845,6 @@ func (m *awsRestjson1_deserializeOpListMatchingJobs) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5050,10 +4974,8 @@ func (m *awsRestjson1_deserializeOpListMatchingWorkflo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5093,7 +5015,6 @@ func (m *awsRestjson1_deserializeOpListMatchingWorkflo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5220,10 +5141,8 @@ func (m *awsRestjson1_deserializeOpListProviderService
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5263,7 +5182,6 @@ func (m *awsRestjson1_deserializeOpListProviderService
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5390,10 +5308,8 @@ func (m *awsRestjson1_deserializeOpListSchemaMappings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5433,7 +5349,6 @@ func (m *awsRestjson1_deserializeOpListSchemaMappings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5560,10 +5475,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5603,7 +5516,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5718,10 +5630,8 @@ func (m *awsRestjson1_deserializeOpPutPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5761,7 +5671,6 @@ func (m *awsRestjson1_deserializeOpPutPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5907,10 +5816,8 @@ func (m *awsRestjson1_deserializeOpStartIdMappingJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5950,7 +5857,6 @@ func (m *awsRestjson1_deserializeOpStartIdMappingJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6086,10 +5992,8 @@ func (m *awsRestjson1_deserializeOpStartMatchingJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6129,7 +6033,6 @@ func (m *awsRestjson1_deserializeOpStartMatchingJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6260,10 +6163,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6275,7 +6176,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6354,10 +6254,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6369,7 +6267,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6445,10 +6342,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdMappingWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6488,7 +6383,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdMappingWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6655,10 +6549,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdNamespace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6698,7 +6590,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdNamespace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6901,10 +6792,8 @@ func (m *awsRestjson1_deserializeOpUpdateMatchingWorkf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6944,7 +6833,6 @@ func (m *awsRestjson1_deserializeOpUpdateMatchingWorkf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7107,10 +6995,8 @@ func (m *awsRestjson1_deserializeOpUpdateSchemaMapping
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7150,7 +7036,6 @@ func (m *awsRestjson1_deserializeOpUpdateSchemaMapping
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
