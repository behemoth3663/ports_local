--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagent/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagent/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAssociateAgentColla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpAssociateAgentColla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpAssociateAgentKnowl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpAssociateAgentKnowl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -387,10 +380,8 @@ func (m *awsRestjson1_deserializeOpCreateAgent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -430,7 +421,6 @@ func (m *awsRestjson1_deserializeOpCreateAgent) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -554,10 +544,8 @@ func (m *awsRestjson1_deserializeOpCreateAgentActionGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -597,7 +585,6 @@ func (m *awsRestjson1_deserializeOpCreateAgentActionGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -724,10 +711,8 @@ func (m *awsRestjson1_deserializeOpCreateAgentAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -767,7 +752,6 @@ func (m *awsRestjson1_deserializeOpCreateAgentAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -894,10 +878,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -937,7 +919,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1064,10 +1045,8 @@ func (m *awsRestjson1_deserializeOpCreateFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1107,7 +1086,6 @@ func (m *awsRestjson1_deserializeOpCreateFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1329,10 +1307,8 @@ func (m *awsRestjson1_deserializeOpCreateFlowAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1372,7 +1348,6 @@ func (m *awsRestjson1_deserializeOpCreateFlowAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1575,10 +1550,8 @@ func (m *awsRestjson1_deserializeOpCreateFlowVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1618,7 +1591,6 @@ func (m *awsRestjson1_deserializeOpCreateFlowVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1830,10 +1802,8 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1873,7 +1843,6 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1997,10 +1966,8 @@ func (m *awsRestjson1_deserializeOpCreatePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2040,7 +2007,6 @@ func (m *awsRestjson1_deserializeOpCreatePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2253,10 +2219,8 @@ func (m *awsRestjson1_deserializeOpCreatePromptVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2296,7 +2260,6 @@ func (m *awsRestjson1_deserializeOpCreatePromptVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2512,10 +2475,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2555,7 +2516,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgent) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2692,10 +2652,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentActionGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2707,7 +2665,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentActionGr
 	output := &DeleteAgentActionGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2795,10 +2752,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2838,7 +2793,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2981,10 +2935,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3024,7 +2976,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3170,10 +3121,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3213,7 +3162,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3359,10 +3307,8 @@ func (m *awsRestjson1_deserializeOpDeleteFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3402,7 +3348,6 @@ func (m *awsRestjson1_deserializeOpDeleteFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3530,10 +3475,8 @@ func (m *awsRestjson1_deserializeOpDeleteFlowAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3573,7 +3516,6 @@ func (m *awsRestjson1_deserializeOpDeleteFlowAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3710,10 +3652,8 @@ func (m *awsRestjson1_deserializeOpDeleteFlowVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3753,7 +3693,6 @@ func (m *awsRestjson1_deserializeOpDeleteFlowVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3890,10 +3829,8 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3933,7 +3870,6 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4070,10 +4006,8 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4113,7 +4047,6 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4237,10 +4170,8 @@ func (m *awsRestjson1_deserializeOpDeletePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4280,7 +4211,6 @@ func (m *awsRestjson1_deserializeOpDeletePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4417,10 +4347,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAgentCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4432,7 +4360,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAgentCo
 	output := &DisassociateAgentCollaboratorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4520,10 +4447,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAgentKn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4535,7 +4460,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAgentKn
 	output := &DisassociateAgentKnowledgeBaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4623,10 +4547,8 @@ func (m *awsRestjson1_deserializeOpGetAgent) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4666,7 +4588,6 @@ func (m *awsRestjson1_deserializeOpGetAgent) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4787,10 +4708,8 @@ func (m *awsRestjson1_deserializeOpGetAgentActionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4830,7 +4749,6 @@ func (m *awsRestjson1_deserializeOpGetAgentActionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4951,10 +4869,8 @@ func (m *awsRestjson1_deserializeOpGetAgentAlias) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4994,7 +4910,6 @@ func (m *awsRestjson1_deserializeOpGetAgentAlias) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5115,10 +5030,8 @@ func (m *awsRestjson1_deserializeOpGetAgentCollaborato
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5158,7 +5071,6 @@ func (m *awsRestjson1_deserializeOpGetAgentCollaborato
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5279,10 +5191,8 @@ func (m *awsRestjson1_deserializeOpGetAgentKnowledgeBa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5322,7 +5232,6 @@ func (m *awsRestjson1_deserializeOpGetAgentKnowledgeBa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5443,10 +5352,8 @@ func (m *awsRestjson1_deserializeOpGetAgentVersion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5486,7 +5393,6 @@ func (m *awsRestjson1_deserializeOpGetAgentVersion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5607,10 +5513,8 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5650,7 +5554,6 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5771,10 +5674,8 @@ func (m *awsRestjson1_deserializeOpGetFlow) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5814,7 +5715,6 @@ func (m *awsRestjson1_deserializeOpGetFlow) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6038,10 +5938,8 @@ func (m *awsRestjson1_deserializeOpGetFlowAlias) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6081,7 +5979,6 @@ func (m *awsRestjson1_deserializeOpGetFlowAlias) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6278,10 +6175,8 @@ func (m *awsRestjson1_deserializeOpGetFlowVersion) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6321,7 +6216,6 @@ func (m *awsRestjson1_deserializeOpGetFlowVersion) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6527,10 +6421,8 @@ func (m *awsRestjson1_deserializeOpGetIngestionJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6570,7 +6462,6 @@ func (m *awsRestjson1_deserializeOpGetIngestionJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6691,10 +6582,8 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6734,7 +6623,6 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6855,10 +6743,8 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBaseDoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6898,7 +6784,6 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBaseDoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7022,10 +6907,8 @@ func (m *awsRestjson1_deserializeOpGetPrompt) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7065,7 +6948,6 @@ func (m *awsRestjson1_deserializeOpGetPrompt) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7275,10 +7157,8 @@ func (m *awsRestjson1_deserializeOpIngestKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7318,7 +7198,6 @@ func (m *awsRestjson1_deserializeOpIngestKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7442,10 +7321,8 @@ func (m *awsRestjson1_deserializeOpListAgentActionGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7485,7 +7362,6 @@ func (m *awsRestjson1_deserializeOpListAgentActionGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7615,10 +7491,8 @@ func (m *awsRestjson1_deserializeOpListAgentAliases) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7658,7 +7532,6 @@ func (m *awsRestjson1_deserializeOpListAgentAliases) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7788,10 +7661,8 @@ func (m *awsRestjson1_deserializeOpListAgentCollaborat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7831,7 +7702,6 @@ func (m *awsRestjson1_deserializeOpListAgentCollaborat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7961,10 +7831,8 @@ func (m *awsRestjson1_deserializeOpListAgentKnowledgeB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8004,7 +7872,6 @@ func (m *awsRestjson1_deserializeOpListAgentKnowledgeB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8134,10 +8001,8 @@ func (m *awsRestjson1_deserializeOpListAgents) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8177,7 +8042,6 @@ func (m *awsRestjson1_deserializeOpListAgents) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8304,10 +8168,8 @@ func (m *awsRestjson1_deserializeOpListAgentVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8347,7 +8209,6 @@ func (m *awsRestjson1_deserializeOpListAgentVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8477,10 +8338,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8520,7 +8379,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8650,10 +8508,8 @@ func (m *awsRestjson1_deserializeOpListFlowAliases) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8693,7 +8549,6 @@ func (m *awsRestjson1_deserializeOpListFlowAliases) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8823,10 +8678,8 @@ func (m *awsRestjson1_deserializeOpListFlows) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8866,7 +8719,6 @@ func (m *awsRestjson1_deserializeOpListFlows) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8993,10 +8845,8 @@ func (m *awsRestjson1_deserializeOpListFlowVersions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9036,7 +8886,6 @@ func (m *awsRestjson1_deserializeOpListFlowVersions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9166,10 +9015,8 @@ func (m *awsRestjson1_deserializeOpListIngestionJobs) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9209,7 +9056,6 @@ func (m *awsRestjson1_deserializeOpListIngestionJobs) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9339,10 +9185,8 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBaseDo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9382,7 +9226,6 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBaseDo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9515,10 +9358,8 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9558,7 +9399,6 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9685,10 +9525,8 @@ func (m *awsRestjson1_deserializeOpListPrompts) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9728,7 +9566,6 @@ func (m *awsRestjson1_deserializeOpListPrompts) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9858,10 +9695,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9901,7 +9736,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10022,10 +9856,8 @@ func (m *awsRestjson1_deserializeOpPrepareAgent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10065,7 +9897,6 @@ func (m *awsRestjson1_deserializeOpPrepareAgent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10227,10 +10058,8 @@ func (m *awsRestjson1_deserializeOpPrepareFlow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10270,7 +10099,6 @@ func (m *awsRestjson1_deserializeOpPrepareFlow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10410,10 +10238,8 @@ func (m *awsRestjson1_deserializeOpStartIngestionJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10453,7 +10279,6 @@ func (m *awsRestjson1_deserializeOpStartIngestionJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10580,10 +10405,8 @@ func (m *awsRestjson1_deserializeOpStopIngestionJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10623,7 +10446,6 @@ func (m *awsRestjson1_deserializeOpStopIngestionJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10747,10 +10569,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10762,7 +10582,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10850,10 +10669,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10865,7 +10682,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10950,10 +10766,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgent) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10993,7 +10807,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgent) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11120,10 +10933,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentActionGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11163,7 +10974,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentActionGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11290,10 +11100,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11333,7 +11141,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11460,10 +11267,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentCollabor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11503,7 +11308,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentCollabor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11630,10 +11434,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentKnowledg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11673,7 +11475,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentKnowledg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11797,10 +11598,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11840,7 +11639,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11964,10 +11762,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlow) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12007,7 +11803,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlow) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12232,10 +12027,8 @@ func (m *awsRestjson1_deserializeOpUpdateFlowAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12275,7 +12068,6 @@ func (m *awsRestjson1_deserializeOpUpdateFlowAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12478,10 +12270,8 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12521,7 +12311,6 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12645,10 +12434,8 @@ func (m *awsRestjson1_deserializeOpUpdatePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12688,7 +12475,6 @@ func (m *awsRestjson1_deserializeOpUpdatePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12904,10 +12690,8 @@ func (m *awsRestjson1_deserializeOpValidateFlowDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12947,7 +12731,6 @@ func (m *awsRestjson1_deserializeOpValidateFlowDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
