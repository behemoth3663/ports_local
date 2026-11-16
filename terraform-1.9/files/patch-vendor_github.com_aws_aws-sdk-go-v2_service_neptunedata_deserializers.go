--- vendor/github.com/aws/aws-sdk-go-v2/service/neptunedata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/neptunedata/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCancelGremlinQuery)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCancelGremlinQuery)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -241,10 +237,8 @@ func (m *awsRestjson1_deserializeOpCancelLoaderJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -284,7 +278,6 @@ func (m *awsRestjson1_deserializeOpCancelLoaderJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -433,10 +426,8 @@ func (m *awsRestjson1_deserializeOpCancelMLDataProcess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -476,7 +467,6 @@ func (m *awsRestjson1_deserializeOpCancelMLDataProcess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -619,10 +609,8 @@ func (m *awsRestjson1_deserializeOpCancelMLModelTraini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -662,7 +650,6 @@ func (m *awsRestjson1_deserializeOpCancelMLModelTraini
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -805,10 +792,8 @@ func (m *awsRestjson1_deserializeOpCancelMLModelTransf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -848,7 +833,6 @@ func (m *awsRestjson1_deserializeOpCancelMLModelTransf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -991,10 +975,8 @@ func (m *awsRestjson1_deserializeOpCancelOpenCypherQue
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1034,7 +1016,6 @@ func (m *awsRestjson1_deserializeOpCancelOpenCypherQue
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1198,10 +1179,8 @@ func (m *awsRestjson1_deserializeOpCreateMLEndpoint) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1241,7 +1220,6 @@ func (m *awsRestjson1_deserializeOpCreateMLEndpoint) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1406,10 +1384,8 @@ func (m *awsRestjson1_deserializeOpDeleteMLEndpoint) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1449,7 +1425,6 @@ func (m *awsRestjson1_deserializeOpDeleteMLEndpoint) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1592,10 +1567,8 @@ func (m *awsRestjson1_deserializeOpDeletePropertygraph
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1640,7 +1613,6 @@ func (m *awsRestjson1_deserializeOpDeletePropertygraph
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1803,10 +1775,8 @@ func (m *awsRestjson1_deserializeOpDeleteSparqlStatist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1851,7 +1821,6 @@ func (m *awsRestjson1_deserializeOpDeleteSparqlStatist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2014,10 +1983,8 @@ func (m *awsRestjson1_deserializeOpExecuteFastReset) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2057,7 +2024,6 @@ func (m *awsRestjson1_deserializeOpExecuteFastReset) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2211,10 +2177,8 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinExpla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2231,7 +2195,6 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinExpla
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2383,10 +2346,8 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2403,7 +2364,6 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinProfi
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2555,10 +2515,8 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinQuery
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2598,7 +2556,6 @@ func (m *awsRestjson1_deserializeOpExecuteGremlinQuery
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2783,10 +2740,8 @@ func (m *awsRestjson1_deserializeOpExecuteOpenCypherEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2803,7 +2758,6 @@ func (m *awsRestjson1_deserializeOpExecuteOpenCypherEx
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2958,10 +2912,8 @@ func (m *awsRestjson1_deserializeOpExecuteOpenCypherQu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3001,7 +2953,6 @@ func (m *awsRestjson1_deserializeOpExecuteOpenCypherQu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3170,10 +3121,8 @@ func (m *awsRestjson1_deserializeOpGetEngineStatus) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3213,7 +3162,6 @@ func (m *awsRestjson1_deserializeOpGetEngineStatus) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3435,10 +3383,8 @@ func (m *awsRestjson1_deserializeOpGetGremlinQueryStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3478,7 +3424,6 @@ func (m *awsRestjson1_deserializeOpGetGremlinQueryStat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3650,10 +3595,8 @@ func (m *awsRestjson1_deserializeOpGetLoaderJobStatus)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3693,7 +3636,6 @@ func (m *awsRestjson1_deserializeOpGetLoaderJobStatus)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3847,10 +3789,8 @@ func (m *awsRestjson1_deserializeOpGetMLDataProcessing
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3890,7 +3830,6 @@ func (m *awsRestjson1_deserializeOpGetMLDataProcessing
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4047,10 +3986,8 @@ func (m *awsRestjson1_deserializeOpGetMLEndpoint) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4090,7 +4027,6 @@ func (m *awsRestjson1_deserializeOpGetMLEndpoint) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4252,10 +4188,8 @@ func (m *awsRestjson1_deserializeOpGetMLModelTrainingJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4295,7 +4229,6 @@ func (m *awsRestjson1_deserializeOpGetMLModelTrainingJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4467,10 +4400,8 @@ func (m *awsRestjson1_deserializeOpGetMLModelTransform
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4510,7 +4441,6 @@ func (m *awsRestjson1_deserializeOpGetMLModelTransform
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4677,10 +4607,8 @@ func (m *awsRestjson1_deserializeOpGetOpenCypherQueryS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4720,7 +4648,6 @@ func (m *awsRestjson1_deserializeOpGetOpenCypherQueryS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4895,10 +4822,8 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4938,7 +4863,6 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphSta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5092,10 +5016,8 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphStr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5135,7 +5057,6 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphStr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5317,10 +5238,8 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphSum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5365,7 +5284,6 @@ func (m *awsRestjson1_deserializeOpGetPropertygraphSum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5519,10 +5437,8 @@ func (m *awsRestjson1_deserializeOpGetRDFGraphSummary)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5567,7 +5483,6 @@ func (m *awsRestjson1_deserializeOpGetRDFGraphSummary)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5721,10 +5636,8 @@ func (m *awsRestjson1_deserializeOpGetSparqlStatistics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5764,7 +5677,6 @@ func (m *awsRestjson1_deserializeOpGetSparqlStatistics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5918,10 +5830,8 @@ func (m *awsRestjson1_deserializeOpGetSparqlStream) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5961,7 +5871,6 @@ func (m *awsRestjson1_deserializeOpGetSparqlStream) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6143,10 +6052,8 @@ func (m *awsRestjson1_deserializeOpListGremlinQueries)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6186,7 +6093,6 @@ func (m *awsRestjson1_deserializeOpListGremlinQueries)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6366,10 +6272,8 @@ func (m *awsRestjson1_deserializeOpListLoaderJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6409,7 +6313,6 @@ func (m *awsRestjson1_deserializeOpListLoaderJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6560,10 +6463,8 @@ func (m *awsRestjson1_deserializeOpListMLDataProcessin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6603,7 +6504,6 @@ func (m *awsRestjson1_deserializeOpListMLDataProcessin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6742,10 +6642,8 @@ func (m *awsRestjson1_deserializeOpListMLEndpoints) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6785,7 +6683,6 @@ func (m *awsRestjson1_deserializeOpListMLEndpoints) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6924,10 +6821,8 @@ func (m *awsRestjson1_deserializeOpListMLModelTraining
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6967,7 +6862,6 @@ func (m *awsRestjson1_deserializeOpListMLModelTraining
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7106,10 +7000,8 @@ func (m *awsRestjson1_deserializeOpListMLModelTransfor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7149,7 +7041,6 @@ func (m *awsRestjson1_deserializeOpListMLModelTransfor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7288,10 +7179,8 @@ func (m *awsRestjson1_deserializeOpListOpenCypherQueri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7331,7 +7220,6 @@ func (m *awsRestjson1_deserializeOpListOpenCypherQueri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7514,10 +7402,8 @@ func (m *awsRestjson1_deserializeOpManagePropertygraph
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7557,7 +7443,6 @@ func (m *awsRestjson1_deserializeOpManagePropertygraph
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7711,10 +7596,8 @@ func (m *awsRestjson1_deserializeOpManageSparqlStatist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7754,7 +7637,6 @@ func (m *awsRestjson1_deserializeOpManageSparqlStatist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7908,10 +7790,8 @@ func (m *awsRestjson1_deserializeOpStartLoaderJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7951,7 +7831,6 @@ func (m *awsRestjson1_deserializeOpStartLoaderJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8108,10 +7987,8 @@ func (m *awsRestjson1_deserializeOpStartMLDataProcessi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8151,7 +8028,6 @@ func (m *awsRestjson1_deserializeOpStartMLDataProcessi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8316,10 +8192,8 @@ func (m *awsRestjson1_deserializeOpStartMLModelTrainin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8359,7 +8233,6 @@ func (m *awsRestjson1_deserializeOpStartMLModelTrainin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8524,10 +8397,8 @@ func (m *awsRestjson1_deserializeOpStartMLModelTransfo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8567,7 +8438,6 @@ func (m *awsRestjson1_deserializeOpStartMLModelTransfo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
