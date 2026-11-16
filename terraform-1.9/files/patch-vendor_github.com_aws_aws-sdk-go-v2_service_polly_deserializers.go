--- vendor/github.com/aws/aws-sdk-go-v2/service/polly/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/polly/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpDeleteLexicon) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpDeleteLexicon) Hand
 	output := &DeleteLexiconOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -136,10 +132,8 @@ func (m *awsRestjson1_deserializeOpDescribeVoices) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -179,7 +173,6 @@ func (m *awsRestjson1_deserializeOpDescribeVoices) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -300,10 +293,8 @@ func (m *awsRestjson1_deserializeOpGetLexicon) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -343,7 +334,6 @@ func (m *awsRestjson1_deserializeOpGetLexicon) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -460,10 +450,8 @@ func (m *awsRestjson1_deserializeOpGetSpeechSynthesisT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -503,7 +491,6 @@ func (m *awsRestjson1_deserializeOpGetSpeechSynthesisT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -618,10 +605,8 @@ func (m *awsRestjson1_deserializeOpListLexicons) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -661,7 +646,6 @@ func (m *awsRestjson1_deserializeOpListLexicons) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -782,10 +766,8 @@ func (m *awsRestjson1_deserializeOpListSpeechSynthesis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -825,7 +807,6 @@ func (m *awsRestjson1_deserializeOpListSpeechSynthesis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -946,10 +927,8 @@ func (m *awsRestjson1_deserializeOpPutLexicon) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -961,7 +940,6 @@ func (m *awsRestjson1_deserializeOpPutLexicon) HandleD
 	output := &PutLexiconOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1052,10 +1030,8 @@ func (m *awsRestjson1_deserializeOpStartSpeechSynthesi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1095,7 +1071,6 @@ func (m *awsRestjson1_deserializeOpStartSpeechSynthesi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1237,10 +1212,8 @@ func (m *awsRestjson1_deserializeOpSynthesizeSpeech) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1262,7 +1235,6 @@ func (m *awsRestjson1_deserializeOpSynthesizeSpeech) H
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
