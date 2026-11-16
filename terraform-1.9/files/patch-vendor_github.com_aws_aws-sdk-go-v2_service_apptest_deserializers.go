--- vendor/github.com/aws/aws-sdk-go-v2/service/apptest/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apptest/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateTestCase) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateTestCase) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -229,10 +225,8 @@ func (m *awsRestjson1_deserializeOpCreateTestConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -272,7 +266,6 @@ func (m *awsRestjson1_deserializeOpCreateTestConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -413,10 +406,8 @@ func (m *awsRestjson1_deserializeOpCreateTestSuite) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -456,7 +447,6 @@ func (m *awsRestjson1_deserializeOpCreateTestSuite) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -597,10 +587,8 @@ func (m *awsRestjson1_deserializeOpDeleteTestCase) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -612,7 +600,6 @@ func (m *awsRestjson1_deserializeOpDeleteTestCase) Han
 	output := &DeleteTestCaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -700,10 +687,8 @@ func (m *awsRestjson1_deserializeOpDeleteTestConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -715,7 +700,6 @@ func (m *awsRestjson1_deserializeOpDeleteTestConfigura
 	output := &DeleteTestConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -803,10 +787,8 @@ func (m *awsRestjson1_deserializeOpDeleteTestRun) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -818,7 +800,6 @@ func (m *awsRestjson1_deserializeOpDeleteTestRun) Hand
 	output := &DeleteTestRunOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -903,10 +884,8 @@ func (m *awsRestjson1_deserializeOpDeleteTestSuite) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -918,7 +897,6 @@ func (m *awsRestjson1_deserializeOpDeleteTestSuite) Ha
 	output := &DeleteTestSuiteOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1006,10 +984,8 @@ func (m *awsRestjson1_deserializeOpGetTestCase) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1049,7 +1025,6 @@ func (m *awsRestjson1_deserializeOpGetTestCase) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1279,10 +1254,8 @@ func (m *awsRestjson1_deserializeOpGetTestConfiguratio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1322,7 +1295,6 @@ func (m *awsRestjson1_deserializeOpGetTestConfiguratio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1562,10 +1534,8 @@ func (m *awsRestjson1_deserializeOpGetTestRunStep) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1605,7 +1575,6 @@ func (m *awsRestjson1_deserializeOpGetTestRunStep) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1856,10 +1825,8 @@ func (m *awsRestjson1_deserializeOpGetTestSuite) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1899,7 +1866,6 @@ func (m *awsRestjson1_deserializeOpGetTestSuite) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2139,10 +2105,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2182,7 +2146,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2303,10 +2266,8 @@ func (m *awsRestjson1_deserializeOpListTestCases) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2346,7 +2307,6 @@ func (m *awsRestjson1_deserializeOpListTestCases) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2476,10 +2436,8 @@ func (m *awsRestjson1_deserializeOpListTestConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2519,7 +2477,6 @@ func (m *awsRestjson1_deserializeOpListTestConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2649,10 +2606,8 @@ func (m *awsRestjson1_deserializeOpListTestRuns) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2692,7 +2647,6 @@ func (m *awsRestjson1_deserializeOpListTestRuns) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2822,10 +2776,8 @@ func (m *awsRestjson1_deserializeOpListTestRunSteps) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2865,7 +2817,6 @@ func (m *awsRestjson1_deserializeOpListTestRunSteps) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2995,10 +2946,8 @@ func (m *awsRestjson1_deserializeOpListTestRunTestCase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3038,7 +2987,6 @@ func (m *awsRestjson1_deserializeOpListTestRunTestCase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3168,10 +3116,8 @@ func (m *awsRestjson1_deserializeOpListTestSuites) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3211,7 +3157,6 @@ func (m *awsRestjson1_deserializeOpListTestSuites) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3341,10 +3286,8 @@ func (m *awsRestjson1_deserializeOpStartTestRun) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3384,7 +3327,6 @@ func (m *awsRestjson1_deserializeOpStartTestRun) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3524,10 +3466,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3539,7 +3479,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3627,10 +3566,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3642,7 +3579,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3727,10 +3663,8 @@ func (m *awsRestjson1_deserializeOpUpdateTestCase) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3770,7 +3704,6 @@ func (m *awsRestjson1_deserializeOpUpdateTestCase) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3911,10 +3844,8 @@ func (m *awsRestjson1_deserializeOpUpdateTestConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3954,7 +3885,6 @@ func (m *awsRestjson1_deserializeOpUpdateTestConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4095,10 +4025,8 @@ func (m *awsRestjson1_deserializeOpUpdateTestSuite) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4138,7 +4066,6 @@ func (m *awsRestjson1_deserializeOpUpdateTestSuite) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
