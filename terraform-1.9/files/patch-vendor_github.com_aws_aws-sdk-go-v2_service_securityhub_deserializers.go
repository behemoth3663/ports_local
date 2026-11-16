--- vendor/github.com/aws/aws-sdk-go-v2/service/securityhub/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/securityhub/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAcceptAdministrator
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -62,7 +59,6 @@ func (m *awsRestjson1_deserializeOpAcceptAdministrator
 	output := &AcceptAdministratorInvitationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -147,10 +143,8 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -162,7 +156,6 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 	output := &AcceptInvitationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -247,10 +240,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteAutomati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -290,7 +281,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteAutomati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -416,10 +406,8 @@ func (m *awsRestjson1_deserializeOpBatchDisableStandar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -459,7 +447,6 @@ func (m *awsRestjson1_deserializeOpBatchDisableStandar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -580,10 +567,8 @@ func (m *awsRestjson1_deserializeOpBatchEnableStandard
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -623,7 +608,6 @@ func (m *awsRestjson1_deserializeOpBatchEnableStandard
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -744,10 +728,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAutomationR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -787,7 +769,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAutomationR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -916,10 +897,8 @@ func (m *awsRestjson1_deserializeOpBatchGetConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -959,7 +938,6 @@ func (m *awsRestjson1_deserializeOpBatchGetConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1088,10 +1066,8 @@ func (m *awsRestjson1_deserializeOpBatchGetSecurityCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1131,7 +1107,6 @@ func (m *awsRestjson1_deserializeOpBatchGetSecurityCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1254,10 +1229,8 @@ func (m *awsRestjson1_deserializeOpBatchGetStandardsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1297,7 +1270,6 @@ func (m *awsRestjson1_deserializeOpBatchGetStandardsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1420,10 +1392,8 @@ func (m *awsRestjson1_deserializeOpBatchImportFindings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1463,7 +1433,6 @@ func (m *awsRestjson1_deserializeOpBatchImportFindings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1607,10 +1576,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateAutomati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1650,7 +1617,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateAutomati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1776,10 +1742,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateFindings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1819,7 +1783,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateFindings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1942,10 +1905,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateFindings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1985,7 +1946,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateFindings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2111,10 +2071,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateStandard
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2154,7 +2112,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateStandard
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2275,10 +2232,8 @@ func (m *awsRestjson1_deserializeOpConnectorRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2318,7 +2273,6 @@ func (m *awsRestjson1_deserializeOpConnectorRegistrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2455,10 +2409,8 @@ func (m *awsRestjson1_deserializeOpCreateActionTarget)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2498,7 +2450,6 @@ func (m *awsRestjson1_deserializeOpCreateActionTarget)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2623,10 +2574,8 @@ func (m *awsRestjson1_deserializeOpCreateAggregatorV2)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2666,7 +2615,6 @@ func (m *awsRestjson1_deserializeOpCreateAggregatorV2)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2817,10 +2765,8 @@ func (m *awsRestjson1_deserializeOpCreateAutomationRul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2860,7 +2806,6 @@ func (m *awsRestjson1_deserializeOpCreateAutomationRul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2985,10 +2930,8 @@ func (m *awsRestjson1_deserializeOpCreateAutomationRul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3028,7 +2971,6 @@ func (m *awsRestjson1_deserializeOpCreateAutomationRul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3162,10 +3104,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3205,7 +3145,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3391,10 +3330,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectorV2) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3434,7 +3371,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectorV2) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3580,10 +3516,8 @@ func (m *awsRestjson1_deserializeOpCreateFindingAggreg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3623,7 +3557,6 @@ func (m *awsRestjson1_deserializeOpCreateFindingAggreg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3771,10 +3704,8 @@ func (m *awsRestjson1_deserializeOpCreateInsight) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3814,7 +3745,6 @@ func (m *awsRestjson1_deserializeOpCreateInsight) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3939,10 +3869,8 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3982,7 +3910,6 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4106,10 +4033,8 @@ func (m *awsRestjson1_deserializeOpCreateTicketV2) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4149,7 +4074,6 @@ func (m *awsRestjson1_deserializeOpCreateTicketV2) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4286,10 +4210,8 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4329,7 +4251,6 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4447,10 +4368,8 @@ func (m *awsRestjson1_deserializeOpDeleteActionTarget)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4490,7 +4409,6 @@ func (m *awsRestjson1_deserializeOpDeleteActionTarget)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4612,10 +4530,8 @@ func (m *awsRestjson1_deserializeOpDeleteAggregatorV2)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4627,7 +4543,6 @@ func (m *awsRestjson1_deserializeOpDeleteAggregatorV2)
 	output := &DeleteAggregatorV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4715,10 +4630,8 @@ func (m *awsRestjson1_deserializeOpDeleteAutomationRul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4730,7 +4643,6 @@ func (m *awsRestjson1_deserializeOpDeleteAutomationRul
 	output := &DeleteAutomationRuleV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4818,10 +4730,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4833,7 +4743,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 	output := &DeleteConfigurationPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4924,10 +4833,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorV2) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4939,7 +4846,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectorV2) 
 	output := &DeleteConnectorV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5027,10 +4933,8 @@ func (m *awsRestjson1_deserializeOpDeleteFindingAggreg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5042,7 +4946,6 @@ func (m *awsRestjson1_deserializeOpDeleteFindingAggreg
 	output := &DeleteFindingAggregatorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5130,10 +5033,8 @@ func (m *awsRestjson1_deserializeOpDeleteInsight) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5173,7 +5074,6 @@ func (m *awsRestjson1_deserializeOpDeleteInsight) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5298,10 +5198,8 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5341,7 +5239,6 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5462,10 +5359,8 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5505,7 +5400,6 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5626,10 +5520,8 @@ func (m *awsRestjson1_deserializeOpDescribeActionTarge
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5669,7 +5561,6 @@ func (m *awsRestjson1_deserializeOpDescribeActionTarge
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5796,10 +5687,8 @@ func (m *awsRestjson1_deserializeOpDescribeHub) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5839,7 +5728,6 @@ func (m *awsRestjson1_deserializeOpDescribeHub) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5991,10 +5879,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6034,7 +5920,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6179,10 +6064,8 @@ func (m *awsRestjson1_deserializeOpDescribeProducts) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6222,7 +6105,6 @@ func (m *awsRestjson1_deserializeOpDescribeProducts) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6349,10 +6231,8 @@ func (m *awsRestjson1_deserializeOpDescribeProductsV2)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6392,7 +6272,6 @@ func (m *awsRestjson1_deserializeOpDescribeProductsV2)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6522,10 +6401,8 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityHub
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6565,7 +6442,6 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityHub
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6696,10 +6572,8 @@ func (m *awsRestjson1_deserializeOpDescribeStandards) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6739,7 +6613,6 @@ func (m *awsRestjson1_deserializeOpDescribeStandards) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6863,10 +6736,8 @@ func (m *awsRestjson1_deserializeOpDescribeStandardsCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6906,7 +6777,6 @@ func (m *awsRestjson1_deserializeOpDescribeStandardsCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7033,10 +6903,8 @@ func (m *awsRestjson1_deserializeOpDisableImportFindin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7048,7 +6916,6 @@ func (m *awsRestjson1_deserializeOpDisableImportFindin
 	output := &DisableImportFindingsForProductOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7133,10 +7000,8 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7148,7 +7013,6 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 	output := &DisableOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7233,10 +7097,8 @@ func (m *awsRestjson1_deserializeOpDisableSecurityHub)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7248,7 +7110,6 @@ func (m *awsRestjson1_deserializeOpDisableSecurityHub)
 	output := &DisableSecurityHubOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7333,10 +7194,8 @@ func (m *awsRestjson1_deserializeOpDisableSecurityHubV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7348,7 +7207,6 @@ func (m *awsRestjson1_deserializeOpDisableSecurityHubV
 	output := &DisableSecurityHubV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7430,10 +7288,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7445,7 +7301,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 	output := &DisassociateFromAdministratorAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7530,10 +7385,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7545,7 +7398,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 	output := &DisassociateFromMasterAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7630,10 +7482,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7645,7 +7495,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 	output := &DisassociateMembersOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7733,10 +7582,8 @@ func (m *awsRestjson1_deserializeOpEnableImportFinding
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7776,7 +7623,6 @@ func (m *awsRestjson1_deserializeOpEnableImportFinding
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7901,10 +7747,8 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7944,7 +7788,6 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8078,10 +7921,8 @@ func (m *awsRestjson1_deserializeOpEnableSecurityHub) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8093,7 +7934,6 @@ func (m *awsRestjson1_deserializeOpEnableSecurityHub) 
 	output := &EnableSecurityHubOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8178,10 +8018,8 @@ func (m *awsRestjson1_deserializeOpEnableSecurityHubV2
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8221,7 +8059,6 @@ func (m *awsRestjson1_deserializeOpEnableSecurityHubV2
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8343,10 +8180,8 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8386,7 +8221,6 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8507,10 +8341,8 @@ func (m *awsRestjson1_deserializeOpGetAggregatorV2) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8550,7 +8382,6 @@ func (m *awsRestjson1_deserializeOpGetAggregatorV2) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8701,10 +8532,8 @@ func (m *awsRestjson1_deserializeOpGetAutomationRuleV2
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8744,7 +8573,6 @@ func (m *awsRestjson1_deserializeOpGetAutomationRuleV2
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8978,10 +8806,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9021,7 +8847,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9207,10 +9032,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9250,7 +9073,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9436,10 +9258,8 @@ func (m *awsRestjson1_deserializeOpGetConnectorV2) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9479,7 +9299,6 @@ func (m *awsRestjson1_deserializeOpGetConnectorV2) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9679,10 +9498,8 @@ func (m *awsRestjson1_deserializeOpGetEnabledStandards
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9722,7 +9539,6 @@ func (m *awsRestjson1_deserializeOpGetEnabledStandards
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9849,10 +9665,8 @@ func (m *awsRestjson1_deserializeOpGetFindingAggregato
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9892,7 +9706,6 @@ func (m *awsRestjson1_deserializeOpGetFindingAggregato
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10043,10 +9856,8 @@ func (m *awsRestjson1_deserializeOpGetFindingHistory) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10086,7 +9897,6 @@ func (m *awsRestjson1_deserializeOpGetFindingHistory) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10213,10 +10023,8 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10256,7 +10064,6 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10383,10 +10190,8 @@ func (m *awsRestjson1_deserializeOpGetFindingStatistic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10426,7 +10231,6 @@ func (m *awsRestjson1_deserializeOpGetFindingStatistic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10547,10 +10351,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsV2) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10590,7 +10392,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsV2) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10720,10 +10521,8 @@ func (m *awsRestjson1_deserializeOpGetInsightResults) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10763,7 +10562,6 @@ func (m *awsRestjson1_deserializeOpGetInsightResults) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10884,10 +10682,8 @@ func (m *awsRestjson1_deserializeOpGetInsights) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10927,7 +10723,6 @@ func (m *awsRestjson1_deserializeOpGetInsights) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11057,10 +10852,8 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11100,7 +10893,6 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11226,10 +11018,8 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11269,7 +11059,6 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11390,10 +11179,8 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11433,7 +11220,6 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11559,10 +11345,8 @@ func (m *awsRestjson1_deserializeOpGetResourcesStatist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11602,7 +11386,6 @@ func (m *awsRestjson1_deserializeOpGetResourcesStatist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11726,10 +11509,8 @@ func (m *awsRestjson1_deserializeOpGetResourcesV2) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11769,7 +11550,6 @@ func (m *awsRestjson1_deserializeOpGetResourcesV2) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11902,10 +11682,8 @@ func (m *awsRestjson1_deserializeOpGetSecurityControlD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11945,7 +11723,6 @@ func (m *awsRestjson1_deserializeOpGetSecurityControlD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12066,10 +11843,8 @@ func (m *awsRestjson1_deserializeOpInviteMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12109,7 +11884,6 @@ func (m *awsRestjson1_deserializeOpInviteMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12230,10 +12004,8 @@ func (m *awsRestjson1_deserializeOpListAggregatorsV2) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12273,7 +12045,6 @@ func (m *awsRestjson1_deserializeOpListAggregatorsV2) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12406,10 +12177,8 @@ func (m *awsRestjson1_deserializeOpListAutomationRules
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12449,7 +12218,6 @@ func (m *awsRestjson1_deserializeOpListAutomationRules
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12579,10 +12347,8 @@ func (m *awsRestjson1_deserializeOpListAutomationRules
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12622,7 +12388,6 @@ func (m *awsRestjson1_deserializeOpListAutomationRules
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12752,10 +12517,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12795,7 +12558,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12925,10 +12687,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12968,7 +12728,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13098,10 +12857,8 @@ func (m *awsRestjson1_deserializeOpListConnectorsV2) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13141,7 +12898,6 @@ func (m *awsRestjson1_deserializeOpListConnectorsV2) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13274,10 +13030,8 @@ func (m *awsRestjson1_deserializeOpListEnabledProducts
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13317,7 +13071,6 @@ func (m *awsRestjson1_deserializeOpListEnabledProducts
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13441,10 +13194,8 @@ func (m *awsRestjson1_deserializeOpListFindingAggregat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13484,7 +13235,6 @@ func (m *awsRestjson1_deserializeOpListFindingAggregat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13614,10 +13364,8 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13657,7 +13405,6 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13784,10 +13531,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13827,7 +13572,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13954,10 +13698,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13997,7 +13739,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14133,10 +13874,8 @@ func (m *awsRestjson1_deserializeOpListSecurityControl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14176,7 +13915,6 @@ func (m *awsRestjson1_deserializeOpListSecurityControl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14303,10 +14041,8 @@ func (m *awsRestjson1_deserializeOpListStandardsContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14346,7 +14082,6 @@ func (m *awsRestjson1_deserializeOpListStandardsContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14473,10 +14208,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14516,7 +14249,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14631,10 +14363,8 @@ func (m *awsRestjson1_deserializeOpStartConfigurationP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14674,7 +14404,6 @@ func (m *awsRestjson1_deserializeOpStartConfigurationP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14860,10 +14589,8 @@ func (m *awsRestjson1_deserializeOpStartConfigurationP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14875,7 +14602,6 @@ func (m *awsRestjson1_deserializeOpStartConfigurationP
 	output := &StartConfigurationPolicyDisassociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14963,10 +14689,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14978,7 +14702,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15057,10 +14780,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15072,7 +14793,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15151,10 +14871,8 @@ func (m *awsRestjson1_deserializeOpUpdateActionTarget)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15166,7 +14884,6 @@ func (m *awsRestjson1_deserializeOpUpdateActionTarget)
 	output := &UpdateActionTargetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15248,10 +14965,8 @@ func (m *awsRestjson1_deserializeOpUpdateAggregatorV2)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15291,7 +15006,6 @@ func (m *awsRestjson1_deserializeOpUpdateAggregatorV2)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15442,10 +15156,8 @@ func (m *awsRestjson1_deserializeOpUpdateAutomationRul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15457,7 +15169,6 @@ func (m *awsRestjson1_deserializeOpUpdateAutomationRul
 	output := &UpdateAutomationRuleV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15545,10 +15256,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15588,7 +15297,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15777,10 +15485,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorV2) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15792,7 +15498,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnectorV2) 
 	output := &UpdateConnectorV2Output{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15880,10 +15585,8 @@ func (m *awsRestjson1_deserializeOpUpdateFindingAggreg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15923,7 +15626,6 @@ func (m *awsRestjson1_deserializeOpUpdateFindingAggreg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16074,10 +15776,8 @@ func (m *awsRestjson1_deserializeOpUpdateFindings) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16089,7 +15789,6 @@ func (m *awsRestjson1_deserializeOpUpdateFindings) Han
 	output := &UpdateFindingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16174,10 +15873,8 @@ func (m *awsRestjson1_deserializeOpUpdateInsight) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16189,7 +15886,6 @@ func (m *awsRestjson1_deserializeOpUpdateInsight) Hand
 	output := &UpdateInsightOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16274,10 +15970,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16289,7 +15983,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 	output := &UpdateOrganizationConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16380,10 +16073,8 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityContr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16395,7 +16086,6 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityContr
 	output := &UpdateSecurityControlOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16486,10 +16176,8 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityHubCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16501,7 +16189,6 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityHubCo
 	output := &UpdateSecurityHubConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16589,10 +16276,8 @@ func (m *awsRestjson1_deserializeOpUpdateStandardsCont
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16604,7 +16289,6 @@ func (m *awsRestjson1_deserializeOpUpdateStandardsCont
 	output := &UpdateStandardsControlOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
