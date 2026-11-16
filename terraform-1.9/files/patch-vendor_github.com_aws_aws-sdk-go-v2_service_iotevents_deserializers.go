--- vendor/github.com/aws/aws-sdk-go-v2/service/iotevents/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotevents/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateAlarmModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateAlarmModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -269,10 +265,8 @@ func (m *awsRestjson1_deserializeOpCreateDetectorModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -312,7 +306,6 @@ func (m *awsRestjson1_deserializeOpCreateDetectorModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -439,10 +432,8 @@ func (m *awsRestjson1_deserializeOpCreateInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -482,7 +473,6 @@ func (m *awsRestjson1_deserializeOpCreateInput) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -603,10 +593,8 @@ func (m *awsRestjson1_deserializeOpDeleteAlarmModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -618,7 +606,6 @@ func (m *awsRestjson1_deserializeOpDeleteAlarmModel) H
 	output := &DeleteAlarmModelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -706,10 +693,8 @@ func (m *awsRestjson1_deserializeOpDeleteDetectorModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -721,7 +706,6 @@ func (m *awsRestjson1_deserializeOpDeleteDetectorModel
 	output := &DeleteDetectorModelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -809,10 +793,8 @@ func (m *awsRestjson1_deserializeOpDeleteInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -824,7 +806,6 @@ func (m *awsRestjson1_deserializeOpDeleteInput) Handle
 	output := &DeleteInputOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -912,10 +893,8 @@ func (m *awsRestjson1_deserializeOpDescribeAlarmModel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -955,7 +934,6 @@ func (m *awsRestjson1_deserializeOpDescribeAlarmModel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1208,10 +1186,8 @@ func (m *awsRestjson1_deserializeOpDescribeDetectorMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1251,7 +1227,6 @@ func (m *awsRestjson1_deserializeOpDescribeDetectorMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1372,10 +1347,8 @@ func (m *awsRestjson1_deserializeOpDescribeDetectorMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1415,7 +1388,6 @@ func (m *awsRestjson1_deserializeOpDescribeDetectorMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1540,10 +1512,8 @@ func (m *awsRestjson1_deserializeOpDescribeInput) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1583,7 +1553,6 @@ func (m *awsRestjson1_deserializeOpDescribeInput) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1704,10 +1673,8 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1747,7 +1714,6 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1871,10 +1837,8 @@ func (m *awsRestjson1_deserializeOpGetDetectorModelAna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1914,7 +1878,6 @@ func (m *awsRestjson1_deserializeOpGetDetectorModelAna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2044,10 +2007,8 @@ func (m *awsRestjson1_deserializeOpListAlarmModels) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2087,7 +2048,6 @@ func (m *awsRestjson1_deserializeOpListAlarmModels) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2214,10 +2174,8 @@ func (m *awsRestjson1_deserializeOpListAlarmModelVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2257,7 +2215,6 @@ func (m *awsRestjson1_deserializeOpListAlarmModelVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2387,10 +2344,8 @@ func (m *awsRestjson1_deserializeOpListDetectorModels)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2430,7 +2385,6 @@ func (m *awsRestjson1_deserializeOpListDetectorModels)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2557,10 +2511,8 @@ func (m *awsRestjson1_deserializeOpListDetectorModelVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2600,7 +2552,6 @@ func (m *awsRestjson1_deserializeOpListDetectorModelVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2730,10 +2681,8 @@ func (m *awsRestjson1_deserializeOpListInputRoutings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2773,7 +2722,6 @@ func (m *awsRestjson1_deserializeOpListInputRoutings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2903,10 +2851,8 @@ func (m *awsRestjson1_deserializeOpListInputs) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2946,7 +2892,6 @@ func (m *awsRestjson1_deserializeOpListInputs) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3073,10 +3018,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3116,7 +3059,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3237,10 +3179,8 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3258,7 +3198,6 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3346,10 +3285,8 @@ func (m *awsRestjson1_deserializeOpStartDetectorModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3389,7 +3326,6 @@ func (m *awsRestjson1_deserializeOpStartDetectorModelA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3514,10 +3450,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3529,7 +3463,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3617,10 +3550,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3632,7 +3563,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3717,10 +3647,8 @@ func (m *awsRestjson1_deserializeOpUpdateAlarmModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3760,7 +3688,6 @@ func (m *awsRestjson1_deserializeOpUpdateAlarmModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3938,10 +3865,8 @@ func (m *awsRestjson1_deserializeOpUpdateDetectorModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3981,7 +3906,6 @@ func (m *awsRestjson1_deserializeOpUpdateDetectorModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4105,10 +4029,8 @@ func (m *awsRestjson1_deserializeOpUpdateInput) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4148,7 +4070,6 @@ func (m *awsRestjson1_deserializeOpUpdateInput) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
