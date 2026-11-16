--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/deserializers.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -50,10 +49,8 @@ func (m *awsRestjson1_deserializeOpCreateEvent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpCreateEvent) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -220,10 +216,8 @@ func (m *awsRestjson1_deserializeOpDeleteEvent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -263,7 +257,6 @@ func (m *awsRestjson1_deserializeOpDeleteEvent) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -394,10 +387,8 @@ func (m *awsRestjson1_deserializeOpDeleteMemoryRecord)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -437,7 +428,6 @@ func (m *awsRestjson1_deserializeOpDeleteMemoryRecord)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -568,10 +558,8 @@ func (m *awsRestjson1_deserializeOpGetBrowserSession) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -611,7 +599,6 @@ func (m *awsRestjson1_deserializeOpGetBrowserSession) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -821,10 +808,8 @@ func (m *awsRestjson1_deserializeOpGetCodeInterpreterS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -864,7 +849,6 @@ func (m *awsRestjson1_deserializeOpGetCodeInterpreterS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1042,10 +1026,8 @@ func (m *awsRestjson1_deserializeOpGetEvent) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1085,7 +1067,6 @@ func (m *awsRestjson1_deserializeOpGetEvent) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1212,10 +1193,8 @@ func (m *awsRestjson1_deserializeOpGetMemoryRecord) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1255,7 +1234,6 @@ func (m *awsRestjson1_deserializeOpGetMemoryRecord) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1382,10 +1360,8 @@ func (m *awsRestjson1_deserializeOpGetResourceApiKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1425,7 +1401,6 @@ func (m *awsRestjson1_deserializeOpGetResourceApiKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1553,10 +1528,8 @@ func (m *awsRestjson1_deserializeOpGetResourceOauth2To
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1596,7 +1569,6 @@ func (m *awsRestjson1_deserializeOpGetResourceOauth2To
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1733,10 +1705,8 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1776,7 +1746,6 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1904,10 +1873,8 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1947,7 +1914,6 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2075,10 +2041,8 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2118,7 +2082,6 @@ func (m *awsRestjson1_deserializeOpGetWorkloadAccessTo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2246,10 +2209,8 @@ func (m *awsRestjson1_deserializeOpInvokeAgentRuntime)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2271,7 +2232,6 @@ func (m *awsRestjson1_deserializeOpInvokeAgentRuntime)
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2420,10 +2380,8 @@ func (m *awsRestjson1_deserializeOpInvokeCodeInterpret
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2440,7 +2398,6 @@ func (m *awsRestjson1_deserializeOpInvokeCodeInterpret
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2544,10 +2501,8 @@ func (m *awsRestjson1_deserializeOpListActors) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2587,7 +2542,6 @@ func (m *awsRestjson1_deserializeOpListActors) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2723,10 +2677,8 @@ func (m *awsRestjson1_deserializeOpListBrowserSessions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2766,7 +2718,6 @@ func (m *awsRestjson1_deserializeOpListBrowserSessions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2896,10 +2847,8 @@ func (m *awsRestjson1_deserializeOpListCodeInterpreter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2939,7 +2888,6 @@ func (m *awsRestjson1_deserializeOpListCodeInterpreter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3069,10 +3017,8 @@ func (m *awsRestjson1_deserializeOpListEvents) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3112,7 +3058,6 @@ func (m *awsRestjson1_deserializeOpListEvents) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3248,10 +3193,8 @@ func (m *awsRestjson1_deserializeOpListMemoryRecords) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3291,7 +3234,6 @@ func (m *awsRestjson1_deserializeOpListMemoryRecords) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3427,10 +3369,8 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3470,7 +3410,6 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3606,10 +3545,8 @@ func (m *awsRestjson1_deserializeOpRetrieveMemoryRecor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3649,7 +3586,6 @@ func (m *awsRestjson1_deserializeOpRetrieveMemoryRecor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3785,10 +3721,8 @@ func (m *awsRestjson1_deserializeOpStartBrowserSession
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3828,7 +3762,6 @@ func (m *awsRestjson1_deserializeOpStartBrowserSession
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3986,10 +3919,8 @@ func (m *awsRestjson1_deserializeOpStartCodeInterprete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4029,7 +3960,6 @@ func (m *awsRestjson1_deserializeOpStartCodeInterprete
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4182,10 +4112,8 @@ func (m *awsRestjson1_deserializeOpStopBrowserSession)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4225,7 +4153,6 @@ func (m *awsRestjson1_deserializeOpStopBrowserSession)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4378,10 +4305,8 @@ func (m *awsRestjson1_deserializeOpStopCodeInterpreter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4421,7 +4346,6 @@ func (m *awsRestjson1_deserializeOpStopCodeInterpreter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4574,10 +4498,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrowserStream
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4617,7 +4539,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrowserStream
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
