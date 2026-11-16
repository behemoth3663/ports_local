--- vendor/github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpGetPendingJobExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpGetPendingJobExecut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -381,10 +374,8 @@ func (m *awsRestjson1_deserializeOpStartCommandExecuti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -424,7 +415,6 @@ func (m *awsRestjson1_deserializeOpStartCommandExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +542,8 @@ func (m *awsRestjson1_deserializeOpStartNextPendingJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -595,7 +583,6 @@ func (m *awsRestjson1_deserializeOpStartNextPendingJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -716,10 +703,8 @@ func (m *awsRestjson1_deserializeOpUpdateJobExecution)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -759,7 +744,6 @@ func (m *awsRestjson1_deserializeOpUpdateJobExecution)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
