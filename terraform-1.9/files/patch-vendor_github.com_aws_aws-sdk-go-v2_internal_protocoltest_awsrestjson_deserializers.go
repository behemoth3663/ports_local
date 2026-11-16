--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/deserializers.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -50,10 +49,8 @@ func (m *awsRestjson1_deserializeOpAllQueryStringTypes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -71,7 +68,6 @@ func (m *awsRestjson1_deserializeOpAllQueryStringTypes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -141,10 +137,8 @@ func (m *awsRestjson1_deserializeOpConstantAndVariable
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -162,7 +156,6 @@ func (m *awsRestjson1_deserializeOpConstantAndVariable
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -232,10 +225,8 @@ func (m *awsRestjson1_deserializeOpConstantQueryString
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -253,7 +244,6 @@ func (m *awsRestjson1_deserializeOpConstantQueryString
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -323,10 +313,8 @@ func (m *awsRestjson1_deserializeOpContentTypeParamete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -338,7 +326,6 @@ func (m *awsRestjson1_deserializeOpContentTypeParamete
 	output := &ContentTypeParametersOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -408,10 +395,8 @@ func (m *awsRestjson1_deserializeOpDatetimeOffsets) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -451,7 +436,6 @@ func (m *awsRestjson1_deserializeOpDatetimeOffsets) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -565,10 +549,8 @@ func (m *awsRestjson1_deserializeOpDocumentType) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -608,7 +590,6 @@ func (m *awsRestjson1_deserializeOpDocumentType) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -723,10 +704,8 @@ func (m *awsRestjson1_deserializeOpDocumentTypeAsMapVa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -766,7 +745,6 @@ func (m *awsRestjson1_deserializeOpDocumentTypeAsMapVa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -872,10 +850,8 @@ func (m *awsRestjson1_deserializeOpDocumentTypeAsPaylo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -915,7 +891,6 @@ func (m *awsRestjson1_deserializeOpDocumentTypeAsPaylo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1021,10 +996,8 @@ func (m *awsRestjson1_deserializeOpEmptyInputAndEmptyO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1036,7 +1009,6 @@ func (m *awsRestjson1_deserializeOpEmptyInputAndEmptyO
 	output := &EmptyInputAndEmptyOutputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1106,10 +1078,8 @@ func (m *awsRestjson1_deserializeOpEndpointOperation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1127,7 +1097,6 @@ func (m *awsRestjson1_deserializeOpEndpointOperation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1197,10 +1166,8 @@ func (m *awsRestjson1_deserializeOpEndpointWithHostLab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1218,7 +1185,6 @@ func (m *awsRestjson1_deserializeOpEndpointWithHostLab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1288,10 +1254,8 @@ func (m *awsRestjson1_deserializeOpFractionalSeconds) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1331,7 +1295,6 @@ func (m *awsRestjson1_deserializeOpFractionalSeconds) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1445,10 +1408,8 @@ func (m *awsRestjson1_deserializeOpGreetingWithErrors)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1465,7 +1426,6 @@ func (m *awsRestjson1_deserializeOpGreetingWithErrors)
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1557,10 +1517,8 @@ func (m *awsRestjson1_deserializeOpHostWithPathOperati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1578,7 +1536,6 @@ func (m *awsRestjson1_deserializeOpHostWithPathOperati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1648,10 +1605,8 @@ func (m *awsRestjson1_deserializeOpHttpChecksumRequire
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1691,7 +1646,6 @@ func (m *awsRestjson1_deserializeOpHttpChecksumRequire
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1801,10 +1755,8 @@ func (m *awsRestjson1_deserializeOpHttpEmptyPrefixHead
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1821,7 +1773,6 @@ func (m *awsRestjson1_deserializeOpHttpEmptyPrefixHead
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1914,10 +1865,8 @@ func (m *awsRestjson1_deserializeOpHttpEnumPayload) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1934,7 +1883,6 @@ func (m *awsRestjson1_deserializeOpHttpEnumPayload) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2026,10 +1974,8 @@ func (m *awsRestjson1_deserializeOpHttpPayloadTraits) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2051,7 +1997,6 @@ func (m *awsRestjson1_deserializeOpHttpPayloadTraits) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2155,10 +2100,8 @@ func (m *awsRestjson1_deserializeOpHttpPayloadTraitsWi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2180,7 +2123,6 @@ func (m *awsRestjson1_deserializeOpHttpPayloadTraitsWi
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2284,10 +2226,8 @@ func (m *awsRestjson1_deserializeOpHttpPayloadWithStru
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2327,7 +2267,6 @@ func (m *awsRestjson1_deserializeOpHttpPayloadWithStru
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2433,10 +2372,8 @@ func (m *awsRestjson1_deserializeOpHttpPayloadWithUnio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2476,7 +2413,6 @@ func (m *awsRestjson1_deserializeOpHttpPayloadWithUnio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2582,10 +2518,8 @@ func (m *awsRestjson1_deserializeOpHttpPrefixHeaders) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2602,7 +2536,6 @@ func (m *awsRestjson1_deserializeOpHttpPrefixHeaders) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2695,10 +2628,8 @@ func (m *awsRestjson1_deserializeOpHttpPrefixHeadersIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2715,7 +2646,6 @@ func (m *awsRestjson1_deserializeOpHttpPrefixHeadersIn
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2803,10 +2733,8 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithFloa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2824,7 +2752,6 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithFloa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2894,10 +2821,8 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithGree
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2915,7 +2840,6 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithGree
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2985,10 +2909,8 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithLabe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3006,7 +2928,6 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithLabe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3076,10 +2997,8 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithLabe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3097,7 +3016,6 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithLabe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3167,10 +3085,8 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithRege
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3188,7 +3104,6 @@ func (m *awsRestjson1_deserializeOpHttpRequestWithRege
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3258,10 +3173,8 @@ func (m *awsRestjson1_deserializeOpHttpResponseCode) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3278,7 +3191,6 @@ func (m *awsRestjson1_deserializeOpHttpResponseCode) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3358,10 +3270,8 @@ func (m *awsRestjson1_deserializeOpHttpStringPayload) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3378,7 +3288,6 @@ func (m *awsRestjson1_deserializeOpHttpStringPayload) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3470,10 +3379,8 @@ func (m *awsRestjson1_deserializeOpIgnoreQueryParamsIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3513,7 +3420,6 @@ func (m *awsRestjson1_deserializeOpIgnoreQueryParamsIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3623,10 +3529,8 @@ func (m *awsRestjson1_deserializeOpInputAndOutputWithH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3643,7 +3547,6 @@ func (m *awsRestjson1_deserializeOpInputAndOutputWithH
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3940,10 +3843,8 @@ func (m *awsRestjson1_deserializeOpJsonBlobs) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3983,7 +3884,6 @@ func (m *awsRestjson1_deserializeOpJsonBlobs) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4097,10 +3997,8 @@ func (m *awsRestjson1_deserializeOpJsonEnums) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4140,7 +4038,6 @@ func (m *awsRestjson1_deserializeOpJsonEnums) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4283,10 +4180,8 @@ func (m *awsRestjson1_deserializeOpJsonIntEnums) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4326,7 +4221,6 @@ func (m *awsRestjson1_deserializeOpJsonIntEnums) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4481,10 +4375,8 @@ func (m *awsRestjson1_deserializeOpJsonLists) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4524,7 +4416,6 @@ func (m *awsRestjson1_deserializeOpJsonLists) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4670,10 +4561,8 @@ func (m *awsRestjson1_deserializeOpJsonMaps) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4713,7 +4602,6 @@ func (m *awsRestjson1_deserializeOpJsonMaps) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4839,10 +4727,8 @@ func (m *awsRestjson1_deserializeOpJsonTimestamps) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4882,7 +4768,6 @@ func (m *awsRestjson1_deserializeOpJsonTimestamps) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5083,10 +4968,8 @@ func (m *awsRestjson1_deserializeOpJsonUnions) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5126,7 +5009,6 @@ func (m *awsRestjson1_deserializeOpJsonUnions) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5232,10 +5114,8 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5275,7 +5155,6 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5385,10 +5264,8 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5405,7 +5282,6 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5497,10 +5373,8 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5517,7 +5391,6 @@ func (m *awsRestjson1_deserializeOpMalformedAcceptWith
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5609,10 +5482,8 @@ func (m *awsRestjson1_deserializeOpMalformedBlob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5630,7 +5501,6 @@ func (m *awsRestjson1_deserializeOpMalformedBlob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5700,10 +5570,8 @@ func (m *awsRestjson1_deserializeOpMalformedBoolean) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5721,7 +5589,6 @@ func (m *awsRestjson1_deserializeOpMalformedBoolean) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5791,10 +5658,8 @@ func (m *awsRestjson1_deserializeOpMalformedByte) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5812,7 +5677,6 @@ func (m *awsRestjson1_deserializeOpMalformedByte) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5882,10 +5746,8 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5903,7 +5765,6 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5973,10 +5834,8 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5994,7 +5853,6 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6064,10 +5922,8 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6085,7 +5941,6 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6155,10 +6010,8 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6176,7 +6029,6 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6246,10 +6098,8 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6267,7 +6117,6 @@ func (m *awsRestjson1_deserializeOpMalformedContentTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6337,10 +6186,8 @@ func (m *awsRestjson1_deserializeOpMalformedDouble) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6358,7 +6205,6 @@ func (m *awsRestjson1_deserializeOpMalformedDouble) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6428,10 +6274,8 @@ func (m *awsRestjson1_deserializeOpMalformedFloat) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6449,7 +6293,6 @@ func (m *awsRestjson1_deserializeOpMalformedFloat) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6519,10 +6362,8 @@ func (m *awsRestjson1_deserializeOpMalformedInteger) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6540,7 +6381,6 @@ func (m *awsRestjson1_deserializeOpMalformedInteger) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6610,10 +6450,8 @@ func (m *awsRestjson1_deserializeOpMalformedList) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6631,7 +6469,6 @@ func (m *awsRestjson1_deserializeOpMalformedList) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6701,10 +6538,8 @@ func (m *awsRestjson1_deserializeOpMalformedLong) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6722,7 +6557,6 @@ func (m *awsRestjson1_deserializeOpMalformedLong) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6792,10 +6626,8 @@ func (m *awsRestjson1_deserializeOpMalformedMap) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6813,7 +6645,6 @@ func (m *awsRestjson1_deserializeOpMalformedMap) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6883,10 +6714,8 @@ func (m *awsRestjson1_deserializeOpMalformedRequestBod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6904,7 +6733,6 @@ func (m *awsRestjson1_deserializeOpMalformedRequestBod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6974,10 +6802,8 @@ func (m *awsRestjson1_deserializeOpMalformedShort) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6995,7 +6821,6 @@ func (m *awsRestjson1_deserializeOpMalformedShort) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7065,10 +6890,8 @@ func (m *awsRestjson1_deserializeOpMalformedString) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7086,7 +6909,6 @@ func (m *awsRestjson1_deserializeOpMalformedString) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7156,10 +6978,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7177,7 +6997,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7247,10 +7066,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7268,7 +7085,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7338,10 +7154,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7359,7 +7173,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7429,10 +7242,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7450,7 +7261,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7520,10 +7330,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7541,7 +7349,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7611,10 +7418,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7632,7 +7437,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampH
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7702,10 +7506,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7723,7 +7525,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7793,10 +7594,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7814,7 +7613,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7884,10 +7682,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7905,7 +7701,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7975,10 +7770,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7996,7 +7789,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8066,10 +7858,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8087,7 +7877,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8157,10 +7946,8 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8178,7 +7965,6 @@ func (m *awsRestjson1_deserializeOpMalformedTimestampQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8248,10 +8034,8 @@ func (m *awsRestjson1_deserializeOpMalformedUnion) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8269,7 +8053,6 @@ func (m *awsRestjson1_deserializeOpMalformedUnion) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8339,10 +8122,8 @@ func (m *awsRestjson1_deserializeOpMediaTypeHeader) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8359,7 +8140,6 @@ func (m *awsRestjson1_deserializeOpMediaTypeHeader) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8446,10 +8226,8 @@ func (m *awsRestjson1_deserializeOpNoInputAndNoOutput)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8467,7 +8245,6 @@ func (m *awsRestjson1_deserializeOpNoInputAndNoOutput)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8537,10 +8314,8 @@ func (m *awsRestjson1_deserializeOpNoInputAndOutput) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8552,7 +8327,6 @@ func (m *awsRestjson1_deserializeOpNoInputAndOutput) H
 	output := &NoInputAndOutputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8622,10 +8396,8 @@ func (m *awsRestjson1_deserializeOpNullAndEmptyHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8642,7 +8414,6 @@ func (m *awsRestjson1_deserializeOpNullAndEmptyHeaders
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8746,10 +8517,8 @@ func (m *awsRestjson1_deserializeOpNullAndEmptyHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8766,7 +8535,6 @@ func (m *awsRestjson1_deserializeOpNullAndEmptyHeaders
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8870,10 +8638,8 @@ func (m *awsRestjson1_deserializeOpOmitsNullSerializes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8891,7 +8657,6 @@ func (m *awsRestjson1_deserializeOpOmitsNullSerializes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8961,10 +8726,8 @@ func (m *awsRestjson1_deserializeOpOmitsSerializingEmp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8982,7 +8745,6 @@ func (m *awsRestjson1_deserializeOpOmitsSerializingEmp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9052,10 +8814,8 @@ func (m *awsRestjson1_deserializeOpOperationWithDefaul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9095,7 +8855,6 @@ func (m *awsRestjson1_deserializeOpOperationWithDefaul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9571,10 +9330,8 @@ func (m *awsRestjson1_deserializeOpOperationWithNested
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9614,7 +9371,6 @@ func (m *awsRestjson1_deserializeOpOperationWithNested
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9730,10 +9486,8 @@ func (m *awsRestjson1_deserializeOpPostPlayerAction) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9773,7 +9527,6 @@ func (m *awsRestjson1_deserializeOpPostPlayerAction) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9879,10 +9632,8 @@ func (m *awsRestjson1_deserializeOpPostUnionWithJsonNa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9922,7 +9673,6 @@ func (m *awsRestjson1_deserializeOpPostUnionWithJsonNa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10028,10 +9778,8 @@ func (m *awsRestjson1_deserializeOpPutWithContentEncod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10049,7 +9797,6 @@ func (m *awsRestjson1_deserializeOpPutWithContentEncod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10119,10 +9866,8 @@ func (m *awsRestjson1_deserializeOpQueryIdempotencyTok
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10140,7 +9885,6 @@ func (m *awsRestjson1_deserializeOpQueryIdempotencyTok
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10210,10 +9954,8 @@ func (m *awsRestjson1_deserializeOpQueryParamsAsString
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10231,7 +9973,6 @@ func (m *awsRestjson1_deserializeOpQueryParamsAsString
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10301,10 +10042,8 @@ func (m *awsRestjson1_deserializeOpQueryPrecedence) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10322,7 +10061,6 @@ func (m *awsRestjson1_deserializeOpQueryPrecedence) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10392,10 +10130,8 @@ func (m *awsRestjson1_deserializeOpRecursiveShapes) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10435,7 +10171,6 @@ func (m *awsRestjson1_deserializeOpRecursiveShapes) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10541,10 +10276,8 @@ func (m *awsRestjson1_deserializeOpResponseCodeHttpFal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10556,7 +10289,6 @@ func (m *awsRestjson1_deserializeOpResponseCodeHttpFal
 	output := &ResponseCodeHttpFallbackOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10626,10 +10358,8 @@ func (m *awsRestjson1_deserializeOpResponseCodeRequire
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10646,7 +10376,6 @@ func (m *awsRestjson1_deserializeOpResponseCodeRequire
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10726,10 +10455,8 @@ func (m *awsRestjson1_deserializeOpSimpleScalarPropert
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10774,7 +10501,6 @@ func (m *awsRestjson1_deserializeOpSimpleScalarPropert
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11034,10 +10760,8 @@ func (m *awsRestjson1_deserializeOpSparseJsonLists) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11077,7 +10801,6 @@ func (m *awsRestjson1_deserializeOpSparseJsonLists) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11188,10 +10911,8 @@ func (m *awsRestjson1_deserializeOpSparseJsonMaps) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11231,7 +10952,6 @@ func (m *awsRestjson1_deserializeOpSparseJsonMaps) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11357,10 +11077,8 @@ func (m *awsRestjson1_deserializeOpStreamingTraits) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11382,7 +11100,6 @@ func (m *awsRestjson1_deserializeOpStreamingTraits) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11473,10 +11190,8 @@ func (m *awsRestjson1_deserializeOpStreamingTraitsRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11494,7 +11209,6 @@ func (m *awsRestjson1_deserializeOpStreamingTraitsRequ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11564,10 +11278,8 @@ func (m *awsRestjson1_deserializeOpStreamingTraitsWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11589,7 +11301,6 @@ func (m *awsRestjson1_deserializeOpStreamingTraitsWith
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11680,10 +11391,8 @@ func (m *awsRestjson1_deserializeOpTestBodyStructure) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11728,7 +11437,6 @@ func (m *awsRestjson1_deserializeOpTestBodyStructure) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11846,10 +11554,8 @@ func (m *awsRestjson1_deserializeOpTestGetNoInputNoPay
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11866,7 +11572,6 @@ func (m *awsRestjson1_deserializeOpTestGetNoInputNoPay
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11949,10 +11654,8 @@ func (m *awsRestjson1_deserializeOpTestGetNoPayload) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11969,7 +11672,6 @@ func (m *awsRestjson1_deserializeOpTestGetNoPayload) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12052,10 +11754,8 @@ func (m *awsRestjson1_deserializeOpTestPayloadBlob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12077,7 +11777,6 @@ func (m *awsRestjson1_deserializeOpTestPayloadBlob) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12181,10 +11880,8 @@ func (m *awsRestjson1_deserializeOpTestPayloadStructur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12229,7 +11926,6 @@ func (m *awsRestjson1_deserializeOpTestPayloadStructur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12347,10 +12043,8 @@ func (m *awsRestjson1_deserializeOpTestPostNoInputNoPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12367,7 +12061,6 @@ func (m *awsRestjson1_deserializeOpTestPostNoInputNoPa
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12450,10 +12143,8 @@ func (m *awsRestjson1_deserializeOpTestPostNoPayload) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12470,7 +12161,6 @@ func (m *awsRestjson1_deserializeOpTestPostNoPayload) 
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12553,10 +12243,8 @@ func (m *awsRestjson1_deserializeOpTimestampFormatHead
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12573,7 +12261,6 @@ func (m *awsRestjson1_deserializeOpTimestampFormatHead
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12716,10 +12403,8 @@ func (m *awsRestjson1_deserializeOpUnitInputAndOutput)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12737,7 +12422,6 @@ func (m *awsRestjson1_deserializeOpUnitInputAndOutput)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
