--- vendor/github.com/aws/aws-sdk-go-v2/service/connectparticipant/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connectparticipant/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCancelParticipantAu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpCancelParticipantAu
 	output := &CancelParticipantAuthenticationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -141,10 +137,8 @@ func (m *awsRestjson1_deserializeOpCompleteAttachmentU
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -156,7 +150,6 @@ func (m *awsRestjson1_deserializeOpCompleteAttachmentU
 	output := &CompleteAttachmentUploadOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -244,10 +237,8 @@ func (m *awsRestjson1_deserializeOpCreateParticipantCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -287,7 +278,6 @@ func (m *awsRestjson1_deserializeOpCreateParticipantCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -410,10 +400,8 @@ func (m *awsRestjson1_deserializeOpDescribeView) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -453,7 +441,6 @@ func (m *awsRestjson1_deserializeOpDescribeView) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -574,10 +561,8 @@ func (m *awsRestjson1_deserializeOpDisconnectParticipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -589,7 +574,6 @@ func (m *awsRestjson1_deserializeOpDisconnectParticipa
 	output := &DisconnectParticipantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -671,10 +655,8 @@ func (m *awsRestjson1_deserializeOpGetAttachment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -714,7 +696,6 @@ func (m *awsRestjson1_deserializeOpGetAttachment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -858,10 +839,8 @@ func (m *awsRestjson1_deserializeOpGetAuthenticationUr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -901,7 +880,6 @@ func (m *awsRestjson1_deserializeOpGetAuthenticationUr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1023,10 +1001,8 @@ func (m *awsRestjson1_deserializeOpGetTranscript) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1066,7 +1042,6 @@ func (m *awsRestjson1_deserializeOpGetTranscript) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1202,10 +1177,8 @@ func (m *awsRestjson1_deserializeOpSendEvent) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1245,7 +1218,6 @@ func (m *awsRestjson1_deserializeOpSendEvent) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1379,10 +1351,8 @@ func (m *awsRestjson1_deserializeOpSendMessage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1422,7 +1392,6 @@ func (m *awsRestjson1_deserializeOpSendMessage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1553,10 +1522,8 @@ func (m *awsRestjson1_deserializeOpStartAttachmentUplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1596,7 +1563,6 @@ func (m *awsRestjson1_deserializeOpStartAttachmentUplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
