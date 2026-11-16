--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/deserializers.go
@@ -17,7 +17,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -49,10 +48,8 @@ func (m *awsRestxml_deserializeOpAllQueryStringTypes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -70,7 +67,6 @@ func (m *awsRestxml_deserializeOpAllQueryStringTypes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -124,10 +120,8 @@ func (m *awsRestxml_deserializeOpBodyWithXmlName) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -167,7 +161,6 @@ func (m *awsRestxml_deserializeOpBodyWithXmlName) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -263,10 +256,8 @@ func (m *awsRestxml_deserializeOpConstantAndVariableQu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -284,7 +275,6 @@ func (m *awsRestxml_deserializeOpConstantAndVariableQu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -338,10 +328,8 @@ func (m *awsRestxml_deserializeOpConstantQueryString) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -359,7 +347,6 @@ func (m *awsRestxml_deserializeOpConstantQueryString) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -413,10 +400,8 @@ func (m *awsRestxml_deserializeOpContentTypeParameters
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -428,7 +413,6 @@ func (m *awsRestxml_deserializeOpContentTypeParameters
 	output := &ContentTypeParametersOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -482,10 +466,8 @@ func (m *awsRestxml_deserializeOpDatetimeOffsets) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -525,7 +507,6 @@ func (m *awsRestxml_deserializeOpDatetimeOffsets) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -632,10 +613,8 @@ func (m *awsRestxml_deserializeOpEmptyInputAndEmptyOut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -647,7 +626,6 @@ func (m *awsRestxml_deserializeOpEmptyInputAndEmptyOut
 	output := &EmptyInputAndEmptyOutputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -701,10 +679,8 @@ func (m *awsRestxml_deserializeOpEndpointOperation) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -722,7 +698,6 @@ func (m *awsRestxml_deserializeOpEndpointOperation) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -776,10 +751,8 @@ func (m *awsRestxml_deserializeOpEndpointWithHostLabel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -797,7 +770,6 @@ func (m *awsRestxml_deserializeOpEndpointWithHostLabel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -851,10 +823,8 @@ func (m *awsRestxml_deserializeOpEndpointWithHostLabel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -872,7 +842,6 @@ func (m *awsRestxml_deserializeOpEndpointWithHostLabel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -926,10 +895,8 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMap) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -969,7 +936,6 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMap) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1065,10 +1031,8 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMapWithXm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1108,7 +1072,6 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMapWithXm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1204,10 +1167,8 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMapWithXm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1247,7 +1208,6 @@ func (m *awsRestxml_deserializeOpFlattenedXmlMapWithXm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1343,10 +1303,8 @@ func (m *awsRestxml_deserializeOpFractionalSeconds) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1386,7 +1344,6 @@ func (m *awsRestxml_deserializeOpFractionalSeconds) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1493,10 +1450,8 @@ func (m *awsRestxml_deserializeOpGreetingWithErrors) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1513,7 +1468,6 @@ func (m *awsRestxml_deserializeOpGreetingWithErrors) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1586,10 +1540,8 @@ func (m *awsRestxml_deserializeOpHttpEmptyPrefixHeader
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1606,7 +1558,6 @@ func (m *awsRestxml_deserializeOpHttpEmptyPrefixHeader
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1683,10 +1634,8 @@ func (m *awsRestxml_deserializeOpHttpEnumPayload) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1703,7 +1652,6 @@ func (m *awsRestxml_deserializeOpHttpEnumPayload) Hand
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1778,10 +1726,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadTraits) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1803,7 +1749,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadTraits) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1890,10 +1835,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadTraitsWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1915,7 +1858,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadTraitsWith
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2002,10 +1944,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2045,7 +1985,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2141,10 +2080,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithStruct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2184,7 +2121,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithStruct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2280,10 +2216,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithUnion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2323,7 +2257,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithUnion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2419,10 +2352,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2462,7 +2393,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2558,10 +2488,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2601,7 +2529,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2697,10 +2624,8 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2740,7 +2665,6 @@ func (m *awsRestxml_deserializeOpHttpPayloadWithXmlNam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2836,10 +2760,8 @@ func (m *awsRestxml_deserializeOpHttpPrefixHeaders) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2856,7 +2778,6 @@ func (m *awsRestxml_deserializeOpHttpPrefixHeaders) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2933,10 +2854,8 @@ func (m *awsRestxml_deserializeOpHttpRequestWithFloatL
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2954,7 +2873,6 @@ func (m *awsRestxml_deserializeOpHttpRequestWithFloatL
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3008,10 +2926,8 @@ func (m *awsRestxml_deserializeOpHttpRequestWithGreedy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3029,7 +2945,6 @@ func (m *awsRestxml_deserializeOpHttpRequestWithGreedy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3083,10 +2998,8 @@ func (m *awsRestxml_deserializeOpHttpRequestWithLabels
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3104,7 +3017,6 @@ func (m *awsRestxml_deserializeOpHttpRequestWithLabels
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3158,10 +3070,8 @@ func (m *awsRestxml_deserializeOpHttpRequestWithLabels
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3179,7 +3089,6 @@ func (m *awsRestxml_deserializeOpHttpRequestWithLabels
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3233,10 +3142,8 @@ func (m *awsRestxml_deserializeOpHttpResponseCode) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3253,7 +3160,6 @@ func (m *awsRestxml_deserializeOpHttpResponseCode) Han
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3317,10 +3223,8 @@ func (m *awsRestxml_deserializeOpHttpStringPayload) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3337,7 +3241,6 @@ func (m *awsRestxml_deserializeOpHttpStringPayload) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3412,10 +3315,8 @@ func (m *awsRestxml_deserializeOpIgnoreQueryParamsInRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3455,7 +3356,6 @@ func (m *awsRestxml_deserializeOpIgnoreQueryParamsInRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3558,10 +3458,8 @@ func (m *awsRestxml_deserializeOpInputAndOutputWithHea
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3578,7 +3476,6 @@ func (m *awsRestxml_deserializeOpInputAndOutputWithHea
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3830,10 +3727,8 @@ func (m *awsRestxml_deserializeOpNestedXmlMaps) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3873,7 +3768,6 @@ func (m *awsRestxml_deserializeOpNestedXmlMaps) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3975,10 +3869,8 @@ func (m *awsRestxml_deserializeOpNestedXmlMapWithXmlNa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4018,7 +3910,6 @@ func (m *awsRestxml_deserializeOpNestedXmlMapWithXmlNa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4114,10 +4005,8 @@ func (m *awsRestxml_deserializeOpNoInputAndNoOutput) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4135,7 +4024,6 @@ func (m *awsRestxml_deserializeOpNoInputAndNoOutput) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4189,10 +4077,8 @@ func (m *awsRestxml_deserializeOpNoInputAndOutput) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4204,7 +4090,6 @@ func (m *awsRestxml_deserializeOpNoInputAndOutput) Han
 	output := &NoInputAndOutputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4258,10 +4143,8 @@ func (m *awsRestxml_deserializeOpNullAndEmptyHeadersCl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4278,7 +4161,6 @@ func (m *awsRestxml_deserializeOpNullAndEmptyHeadersCl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4366,10 +4248,8 @@ func (m *awsRestxml_deserializeOpNullAndEmptyHeadersSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4386,7 +4266,6 @@ func (m *awsRestxml_deserializeOpNullAndEmptyHeadersSe
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4474,10 +4353,8 @@ func (m *awsRestxml_deserializeOpOmitsNullSerializesEm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4495,7 +4372,6 @@ func (m *awsRestxml_deserializeOpOmitsNullSerializesEm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4549,10 +4425,8 @@ func (m *awsRestxml_deserializeOpPutWithContentEncodin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4570,7 +4444,6 @@ func (m *awsRestxml_deserializeOpPutWithContentEncodin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4624,10 +4497,8 @@ func (m *awsRestxml_deserializeOpQueryIdempotencyToken
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4645,7 +4516,6 @@ func (m *awsRestxml_deserializeOpQueryIdempotencyToken
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4699,10 +4569,8 @@ func (m *awsRestxml_deserializeOpQueryParamsAsStringLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4720,7 +4588,6 @@ func (m *awsRestxml_deserializeOpQueryParamsAsStringLi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4774,10 +4641,8 @@ func (m *awsRestxml_deserializeOpQueryPrecedence) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4795,7 +4660,6 @@ func (m *awsRestxml_deserializeOpQueryPrecedence) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4849,10 +4713,8 @@ func (m *awsRestxml_deserializeOpRecursiveShapes) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4892,7 +4754,6 @@ func (m *awsRestxml_deserializeOpRecursiveShapes) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4988,10 +4849,8 @@ func (m *awsRestxml_deserializeOpSimpleScalarPropertie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5036,7 +4895,6 @@ func (m *awsRestxml_deserializeOpSimpleScalarPropertie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5285,10 +5143,8 @@ func (m *awsRestxml_deserializeOpTimestampFormatHeader
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5305,7 +5161,6 @@ func (m *awsRestxml_deserializeOpTimestampFormatHeader
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5432,10 +5287,8 @@ func (m *awsRestxml_deserializeOpXmlAttributes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5475,7 +5328,6 @@ func (m *awsRestxml_deserializeOpXmlAttributes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5593,10 +5445,8 @@ func (m *awsRestxml_deserializeOpXmlAttributesOnPayloa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5636,7 +5486,6 @@ func (m *awsRestxml_deserializeOpXmlAttributesOnPayloa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5732,10 +5581,8 @@ func (m *awsRestxml_deserializeOpXmlBlobs) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5775,7 +5622,6 @@ func (m *awsRestxml_deserializeOpXmlBlobs) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5883,10 +5729,8 @@ func (m *awsRestxml_deserializeOpXmlEmptyBlobs) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5926,7 +5770,6 @@ func (m *awsRestxml_deserializeOpXmlEmptyBlobs) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6034,10 +5877,8 @@ func (m *awsRestxml_deserializeOpXmlEmptyLists) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6077,7 +5918,6 @@ func (m *awsRestxml_deserializeOpXmlEmptyLists) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6257,10 +6097,8 @@ func (m *awsRestxml_deserializeOpXmlEmptyMaps) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6300,7 +6138,6 @@ func (m *awsRestxml_deserializeOpXmlEmptyMaps) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6396,10 +6233,8 @@ func (m *awsRestxml_deserializeOpXmlEmptyStrings) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6439,7 +6274,6 @@ func (m *awsRestxml_deserializeOpXmlEmptyStrings) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6542,10 +6376,8 @@ func (m *awsRestxml_deserializeOpXmlEnums) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6585,7 +6417,6 @@ func (m *awsRestxml_deserializeOpXmlEnums) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6732,10 +6563,8 @@ func (m *awsRestxml_deserializeOpXmlIntEnums) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6775,7 +6604,6 @@ func (m *awsRestxml_deserializeOpXmlIntEnums) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6934,10 +6762,8 @@ func (m *awsRestxml_deserializeOpXmlLists) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6977,7 +6803,6 @@ func (m *awsRestxml_deserializeOpXmlLists) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7157,10 +6982,8 @@ func (m *awsRestxml_deserializeOpXmlMaps) HandleDeseri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7200,7 +7023,6 @@ func (m *awsRestxml_deserializeOpXmlMaps) HandleDeseri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7296,10 +7118,8 @@ func (m *awsRestxml_deserializeOpXmlMapsXmlName) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7339,7 +7159,6 @@ func (m *awsRestxml_deserializeOpXmlMapsXmlName) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7435,10 +7254,8 @@ func (m *awsRestxml_deserializeOpXmlMapWithXmlNamespac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7478,7 +7295,6 @@ func (m *awsRestxml_deserializeOpXmlMapWithXmlNamespac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7574,10 +7390,8 @@ func (m *awsRestxml_deserializeOpXmlNamespaces) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7617,7 +7431,6 @@ func (m *awsRestxml_deserializeOpXmlNamespaces) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7713,10 +7526,8 @@ func (m *awsRestxml_deserializeOpXmlTimestamps) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7756,7 +7567,6 @@ func (m *awsRestxml_deserializeOpXmlTimestamps) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7965,10 +7775,8 @@ func (m *awsRestxml_deserializeOpXmlUnions) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8008,7 +7816,6 @@ func (m *awsRestxml_deserializeOpXmlUnions) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
