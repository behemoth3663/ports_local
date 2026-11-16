--- vendor/github.com/aws/aws-sdk-go-v2/service/kafkaconnect/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/kafkaconnect/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -239,10 +235,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomPlugin)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -282,7 +276,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomPlugin)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -447,10 +440,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkerConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -490,7 +481,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkerConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -660,10 +650,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -703,7 +691,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -843,10 +830,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomPlugin)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -886,7 +871,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomPlugin)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1026,10 +1010,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkerConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1069,7 +1051,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkerConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1209,10 +1190,8 @@ func (m *awsRestjson1_deserializeOpDescribeConnector) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1252,7 +1231,6 @@ func (m *awsRestjson1_deserializeOpDescribeConnector) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1495,10 +1473,8 @@ func (m *awsRestjson1_deserializeOpDescribeConnectorOp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1538,7 +1514,6 @@ func (m *awsRestjson1_deserializeOpDescribeConnectorOp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1752,10 +1727,8 @@ func (m *awsRestjson1_deserializeOpDescribeCustomPlugi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1795,7 +1768,6 @@ func (m *awsRestjson1_deserializeOpDescribeCustomPlugi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1976,10 +1948,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkerConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2019,7 +1989,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkerConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2195,10 +2164,8 @@ func (m *awsRestjson1_deserializeOpListConnectorOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2238,7 +2205,6 @@ func (m *awsRestjson1_deserializeOpListConnectorOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2374,10 +2340,8 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2417,7 +2381,6 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2553,10 +2516,8 @@ func (m *awsRestjson1_deserializeOpListCustomPlugins) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2596,7 +2557,6 @@ func (m *awsRestjson1_deserializeOpListCustomPlugins) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2732,10 +2692,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2775,7 +2733,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2902,10 +2859,8 @@ func (m *awsRestjson1_deserializeOpListWorkerConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2945,7 +2900,6 @@ func (m *awsRestjson1_deserializeOpListWorkerConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3081,10 +3035,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3096,7 +3048,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3190,10 +3141,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3205,7 +3154,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3296,10 +3244,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3339,7 +3285,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
