--- vendor/github.com/aws/aws-sdk-go-v2/service/databrew/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/databrew/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteRecipeVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteRecipeVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -377,10 +370,8 @@ func (m *awsRestjson1_deserializeOpCreateProfileJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -420,7 +411,6 @@ func (m *awsRestjson1_deserializeOpCreateProfileJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -545,10 +535,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -588,7 +576,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -710,10 +697,8 @@ func (m *awsRestjson1_deserializeOpCreateRecipe) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -753,7 +738,6 @@ func (m *awsRestjson1_deserializeOpCreateRecipe) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -872,10 +856,8 @@ func (m *awsRestjson1_deserializeOpCreateRecipeJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -915,7 +897,6 @@ func (m *awsRestjson1_deserializeOpCreateRecipeJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1040,10 +1021,8 @@ func (m *awsRestjson1_deserializeOpCreateRuleset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1083,7 +1062,6 @@ func (m *awsRestjson1_deserializeOpCreateRuleset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1202,10 +1180,8 @@ func (m *awsRestjson1_deserializeOpCreateSchedule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1245,7 +1221,6 @@ func (m *awsRestjson1_deserializeOpCreateSchedule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1364,10 +1339,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1407,7 +1380,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1526,10 +1498,8 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1569,7 +1539,6 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1688,10 +1657,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1731,7 +1698,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1850,10 +1816,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecipeVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1893,7 +1857,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecipeVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2021,10 +1984,8 @@ func (m *awsRestjson1_deserializeOpDeleteRuleset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2064,7 +2025,6 @@ func (m *awsRestjson1_deserializeOpDeleteRuleset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2183,10 +2143,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchedule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2226,7 +2184,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchedule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2342,10 +2299,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2385,7 +2340,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2598,10 +2552,8 @@ func (m *awsRestjson1_deserializeOpDescribeJob) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2641,7 +2593,6 @@ func (m *awsRestjson1_deserializeOpDescribeJob) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2958,10 +2909,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobRun) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3001,7 +2950,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobRun) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3273,10 +3221,8 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3316,7 +3262,6 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3562,10 +3507,8 @@ func (m *awsRestjson1_deserializeOpDescribeRecipe) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3605,7 +3548,6 @@ func (m *awsRestjson1_deserializeOpDescribeRecipe) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3842,10 +3784,8 @@ func (m *awsRestjson1_deserializeOpDescribeRuleset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3885,7 +3825,6 @@ func (m *awsRestjson1_deserializeOpDescribeRuleset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4088,10 +4027,8 @@ func (m *awsRestjson1_deserializeOpDescribeSchedule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4131,7 +4068,6 @@ func (m *awsRestjson1_deserializeOpDescribeSchedule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4325,10 +4261,8 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4368,7 +4302,6 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4486,10 +4419,8 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4529,7 +4460,6 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4650,10 +4580,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4693,7 +4621,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4811,10 +4738,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4854,7 +4779,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4972,10 +4896,8 @@ func (m *awsRestjson1_deserializeOpListRecipes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5015,7 +4937,6 @@ func (m *awsRestjson1_deserializeOpListRecipes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5133,10 +5054,8 @@ func (m *awsRestjson1_deserializeOpListRecipeVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5176,7 +5095,6 @@ func (m *awsRestjson1_deserializeOpListRecipeVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5294,10 +5212,8 @@ func (m *awsRestjson1_deserializeOpListRulesets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5337,7 +5253,6 @@ func (m *awsRestjson1_deserializeOpListRulesets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5458,10 +5373,8 @@ func (m *awsRestjson1_deserializeOpListSchedules) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5501,7 +5414,6 @@ func (m *awsRestjson1_deserializeOpListSchedules) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5619,10 +5531,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5662,7 +5572,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5777,10 +5686,8 @@ func (m *awsRestjson1_deserializeOpPublishRecipe) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5820,7 +5727,6 @@ func (m *awsRestjson1_deserializeOpPublishRecipe) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5939,10 +5845,8 @@ func (m *awsRestjson1_deserializeOpSendProjectSessionA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5982,7 +5886,6 @@ func (m *awsRestjson1_deserializeOpSendProjectSessionA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6123,10 +6026,8 @@ func (m *awsRestjson1_deserializeOpStartJobRun) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6166,7 +6067,6 @@ func (m *awsRestjson1_deserializeOpStartJobRun) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6288,10 +6188,8 @@ func (m *awsRestjson1_deserializeOpStartProjectSession
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6331,7 +6229,6 @@ func (m *awsRestjson1_deserializeOpStartProjectSession
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6462,10 +6359,8 @@ func (m *awsRestjson1_deserializeOpStopJobRun) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6505,7 +6400,6 @@ func (m *awsRestjson1_deserializeOpStopJobRun) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6621,10 +6515,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6636,7 +6528,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6715,10 +6606,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6730,7 +6619,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6809,10 +6697,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6852,7 +6738,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6971,10 +6856,8 @@ func (m *awsRestjson1_deserializeOpUpdateProfileJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7014,7 +6897,6 @@ func (m *awsRestjson1_deserializeOpUpdateProfileJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7133,10 +7015,8 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7176,7 +7056,6 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7308,10 +7187,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecipe) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7351,7 +7228,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecipe) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7467,10 +7343,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecipeJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7510,7 +7384,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecipeJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7629,10 +7502,8 @@ func (m *awsRestjson1_deserializeOpUpdateRuleset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7672,7 +7543,6 @@ func (m *awsRestjson1_deserializeOpUpdateRuleset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7788,10 +7658,8 @@ func (m *awsRestjson1_deserializeOpUpdateSchedule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7831,7 +7699,6 @@ func (m *awsRestjson1_deserializeOpUpdateSchedule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
