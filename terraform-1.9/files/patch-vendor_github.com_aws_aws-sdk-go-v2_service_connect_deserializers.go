--- vendor/github.com/aws/aws-sdk-go-v2/service/connect/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connect/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpActivateEvaluationF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpActivateEvaluationF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -236,10 +232,8 @@ func (m *awsRestjson1_deserializeOpAssociateAnalyticsD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -279,7 +273,6 @@ func (m *awsRestjson1_deserializeOpAssociateAnalyticsD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -431,10 +424,8 @@ func (m *awsRestjson1_deserializeOpAssociateApprovedOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -452,7 +443,6 @@ func (m *awsRestjson1_deserializeOpAssociateApprovedOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -543,10 +533,8 @@ func (m *awsRestjson1_deserializeOpAssociateBot) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -564,7 +552,6 @@ func (m *awsRestjson1_deserializeOpAssociateBot) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -655,10 +642,8 @@ func (m *awsRestjson1_deserializeOpAssociateDefaultVoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -670,7 +655,6 @@ func (m *awsRestjson1_deserializeOpAssociateDefaultVoc
 	output := &AssociateDefaultVocabularyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -755,10 +739,8 @@ func (m *awsRestjson1_deserializeOpAssociateFlow) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -770,7 +752,6 @@ func (m *awsRestjson1_deserializeOpAssociateFlow) Hand
 	output := &AssociateFlowOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -858,10 +839,8 @@ func (m *awsRestjson1_deserializeOpAssociateInstanceSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -901,7 +880,6 @@ func (m *awsRestjson1_deserializeOpAssociateInstanceSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1029,10 +1007,8 @@ func (m *awsRestjson1_deserializeOpAssociateLambdaFunc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1050,7 +1026,6 @@ func (m *awsRestjson1_deserializeOpAssociateLambdaFunc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1141,10 +1116,8 @@ func (m *awsRestjson1_deserializeOpAssociateLexBot) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1162,7 +1135,6 @@ func (m *awsRestjson1_deserializeOpAssociateLexBot) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1253,10 +1225,8 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1274,7 +1244,6 @@ func (m *awsRestjson1_deserializeOpAssociatePhoneNumbe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1359,10 +1328,8 @@ func (m *awsRestjson1_deserializeOpAssociateQueueQuick
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1380,7 +1347,6 @@ func (m *awsRestjson1_deserializeOpAssociateQueueQuick
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1468,10 +1434,8 @@ func (m *awsRestjson1_deserializeOpAssociateRoutingPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1489,7 +1453,6 @@ func (m *awsRestjson1_deserializeOpAssociateRoutingPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1574,10 +1537,8 @@ func (m *awsRestjson1_deserializeOpAssociateSecurityKe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1617,7 +1578,6 @@ func (m *awsRestjson1_deserializeOpAssociateSecurityKe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1748,10 +1708,8 @@ func (m *awsRestjson1_deserializeOpAssociateTrafficDis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1763,7 +1721,6 @@ func (m *awsRestjson1_deserializeOpAssociateTrafficDis
 	output := &AssociateTrafficDistributionGroupUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1851,10 +1808,8 @@ func (m *awsRestjson1_deserializeOpAssociateUserProfic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1872,7 +1827,6 @@ func (m *awsRestjson1_deserializeOpAssociateUserProfic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1957,10 +1911,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateAnaly
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2000,7 +1952,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateAnaly
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2126,10 +2077,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2169,7 +2118,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2295,10 +2243,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAttachedFil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2338,7 +2284,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAttachedFil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2464,10 +2409,8 @@ func (m *awsRestjson1_deserializeOpBatchGetFlowAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2507,7 +2450,6 @@ func (m *awsRestjson1_deserializeOpBatchGetFlowAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2631,10 +2573,8 @@ func (m *awsRestjson1_deserializeOpBatchPutContact) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2674,7 +2614,6 @@ func (m *awsRestjson1_deserializeOpBatchPutContact) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2803,10 +2742,8 @@ func (m *awsRestjson1_deserializeOpClaimPhoneNumber) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2846,7 +2783,6 @@ func (m *awsRestjson1_deserializeOpClaimPhoneNumber) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2983,10 +2919,8 @@ func (m *awsRestjson1_deserializeOpCompleteAttachedFil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2998,7 +2932,6 @@ func (m *awsRestjson1_deserializeOpCompleteAttachedFil
 	output := &CompleteAttachedFileUploadOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3083,10 +3016,8 @@ func (m *awsRestjson1_deserializeOpCreateAgentStatus) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3126,7 +3057,6 @@ func (m *awsRestjson1_deserializeOpCreateAgentStatus) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3266,10 +3196,8 @@ func (m *awsRestjson1_deserializeOpCreateContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3309,7 +3237,6 @@ func (m *awsRestjson1_deserializeOpCreateContact) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3455,10 +3382,8 @@ func (m *awsRestjson1_deserializeOpCreateContactFlow) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3498,7 +3423,6 @@ func (m *awsRestjson1_deserializeOpCreateContactFlow) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3650,10 +3574,8 @@ func (m *awsRestjson1_deserializeOpCreateContactFlowMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3693,7 +3615,6 @@ func (m *awsRestjson1_deserializeOpCreateContactFlowMo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3842,10 +3763,8 @@ func (m *awsRestjson1_deserializeOpCreateContactFlowVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3885,7 +3804,6 @@ func (m *awsRestjson1_deserializeOpCreateContactFlowVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4029,10 +3947,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailAddress)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4072,7 +3988,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailAddress)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4221,10 +4136,8 @@ func (m *awsRestjson1_deserializeOpCreateEvaluationFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4264,7 +4177,6 @@ func (m *awsRestjson1_deserializeOpCreateEvaluationFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4401,10 +4313,8 @@ func (m *awsRestjson1_deserializeOpCreateHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4444,7 +4354,6 @@ func (m *awsRestjson1_deserializeOpCreateHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4584,10 +4493,8 @@ func (m *awsRestjson1_deserializeOpCreateHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4627,7 +4534,6 @@ func (m *awsRestjson1_deserializeOpCreateHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4758,10 +4664,8 @@ func (m *awsRestjson1_deserializeOpCreateInstance) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4801,7 +4705,6 @@ func (m *awsRestjson1_deserializeOpCreateInstance) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4935,10 +4838,8 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4978,7 +4879,6 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5112,10 +5012,8 @@ func (m *awsRestjson1_deserializeOpCreateParticipant) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5155,7 +5053,6 @@ func (m *awsRestjson1_deserializeOpCreateParticipant) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5285,10 +5182,8 @@ func (m *awsRestjson1_deserializeOpCreatePersistentCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5328,7 +5223,6 @@ func (m *awsRestjson1_deserializeOpCreatePersistentCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5456,10 +5350,8 @@ func (m *awsRestjson1_deserializeOpCreatePredefinedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5477,7 +5369,6 @@ func (m *awsRestjson1_deserializeOpCreatePredefinedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5568,10 +5459,8 @@ func (m *awsRestjson1_deserializeOpCreatePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5611,7 +5500,6 @@ func (m *awsRestjson1_deserializeOpCreatePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5748,10 +5636,8 @@ func (m *awsRestjson1_deserializeOpCreatePushNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5791,7 +5677,6 @@ func (m *awsRestjson1_deserializeOpCreatePushNotificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5919,10 +5804,8 @@ func (m *awsRestjson1_deserializeOpCreateQueue) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5962,7 +5845,6 @@ func (m *awsRestjson1_deserializeOpCreateQueue) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6102,10 +5984,8 @@ func (m *awsRestjson1_deserializeOpCreateQuickConnect)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6145,7 +6025,6 @@ func (m *awsRestjson1_deserializeOpCreateQuickConnect)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6285,10 +6164,8 @@ func (m *awsRestjson1_deserializeOpCreateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6328,7 +6205,6 @@ func (m *awsRestjson1_deserializeOpCreateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6468,10 +6344,8 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6511,7 +6385,6 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6651,10 +6524,8 @@ func (m *awsRestjson1_deserializeOpCreateSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6694,7 +6565,6 @@ func (m *awsRestjson1_deserializeOpCreateSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6834,10 +6704,8 @@ func (m *awsRestjson1_deserializeOpCreateTaskTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6877,7 +6745,6 @@ func (m *awsRestjson1_deserializeOpCreateTaskTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7014,10 +6881,8 @@ func (m *awsRestjson1_deserializeOpCreateTrafficDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7057,7 +6922,6 @@ func (m *awsRestjson1_deserializeOpCreateTrafficDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7200,10 +7064,8 @@ func (m *awsRestjson1_deserializeOpCreateUseCase) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7243,7 +7105,6 @@ func (m *awsRestjson1_deserializeOpCreateUseCase) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7377,10 +7238,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7420,7 +7279,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7560,10 +7418,8 @@ func (m *awsRestjson1_deserializeOpCreateUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7603,7 +7459,6 @@ func (m *awsRestjson1_deserializeOpCreateUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7743,10 +7598,8 @@ func (m *awsRestjson1_deserializeOpCreateView) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7786,7 +7639,6 @@ func (m *awsRestjson1_deserializeOpCreateView) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7919,10 +7771,8 @@ func (m *awsRestjson1_deserializeOpCreateViewVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7962,7 +7812,6 @@ func (m *awsRestjson1_deserializeOpCreateViewVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8092,10 +7941,8 @@ func (m *awsRestjson1_deserializeOpCreateVocabulary) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8135,7 +7982,6 @@ func (m *awsRestjson1_deserializeOpCreateVocabulary) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8284,10 +8130,8 @@ func (m *awsRestjson1_deserializeOpDeactivateEvaluatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8327,7 +8171,6 @@ func (m *awsRestjson1_deserializeOpDeactivateEvaluatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8474,10 +8317,8 @@ func (m *awsRestjson1_deserializeOpDeleteAttachedFile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8489,7 +8330,6 @@ func (m *awsRestjson1_deserializeOpDeleteAttachedFile)
 	output := &DeleteAttachedFileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8574,10 +8414,8 @@ func (m *awsRestjson1_deserializeOpDeleteContactEvalua
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8595,7 +8433,6 @@ func (m *awsRestjson1_deserializeOpDeleteContactEvalua
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8680,10 +8517,8 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlow) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8695,7 +8530,6 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlow) 
 	output := &DeleteContactFlowOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8783,10 +8617,8 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlowMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8798,7 +8630,6 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlowMo
 	output := &DeleteContactFlowModuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8886,10 +8717,8 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlowVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8901,7 +8730,6 @@ func (m *awsRestjson1_deserializeOpDeleteContactFlowVe
 	output := &DeleteContactFlowVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8989,10 +8817,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailAddress)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9004,7 +8830,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailAddress)
 	output := &DeleteEmailAddressOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9095,10 +8920,8 @@ func (m *awsRestjson1_deserializeOpDeleteEvaluationFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9116,7 +8939,6 @@ func (m *awsRestjson1_deserializeOpDeleteEvaluationFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9201,10 +9023,8 @@ func (m *awsRestjson1_deserializeOpDeleteHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9222,7 +9042,6 @@ func (m *awsRestjson1_deserializeOpDeleteHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9307,10 +9126,8 @@ func (m *awsRestjson1_deserializeOpDeleteHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9328,7 +9145,6 @@ func (m *awsRestjson1_deserializeOpDeleteHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9413,10 +9229,8 @@ func (m *awsRestjson1_deserializeOpDeleteInstance) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9434,7 +9248,6 @@ func (m *awsRestjson1_deserializeOpDeleteInstance) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9513,10 +9326,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9534,7 +9345,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9616,10 +9426,8 @@ func (m *awsRestjson1_deserializeOpDeletePredefinedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9637,7 +9445,6 @@ func (m *awsRestjson1_deserializeOpDeletePredefinedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9725,10 +9532,8 @@ func (m *awsRestjson1_deserializeOpDeletePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9746,7 +9551,6 @@ func (m *awsRestjson1_deserializeOpDeletePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9831,10 +9635,8 @@ func (m *awsRestjson1_deserializeOpDeletePushNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9846,7 +9648,6 @@ func (m *awsRestjson1_deserializeOpDeletePushNotificat
 	output := &DeletePushNotificationRegistrationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9931,10 +9732,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueue) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9952,7 +9751,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueue) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10040,10 +9838,8 @@ func (m *awsRestjson1_deserializeOpDeleteQuickConnect)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10061,7 +9857,6 @@ func (m *awsRestjson1_deserializeOpDeleteQuickConnect)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10146,10 +9941,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10167,7 +9960,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10255,10 +10047,8 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10276,7 +10066,6 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10361,10 +10150,8 @@ func (m *awsRestjson1_deserializeOpDeleteSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10382,7 +10169,6 @@ func (m *awsRestjson1_deserializeOpDeleteSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10473,10 +10259,8 @@ func (m *awsRestjson1_deserializeOpDeleteTaskTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10488,7 +10272,6 @@ func (m *awsRestjson1_deserializeOpDeleteTaskTemplate)
 	output := &DeleteTaskTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10573,10 +10356,8 @@ func (m *awsRestjson1_deserializeOpDeleteTrafficDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10588,7 +10369,6 @@ func (m *awsRestjson1_deserializeOpDeleteTrafficDistri
 	output := &DeleteTrafficDistributionGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10673,10 +10453,8 @@ func (m *awsRestjson1_deserializeOpDeleteUseCase) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10694,7 +10472,6 @@ func (m *awsRestjson1_deserializeOpDeleteUseCase) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10776,10 +10553,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10797,7 +10572,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10882,10 +10656,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10903,7 +10675,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10991,10 +10762,8 @@ func (m *awsRestjson1_deserializeOpDeleteView) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11006,7 +10775,6 @@ func (m *awsRestjson1_deserializeOpDeleteView) HandleD
 	output := &DeleteViewOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11097,10 +10865,8 @@ func (m *awsRestjson1_deserializeOpDeleteViewVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11112,7 +10878,6 @@ func (m *awsRestjson1_deserializeOpDeleteViewVersion) 
 	output := &DeleteViewVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11203,10 +10968,8 @@ func (m *awsRestjson1_deserializeOpDeleteVocabulary) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11246,7 +11009,6 @@ func (m *awsRestjson1_deserializeOpDeleteVocabulary) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11392,10 +11154,8 @@ func (m *awsRestjson1_deserializeOpDescribeAgentStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11435,7 +11195,6 @@ func (m *awsRestjson1_deserializeOpDescribeAgentStatus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11556,10 +11315,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuthenticat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11599,7 +11356,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuthenticat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11720,10 +11476,8 @@ func (m *awsRestjson1_deserializeOpDescribeContact) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11763,7 +11517,6 @@ func (m *awsRestjson1_deserializeOpDescribeContact) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11884,10 +11637,8 @@ func (m *awsRestjson1_deserializeOpDescribeContactEval
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11927,7 +11678,6 @@ func (m *awsRestjson1_deserializeOpDescribeContactEval
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12050,10 +11800,8 @@ func (m *awsRestjson1_deserializeOpDescribeContactFlow
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12093,7 +11841,6 @@ func (m *awsRestjson1_deserializeOpDescribeContactFlow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12217,10 +11964,8 @@ func (m *awsRestjson1_deserializeOpDescribeContactFlow
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12260,7 +12005,6 @@ func (m *awsRestjson1_deserializeOpDescribeContactFlow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12384,10 +12128,8 @@ func (m *awsRestjson1_deserializeOpDescribeEmailAddres
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12427,7 +12169,6 @@ func (m *awsRestjson1_deserializeOpDescribeEmailAddres
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12614,10 +12355,8 @@ func (m *awsRestjson1_deserializeOpDescribeEvaluationF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12657,7 +12396,6 @@ func (m *awsRestjson1_deserializeOpDescribeEvaluationF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12775,10 +12513,8 @@ func (m *awsRestjson1_deserializeOpDescribeHoursOfOper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12818,7 +12554,6 @@ func (m *awsRestjson1_deserializeOpDescribeHoursOfOper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12939,10 +12674,8 @@ func (m *awsRestjson1_deserializeOpDescribeHoursOfOper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12982,7 +12715,6 @@ func (m *awsRestjson1_deserializeOpDescribeHoursOfOper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13103,10 +12835,8 @@ func (m *awsRestjson1_deserializeOpDescribeInstance) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13146,7 +12876,6 @@ func (m *awsRestjson1_deserializeOpDescribeInstance) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13266,10 +12995,8 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13309,7 +13036,6 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13430,10 +13156,8 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13473,7 +13197,6 @@ func (m *awsRestjson1_deserializeOpDescribeInstanceSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13594,10 +13317,8 @@ func (m *awsRestjson1_deserializeOpDescribePhoneNumber
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13637,7 +13358,6 @@ func (m *awsRestjson1_deserializeOpDescribePhoneNumber
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13758,10 +13478,8 @@ func (m *awsRestjson1_deserializeOpDescribePredefinedA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13801,7 +13519,6 @@ func (m *awsRestjson1_deserializeOpDescribePredefinedA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13922,10 +13639,8 @@ func (m *awsRestjson1_deserializeOpDescribePrompt) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13965,7 +13680,6 @@ func (m *awsRestjson1_deserializeOpDescribePrompt) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14086,10 +13800,8 @@ func (m *awsRestjson1_deserializeOpDescribeQueue) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14129,7 +13841,6 @@ func (m *awsRestjson1_deserializeOpDescribeQueue) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14250,10 +13961,8 @@ func (m *awsRestjson1_deserializeOpDescribeQuickConnec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14293,7 +14002,6 @@ func (m *awsRestjson1_deserializeOpDescribeQuickConnec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14414,10 +14122,8 @@ func (m *awsRestjson1_deserializeOpDescribeRoutingProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14457,7 +14163,6 @@ func (m *awsRestjson1_deserializeOpDescribeRoutingProf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14578,10 +14283,8 @@ func (m *awsRestjson1_deserializeOpDescribeRule) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14621,7 +14324,6 @@ func (m *awsRestjson1_deserializeOpDescribeRule) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14742,10 +14444,8 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14785,7 +14485,6 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14906,10 +14605,8 @@ func (m *awsRestjson1_deserializeOpDescribeTrafficDist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14949,7 +14646,6 @@ func (m *awsRestjson1_deserializeOpDescribeTrafficDist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15070,10 +14766,8 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15113,7 +14807,6 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15234,10 +14927,8 @@ func (m *awsRestjson1_deserializeOpDescribeUserHierarc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15277,7 +14968,6 @@ func (m *awsRestjson1_deserializeOpDescribeUserHierarc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15398,10 +15088,8 @@ func (m *awsRestjson1_deserializeOpDescribeUserHierarc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15441,7 +15129,6 @@ func (m *awsRestjson1_deserializeOpDescribeUserHierarc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15562,10 +15249,8 @@ func (m *awsRestjson1_deserializeOpDescribeView) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15605,7 +15290,6 @@ func (m *awsRestjson1_deserializeOpDescribeView) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15729,10 +15413,8 @@ func (m *awsRestjson1_deserializeOpDescribeVocabulary)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15772,7 +15454,6 @@ func (m *awsRestjson1_deserializeOpDescribeVocabulary)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15893,10 +15574,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAnalyti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15914,7 +15593,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAnalyti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15999,10 +15677,8 @@ func (m *awsRestjson1_deserializeOpDisassociateApprove
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16020,7 +15696,6 @@ func (m *awsRestjson1_deserializeOpDisassociateApprove
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16105,10 +15780,8 @@ func (m *awsRestjson1_deserializeOpDisassociateBot) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16126,7 +15799,6 @@ func (m *awsRestjson1_deserializeOpDisassociateBot) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16208,10 +15880,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFlow) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16223,7 +15893,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFlow) H
 	output := &DisassociateFlowOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16311,10 +15980,8 @@ func (m *awsRestjson1_deserializeOpDisassociateInstanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16332,7 +15999,6 @@ func (m *awsRestjson1_deserializeOpDisassociateInstanc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16417,10 +16083,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLambdaF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16438,7 +16102,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLambdaF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16523,10 +16186,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLexBot)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16544,7 +16205,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLexBot)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16629,10 +16289,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16650,7 +16308,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePhoneNu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16735,10 +16392,8 @@ func (m *awsRestjson1_deserializeOpDisassociateQueueQu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16756,7 +16411,6 @@ func (m *awsRestjson1_deserializeOpDisassociateQueueQu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16841,10 +16495,8 @@ func (m *awsRestjson1_deserializeOpDisassociateRouting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16862,7 +16514,6 @@ func (m *awsRestjson1_deserializeOpDisassociateRouting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16947,10 +16598,8 @@ func (m *awsRestjson1_deserializeOpDisassociateSecurit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16968,7 +16617,6 @@ func (m *awsRestjson1_deserializeOpDisassociateSecurit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17053,10 +16701,8 @@ func (m *awsRestjson1_deserializeOpDisassociateTraffic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17068,7 +16714,6 @@ func (m *awsRestjson1_deserializeOpDisassociateTraffic
 	output := &DisassociateTrafficDistributionGroupUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17156,10 +16801,8 @@ func (m *awsRestjson1_deserializeOpDisassociateUserPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17177,7 +16820,6 @@ func (m *awsRestjson1_deserializeOpDisassociateUserPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17262,10 +16904,8 @@ func (m *awsRestjson1_deserializeOpDismissUserContact)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17277,7 +16917,6 @@ func (m *awsRestjson1_deserializeOpDismissUserContact)
 	output := &DismissUserContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17365,10 +17004,8 @@ func (m *awsRestjson1_deserializeOpGetAttachedFile) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17408,7 +17045,6 @@ func (m *awsRestjson1_deserializeOpGetAttachedFile) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17615,10 +17251,8 @@ func (m *awsRestjson1_deserializeOpGetContactAttribute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17658,7 +17292,6 @@ func (m *awsRestjson1_deserializeOpGetContactAttribute
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17773,10 +17406,8 @@ func (m *awsRestjson1_deserializeOpGetCurrentMetricDat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17816,7 +17447,6 @@ func (m *awsRestjson1_deserializeOpGetCurrentMetricDat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17975,10 +17605,8 @@ func (m *awsRestjson1_deserializeOpGetCurrentUserData)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18018,7 +17646,6 @@ func (m *awsRestjson1_deserializeOpGetCurrentUserData)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18161,10 +17788,8 @@ func (m *awsRestjson1_deserializeOpGetEffectiveHoursOf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18204,7 +17829,6 @@ func (m *awsRestjson1_deserializeOpGetEffectiveHoursOf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18334,10 +17958,8 @@ func (m *awsRestjson1_deserializeOpGetFederationToken)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18377,7 +17999,6 @@ func (m *awsRestjson1_deserializeOpGetFederationToken)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18528,10 +18149,8 @@ func (m *awsRestjson1_deserializeOpGetFlowAssociation)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18571,7 +18190,6 @@ func (m *awsRestjson1_deserializeOpGetFlowAssociation)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18717,10 +18335,8 @@ func (m *awsRestjson1_deserializeOpGetMetricData) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18760,7 +18376,6 @@ func (m *awsRestjson1_deserializeOpGetMetricData) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18890,10 +18505,8 @@ func (m *awsRestjson1_deserializeOpGetMetricDataV2) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18933,7 +18546,6 @@ func (m *awsRestjson1_deserializeOpGetMetricDataV2) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19063,10 +18675,8 @@ func (m *awsRestjson1_deserializeOpGetPromptFile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19106,7 +18716,6 @@ func (m *awsRestjson1_deserializeOpGetPromptFile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19256,10 +18865,8 @@ func (m *awsRestjson1_deserializeOpGetTaskTemplate) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19299,7 +18906,6 @@ func (m *awsRestjson1_deserializeOpGetTaskTemplate) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19539,10 +19145,8 @@ func (m *awsRestjson1_deserializeOpGetTrafficDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19582,7 +19186,6 @@ func (m *awsRestjson1_deserializeOpGetTrafficDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19731,10 +19334,8 @@ func (m *awsRestjson1_deserializeOpImportPhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19774,7 +19375,6 @@ func (m *awsRestjson1_deserializeOpImportPhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19911,10 +19511,8 @@ func (m *awsRestjson1_deserializeOpListAgentStatuses) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19954,7 +19552,6 @@ func (m *awsRestjson1_deserializeOpListAgentStatuses) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20084,10 +19681,8 @@ func (m *awsRestjson1_deserializeOpListAnalyticsDataAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20127,7 +19722,6 @@ func (m *awsRestjson1_deserializeOpListAnalyticsDataAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20257,10 +19851,8 @@ func (m *awsRestjson1_deserializeOpListAnalyticsDataLa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20300,7 +19892,6 @@ func (m *awsRestjson1_deserializeOpListAnalyticsDataLa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20430,10 +20021,8 @@ func (m *awsRestjson1_deserializeOpListApprovedOrigins
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20473,7 +20062,6 @@ func (m *awsRestjson1_deserializeOpListApprovedOrigins
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20603,10 +20191,8 @@ func (m *awsRestjson1_deserializeOpListAssociatedConta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20646,7 +20232,6 @@ func (m *awsRestjson1_deserializeOpListAssociatedConta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20776,10 +20361,8 @@ func (m *awsRestjson1_deserializeOpListAuthenticationP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20819,7 +20402,6 @@ func (m *awsRestjson1_deserializeOpListAuthenticationP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20949,10 +20531,8 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20992,7 +20572,6 @@ func (m *awsRestjson1_deserializeOpListBots) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21119,10 +20698,8 @@ func (m *awsRestjson1_deserializeOpListContactEvaluati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21162,7 +20739,6 @@ func (m *awsRestjson1_deserializeOpListContactEvaluati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21289,10 +20865,8 @@ func (m *awsRestjson1_deserializeOpListContactFlowModu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21332,7 +20906,6 @@ func (m *awsRestjson1_deserializeOpListContactFlowModu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21465,10 +21038,8 @@ func (m *awsRestjson1_deserializeOpListContactFlows) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21508,7 +21079,6 @@ func (m *awsRestjson1_deserializeOpListContactFlows) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21638,10 +21208,8 @@ func (m *awsRestjson1_deserializeOpListContactFlowVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21681,7 +21249,6 @@ func (m *awsRestjson1_deserializeOpListContactFlowVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21814,10 +21381,8 @@ func (m *awsRestjson1_deserializeOpListContactReferenc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21857,7 +21422,6 @@ func (m *awsRestjson1_deserializeOpListContactReferenc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21987,10 +21551,8 @@ func (m *awsRestjson1_deserializeOpListDefaultVocabula
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22030,7 +21592,6 @@ func (m *awsRestjson1_deserializeOpListDefaultVocabula
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22157,10 +21718,8 @@ func (m *awsRestjson1_deserializeOpListEvaluationForms
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22200,7 +21759,6 @@ func (m *awsRestjson1_deserializeOpListEvaluationForms
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22327,10 +21885,8 @@ func (m *awsRestjson1_deserializeOpListEvaluationFormV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22370,7 +21926,6 @@ func (m *awsRestjson1_deserializeOpListEvaluationFormV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22497,10 +22052,8 @@ func (m *awsRestjson1_deserializeOpListFlowAssociation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22540,7 +22093,6 @@ func (m *awsRestjson1_deserializeOpListFlowAssociation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22673,10 +22225,8 @@ func (m *awsRestjson1_deserializeOpListHoursOfOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22716,7 +22266,6 @@ func (m *awsRestjson1_deserializeOpListHoursOfOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22871,10 +22420,8 @@ func (m *awsRestjson1_deserializeOpListHoursOfOperatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22914,7 +22461,6 @@ func (m *awsRestjson1_deserializeOpListHoursOfOperatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23044,10 +22590,8 @@ func (m *awsRestjson1_deserializeOpListInstanceAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23087,7 +22631,6 @@ func (m *awsRestjson1_deserializeOpListInstanceAttribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23217,10 +22760,8 @@ func (m *awsRestjson1_deserializeOpListInstances) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23260,7 +22801,6 @@ func (m *awsRestjson1_deserializeOpListInstances) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23381,10 +22921,8 @@ func (m *awsRestjson1_deserializeOpListInstanceStorage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23424,7 +22962,6 @@ func (m *awsRestjson1_deserializeOpListInstanceStorage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23554,10 +23091,8 @@ func (m *awsRestjson1_deserializeOpListIntegrationAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23597,7 +23132,6 @@ func (m *awsRestjson1_deserializeOpListIntegrationAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23724,10 +23258,8 @@ func (m *awsRestjson1_deserializeOpListLambdaFunctions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23767,7 +23299,6 @@ func (m *awsRestjson1_deserializeOpListLambdaFunctions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23897,10 +23428,8 @@ func (m *awsRestjson1_deserializeOpListLexBots) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23940,7 +23469,6 @@ func (m *awsRestjson1_deserializeOpListLexBots) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24070,10 +23598,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24113,7 +23639,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24243,10 +23768,8 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbersV2)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24286,7 +23809,6 @@ func (m *awsRestjson1_deserializeOpListPhoneNumbersV2)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24416,10 +23938,8 @@ func (m *awsRestjson1_deserializeOpListPredefinedAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24459,7 +23979,6 @@ func (m *awsRestjson1_deserializeOpListPredefinedAttri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24589,10 +24108,8 @@ func (m *awsRestjson1_deserializeOpListPrompts) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24632,7 +24149,6 @@ func (m *awsRestjson1_deserializeOpListPrompts) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24762,10 +24278,8 @@ func (m *awsRestjson1_deserializeOpListQueueQuickConne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24805,7 +24319,6 @@ func (m *awsRestjson1_deserializeOpListQueueQuickConne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24960,10 +24473,8 @@ func (m *awsRestjson1_deserializeOpListQueues) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25003,7 +24514,6 @@ func (m *awsRestjson1_deserializeOpListQueues) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25133,10 +24643,8 @@ func (m *awsRestjson1_deserializeOpListQuickConnects) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25176,7 +24684,6 @@ func (m *awsRestjson1_deserializeOpListQuickConnects) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25306,10 +24813,8 @@ func (m *awsRestjson1_deserializeOpListRealtimeContact
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25349,7 +24854,6 @@ func (m *awsRestjson1_deserializeOpListRealtimeContact
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25500,10 +25004,8 @@ func (m *awsRestjson1_deserializeOpListRoutingProfileQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25543,7 +25045,6 @@ func (m *awsRestjson1_deserializeOpListRoutingProfileQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25698,10 +25199,8 @@ func (m *awsRestjson1_deserializeOpListRoutingProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25741,7 +25240,6 @@ func (m *awsRestjson1_deserializeOpListRoutingProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25871,10 +25369,8 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25914,7 +25410,6 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26044,10 +25539,8 @@ func (m *awsRestjson1_deserializeOpListSecurityKeys) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26087,7 +25580,6 @@ func (m *awsRestjson1_deserializeOpListSecurityKeys) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26217,10 +25709,8 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26260,7 +25750,6 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26415,10 +25904,8 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26458,7 +25945,6 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26613,10 +26099,8 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26656,7 +26140,6 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26786,10 +26269,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26829,7 +26310,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26950,10 +26430,8 @@ func (m *awsRestjson1_deserializeOpListTaskTemplates) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26993,7 +26471,6 @@ func (m *awsRestjson1_deserializeOpListTaskTemplates) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27123,10 +26600,8 @@ func (m *awsRestjson1_deserializeOpListTrafficDistribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27166,7 +26641,6 @@ func (m *awsRestjson1_deserializeOpListTrafficDistribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27293,10 +26767,8 @@ func (m *awsRestjson1_deserializeOpListTrafficDistribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27336,7 +26808,6 @@ func (m *awsRestjson1_deserializeOpListTrafficDistribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27466,10 +26937,8 @@ func (m *awsRestjson1_deserializeOpListUseCases) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27509,7 +26978,6 @@ func (m *awsRestjson1_deserializeOpListUseCases) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27636,10 +27104,8 @@ func (m *awsRestjson1_deserializeOpListUserHierarchyGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27679,7 +27145,6 @@ func (m *awsRestjson1_deserializeOpListUserHierarchyGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27809,10 +27274,8 @@ func (m *awsRestjson1_deserializeOpListUserProficienci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27852,7 +27315,6 @@ func (m *awsRestjson1_deserializeOpListUserProficienci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28007,10 +27469,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28050,7 +27510,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28180,10 +27639,8 @@ func (m *awsRestjson1_deserializeOpListViews) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28223,7 +27680,6 @@ func (m *awsRestjson1_deserializeOpListViews) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28356,10 +27812,8 @@ func (m *awsRestjson1_deserializeOpListViewVersions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28399,7 +27853,6 @@ func (m *awsRestjson1_deserializeOpListViewVersions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28532,10 +27985,8 @@ func (m *awsRestjson1_deserializeOpMonitorContact) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28575,7 +28026,6 @@ func (m *awsRestjson1_deserializeOpMonitorContact) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28715,10 +28165,8 @@ func (m *awsRestjson1_deserializeOpPauseContact) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28730,7 +28178,6 @@ func (m *awsRestjson1_deserializeOpPauseContact) Handl
 	output := &PauseContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28824,10 +28271,8 @@ func (m *awsRestjson1_deserializeOpPutUserStatus) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28839,7 +28284,6 @@ func (m *awsRestjson1_deserializeOpPutUserStatus) Hand
 	output := &PutUserStatusOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28927,10 +28371,8 @@ func (m *awsRestjson1_deserializeOpReleasePhoneNumber)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28948,7 +28390,6 @@ func (m *awsRestjson1_deserializeOpReleasePhoneNumber)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29039,10 +28480,8 @@ func (m *awsRestjson1_deserializeOpReplicateInstance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29082,7 +28521,6 @@ func (m *awsRestjson1_deserializeOpReplicateInstance) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29225,10 +28663,8 @@ func (m *awsRestjson1_deserializeOpResumeContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29240,7 +28676,6 @@ func (m *awsRestjson1_deserializeOpResumeContact) Hand
 	output := &ResumeContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29331,10 +28766,8 @@ func (m *awsRestjson1_deserializeOpResumeContactRecord
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29346,7 +28779,6 @@ func (m *awsRestjson1_deserializeOpResumeContactRecord
 	output := &ResumeContactRecordingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29425,10 +28857,8 @@ func (m *awsRestjson1_deserializeOpSearchAgentStatuses
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29468,7 +28898,6 @@ func (m *awsRestjson1_deserializeOpSearchAgentStatuses
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29611,10 +29040,8 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29654,7 +29081,6 @@ func (m *awsRestjson1_deserializeOpSearchAvailablePhon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29781,10 +29207,8 @@ func (m *awsRestjson1_deserializeOpSearchContactFlowMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29824,7 +29248,6 @@ func (m *awsRestjson1_deserializeOpSearchContactFlowMo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29967,10 +29390,8 @@ func (m *awsRestjson1_deserializeOpSearchContactFlows)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30010,7 +29431,6 @@ func (m *awsRestjson1_deserializeOpSearchContactFlows)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30153,10 +29573,8 @@ func (m *awsRestjson1_deserializeOpSearchContacts) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30196,7 +29614,6 @@ func (m *awsRestjson1_deserializeOpSearchContacts) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30339,10 +29756,8 @@ func (m *awsRestjson1_deserializeOpSearchEmailAddresse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30382,7 +29797,6 @@ func (m *awsRestjson1_deserializeOpSearchEmailAddresse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30528,10 +29942,8 @@ func (m *awsRestjson1_deserializeOpSearchHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30571,7 +29983,6 @@ func (m *awsRestjson1_deserializeOpSearchHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30714,10 +30125,8 @@ func (m *awsRestjson1_deserializeOpSearchHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30757,7 +30166,6 @@ func (m *awsRestjson1_deserializeOpSearchHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30900,10 +30308,8 @@ func (m *awsRestjson1_deserializeOpSearchPredefinedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30943,7 +30349,6 @@ func (m *awsRestjson1_deserializeOpSearchPredefinedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31086,10 +30491,8 @@ func (m *awsRestjson1_deserializeOpSearchPrompts) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31129,7 +30532,6 @@ func (m *awsRestjson1_deserializeOpSearchPrompts) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31272,10 +30674,8 @@ func (m *awsRestjson1_deserializeOpSearchQueues) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31315,7 +30715,6 @@ func (m *awsRestjson1_deserializeOpSearchQueues) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31458,10 +30857,8 @@ func (m *awsRestjson1_deserializeOpSearchQuickConnects
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31501,7 +30898,6 @@ func (m *awsRestjson1_deserializeOpSearchQuickConnects
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31644,10 +31040,8 @@ func (m *awsRestjson1_deserializeOpSearchResourceTags)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31687,7 +31081,6 @@ func (m *awsRestjson1_deserializeOpSearchResourceTags)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31820,10 +31213,8 @@ func (m *awsRestjson1_deserializeOpSearchRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31863,7 +31254,6 @@ func (m *awsRestjson1_deserializeOpSearchRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32006,10 +31396,8 @@ func (m *awsRestjson1_deserializeOpSearchSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32049,7 +31437,6 @@ func (m *awsRestjson1_deserializeOpSearchSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32192,10 +31579,8 @@ func (m *awsRestjson1_deserializeOpSearchUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32235,7 +31620,6 @@ func (m *awsRestjson1_deserializeOpSearchUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32378,10 +31762,8 @@ func (m *awsRestjson1_deserializeOpSearchUsers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32421,7 +31803,6 @@ func (m *awsRestjson1_deserializeOpSearchUsers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32564,10 +31945,8 @@ func (m *awsRestjson1_deserializeOpSearchVocabularies)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32607,7 +31986,6 @@ func (m *awsRestjson1_deserializeOpSearchVocabularies)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32734,10 +32112,8 @@ func (m *awsRestjson1_deserializeOpSendChatIntegration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32777,7 +32153,6 @@ func (m *awsRestjson1_deserializeOpSendChatIntegration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32911,10 +32286,8 @@ func (m *awsRestjson1_deserializeOpSendOutboundEmail) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32926,7 +32299,6 @@ func (m *awsRestjson1_deserializeOpSendOutboundEmail) 
 	output := &SendOutboundEmailOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33017,10 +32389,8 @@ func (m *awsRestjson1_deserializeOpStartAttachedFileUp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33060,7 +32430,6 @@ func (m *awsRestjson1_deserializeOpStartAttachedFileUp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33225,10 +32594,8 @@ func (m *awsRestjson1_deserializeOpStartChatContact) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33268,7 +32635,6 @@ func (m *awsRestjson1_deserializeOpStartChatContact) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33420,10 +32786,8 @@ func (m *awsRestjson1_deserializeOpStartContactEvaluat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33463,7 +32827,6 @@ func (m *awsRestjson1_deserializeOpStartContactEvaluat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33600,10 +32963,8 @@ func (m *awsRestjson1_deserializeOpStartContactRecordi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33615,7 +32976,6 @@ func (m *awsRestjson1_deserializeOpStartContactRecordi
 	output := &StartContactRecordingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33697,10 +33057,8 @@ func (m *awsRestjson1_deserializeOpStartContactStreami
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33740,7 +33098,6 @@ func (m *awsRestjson1_deserializeOpStartContactStreami
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33865,10 +33222,8 @@ func (m *awsRestjson1_deserializeOpStartEmailContact) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33908,7 +33263,6 @@ func (m *awsRestjson1_deserializeOpStartEmailContact) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34039,10 +33393,8 @@ func (m *awsRestjson1_deserializeOpStartOutboundChatCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34082,7 +33434,6 @@ func (m *awsRestjson1_deserializeOpStartOutboundChatCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34213,10 +33564,8 @@ func (m *awsRestjson1_deserializeOpStartOutboundEmailC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34256,7 +33605,6 @@ func (m *awsRestjson1_deserializeOpStartOutboundEmailC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34387,10 +33735,8 @@ func (m *awsRestjson1_deserializeOpStartOutboundVoiceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34430,7 +33776,6 @@ func (m *awsRestjson1_deserializeOpStartOutboundVoiceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34561,10 +33906,8 @@ func (m *awsRestjson1_deserializeOpStartScreenSharing)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34576,7 +33919,6 @@ func (m *awsRestjson1_deserializeOpStartScreenSharing)
 	output := &StartScreenSharingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34664,10 +34006,8 @@ func (m *awsRestjson1_deserializeOpStartTaskContact) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34707,7 +34047,6 @@ func (m *awsRestjson1_deserializeOpStartTaskContact) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34835,10 +34174,8 @@ func (m *awsRestjson1_deserializeOpStartWebRTCContact)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34878,7 +34215,6 @@ func (m *awsRestjson1_deserializeOpStartWebRTCContact)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35026,10 +34362,8 @@ func (m *awsRestjson1_deserializeOpStopContact) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35041,7 +34375,6 @@ func (m *awsRestjson1_deserializeOpStopContact) Handle
 	output := &StopContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35126,10 +34459,8 @@ func (m *awsRestjson1_deserializeOpStopContactRecordin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35141,7 +34472,6 @@ func (m *awsRestjson1_deserializeOpStopContactRecordin
 	output := &StopContactRecordingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35220,10 +34550,8 @@ func (m *awsRestjson1_deserializeOpStopContactStreamin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35235,7 +34563,6 @@ func (m *awsRestjson1_deserializeOpStopContactStreamin
 	output := &StopContactStreamingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35317,10 +34644,8 @@ func (m *awsRestjson1_deserializeOpSubmitContactEvalua
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35360,7 +34685,6 @@ func (m *awsRestjson1_deserializeOpSubmitContactEvalua
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35494,10 +34818,8 @@ func (m *awsRestjson1_deserializeOpSuspendContactRecor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35509,7 +34831,6 @@ func (m *awsRestjson1_deserializeOpSuspendContactRecor
 	output := &SuspendContactRecordingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35588,10 +34909,8 @@ func (m *awsRestjson1_deserializeOpTagContact) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35603,7 +34922,6 @@ func (m *awsRestjson1_deserializeOpTagContact) HandleD
 	output := &TagContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35688,10 +35006,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35709,7 +35025,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35794,10 +35109,8 @@ func (m *awsRestjson1_deserializeOpTransferContact) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35837,7 +35150,6 @@ func (m *awsRestjson1_deserializeOpTransferContact) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35977,10 +35289,8 @@ func (m *awsRestjson1_deserializeOpUntagContact) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35992,7 +35302,6 @@ func (m *awsRestjson1_deserializeOpUntagContact) Handl
 	output := &UntagContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36077,10 +35386,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36098,7 +35405,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36183,10 +35489,8 @@ func (m *awsRestjson1_deserializeOpUpdateAgentStatus) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36204,7 +35508,6 @@ func (m *awsRestjson1_deserializeOpUpdateAgentStatus) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36295,10 +35598,8 @@ func (m *awsRestjson1_deserializeOpUpdateAuthenticatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36316,7 +35617,6 @@ func (m *awsRestjson1_deserializeOpUpdateAuthenticatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36401,10 +35701,8 @@ func (m *awsRestjson1_deserializeOpUpdateContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36416,7 +35714,6 @@ func (m *awsRestjson1_deserializeOpUpdateContact) Hand
 	output := &UpdateContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36507,10 +35804,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36522,7 +35817,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactAttrib
 	output := &UpdateContactAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36604,10 +35898,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactEvalua
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36647,7 +35939,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactEvalua
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36781,10 +36072,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36796,7 +36085,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowCo
 	output := &UpdateContactFlowContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36884,10 +36172,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36899,7 +36185,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMe
 	output := &UpdateContactFlowMetadataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36987,10 +36272,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37002,7 +36285,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMo
 	output := &UpdateContactFlowModuleContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37090,10 +36372,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37105,7 +36385,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowMo
 	output := &UpdateContactFlowModuleMetadataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37196,10 +36475,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowNa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37211,7 +36488,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactFlowNa
 	output := &UpdateContactFlowNameOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37299,10 +36575,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactRoutin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37314,7 +36588,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactRoutin
 	output := &UpdateContactRoutingDataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37402,10 +36675,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactSchedu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37417,7 +36688,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactSchedu
 	output := &UpdateContactScheduleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37505,10 +36775,8 @@ func (m *awsRestjson1_deserializeOpUpdateEmailAddressM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37548,7 +36816,6 @@ func (m *awsRestjson1_deserializeOpUpdateEmailAddressM
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37688,10 +36955,8 @@ func (m *awsRestjson1_deserializeOpUpdateEvaluationFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37731,7 +36996,6 @@ func (m *awsRestjson1_deserializeOpUpdateEvaluationFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37881,10 +37145,8 @@ func (m *awsRestjson1_deserializeOpUpdateHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37902,7 +37164,6 @@ func (m *awsRestjson1_deserializeOpUpdateHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37990,10 +37251,8 @@ func (m *awsRestjson1_deserializeOpUpdateHoursOfOperat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38011,7 +37270,6 @@ func (m *awsRestjson1_deserializeOpUpdateHoursOfOperat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38102,10 +37360,8 @@ func (m *awsRestjson1_deserializeOpUpdateInstanceAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38123,7 +37379,6 @@ func (m *awsRestjson1_deserializeOpUpdateInstanceAttri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38208,10 +37463,8 @@ func (m *awsRestjson1_deserializeOpUpdateInstanceStora
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38229,7 +37482,6 @@ func (m *awsRestjson1_deserializeOpUpdateInstanceStora
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38314,10 +37566,8 @@ func (m *awsRestjson1_deserializeOpUpdateParticipantAu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38329,7 +37579,6 @@ func (m *awsRestjson1_deserializeOpUpdateParticipantAu
 	output := &UpdateParticipantAuthenticationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38417,10 +37666,8 @@ func (m *awsRestjson1_deserializeOpUpdateParticipantRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38432,7 +37679,6 @@ func (m *awsRestjson1_deserializeOpUpdateParticipantRo
 	output := &UpdateParticipantRoleConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38520,10 +37766,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38563,7 +37807,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumber) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38703,10 +37946,8 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38724,7 +37965,6 @@ func (m *awsRestjson1_deserializeOpUpdatePhoneNumberMe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38818,10 +38058,8 @@ func (m *awsRestjson1_deserializeOpUpdatePredefinedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38839,7 +38077,6 @@ func (m *awsRestjson1_deserializeOpUpdatePredefinedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38924,10 +38161,8 @@ func (m *awsRestjson1_deserializeOpUpdatePrompt) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38967,7 +38202,6 @@ func (m *awsRestjson1_deserializeOpUpdatePrompt) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39101,10 +38335,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueHoursOfO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39122,7 +38354,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueHoursOfO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39207,10 +38438,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueMaxConta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39228,7 +38457,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueMaxConta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39313,10 +38541,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueName) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39334,7 +38560,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueName) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39422,10 +38647,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueOutbound
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39443,7 +38666,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueOutbound
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39528,10 +38750,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueOutbound
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39549,7 +38769,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueOutbound
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39640,10 +38859,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueStatus) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39661,7 +38878,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueStatus) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39746,10 +38962,8 @@ func (m *awsRestjson1_deserializeOpUpdateQuickConnectC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39767,7 +38981,6 @@ func (m *awsRestjson1_deserializeOpUpdateQuickConnectC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39852,10 +39065,8 @@ func (m *awsRestjson1_deserializeOpUpdateQuickConnectN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39873,7 +39084,6 @@ func (m *awsRestjson1_deserializeOpUpdateQuickConnectN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39958,10 +39168,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39979,7 +39187,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40064,10 +39271,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40085,7 +39290,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40170,10 +39374,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40191,7 +39393,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40276,10 +39477,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40297,7 +39496,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40385,10 +39583,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40406,7 +39602,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40491,10 +39686,8 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40512,7 +39705,6 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40600,10 +39792,8 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40621,7 +39811,6 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40706,10 +39895,8 @@ func (m *awsRestjson1_deserializeOpUpdateTaskTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40749,7 +39936,6 @@ func (m *awsRestjson1_deserializeOpUpdateTaskTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40987,10 +40173,8 @@ func (m *awsRestjson1_deserializeOpUpdateTrafficDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41002,7 +40186,6 @@ func (m *awsRestjson1_deserializeOpUpdateTrafficDistri
 	output := &UpdateTrafficDistributionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41090,10 +40273,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41111,7 +40292,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41196,10 +40376,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41217,7 +40395,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41305,10 +40482,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41326,7 +40501,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserHierarchy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41414,10 +40588,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserIdentityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41435,7 +40607,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserIdentityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41520,10 +40691,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserPhoneConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41541,7 +40710,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserPhoneConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41626,10 +40794,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserProficien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41647,7 +40813,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserProficien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41732,10 +40897,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserRoutingPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41753,7 +40916,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserRoutingPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41838,10 +41000,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserSecurityP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41859,7 +41019,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserSecurityP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41944,10 +41103,8 @@ func (m *awsRestjson1_deserializeOpUpdateViewContent) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41987,7 +41144,6 @@ func (m *awsRestjson1_deserializeOpUpdateViewContent) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42114,10 +41270,8 @@ func (m *awsRestjson1_deserializeOpUpdateViewMetadata)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42129,7 +41283,6 @@ func (m *awsRestjson1_deserializeOpUpdateViewMetadata)
 	output := &UpdateViewMetadataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
