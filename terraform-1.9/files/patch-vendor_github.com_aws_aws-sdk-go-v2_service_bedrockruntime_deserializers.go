--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/deserializers.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -50,10 +49,8 @@ func (m *awsRestjson1_deserializeOpApplyGuardrail) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpApplyGuardrail) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -250,10 +246,8 @@ func (m *awsRestjson1_deserializeOpConverse) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -293,7 +287,6 @@ func (m *awsRestjson1_deserializeOpConverse) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -460,10 +453,8 @@ func (m *awsRestjson1_deserializeOpConverseStream) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -475,7 +466,6 @@ func (m *awsRestjson1_deserializeOpConverseStream) Han
 	output := &ConverseStreamOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -572,10 +562,8 @@ func (m *awsRestjson1_deserializeOpGetAsyncInvoke) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -615,7 +603,6 @@ func (m *awsRestjson1_deserializeOpGetAsyncInvoke) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -817,10 +804,8 @@ func (m *awsRestjson1_deserializeOpInvokeModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -842,7 +827,6 @@ func (m *awsRestjson1_deserializeOpInvokeModel) Handle
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -981,10 +965,8 @@ func (m *awsRestjson1_deserializeOpInvokeModelWithResp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1001,7 +983,6 @@ func (m *awsRestjson1_deserializeOpInvokeModelWithResp
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1122,10 +1103,8 @@ func (m *awsRestjson1_deserializeOpListAsyncInvokes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1165,7 +1144,6 @@ func (m *awsRestjson1_deserializeOpListAsyncInvokes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1292,10 +1270,8 @@ func (m *awsRestjson1_deserializeOpStartAsyncInvoke) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1335,7 +1311,6 @@ func (m *awsRestjson1_deserializeOpStartAsyncInvoke) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
