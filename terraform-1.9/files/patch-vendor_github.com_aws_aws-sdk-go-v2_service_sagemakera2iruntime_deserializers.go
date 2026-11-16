--- vendor/github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpDeleteHumanLoop) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpDeleteHumanLoop) Ha
 	output := &DeleteHumanLoopOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -141,10 +137,8 @@ func (m *awsRestjson1_deserializeOpDescribeHumanLoop) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -184,7 +178,6 @@ func (m *awsRestjson1_deserializeOpDescribeHumanLoop) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -369,10 +362,8 @@ func (m *awsRestjson1_deserializeOpListHumanLoops) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -412,7 +403,6 @@ func (m *awsRestjson1_deserializeOpListHumanLoops) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -539,10 +529,8 @@ func (m *awsRestjson1_deserializeOpStartHumanLoop) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -582,7 +570,6 @@ func (m *awsRestjson1_deserializeOpStartHumanLoop) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -707,10 +694,8 @@ func (m *awsRestjson1_deserializeOpStopHumanLoop) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -722,7 +707,6 @@ func (m *awsRestjson1_deserializeOpStopHumanLoop) Hand
 	output := &StopHumanLoopOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
