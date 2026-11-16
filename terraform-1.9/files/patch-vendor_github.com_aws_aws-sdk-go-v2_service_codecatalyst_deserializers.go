--- vendor/github.com/aws/aws-sdk-go-v2/service/codecatalyst/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/codecatalyst/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessToken) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessToken) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -246,10 +242,8 @@ func (m *awsRestjson1_deserializeOpCreateDevEnvironmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -289,7 +283,6 @@ func (m *awsRestjson1_deserializeOpCreateDevEnvironmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -444,10 +437,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -487,7 +478,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -642,10 +632,8 @@ func (m *awsRestjson1_deserializeOpCreateSourceReposit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -685,7 +673,6 @@ func (m *awsRestjson1_deserializeOpCreateSourceReposit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -840,10 +827,8 @@ func (m *awsRestjson1_deserializeOpCreateSourceReposit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -883,7 +868,6 @@ func (m *awsRestjson1_deserializeOpCreateSourceReposit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1042,10 +1026,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessToken) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1057,7 +1039,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessToken) 
 	output := &DeleteAccessTokenOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1145,10 +1126,8 @@ func (m *awsRestjson1_deserializeOpDeleteDevEnvironmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1188,7 +1167,6 @@ func (m *awsRestjson1_deserializeOpDeleteDevEnvironmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1334,10 +1312,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1377,7 +1353,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1523,10 +1498,8 @@ func (m *awsRestjson1_deserializeOpDeleteSourceReposit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1566,7 +1539,6 @@ func (m *awsRestjson1_deserializeOpDeleteSourceReposit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1712,10 +1684,8 @@ func (m *awsRestjson1_deserializeOpDeleteSpace) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1755,7 +1725,6 @@ func (m *awsRestjson1_deserializeOpDeleteSpace) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1892,10 +1861,8 @@ func (m *awsRestjson1_deserializeOpGetDevEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1935,7 +1902,6 @@ func (m *awsRestjson1_deserializeOpGetDevEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2176,10 +2142,8 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2219,7 +2183,6 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2374,10 +2337,8 @@ func (m *awsRestjson1_deserializeOpGetSourceRepository
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2417,7 +2378,6 @@ func (m *awsRestjson1_deserializeOpGetSourceRepository
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2598,10 +2558,8 @@ func (m *awsRestjson1_deserializeOpGetSourceRepository
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2641,7 +2599,6 @@ func (m *awsRestjson1_deserializeOpGetSourceRepository
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2769,10 +2726,8 @@ func (m *awsRestjson1_deserializeOpGetSpace) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2812,7 +2767,6 @@ func (m *awsRestjson1_deserializeOpGetSpace) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2967,10 +2921,8 @@ func (m *awsRestjson1_deserializeOpGetSubscription) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3010,7 +2962,6 @@ func (m *awsRestjson1_deserializeOpGetSubscription) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3169,10 +3120,8 @@ func (m *awsRestjson1_deserializeOpGetUserDetails) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3212,7 +3161,6 @@ func (m *awsRestjson1_deserializeOpGetUserDetails) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3372,10 +3320,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3415,7 +3361,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3637,10 +3582,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowRun) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3680,7 +3623,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowRun) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3888,10 +3830,8 @@ func (m *awsRestjson1_deserializeOpListAccessTokens) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3931,7 +3871,6 @@ func (m *awsRestjson1_deserializeOpListAccessTokens) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4064,10 +4003,8 @@ func (m *awsRestjson1_deserializeOpListDevEnvironments
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4107,7 +4044,6 @@ func (m *awsRestjson1_deserializeOpListDevEnvironments
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4240,10 +4176,8 @@ func (m *awsRestjson1_deserializeOpListDevEnvironmentS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4283,7 +4217,6 @@ func (m *awsRestjson1_deserializeOpListDevEnvironmentS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4416,10 +4349,8 @@ func (m *awsRestjson1_deserializeOpListEventLogs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4459,7 +4390,6 @@ func (m *awsRestjson1_deserializeOpListEventLogs) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4592,10 +4522,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4635,7 +4563,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4768,10 +4695,8 @@ func (m *awsRestjson1_deserializeOpListSourceRepositor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4811,7 +4736,6 @@ func (m *awsRestjson1_deserializeOpListSourceRepositor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4944,10 +4868,8 @@ func (m *awsRestjson1_deserializeOpListSourceRepositor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4987,7 +4909,6 @@ func (m *awsRestjson1_deserializeOpListSourceRepositor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5120,10 +5041,8 @@ func (m *awsRestjson1_deserializeOpListSpaces) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5163,7 +5082,6 @@ func (m *awsRestjson1_deserializeOpListSpaces) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5296,10 +5214,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowRuns) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5339,7 +5255,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowRuns) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5472,10 +5387,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5515,7 +5428,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5648,10 +5560,8 @@ func (m *awsRestjson1_deserializeOpStartDevEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5691,7 +5601,6 @@ func (m *awsRestjson1_deserializeOpStartDevEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5846,10 +5755,8 @@ func (m *awsRestjson1_deserializeOpStartDevEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5889,7 +5796,6 @@ func (m *awsRestjson1_deserializeOpStartDevEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6049,10 +5955,8 @@ func (m *awsRestjson1_deserializeOpStartWorkflowRun) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6092,7 +5996,6 @@ func (m *awsRestjson1_deserializeOpStartWorkflowRun) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6247,10 +6150,8 @@ func (m *awsRestjson1_deserializeOpStopDevEnvironment)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6290,7 +6191,6 @@ func (m *awsRestjson1_deserializeOpStopDevEnvironment)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6445,10 +6345,8 @@ func (m *awsRestjson1_deserializeOpStopDevEnvironmentS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6488,7 +6386,6 @@ func (m *awsRestjson1_deserializeOpStopDevEnvironmentS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6643,10 +6540,8 @@ func (m *awsRestjson1_deserializeOpUpdateDevEnvironmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6686,7 +6581,6 @@ func (m *awsRestjson1_deserializeOpUpdateDevEnvironmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6877,10 +6771,8 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6920,7 +6812,6 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7075,10 +6966,8 @@ func (m *awsRestjson1_deserializeOpUpdateSpace) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7118,7 +7007,6 @@ func (m *awsRestjson1_deserializeOpUpdateSpace) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7264,10 +7152,8 @@ func (m *awsRestjson1_deserializeOpVerifySession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7307,7 +7193,6 @@ func (m *awsRestjson1_deserializeOpVerifySession) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
