--- vendor/github.com/aws/aws-sdk-go-v2/service/cleanrooms/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cleanrooms/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchGetCollaborati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchGetCollaborati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpBatchGetSchema) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -259,7 +253,6 @@ func (m *awsRestjson1_deserializeOpBatchGetSchema) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -385,10 +378,8 @@ func (m *awsRestjson1_deserializeOpBatchGetSchemaAnaly
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -428,7 +419,6 @@ func (m *awsRestjson1_deserializeOpBatchGetSchemaAnaly
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -554,10 +544,8 @@ func (m *awsRestjson1_deserializeOpCreateAnalysisTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -597,7 +585,6 @@ func (m *awsRestjson1_deserializeOpCreateAnalysisTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -724,10 +711,8 @@ func (m *awsRestjson1_deserializeOpCreateCollaboration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -767,7 +752,6 @@ func (m *awsRestjson1_deserializeOpCreateCollaboration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -888,10 +872,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -931,7 +913,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1058,10 +1039,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1101,7 +1080,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1228,10 +1206,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1271,7 +1247,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1398,10 +1373,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1441,7 +1414,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1568,10 +1540,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1611,7 +1581,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1735,10 +1704,8 @@ func (m *awsRestjson1_deserializeOpCreateIdMappingTabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1778,7 +1745,6 @@ func (m *awsRestjson1_deserializeOpCreateIdMappingTabl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1905,10 +1871,8 @@ func (m *awsRestjson1_deserializeOpCreateIdNamespaceAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1948,7 +1912,6 @@ func (m *awsRestjson1_deserializeOpCreateIdNamespaceAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2075,10 +2038,8 @@ func (m *awsRestjson1_deserializeOpCreateMembership) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2118,7 +2079,6 @@ func (m *awsRestjson1_deserializeOpCreateMembership) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2245,10 +2205,8 @@ func (m *awsRestjson1_deserializeOpCreatePrivacyBudget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2288,7 +2246,6 @@ func (m *awsRestjson1_deserializeOpCreatePrivacyBudget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2412,10 +2369,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnalysisTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2427,7 +2382,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnalysisTempl
 	output := &DeleteAnalysisTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2512,10 +2466,8 @@ func (m *awsRestjson1_deserializeOpDeleteCollaboration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2527,7 +2479,6 @@ func (m *awsRestjson1_deserializeOpDeleteCollaboration
 	output := &DeleteCollaborationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2609,10 +2560,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2624,7 +2573,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 	output := &DeleteConfiguredAudienceModelAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2709,10 +2657,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2724,7 +2670,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 	output := &DeleteConfiguredTableOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2812,10 +2757,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2827,7 +2770,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 	output := &DeleteConfiguredTableAnalysisRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2915,10 +2857,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2930,7 +2870,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 	output := &DeleteConfiguredTableAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3018,10 +2957,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3033,7 +2970,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredTab
 	output := &DeleteConfiguredTableAssociationAnalysisRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3121,10 +3057,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdMappingTabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3136,7 +3070,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdMappingTabl
 	output := &DeleteIdMappingTableOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3221,10 +3154,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdNamespaceAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3236,7 +3167,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdNamespaceAs
 	output := &DeleteIdNamespaceAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3321,10 +3251,8 @@ func (m *awsRestjson1_deserializeOpDeleteMember) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3336,7 +3264,6 @@ func (m *awsRestjson1_deserializeOpDeleteMember) Handl
 	output := &DeleteMemberOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3424,10 +3351,8 @@ func (m *awsRestjson1_deserializeOpDeleteMembership) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3439,7 +3364,6 @@ func (m *awsRestjson1_deserializeOpDeleteMembership) H
 	output := &DeleteMembershipOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3527,10 +3451,8 @@ func (m *awsRestjson1_deserializeOpDeletePrivacyBudget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3542,7 +3464,6 @@ func (m *awsRestjson1_deserializeOpDeletePrivacyBudget
 	output := &DeletePrivacyBudgetTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3627,10 +3548,8 @@ func (m *awsRestjson1_deserializeOpGetAnalysisTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3670,7 +3589,6 @@ func (m *awsRestjson1_deserializeOpGetAnalysisTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3791,10 +3709,8 @@ func (m *awsRestjson1_deserializeOpGetCollaboration) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3834,7 +3750,6 @@ func (m *awsRestjson1_deserializeOpGetCollaboration) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3952,10 +3867,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationAna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3995,7 +3908,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationAna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4116,10 +4028,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4159,7 +4069,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4280,10 +4189,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationIdN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4323,7 +4230,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationIdN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4444,10 +4350,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationPri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4487,7 +4391,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationPri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4608,10 +4511,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4651,7 +4552,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4772,10 +4672,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTable)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4815,7 +4713,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTable)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4936,10 +4833,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4979,7 +4874,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5100,10 +4994,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5143,7 +5035,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5264,10 +5155,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5307,7 +5196,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredTableA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5428,10 +5316,8 @@ func (m *awsRestjson1_deserializeOpGetIdMappingTable) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5471,7 +5357,6 @@ func (m *awsRestjson1_deserializeOpGetIdMappingTable) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5592,10 +5477,8 @@ func (m *awsRestjson1_deserializeOpGetIdNamespaceAssoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5635,7 +5518,6 @@ func (m *awsRestjson1_deserializeOpGetIdNamespaceAssoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5756,10 +5638,8 @@ func (m *awsRestjson1_deserializeOpGetMembership) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5799,7 +5679,6 @@ func (m *awsRestjson1_deserializeOpGetMembership) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5920,10 +5799,8 @@ func (m *awsRestjson1_deserializeOpGetPrivacyBudgetTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5963,7 +5840,6 @@ func (m *awsRestjson1_deserializeOpGetPrivacyBudgetTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6084,10 +5960,8 @@ func (m *awsRestjson1_deserializeOpGetProtectedJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6127,7 +6001,6 @@ func (m *awsRestjson1_deserializeOpGetProtectedJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6248,10 +6121,8 @@ func (m *awsRestjson1_deserializeOpGetProtectedQuery) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6291,7 +6162,6 @@ func (m *awsRestjson1_deserializeOpGetProtectedQuery) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6412,10 +6282,8 @@ func (m *awsRestjson1_deserializeOpGetSchema) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6455,7 +6323,6 @@ func (m *awsRestjson1_deserializeOpGetSchema) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6576,10 +6443,8 @@ func (m *awsRestjson1_deserializeOpGetSchemaAnalysisRu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6619,7 +6484,6 @@ func (m *awsRestjson1_deserializeOpGetSchemaAnalysisRu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6740,10 +6604,8 @@ func (m *awsRestjson1_deserializeOpListAnalysisTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6783,7 +6645,6 @@ func (m *awsRestjson1_deserializeOpListAnalysisTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6913,10 +6774,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6956,7 +6815,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7086,10 +6944,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7129,7 +6985,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7259,10 +7114,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationId
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7302,7 +7155,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationId
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7432,10 +7284,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7475,7 +7325,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7605,10 +7454,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7648,7 +7495,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7778,10 +7624,8 @@ func (m *awsRestjson1_deserializeOpListCollaborations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7821,7 +7665,6 @@ func (m *awsRestjson1_deserializeOpListCollaborations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7948,10 +7791,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredAudie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7991,7 +7832,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredAudie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8121,10 +7961,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredTable
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8164,7 +8002,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredTable
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8294,10 +8131,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredTable
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8337,7 +8172,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredTable
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8464,10 +8298,8 @@ func (m *awsRestjson1_deserializeOpListIdMappingTables
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8507,7 +8339,6 @@ func (m *awsRestjson1_deserializeOpListIdMappingTables
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8637,10 +8468,8 @@ func (m *awsRestjson1_deserializeOpListIdNamespaceAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8680,7 +8509,6 @@ func (m *awsRestjson1_deserializeOpListIdNamespaceAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8810,10 +8638,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8853,7 +8679,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8983,10 +8808,8 @@ func (m *awsRestjson1_deserializeOpListMemberships) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9026,7 +8849,6 @@ func (m *awsRestjson1_deserializeOpListMemberships) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9153,10 +8975,8 @@ func (m *awsRestjson1_deserializeOpListPrivacyBudgets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9196,7 +9016,6 @@ func (m *awsRestjson1_deserializeOpListPrivacyBudgets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9326,10 +9145,8 @@ func (m *awsRestjson1_deserializeOpListPrivacyBudgetTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9369,7 +9186,6 @@ func (m *awsRestjson1_deserializeOpListPrivacyBudgetTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9499,10 +9315,8 @@ func (m *awsRestjson1_deserializeOpListProtectedJobs) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9542,7 +9356,6 @@ func (m *awsRestjson1_deserializeOpListProtectedJobs) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9672,10 +9485,8 @@ func (m *awsRestjson1_deserializeOpListProtectedQuerie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9715,7 +9526,6 @@ func (m *awsRestjson1_deserializeOpListProtectedQuerie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9845,10 +9655,8 @@ func (m *awsRestjson1_deserializeOpListSchemas) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9888,7 +9696,6 @@ func (m *awsRestjson1_deserializeOpListSchemas) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10018,10 +9825,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10061,7 +9866,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10173,10 +9977,8 @@ func (m *awsRestjson1_deserializeOpPopulateIdMappingTa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10216,7 +10018,6 @@ func (m *awsRestjson1_deserializeOpPopulateIdMappingTa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10347,10 +10148,8 @@ func (m *awsRestjson1_deserializeOpPreviewPrivacyImpac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10390,7 +10189,6 @@ func (m *awsRestjson1_deserializeOpPreviewPrivacyImpac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10511,10 +10309,8 @@ func (m *awsRestjson1_deserializeOpStartProtectedJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10554,7 +10350,6 @@ func (m *awsRestjson1_deserializeOpStartProtectedJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10678,10 +10473,8 @@ func (m *awsRestjson1_deserializeOpStartProtectedQuery
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10721,7 +10514,6 @@ func (m *awsRestjson1_deserializeOpStartProtectedQuery
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10845,10 +10637,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10860,7 +10650,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10936,10 +10725,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10951,7 +10738,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11027,10 +10813,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysisTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11070,7 +10854,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysisTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11191,10 +10974,8 @@ func (m *awsRestjson1_deserializeOpUpdateCollaboration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11234,7 +11015,6 @@ func (m *awsRestjson1_deserializeOpUpdateCollaboration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11352,10 +11132,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11395,7 +11173,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11516,10 +11293,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11559,7 +11334,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11683,10 +11457,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11726,7 +11498,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11850,10 +11621,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11893,7 +11662,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12017,10 +11785,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12060,7 +11826,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12184,10 +11949,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdMappingTabl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12227,7 +11990,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdMappingTabl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12348,10 +12110,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdNamespaceAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12391,7 +12151,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdNamespaceAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12512,10 +12271,8 @@ func (m *awsRestjson1_deserializeOpUpdateMembership) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12555,7 +12312,6 @@ func (m *awsRestjson1_deserializeOpUpdateMembership) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12679,10 +12435,8 @@ func (m *awsRestjson1_deserializeOpUpdatePrivacyBudget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12722,7 +12476,6 @@ func (m *awsRestjson1_deserializeOpUpdatePrivacyBudget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12846,10 +12599,8 @@ func (m *awsRestjson1_deserializeOpUpdateProtectedJob)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12889,7 +12640,6 @@ func (m *awsRestjson1_deserializeOpUpdateProtectedJob)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13013,10 +12763,8 @@ func (m *awsRestjson1_deserializeOpUpdateProtectedQuer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13056,7 +12804,6 @@ func (m *awsRestjson1_deserializeOpUpdateProtectedQuer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
