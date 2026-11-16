--- vendor/github.com/aws/aws-sdk-go-v2/service/accessanalyzer/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/accessanalyzer/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpApplyArchiveRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_deserializeOpApplyArchiveRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -152,10 +148,8 @@ func (m *awsRestjson1_deserializeOpCancelPolicyGenerat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -167,7 +161,6 @@ func (m *awsRestjson1_deserializeOpCancelPolicyGenerat
 	output := &CancelPolicyGenerationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -249,10 +242,8 @@ func (m *awsRestjson1_deserializeOpCheckAccessNotGrant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -292,7 +283,6 @@ func (m *awsRestjson1_deserializeOpCheckAccessNotGrant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -434,10 +424,8 @@ func (m *awsRestjson1_deserializeOpCheckNoNewAccess) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -477,7 +465,6 @@ func (m *awsRestjson1_deserializeOpCheckNoNewAccess) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -619,10 +606,8 @@ func (m *awsRestjson1_deserializeOpCheckNoPublicAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -662,7 +647,6 @@ func (m *awsRestjson1_deserializeOpCheckNoPublicAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -804,10 +788,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessPreview
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -847,7 +829,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessPreview
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -978,10 +959,8 @@ func (m *awsRestjson1_deserializeOpCreateAnalyzer) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1021,7 +1000,6 @@ func (m *awsRestjson1_deserializeOpCreateAnalyzer) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1149,10 +1127,8 @@ func (m *awsRestjson1_deserializeOpCreateArchiveRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1170,7 +1146,6 @@ func (m *awsRestjson1_deserializeOpCreateArchiveRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1261,10 +1236,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnalyzer) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1282,7 +1255,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnalyzer) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1367,10 +1339,8 @@ func (m *awsRestjson1_deserializeOpDeleteArchiveRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1388,7 +1358,6 @@ func (m *awsRestjson1_deserializeOpDeleteArchiveRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1473,10 +1442,8 @@ func (m *awsRestjson1_deserializeOpGenerateFindingReco
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1494,7 +1461,6 @@ func (m *awsRestjson1_deserializeOpGenerateFindingReco
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1576,10 +1542,8 @@ func (m *awsRestjson1_deserializeOpGetAccessPreview) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1619,7 +1583,6 @@ func (m *awsRestjson1_deserializeOpGetAccessPreview) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1740,10 +1703,8 @@ func (m *awsRestjson1_deserializeOpGetAnalyzedResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1783,7 +1744,6 @@ func (m *awsRestjson1_deserializeOpGetAnalyzedResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1904,10 +1864,8 @@ func (m *awsRestjson1_deserializeOpGetAnalyzer) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1947,7 +1905,6 @@ func (m *awsRestjson1_deserializeOpGetAnalyzer) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2068,10 +2025,8 @@ func (m *awsRestjson1_deserializeOpGetArchiveRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2111,7 +2066,6 @@ func (m *awsRestjson1_deserializeOpGetArchiveRule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2232,10 +2186,8 @@ func (m *awsRestjson1_deserializeOpGetFinding) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2275,7 +2227,6 @@ func (m *awsRestjson1_deserializeOpGetFinding) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2396,10 +2347,8 @@ func (m *awsRestjson1_deserializeOpGetFindingRecommend
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2439,7 +2388,6 @@ func (m *awsRestjson1_deserializeOpGetFindingRecommend
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2627,10 +2575,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsStatisti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2670,7 +2616,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsStatisti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2804,10 +2749,8 @@ func (m *awsRestjson1_deserializeOpGetFindingV2) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2847,7 +2790,6 @@ func (m *awsRestjson1_deserializeOpGetFindingV2) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3079,10 +3021,8 @@ func (m *awsRestjson1_deserializeOpGetGeneratedPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3122,7 +3062,6 @@ func (m *awsRestjson1_deserializeOpGetGeneratedPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3245,10 +3184,8 @@ func (m *awsRestjson1_deserializeOpListAccessPreviewFi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3288,7 +3225,6 @@ func (m *awsRestjson1_deserializeOpListAccessPreviewFi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3421,10 +3357,8 @@ func (m *awsRestjson1_deserializeOpListAccessPreviews)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3464,7 +3398,6 @@ func (m *awsRestjson1_deserializeOpListAccessPreviews)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3594,10 +3527,8 @@ func (m *awsRestjson1_deserializeOpListAnalyzedResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3637,7 +3568,6 @@ func (m *awsRestjson1_deserializeOpListAnalyzedResourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3767,10 +3697,8 @@ func (m *awsRestjson1_deserializeOpListAnalyzers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3810,7 +3738,6 @@ func (m *awsRestjson1_deserializeOpListAnalyzers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3937,10 +3864,8 @@ func (m *awsRestjson1_deserializeOpListArchiveRules) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3980,7 +3905,6 @@ func (m *awsRestjson1_deserializeOpListArchiveRules) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4107,10 +4031,8 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4150,7 +4072,6 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4280,10 +4201,8 @@ func (m *awsRestjson1_deserializeOpListFindingsV2) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4323,7 +4242,6 @@ func (m *awsRestjson1_deserializeOpListFindingsV2) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4453,10 +4371,8 @@ func (m *awsRestjson1_deserializeOpListPolicyGeneratio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4496,7 +4412,6 @@ func (m *awsRestjson1_deserializeOpListPolicyGeneratio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4623,10 +4538,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4666,7 +4579,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4787,10 +4699,8 @@ func (m *awsRestjson1_deserializeOpStartPolicyGenerati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4830,7 +4740,6 @@ func (m *awsRestjson1_deserializeOpStartPolicyGenerati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4958,10 +4867,8 @@ func (m *awsRestjson1_deserializeOpStartResourceScan) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4979,7 +4886,6 @@ func (m *awsRestjson1_deserializeOpStartResourceScan) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5064,10 +4970,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5079,7 +4983,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5164,10 +5067,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5179,7 +5080,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5264,10 +5164,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnalyzer) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5307,7 +5205,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnalyzer) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5431,10 +5328,8 @@ func (m *awsRestjson1_deserializeOpUpdateArchiveRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5452,7 +5347,6 @@ func (m *awsRestjson1_deserializeOpUpdateArchiveRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5537,10 +5431,8 @@ func (m *awsRestjson1_deserializeOpUpdateFindings) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5558,7 +5450,6 @@ func (m *awsRestjson1_deserializeOpUpdateFindings) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5643,10 +5534,8 @@ func (m *awsRestjson1_deserializeOpValidatePolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5686,7 +5575,6 @@ func (m *awsRestjson1_deserializeOpValidatePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
