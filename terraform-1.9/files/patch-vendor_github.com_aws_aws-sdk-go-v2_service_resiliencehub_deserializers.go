--- vendor/github.com/aws/aws-sdk-go-v2/service/resiliencehub/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/resiliencehub/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAcceptResourceGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAcceptResourceGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -218,10 +214,8 @@ func (m *awsRestjson1_deserializeOpAddDraftAppVersionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -261,7 +255,6 @@ func (m *awsRestjson1_deserializeOpAddDraftAppVersionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -406,10 +399,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRecommen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -449,7 +440,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRecommen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -584,10 +574,8 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -627,7 +615,6 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -754,10 +741,8 @@ func (m *awsRestjson1_deserializeOpCreateAppVersionApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -797,7 +782,6 @@ func (m *awsRestjson1_deserializeOpCreateAppVersionApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -942,10 +926,8 @@ func (m *awsRestjson1_deserializeOpCreateAppVersionRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -985,7 +967,6 @@ func (m *awsRestjson1_deserializeOpCreateAppVersionRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1130,10 +1111,8 @@ func (m *awsRestjson1_deserializeOpCreateRecommendatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1173,7 +1152,6 @@ func (m *awsRestjson1_deserializeOpCreateRecommendatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1300,10 +1278,8 @@ func (m *awsRestjson1_deserializeOpCreateResiliencyPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1343,7 +1319,6 @@ func (m *awsRestjson1_deserializeOpCreateResiliencyPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1467,10 +1442,8 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1510,7 +1483,6 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1635,10 +1607,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppAssessment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1678,7 +1648,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppAssessment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1815,10 +1784,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppInputSourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1858,7 +1825,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppInputSourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1991,10 +1957,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppVersionApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2034,7 +1998,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppVersionApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2176,10 +2139,8 @@ func (m *awsRestjson1_deserializeOpDeleteAppVersionRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2219,7 +2180,6 @@ func (m *awsRestjson1_deserializeOpDeleteAppVersionRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2361,10 +2321,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecommendatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2404,7 +2362,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecommendatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2538,10 +2495,8 @@ func (m *awsRestjson1_deserializeOpDeleteResiliencyPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2581,7 +2536,6 @@ func (m *awsRestjson1_deserializeOpDeleteResiliencyPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2709,10 +2663,8 @@ func (m *awsRestjson1_deserializeOpDescribeApp) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2752,7 +2704,6 @@ func (m *awsRestjson1_deserializeOpDescribeApp) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2873,10 +2824,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppAssessme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2916,7 +2865,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppAssessme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3037,10 +2985,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3080,7 +3026,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3219,10 +3164,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3262,7 +3205,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3404,10 +3346,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3447,7 +3387,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3589,10 +3528,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3632,7 +3569,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3793,10 +3729,8 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3836,7 +3770,6 @@ func (m *awsRestjson1_deserializeOpDescribeAppVersionT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3979,10 +3912,8 @@ func (m *awsRestjson1_deserializeOpDescribeDraftAppVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4022,7 +3953,6 @@ func (m *awsRestjson1_deserializeOpDescribeDraftAppVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4195,10 +4125,8 @@ func (m *awsRestjson1_deserializeOpDescribeMetricsExpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4238,7 +4166,6 @@ func (m *awsRestjson1_deserializeOpDescribeMetricsExpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4386,10 +4313,8 @@ func (m *awsRestjson1_deserializeOpDescribeResiliencyP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4429,7 +4354,6 @@ func (m *awsRestjson1_deserializeOpDescribeResiliencyP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4550,10 +4474,8 @@ func (m *awsRestjson1_deserializeOpDescribeResourceGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4593,7 +4515,6 @@ func (m *awsRestjson1_deserializeOpDescribeResourceGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4736,10 +4657,8 @@ func (m *awsRestjson1_deserializeOpImportResourcesToDr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4779,7 +4698,6 @@ func (m *awsRestjson1_deserializeOpImportResourcesToDr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4943,10 +4861,8 @@ func (m *awsRestjson1_deserializeOpListAlarmRecommenda
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4986,7 +4902,6 @@ func (m *awsRestjson1_deserializeOpListAlarmRecommenda
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5116,10 +5031,8 @@ func (m *awsRestjson1_deserializeOpListAppAssessmentCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5159,7 +5072,6 @@ func (m *awsRestjson1_deserializeOpListAppAssessmentCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5286,10 +5198,8 @@ func (m *awsRestjson1_deserializeOpListAppAssessmentRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5329,7 +5239,6 @@ func (m *awsRestjson1_deserializeOpListAppAssessmentRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5456,10 +5365,8 @@ func (m *awsRestjson1_deserializeOpListAppAssessments)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5499,7 +5406,6 @@ func (m *awsRestjson1_deserializeOpListAppAssessments)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5629,10 +5535,8 @@ func (m *awsRestjson1_deserializeOpListAppComponentCom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5672,7 +5576,6 @@ func (m *awsRestjson1_deserializeOpListAppComponentCom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5802,10 +5705,8 @@ func (m *awsRestjson1_deserializeOpListAppComponentRec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5845,7 +5746,6 @@ func (m *awsRestjson1_deserializeOpListAppComponentRec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5975,10 +5875,8 @@ func (m *awsRestjson1_deserializeOpListAppInputSources
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6018,7 +5916,6 @@ func (m *awsRestjson1_deserializeOpListAppInputSources
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6148,10 +6045,8 @@ func (m *awsRestjson1_deserializeOpListApps) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6191,7 +6086,6 @@ func (m *awsRestjson1_deserializeOpListApps) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6318,10 +6212,8 @@ func (m *awsRestjson1_deserializeOpListAppVersionAppCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6361,7 +6253,6 @@ func (m *awsRestjson1_deserializeOpListAppVersionAppCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6512,10 +6403,8 @@ func (m *awsRestjson1_deserializeOpListAppVersionResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6555,7 +6444,6 @@ func (m *awsRestjson1_deserializeOpListAppVersionResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6685,10 +6573,8 @@ func (m *awsRestjson1_deserializeOpListAppVersionResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6728,7 +6614,6 @@ func (m *awsRestjson1_deserializeOpListAppVersionResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6870,10 +6755,8 @@ func (m *awsRestjson1_deserializeOpListAppVersions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6913,7 +6796,6 @@ func (m *awsRestjson1_deserializeOpListAppVersions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7040,10 +6922,8 @@ func (m *awsRestjson1_deserializeOpListMetrics) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7083,7 +6963,6 @@ func (m *awsRestjson1_deserializeOpListMetrics) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7210,10 +7089,8 @@ func (m *awsRestjson1_deserializeOpListRecommendationT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7253,7 +7130,6 @@ func (m *awsRestjson1_deserializeOpListRecommendationT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7380,10 +7256,8 @@ func (m *awsRestjson1_deserializeOpListResiliencyPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7423,7 +7297,6 @@ func (m *awsRestjson1_deserializeOpListResiliencyPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7553,10 +7426,8 @@ func (m *awsRestjson1_deserializeOpListResourceGroupin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7596,7 +7467,6 @@ func (m *awsRestjson1_deserializeOpListResourceGroupin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7726,10 +7596,8 @@ func (m *awsRestjson1_deserializeOpListSopRecommendati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7769,7 +7637,6 @@ func (m *awsRestjson1_deserializeOpListSopRecommendati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7902,10 +7769,8 @@ func (m *awsRestjson1_deserializeOpListSuggestedResili
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7945,7 +7810,6 @@ func (m *awsRestjson1_deserializeOpListSuggestedResili
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8075,10 +7939,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8118,7 +7980,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8239,10 +8100,8 @@ func (m *awsRestjson1_deserializeOpListTestRecommendat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8282,7 +8141,6 @@ func (m *awsRestjson1_deserializeOpListTestRecommendat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8415,10 +8273,8 @@ func (m *awsRestjson1_deserializeOpListUnsupportedAppV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8458,7 +8314,6 @@ func (m *awsRestjson1_deserializeOpListUnsupportedAppV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8600,10 +8455,8 @@ func (m *awsRestjson1_deserializeOpPublishAppVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8643,7 +8496,6 @@ func (m *awsRestjson1_deserializeOpPublishAppVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8802,10 +8654,8 @@ func (m *awsRestjson1_deserializeOpPutDraftAppVersionT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8845,7 +8695,6 @@ func (m *awsRestjson1_deserializeOpPutDraftAppVersionT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8982,10 +8831,8 @@ func (m *awsRestjson1_deserializeOpRejectResourceGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9025,7 +8872,6 @@ func (m *awsRestjson1_deserializeOpRejectResourceGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9155,10 +9001,8 @@ func (m *awsRestjson1_deserializeOpRemoveDraftAppVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9198,7 +9042,6 @@ func (m *awsRestjson1_deserializeOpRemoveDraftAppVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9335,10 +9178,8 @@ func (m *awsRestjson1_deserializeOpResolveAppVersionRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9378,7 +9219,6 @@ func (m *awsRestjson1_deserializeOpResolveAppVersionRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9533,10 +9373,8 @@ func (m *awsRestjson1_deserializeOpStartAppAssessment)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9576,7 +9414,6 @@ func (m *awsRestjson1_deserializeOpStartAppAssessment)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9703,10 +9540,8 @@ func (m *awsRestjson1_deserializeOpStartMetricsExport)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9746,7 +9581,6 @@ func (m *awsRestjson1_deserializeOpStartMetricsExport)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9883,10 +9717,8 @@ func (m *awsRestjson1_deserializeOpStartResourceGroupi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9926,7 +9758,6 @@ func (m *awsRestjson1_deserializeOpStartResourceGroupi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10081,10 +9912,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10096,7 +9925,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10181,10 +10009,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10196,7 +10022,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10281,10 +10106,8 @@ func (m *awsRestjson1_deserializeOpUpdateApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10324,7 +10147,6 @@ func (m *awsRestjson1_deserializeOpUpdateApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10448,10 +10270,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10491,7 +10311,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10633,10 +10452,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersionApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10676,7 +10493,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersionApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10818,10 +10634,8 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersionRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10861,7 +10675,6 @@ func (m *awsRestjson1_deserializeOpUpdateAppVersionRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11006,10 +10819,8 @@ func (m *awsRestjson1_deserializeOpUpdateResiliencyPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11049,7 +10860,6 @@ func (m *awsRestjson1_deserializeOpUpdateResiliencyPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
