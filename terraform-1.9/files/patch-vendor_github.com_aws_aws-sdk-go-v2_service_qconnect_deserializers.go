--- vendor/github.com/aws/aws-sdk-go-v2/service/qconnect/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/qconnect/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpActivateMessageTemp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpActivateMessageTemp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -235,10 +231,8 @@ func (m *awsRestjson1_deserializeOpCreateAIAgent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -278,7 +272,6 @@ func (m *awsRestjson1_deserializeOpCreateAIAgent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -405,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateAIAgentVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -448,7 +439,6 @@ func (m *awsRestjson1_deserializeOpCreateAIAgentVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -588,10 +578,8 @@ func (m *awsRestjson1_deserializeOpCreateAIGuardrail) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -631,7 +619,6 @@ func (m *awsRestjson1_deserializeOpCreateAIGuardrail) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -758,10 +745,8 @@ func (m *awsRestjson1_deserializeOpCreateAIGuardrailVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -801,7 +786,6 @@ func (m *awsRestjson1_deserializeOpCreateAIGuardrailVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -941,10 +925,8 @@ func (m *awsRestjson1_deserializeOpCreateAIPrompt) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -984,7 +966,6 @@ func (m *awsRestjson1_deserializeOpCreateAIPrompt) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1111,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpCreateAIPromptVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1154,7 +1133,6 @@ func (m *awsRestjson1_deserializeOpCreateAIPromptVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1294,10 +1272,8 @@ func (m *awsRestjson1_deserializeOpCreateAssistant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1337,7 +1313,6 @@ func (m *awsRestjson1_deserializeOpCreateAssistant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1458,10 +1433,8 @@ func (m *awsRestjson1_deserializeOpCreateAssistantAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1501,7 +1474,6 @@ func (m *awsRestjson1_deserializeOpCreateAssistantAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1622,10 +1594,8 @@ func (m *awsRestjson1_deserializeOpCreateContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1665,7 +1635,6 @@ func (m *awsRestjson1_deserializeOpCreateContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1789,10 +1758,8 @@ func (m *awsRestjson1_deserializeOpCreateContentAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1832,7 +1799,6 @@ func (m *awsRestjson1_deserializeOpCreateContentAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1959,10 +1925,8 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2002,7 +1966,6 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2123,10 +2086,8 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2166,7 +2127,6 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2290,10 +2250,8 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2333,7 +2291,6 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2460,10 +2417,8 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2503,7 +2458,6 @@ func (m *awsRestjson1_deserializeOpCreateMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2627,10 +2581,8 @@ func (m *awsRestjson1_deserializeOpCreateQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2670,7 +2622,6 @@ func (m *awsRestjson1_deserializeOpCreateQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2794,10 +2745,8 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2837,7 +2786,6 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2958,10 +2906,8 @@ func (m *awsRestjson1_deserializeOpDeactivateMessageTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3001,7 +2947,6 @@ func (m *awsRestjson1_deserializeOpDeactivateMessageTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3148,10 +3093,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIAgent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3163,7 +3106,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIAgent) Hand
 	output := &DeleteAIAgentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3248,10 +3190,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIAgentVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3263,7 +3203,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIAgentVersio
 	output := &DeleteAIAgentVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3351,10 +3290,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIGuardrail) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3366,7 +3303,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIGuardrail) 
 	output := &DeleteAIGuardrailOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3454,10 +3390,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIGuardrailVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3469,7 +3403,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIGuardrailVe
 	output := &DeleteAIGuardrailVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3557,10 +3490,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIPrompt) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3572,7 +3503,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIPrompt) Han
 	output := &DeleteAIPromptOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3657,10 +3587,8 @@ func (m *awsRestjson1_deserializeOpDeleteAIPromptVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3672,7 +3600,6 @@ func (m *awsRestjson1_deserializeOpDeleteAIPromptVersi
 	output := &DeleteAIPromptVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3760,10 +3687,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssistant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3775,7 +3700,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssistant) Ha
 	output := &DeleteAssistantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3857,10 +3781,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssistantAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3872,7 +3794,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssistantAsso
 	output := &DeleteAssistantAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3954,10 +3875,8 @@ func (m *awsRestjson1_deserializeOpDeleteContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3969,7 +3888,6 @@ func (m *awsRestjson1_deserializeOpDeleteContent) Hand
 	output := &DeleteContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3972,8 @@ func (m *awsRestjson1_deserializeOpDeleteContentAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4069,7 +3985,6 @@ func (m *awsRestjson1_deserializeOpDeleteContentAssoci
 	output := &DeleteContentAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4151,10 +4066,8 @@ func (m *awsRestjson1_deserializeOpDeleteImportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4166,7 +4079,6 @@ func (m *awsRestjson1_deserializeOpDeleteImportJob) Ha
 	output := &DeleteImportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4251,10 +4163,8 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4266,7 +4176,6 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 	output := &DeleteKnowledgeBaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4351,10 +4260,8 @@ func (m *awsRestjson1_deserializeOpDeleteMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4366,7 +4273,6 @@ func (m *awsRestjson1_deserializeOpDeleteMessageTempla
 	output := &DeleteMessageTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4451,10 +4357,8 @@ func (m *awsRestjson1_deserializeOpDeleteMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4466,7 +4370,6 @@ func (m *awsRestjson1_deserializeOpDeleteMessageTempla
 	output := &DeleteMessageTemplateAttachmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4551,10 +4454,8 @@ func (m *awsRestjson1_deserializeOpDeleteQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4566,7 +4467,6 @@ func (m *awsRestjson1_deserializeOpDeleteQuickResponse
 	output := &DeleteQuickResponseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4648,10 +4548,8 @@ func (m *awsRestjson1_deserializeOpGetAIAgent) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4691,7 +4589,6 @@ func (m *awsRestjson1_deserializeOpGetAIAgent) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4825,10 +4722,8 @@ func (m *awsRestjson1_deserializeOpGetAIGuardrail) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4868,7 +4763,6 @@ func (m *awsRestjson1_deserializeOpGetAIGuardrail) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5002,10 +4896,8 @@ func (m *awsRestjson1_deserializeOpGetAIPrompt) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5045,7 +4937,6 @@ func (m *awsRestjson1_deserializeOpGetAIPrompt) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5179,10 +5070,8 @@ func (m *awsRestjson1_deserializeOpGetAssistant) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5222,7 +5111,6 @@ func (m *awsRestjson1_deserializeOpGetAssistant) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5340,10 +5228,8 @@ func (m *awsRestjson1_deserializeOpGetAssistantAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5383,7 +5269,6 @@ func (m *awsRestjson1_deserializeOpGetAssistantAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5501,10 +5386,8 @@ func (m *awsRestjson1_deserializeOpGetContent) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5544,7 +5427,6 @@ func (m *awsRestjson1_deserializeOpGetContent) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5662,10 +5544,8 @@ func (m *awsRestjson1_deserializeOpGetContentAssociati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5705,7 +5585,6 @@ func (m *awsRestjson1_deserializeOpGetContentAssociati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5823,10 +5702,8 @@ func (m *awsRestjson1_deserializeOpGetContentSummary) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5866,7 +5743,6 @@ func (m *awsRestjson1_deserializeOpGetContentSummary) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5984,10 +5860,8 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6027,7 +5901,6 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6142,10 +6015,8 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6185,7 +6056,6 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6303,10 +6173,8 @@ func (m *awsRestjson1_deserializeOpGetMessageTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6346,7 +6214,6 @@ func (m *awsRestjson1_deserializeOpGetMessageTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6467,10 +6334,8 @@ func (m *awsRestjson1_deserializeOpGetNextMessage) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6510,7 +6375,6 @@ func (m *awsRestjson1_deserializeOpGetNextMessage) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6662,10 +6526,8 @@ func (m *awsRestjson1_deserializeOpGetQuickResponse) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6705,7 +6567,6 @@ func (m *awsRestjson1_deserializeOpGetQuickResponse) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6823,10 +6684,8 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6866,7 +6725,6 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6986,10 +6844,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7029,7 +6885,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7147,10 +7002,8 @@ func (m *awsRestjson1_deserializeOpListAIAgents) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7190,7 +7043,6 @@ func (m *awsRestjson1_deserializeOpListAIAgents) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7320,10 +7172,8 @@ func (m *awsRestjson1_deserializeOpListAIAgentVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7363,7 +7213,6 @@ func (m *awsRestjson1_deserializeOpListAIAgentVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7493,10 +7342,8 @@ func (m *awsRestjson1_deserializeOpListAIGuardrails) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7536,7 +7383,6 @@ func (m *awsRestjson1_deserializeOpListAIGuardrails) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7666,10 +7512,8 @@ func (m *awsRestjson1_deserializeOpListAIGuardrailVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7709,7 +7553,6 @@ func (m *awsRestjson1_deserializeOpListAIGuardrailVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7839,10 +7682,8 @@ func (m *awsRestjson1_deserializeOpListAIPrompts) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7882,7 +7723,6 @@ func (m *awsRestjson1_deserializeOpListAIPrompts) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8012,10 +7852,8 @@ func (m *awsRestjson1_deserializeOpListAIPromptVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8055,7 +7893,6 @@ func (m *awsRestjson1_deserializeOpListAIPromptVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8185,10 +8022,8 @@ func (m *awsRestjson1_deserializeOpListAssistantAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8228,7 +8063,6 @@ func (m *awsRestjson1_deserializeOpListAssistantAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8352,10 +8186,8 @@ func (m *awsRestjson1_deserializeOpListAssistants) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8395,7 +8227,6 @@ func (m *awsRestjson1_deserializeOpListAssistants) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8519,10 +8350,8 @@ func (m *awsRestjson1_deserializeOpListContentAssociat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8562,7 +8391,6 @@ func (m *awsRestjson1_deserializeOpListContentAssociat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8689,10 +8517,8 @@ func (m *awsRestjson1_deserializeOpListContents) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8732,7 +8558,6 @@ func (m *awsRestjson1_deserializeOpListContents) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8856,10 +8681,8 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8899,7 +8722,6 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9020,10 +8842,8 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9063,7 +8883,6 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9184,10 +9003,8 @@ func (m *awsRestjson1_deserializeOpListMessages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9227,7 +9044,6 @@ func (m *awsRestjson1_deserializeOpListMessages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9351,10 +9167,8 @@ func (m *awsRestjson1_deserializeOpListMessageTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9394,7 +9208,6 @@ func (m *awsRestjson1_deserializeOpListMessageTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9521,10 +9334,8 @@ func (m *awsRestjson1_deserializeOpListMessageTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9564,7 +9375,6 @@ func (m *awsRestjson1_deserializeOpListMessageTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9691,10 +9501,8 @@ func (m *awsRestjson1_deserializeOpListQuickResponses)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9734,7 +9542,6 @@ func (m *awsRestjson1_deserializeOpListQuickResponses)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9858,10 +9665,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9901,7 +9706,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10010,10 +9814,8 @@ func (m *awsRestjson1_deserializeOpNotifyRecommendatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10053,7 +9855,6 @@ func (m *awsRestjson1_deserializeOpNotifyRecommendatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10173,10 +9974,8 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10216,7 +10015,6 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10367,10 +10165,8 @@ func (m *awsRestjson1_deserializeOpQueryAssistant) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10410,7 +10206,6 @@ func (m *awsRestjson1_deserializeOpQueryAssistant) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10537,10 +10332,8 @@ func (m *awsRestjson1_deserializeOpRemoveAssistantAIAg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10552,7 +10345,6 @@ func (m *awsRestjson1_deserializeOpRemoveAssistantAIAg
 	output := &RemoveAssistantAIAgentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10634,10 +10426,8 @@ func (m *awsRestjson1_deserializeOpRemoveKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10649,7 +10439,6 @@ func (m *awsRestjson1_deserializeOpRemoveKnowledgeBase
 	output := &RemoveKnowledgeBaseTemplateUriOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10728,10 +10517,8 @@ func (m *awsRestjson1_deserializeOpRenderMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10771,7 +10558,6 @@ func (m *awsRestjson1_deserializeOpRenderMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10899,10 +10685,8 @@ func (m *awsRestjson1_deserializeOpSearchContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10942,7 +10726,6 @@ func (m *awsRestjson1_deserializeOpSearchContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11069,10 +10852,8 @@ func (m *awsRestjson1_deserializeOpSearchMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11112,7 +10893,6 @@ func (m *awsRestjson1_deserializeOpSearchMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11242,10 +11022,8 @@ func (m *awsRestjson1_deserializeOpSearchQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11285,7 +11063,6 @@ func (m *awsRestjson1_deserializeOpSearchQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11415,10 +11192,8 @@ func (m *awsRestjson1_deserializeOpSearchSessions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11458,7 +11233,6 @@ func (m *awsRestjson1_deserializeOpSearchSessions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11585,10 +11359,8 @@ func (m *awsRestjson1_deserializeOpSendMessage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11628,7 +11400,6 @@ func (m *awsRestjson1_deserializeOpSendMessage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11770,10 +11541,8 @@ func (m *awsRestjson1_deserializeOpStartContentUpload)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11813,7 +11582,6 @@ func (m *awsRestjson1_deserializeOpStartContentUpload)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11965,10 +11733,8 @@ func (m *awsRestjson1_deserializeOpStartImportJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12008,7 +11774,6 @@ func (m *awsRestjson1_deserializeOpStartImportJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12132,10 +11897,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12147,7 +11910,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12223,10 +11985,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12238,7 +11998,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12311,10 +12070,8 @@ func (m *awsRestjson1_deserializeOpUpdateAIAgent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12354,7 +12111,6 @@ func (m *awsRestjson1_deserializeOpUpdateAIAgent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12478,10 +12234,8 @@ func (m *awsRestjson1_deserializeOpUpdateAIGuardrail) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12521,7 +12275,6 @@ func (m *awsRestjson1_deserializeOpUpdateAIGuardrail) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12645,10 +12398,8 @@ func (m *awsRestjson1_deserializeOpUpdateAIPrompt) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12688,7 +12439,6 @@ func (m *awsRestjson1_deserializeOpUpdateAIPrompt) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12812,10 +12562,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssistantAIAg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12855,7 +12603,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssistantAIAg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12973,10 +12720,8 @@ func (m *awsRestjson1_deserializeOpUpdateContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13016,7 +12761,6 @@ func (m *awsRestjson1_deserializeOpUpdateContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13137,10 +12881,8 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13180,7 +12922,6 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13295,10 +13036,8 @@ func (m *awsRestjson1_deserializeOpUpdateMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13338,7 +13077,6 @@ func (m *awsRestjson1_deserializeOpUpdateMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13459,10 +13197,8 @@ func (m *awsRestjson1_deserializeOpUpdateMessageTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13502,7 +13238,6 @@ func (m *awsRestjson1_deserializeOpUpdateMessageTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13623,10 +13358,8 @@ func (m *awsRestjson1_deserializeOpUpdateQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13666,7 +13399,6 @@ func (m *awsRestjson1_deserializeOpUpdateQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13790,10 +13522,8 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13833,7 +13563,6 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13951,10 +13680,8 @@ func (m *awsRestjson1_deserializeOpUpdateSessionData) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13994,7 +13721,6 @@ func (m *awsRestjson1_deserializeOpUpdateSessionData) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
