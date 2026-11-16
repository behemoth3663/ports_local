--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateAgentRuntime)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateAgentRuntime)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -260,10 +256,8 @@ func (m *awsRestjson1_deserializeOpCreateAgentRuntimeE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -303,7 +297,6 @@ func (m *awsRestjson1_deserializeOpCreateAgentRuntimeE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -474,10 +467,8 @@ func (m *awsRestjson1_deserializeOpCreateApiKeyCredent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -517,7 +508,6 @@ func (m *awsRestjson1_deserializeOpCreateApiKeyCredent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -674,10 +664,8 @@ func (m *awsRestjson1_deserializeOpCreateBrowser) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -717,7 +705,6 @@ func (m *awsRestjson1_deserializeOpCreateBrowser) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -876,10 +863,8 @@ func (m *awsRestjson1_deserializeOpCreateCodeInterpret
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -919,7 +904,6 @@ func (m *awsRestjson1_deserializeOpCreateCodeInterpret
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1078,10 +1062,8 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1121,7 +1103,6 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1385,10 +1366,8 @@ func (m *awsRestjson1_deserializeOpCreateGatewayTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1428,7 +1407,6 @@ func (m *awsRestjson1_deserializeOpCreateGatewayTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1636,10 +1614,8 @@ func (m *awsRestjson1_deserializeOpCreateMemory) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1679,7 +1655,6 @@ func (m *awsRestjson1_deserializeOpCreateMemory) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1806,10 +1781,8 @@ func (m *awsRestjson1_deserializeOpCreateOauth2Credent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1849,7 +1822,6 @@ func (m *awsRestjson1_deserializeOpCreateOauth2Credent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2006,10 +1978,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkloadIdent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2049,7 +2019,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkloadIdent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2191,10 +2160,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentRuntime)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2234,7 +2201,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentRuntime)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2359,10 +2325,8 @@ func (m *awsRestjson1_deserializeOpDeleteAgentRuntimeE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2402,7 +2366,6 @@ func (m *awsRestjson1_deserializeOpDeleteAgentRuntimeE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2527,10 +2490,8 @@ func (m *awsRestjson1_deserializeOpDeleteApiKeyCredent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2542,7 +2503,6 @@ func (m *awsRestjson1_deserializeOpDeleteApiKeyCredent
 	output := &DeleteApiKeyCredentialProviderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2630,10 +2590,8 @@ func (m *awsRestjson1_deserializeOpDeleteBrowser) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2673,7 +2631,6 @@ func (m *awsRestjson1_deserializeOpDeleteBrowser) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2826,10 +2783,8 @@ func (m *awsRestjson1_deserializeOpDeleteCodeInterpret
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2869,7 +2824,6 @@ func (m *awsRestjson1_deserializeOpDeleteCodeInterpret
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3022,10 +2976,8 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3065,7 +3017,6 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3207,10 +3158,8 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3250,7 +3199,6 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3401,10 +3349,8 @@ func (m *awsRestjson1_deserializeOpDeleteMemory) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3444,7 +3390,6 @@ func (m *awsRestjson1_deserializeOpDeleteMemory) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3578,10 +3523,8 @@ func (m *awsRestjson1_deserializeOpDeleteOauth2Credent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3593,7 +3536,6 @@ func (m *awsRestjson1_deserializeOpDeleteOauth2Credent
 	output := &DeleteOauth2CredentialProviderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3681,10 +3623,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkloadIdent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3696,7 +3636,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkloadIdent
 	output := &DeleteWorkloadIdentityOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3784,10 +3723,8 @@ func (m *awsRestjson1_deserializeOpGetAgentRuntime) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3827,7 +3764,6 @@ func (m *awsRestjson1_deserializeOpGetAgentRuntime) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4062,10 +3998,8 @@ func (m *awsRestjson1_deserializeOpGetAgentRuntimeEndp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4105,7 +4039,6 @@ func (m *awsRestjson1_deserializeOpGetAgentRuntimeEndp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4328,10 +4261,8 @@ func (m *awsRestjson1_deserializeOpGetApiKeyCredential
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4371,7 +4302,6 @@ func (m *awsRestjson1_deserializeOpGetApiKeyCredential
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4548,10 +4478,8 @@ func (m *awsRestjson1_deserializeOpGetBrowser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4591,7 +4519,6 @@ func (m *awsRestjson1_deserializeOpGetBrowser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4797,10 +4724,8 @@ func (m *awsRestjson1_deserializeOpGetCodeInterpreter)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4840,7 +4765,6 @@ func (m *awsRestjson1_deserializeOpGetCodeInterpreter)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5041,10 +4965,8 @@ func (m *awsRestjson1_deserializeOpGetGateway) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5084,7 +5006,6 @@ func (m *awsRestjson1_deserializeOpGetGateway) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5345,10 +5266,8 @@ func (m *awsRestjson1_deserializeOpGetGatewayTarget) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5388,7 +5307,6 @@ func (m *awsRestjson1_deserializeOpGetGatewayTarget) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5590,10 +5508,8 @@ func (m *awsRestjson1_deserializeOpGetMemory) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5633,7 +5549,6 @@ func (m *awsRestjson1_deserializeOpGetMemory) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5754,10 +5669,8 @@ func (m *awsRestjson1_deserializeOpGetOauth2Credential
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5797,7 +5710,6 @@ func (m *awsRestjson1_deserializeOpGetOauth2Credential
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5988,10 +5900,8 @@ func (m *awsRestjson1_deserializeOpGetTokenVault) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6031,7 +5941,6 @@ func (m *awsRestjson1_deserializeOpGetTokenVault) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6180,10 +6089,8 @@ func (m *awsRestjson1_deserializeOpGetWorkloadIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6223,7 +6130,6 @@ func (m *awsRestjson1_deserializeOpGetWorkloadIdentity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6397,10 +6303,8 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimeEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6440,7 +6344,6 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimeEnd
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6567,10 +6470,8 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6610,7 +6511,6 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6737,10 +6637,8 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimeVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6780,7 +6678,6 @@ func (m *awsRestjson1_deserializeOpListAgentRuntimeVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6910,10 +6807,8 @@ func (m *awsRestjson1_deserializeOpListApiKeyCredentia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6953,7 +6848,6 @@ func (m *awsRestjson1_deserializeOpListApiKeyCredentia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7086,10 +6980,8 @@ func (m *awsRestjson1_deserializeOpListBrowsers) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7129,7 +7021,6 @@ func (m *awsRestjson1_deserializeOpListBrowsers) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7256,10 +7147,8 @@ func (m *awsRestjson1_deserializeOpListCodeInterpreter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7299,7 +7188,6 @@ func (m *awsRestjson1_deserializeOpListCodeInterpreter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7426,10 +7314,8 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7469,7 +7355,6 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7596,10 +7481,8 @@ func (m *awsRestjson1_deserializeOpListGatewayTargets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7639,7 +7522,6 @@ func (m *awsRestjson1_deserializeOpListGatewayTargets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7766,10 +7648,8 @@ func (m *awsRestjson1_deserializeOpListMemories) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7809,7 +7689,6 @@ func (m *awsRestjson1_deserializeOpListMemories) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7939,10 +7818,8 @@ func (m *awsRestjson1_deserializeOpListOauth2Credentia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7982,7 +7859,6 @@ func (m *awsRestjson1_deserializeOpListOauth2Credentia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8115,10 +7991,8 @@ func (m *awsRestjson1_deserializeOpListWorkloadIdentit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8158,7 +8032,6 @@ func (m *awsRestjson1_deserializeOpListWorkloadIdentit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8291,10 +8164,8 @@ func (m *awsRestjson1_deserializeOpSetTokenVaultCMK) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8334,7 +8205,6 @@ func (m *awsRestjson1_deserializeOpSetTokenVaultCMK) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8486,10 +8356,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentRuntime)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8529,7 +8397,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentRuntime)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8718,10 +8585,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentRuntimeE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8761,7 +8626,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentRuntimeE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8954,10 +8818,8 @@ func (m *awsRestjson1_deserializeOpUpdateApiKeyCredent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8997,7 +8859,6 @@ func (m *awsRestjson1_deserializeOpUpdateApiKeyCredent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9183,10 +9044,8 @@ func (m *awsRestjson1_deserializeOpUpdateGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9226,7 +9085,6 @@ func (m *awsRestjson1_deserializeOpUpdateGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9493,10 +9351,8 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9536,7 +9392,6 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9744,10 +9599,8 @@ func (m *awsRestjson1_deserializeOpUpdateMemory) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9787,7 +9640,6 @@ func (m *awsRestjson1_deserializeOpUpdateMemory) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9914,10 +9766,8 @@ func (m *awsRestjson1_deserializeOpUpdateOauth2Credent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9957,7 +9807,6 @@ func (m *awsRestjson1_deserializeOpUpdateOauth2Credent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10157,10 +10006,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkloadIdent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10200,7 +10047,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkloadIdent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
