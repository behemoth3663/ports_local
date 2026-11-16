--- vendor/github.com/aws/aws-sdk-go-v2/service/observabilityadmin/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/observabilityadmin/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpGetTelemetryEvaluat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpGetTelemetryEvaluat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -213,10 +209,8 @@ func (m *awsRestjson1_deserializeOpGetTelemetryEvaluat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -256,7 +250,6 @@ func (m *awsRestjson1_deserializeOpGetTelemetryEvaluat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -384,10 +377,8 @@ func (m *awsRestjson1_deserializeOpListResourceTelemet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_deserializeOpListResourceTelemet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -551,10 +541,8 @@ func (m *awsRestjson1_deserializeOpListResourceTelemet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -594,7 +582,6 @@ func (m *awsRestjson1_deserializeOpListResourceTelemet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -718,10 +705,8 @@ func (m *awsRestjson1_deserializeOpStartTelemetryEvalu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -739,7 +724,6 @@ func (m *awsRestjson1_deserializeOpStartTelemetryEvalu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -818,10 +802,8 @@ func (m *awsRestjson1_deserializeOpStartTelemetryEvalu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -839,7 +821,6 @@ func (m *awsRestjson1_deserializeOpStartTelemetryEvalu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -918,10 +899,8 @@ func (m *awsRestjson1_deserializeOpStopTelemetryEvalua
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -939,7 +918,6 @@ func (m *awsRestjson1_deserializeOpStopTelemetryEvalua
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1018,10 +996,8 @@ func (m *awsRestjson1_deserializeOpStopTelemetryEvalua
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1039,7 +1015,6 @@ func (m *awsRestjson1_deserializeOpStopTelemetryEvalua
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
