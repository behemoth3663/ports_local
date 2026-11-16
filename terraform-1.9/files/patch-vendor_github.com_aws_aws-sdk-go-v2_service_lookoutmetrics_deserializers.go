--- vendor/github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpActivateAnomalyDete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpActivateAnomalyDete
 	output := &ActivateAnomalyDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestjson1_deserializeOpBackTestAnomalyDete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -163,7 +157,6 @@ func (m *awsRestjson1_deserializeOpBackTestAnomalyDete
 	output := &BackTestAnomalyDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -248,10 +241,8 @@ func (m *awsRestjson1_deserializeOpCreateAlert) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -291,7 +282,6 @@ func (m *awsRestjson1_deserializeOpCreateAlert) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -422,10 +412,8 @@ func (m *awsRestjson1_deserializeOpCreateAnomalyDetect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -465,7 +453,6 @@ func (m *awsRestjson1_deserializeOpCreateAnomalyDetect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -593,10 +580,8 @@ func (m *awsRestjson1_deserializeOpCreateMetricSet) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -636,7 +621,6 @@ func (m *awsRestjson1_deserializeOpCreateMetricSet) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -767,10 +751,8 @@ func (m *awsRestjson1_deserializeOpDeactivateAnomalyDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -782,7 +764,6 @@ func (m *awsRestjson1_deserializeOpDeactivateAnomalyDe
 	output := &DeactivateAnomalyDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -870,10 +851,8 @@ func (m *awsRestjson1_deserializeOpDeleteAlert) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -885,7 +864,6 @@ func (m *awsRestjson1_deserializeOpDeleteAlert) Handle
 	output := &DeleteAlertOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -970,10 +948,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnomalyDetect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -985,7 +961,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnomalyDetect
 	output := &DeleteAnomalyDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1073,10 +1048,8 @@ func (m *awsRestjson1_deserializeOpDescribeAlert) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1116,7 +1089,6 @@ func (m *awsRestjson1_deserializeOpDescribeAlert) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1237,10 +1209,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnomalyDete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1280,7 +1250,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnomalyDete
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1410,10 +1379,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnomalyDete
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1453,7 +1420,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnomalyDete
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1669,10 +1635,8 @@ func (m *awsRestjson1_deserializeOpDescribeMetricSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1712,7 +1676,6 @@ func (m *awsRestjson1_deserializeOpDescribeMetricSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1952,10 +1915,8 @@ func (m *awsRestjson1_deserializeOpDetectMetricSetConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1995,7 +1956,6 @@ func (m *awsRestjson1_deserializeOpDetectMetricSetConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2116,10 +2076,8 @@ func (m *awsRestjson1_deserializeOpGetAnomalyGroup) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2159,7 +2117,6 @@ func (m *awsRestjson1_deserializeOpGetAnomalyGroup) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2280,10 +2237,8 @@ func (m *awsRestjson1_deserializeOpGetDataQualityMetri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2323,7 +2278,6 @@ func (m *awsRestjson1_deserializeOpGetDataQualityMetri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2444,10 +2398,8 @@ func (m *awsRestjson1_deserializeOpGetFeedback) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2487,7 +2439,6 @@ func (m *awsRestjson1_deserializeOpGetFeedback) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2617,10 +2568,8 @@ func (m *awsRestjson1_deserializeOpGetSampleData) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2660,7 +2609,6 @@ func (m *awsRestjson1_deserializeOpGetSampleData) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2786,10 +2734,8 @@ func (m *awsRestjson1_deserializeOpListAlerts) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2829,7 +2775,6 @@ func (m *awsRestjson1_deserializeOpListAlerts) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2959,10 +2904,8 @@ func (m *awsRestjson1_deserializeOpListAnomalyDetector
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3002,7 +2945,6 @@ func (m *awsRestjson1_deserializeOpListAnomalyDetector
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3132,10 +3074,8 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupRel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3175,7 +3115,6 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupRel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3305,10 +3244,8 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupSum
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3348,7 +3285,6 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupSum
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3483,10 +3419,8 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupTim
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3526,7 +3460,6 @@ func (m *awsRestjson1_deserializeOpListAnomalyGroupTim
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3679,10 +3612,8 @@ func (m *awsRestjson1_deserializeOpListMetricSets) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3722,7 +3653,6 @@ func (m *awsRestjson1_deserializeOpListMetricSets) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3852,10 +3782,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3895,7 +3823,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4010,10 +3937,8 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4025,7 +3950,6 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 	output := &PutFeedbackOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4110,10 +4034,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4125,7 +4047,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4204,10 +4125,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4219,7 +4138,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4298,10 +4216,8 @@ func (m *awsRestjson1_deserializeOpUpdateAlert) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4341,7 +4257,6 @@ func (m *awsRestjson1_deserializeOpUpdateAlert) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4466,10 +4381,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnomalyDetect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4509,7 +4422,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnomalyDetect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4634,10 +4546,8 @@ func (m *awsRestjson1_deserializeOpUpdateMetricSet) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4677,7 +4587,6 @@ func (m *awsRestjson1_deserializeOpUpdateMetricSet) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
