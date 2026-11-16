--- vendor/github.com/aws/aws-sdk-go-v2/service/devopsguru/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/devopsguru/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAddNotificationChan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAddNotificationChan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -220,10 +216,8 @@ func (m *awsRestjson1_deserializeOpDeleteInsight) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -235,7 +229,6 @@ func (m *awsRestjson1_deserializeOpDeleteInsight) Hand
 	output := &DeleteInsightOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -323,10 +316,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountHeal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -366,7 +357,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountHeal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -544,10 +534,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountOver
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -587,7 +575,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountOver
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -739,10 +726,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnomaly) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -782,7 +767,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnomaly) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -908,10 +892,8 @@ func (m *awsRestjson1_deserializeOpDescribeEventSource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -951,7 +933,6 @@ func (m *awsRestjson1_deserializeOpDescribeEventSource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1069,10 +1050,8 @@ func (m *awsRestjson1_deserializeOpDescribeFeedback) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1112,7 +1091,6 @@ func (m *awsRestjson1_deserializeOpDescribeFeedback) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1233,10 +1211,8 @@ func (m *awsRestjson1_deserializeOpDescribeInsight) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1276,7 +1252,6 @@ func (m *awsRestjson1_deserializeOpDescribeInsight) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1402,10 +1377,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1445,7 +1418,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1610,10 +1582,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1653,7 +1623,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1792,10 +1761,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1835,7 +1802,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1977,10 +1943,8 @@ func (m *awsRestjson1_deserializeOpDescribeResourceCol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2020,7 +1984,6 @@ func (m *awsRestjson1_deserializeOpDescribeResourceCol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2157,10 +2120,8 @@ func (m *awsRestjson1_deserializeOpDescribeServiceInte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2200,7 +2161,6 @@ func (m *awsRestjson1_deserializeOpDescribeServiceInte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2321,10 +2281,8 @@ func (m *awsRestjson1_deserializeOpGetCostEstimation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2364,7 +2322,6 @@ func (m *awsRestjson1_deserializeOpGetCostEstimation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2547,10 +2504,8 @@ func (m *awsRestjson1_deserializeOpGetResourceCollecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2590,7 +2545,6 @@ func (m *awsRestjson1_deserializeOpGetResourceCollecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2720,10 +2674,8 @@ func (m *awsRestjson1_deserializeOpListAnomaliesForIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2763,7 +2715,6 @@ func (m *awsRestjson1_deserializeOpListAnomaliesForIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2898,10 +2849,8 @@ func (m *awsRestjson1_deserializeOpListAnomalousLogGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2941,7 +2890,6 @@ func (m *awsRestjson1_deserializeOpListAnomalousLogGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3080,10 +3028,8 @@ func (m *awsRestjson1_deserializeOpListEvents) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3123,7 +3069,6 @@ func (m *awsRestjson1_deserializeOpListEvents) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3253,10 +3198,8 @@ func (m *awsRestjson1_deserializeOpListInsights) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3296,7 +3239,6 @@ func (m *awsRestjson1_deserializeOpListInsights) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3428,10 +3370,8 @@ func (m *awsRestjson1_deserializeOpListMonitoredResour
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3471,7 +3411,6 @@ func (m *awsRestjson1_deserializeOpListMonitoredResour
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3598,10 +3537,8 @@ func (m *awsRestjson1_deserializeOpListNotificationCha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3641,7 +3578,6 @@ func (m *awsRestjson1_deserializeOpListNotificationCha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3768,10 +3704,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3811,7 +3745,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationIns
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3943,10 +3876,8 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3986,7 +3917,6 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4116,10 +4046,8 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4131,7 +4059,6 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 	output := &PutFeedbackOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4219,10 +4146,8 @@ func (m *awsRestjson1_deserializeOpRemoveNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4234,7 +4159,6 @@ func (m *awsRestjson1_deserializeOpRemoveNotificationC
 	output := &RemoveNotificationChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4322,10 +4246,8 @@ func (m *awsRestjson1_deserializeOpSearchInsights) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4365,7 +4287,6 @@ func (m *awsRestjson1_deserializeOpSearchInsights) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4497,10 +4418,8 @@ func (m *awsRestjson1_deserializeOpSearchOrganizationI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4540,7 +4459,6 @@ func (m *awsRestjson1_deserializeOpSearchOrganizationI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4672,10 +4590,8 @@ func (m *awsRestjson1_deserializeOpStartCostEstimation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4687,7 +4603,6 @@ func (m *awsRestjson1_deserializeOpStartCostEstimation
 	output := &StartCostEstimationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4775,10 +4690,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventSourcesC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4790,7 +4703,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventSourcesC
 	output := &UpdateEventSourcesConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4872,10 +4784,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceColle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4887,7 +4797,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceColle
 	output := &UpdateResourceCollectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4972,10 +4881,8 @@ func (m *awsRestjson1_deserializeOpUpdateServiceIntegr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4987,7 +4894,6 @@ func (m *awsRestjson1_deserializeOpUpdateServiceIntegr
 	output := &UpdateServiceIntegrationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
