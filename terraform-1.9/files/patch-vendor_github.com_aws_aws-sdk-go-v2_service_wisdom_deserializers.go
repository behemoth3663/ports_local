--- vendor/github.com/aws/aws-sdk-go-v2/service/wisdom/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/wisdom/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateAssistant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateAssistant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -206,10 +202,8 @@ func (m *awsRestjson1_deserializeOpCreateAssistantAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -249,7 +243,6 @@ func (m *awsRestjson1_deserializeOpCreateAssistantAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -370,10 +363,8 @@ func (m *awsRestjson1_deserializeOpCreateContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -413,7 +404,6 @@ func (m *awsRestjson1_deserializeOpCreateContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -534,10 +524,8 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -577,7 +565,6 @@ func (m *awsRestjson1_deserializeOpCreateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -695,10 +682,8 @@ func (m *awsRestjson1_deserializeOpCreateQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -738,7 +723,6 @@ func (m *awsRestjson1_deserializeOpCreateQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -859,10 +843,8 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -902,7 +884,6 @@ func (m *awsRestjson1_deserializeOpCreateSession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1017,10 +998,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssistant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1032,7 +1011,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssistant) Ha
 	output := &DeleteAssistantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1111,10 +1089,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssistantAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1126,7 +1102,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssistantAsso
 	output := &DeleteAssistantAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1205,10 +1180,8 @@ func (m *awsRestjson1_deserializeOpDeleteContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1220,7 +1193,6 @@ func (m *awsRestjson1_deserializeOpDeleteContent) Hand
 	output := &DeleteContentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1299,10 +1271,8 @@ func (m *awsRestjson1_deserializeOpDeleteImportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1314,7 +1284,6 @@ func (m *awsRestjson1_deserializeOpDeleteImportJob) Ha
 	output := &DeleteImportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1396,10 +1365,8 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1411,7 +1378,6 @@ func (m *awsRestjson1_deserializeOpDeleteKnowledgeBase
 	output := &DeleteKnowledgeBaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1493,10 +1459,8 @@ func (m *awsRestjson1_deserializeOpDeleteQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1508,7 +1472,6 @@ func (m *awsRestjson1_deserializeOpDeleteQuickResponse
 	output := &DeleteQuickResponseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1587,10 +1550,8 @@ func (m *awsRestjson1_deserializeOpGetAssistant) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1630,7 +1591,6 @@ func (m *awsRestjson1_deserializeOpGetAssistant) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1745,10 +1705,8 @@ func (m *awsRestjson1_deserializeOpGetAssistantAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1788,7 +1746,6 @@ func (m *awsRestjson1_deserializeOpGetAssistantAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1903,10 +1860,8 @@ func (m *awsRestjson1_deserializeOpGetContent) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1946,7 +1901,6 @@ func (m *awsRestjson1_deserializeOpGetContent) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2061,10 +2015,8 @@ func (m *awsRestjson1_deserializeOpGetContentSummary) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2104,7 +2056,6 @@ func (m *awsRestjson1_deserializeOpGetContentSummary) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2219,10 +2170,8 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2262,7 +2211,6 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2377,10 +2325,8 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2420,7 +2366,6 @@ func (m *awsRestjson1_deserializeOpGetKnowledgeBase) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2535,10 +2480,8 @@ func (m *awsRestjson1_deserializeOpGetQuickResponse) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2578,7 +2521,6 @@ func (m *awsRestjson1_deserializeOpGetQuickResponse) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2693,10 +2635,8 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2736,7 +2676,6 @@ func (m *awsRestjson1_deserializeOpGetRecommendations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2856,10 +2795,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2899,7 +2836,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3014,10 +2950,8 @@ func (m *awsRestjson1_deserializeOpListAssistantAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3057,7 +2991,6 @@ func (m *awsRestjson1_deserializeOpListAssistantAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3181,10 +3114,8 @@ func (m *awsRestjson1_deserializeOpListAssistants) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3224,7 +3155,6 @@ func (m *awsRestjson1_deserializeOpListAssistants) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3345,10 +3275,8 @@ func (m *awsRestjson1_deserializeOpListContents) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3388,7 +3316,6 @@ func (m *awsRestjson1_deserializeOpListContents) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3512,10 +3439,8 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3555,7 +3480,6 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3676,10 +3600,8 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3719,7 +3641,6 @@ func (m *awsRestjson1_deserializeOpListKnowledgeBases)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3840,10 +3761,8 @@ func (m *awsRestjson1_deserializeOpListQuickResponses)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3883,7 +3802,6 @@ func (m *awsRestjson1_deserializeOpListQuickResponses)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4007,10 +3925,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4050,7 +3966,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4159,10 +4074,8 @@ func (m *awsRestjson1_deserializeOpNotifyRecommendatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4202,7 +4115,6 @@ func (m *awsRestjson1_deserializeOpNotifyRecommendatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4322,10 +4234,8 @@ func (m *awsRestjson1_deserializeOpQueryAssistant) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4365,7 +4275,6 @@ func (m *awsRestjson1_deserializeOpQueryAssistant) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4492,10 +4401,8 @@ func (m *awsRestjson1_deserializeOpRemoveKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4507,7 +4414,6 @@ func (m *awsRestjson1_deserializeOpRemoveKnowledgeBase
 	output := &RemoveKnowledgeBaseTemplateUriOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4586,10 +4492,8 @@ func (m *awsRestjson1_deserializeOpSearchContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4629,7 +4533,6 @@ func (m *awsRestjson1_deserializeOpSearchContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4753,10 +4656,8 @@ func (m *awsRestjson1_deserializeOpSearchQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4796,7 +4697,6 @@ func (m *awsRestjson1_deserializeOpSearchQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4923,10 +4823,8 @@ func (m *awsRestjson1_deserializeOpSearchSessions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4966,7 +4864,6 @@ func (m *awsRestjson1_deserializeOpSearchSessions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5090,10 +4987,8 @@ func (m *awsRestjson1_deserializeOpStartContentUpload)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5133,7 +5028,6 @@ func (m *awsRestjson1_deserializeOpStartContentUpload)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5282,10 +5176,8 @@ func (m *awsRestjson1_deserializeOpStartImportJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5325,7 +5217,6 @@ func (m *awsRestjson1_deserializeOpStartImportJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5446,10 +5337,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5461,7 +5350,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5537,10 +5425,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5552,7 +5438,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5625,10 +5510,8 @@ func (m *awsRestjson1_deserializeOpUpdateContent) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5668,7 +5551,6 @@ func (m *awsRestjson1_deserializeOpUpdateContent) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5786,10 +5668,8 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5829,7 +5709,6 @@ func (m *awsRestjson1_deserializeOpUpdateKnowledgeBase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5944,10 +5823,8 @@ func (m *awsRestjson1_deserializeOpUpdateQuickResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5987,7 +5864,6 @@ func (m *awsRestjson1_deserializeOpUpdateQuickResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
