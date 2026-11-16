--- vendor/github.com/aws/aws-sdk-go-v2/service/ssmincidents/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ssmincidents/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpBatchGetIncidentFin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpBatchGetIncidentFin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -213,10 +209,8 @@ func (m *awsRestjson1_deserializeOpCreateReplicationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -256,7 +250,6 @@ func (m *awsRestjson1_deserializeOpCreateReplicationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -384,10 +377,8 @@ func (m *awsRestjson1_deserializeOpCreateResponsePlan)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_deserializeOpCreateResponsePlan)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -555,10 +545,8 @@ func (m *awsRestjson1_deserializeOpCreateTimelineEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -598,7 +586,6 @@ func (m *awsRestjson1_deserializeOpCreateTimelineEvent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -735,10 +722,8 @@ func (m *awsRestjson1_deserializeOpDeleteIncidentRecor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -750,7 +735,6 @@ func (m *awsRestjson1_deserializeOpDeleteIncidentRecor
 	output := &DeleteIncidentRecordOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -832,10 +816,8 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -847,7 +829,6 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationSe
 	output := &DeleteReplicationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -932,10 +913,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -947,7 +926,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 	output := &DeleteResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1032,10 +1010,8 @@ func (m *awsRestjson1_deserializeOpDeleteResponsePlan)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1047,7 +1023,6 @@ func (m *awsRestjson1_deserializeOpDeleteResponsePlan)
 	output := &DeleteResponsePlanOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1129,10 +1104,8 @@ func (m *awsRestjson1_deserializeOpDeleteTimelineEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1144,7 +1117,6 @@ func (m *awsRestjson1_deserializeOpDeleteTimelineEvent
 	output := &DeleteTimelineEventOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1226,10 +1198,8 @@ func (m *awsRestjson1_deserializeOpGetIncidentRecord) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1269,7 +1239,6 @@ func (m *awsRestjson1_deserializeOpGetIncidentRecord) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1390,10 +1359,8 @@ func (m *awsRestjson1_deserializeOpGetReplicationSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1433,7 +1400,6 @@ func (m *awsRestjson1_deserializeOpGetReplicationSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1554,10 +1520,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicies
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1597,7 +1561,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicies
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1727,10 +1690,8 @@ func (m *awsRestjson1_deserializeOpGetResponsePlan) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1770,7 +1731,6 @@ func (m *awsRestjson1_deserializeOpGetResponsePlan) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1938,10 +1898,8 @@ func (m *awsRestjson1_deserializeOpGetTimelineEvent) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1981,7 +1939,6 @@ func (m *awsRestjson1_deserializeOpGetTimelineEvent) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2102,10 +2059,8 @@ func (m *awsRestjson1_deserializeOpListIncidentFinding
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2145,7 +2100,6 @@ func (m *awsRestjson1_deserializeOpListIncidentFinding
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2275,10 +2229,8 @@ func (m *awsRestjson1_deserializeOpListIncidentRecords
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2318,7 +2270,6 @@ func (m *awsRestjson1_deserializeOpListIncidentRecords
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2445,10 +2396,8 @@ func (m *awsRestjson1_deserializeOpListRelatedItems) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2488,7 +2437,6 @@ func (m *awsRestjson1_deserializeOpListRelatedItems) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2615,10 +2563,8 @@ func (m *awsRestjson1_deserializeOpListReplicationSets
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2658,7 +2604,6 @@ func (m *awsRestjson1_deserializeOpListReplicationSets
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2785,10 +2730,8 @@ func (m *awsRestjson1_deserializeOpListResponsePlans) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2828,7 +2771,6 @@ func (m *awsRestjson1_deserializeOpListResponsePlans) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2955,10 +2897,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2998,7 +2938,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3119,10 +3058,8 @@ func (m *awsRestjson1_deserializeOpListTimelineEvents)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3162,7 +3099,6 @@ func (m *awsRestjson1_deserializeOpListTimelineEvents)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3289,10 +3225,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3332,7 +3266,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3457,10 +3390,8 @@ func (m *awsRestjson1_deserializeOpStartIncident) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3500,7 +3431,6 @@ func (m *awsRestjson1_deserializeOpStartIncident) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3628,10 +3558,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3643,7 +3571,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3734,10 +3661,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3749,7 +3674,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3837,10 +3761,8 @@ func (m *awsRestjson1_deserializeOpUpdateDeletionProte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3852,7 +3774,6 @@ func (m *awsRestjson1_deserializeOpUpdateDeletionProte
 	output := &UpdateDeletionProtectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3937,10 +3858,8 @@ func (m *awsRestjson1_deserializeOpUpdateIncidentRecor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3952,7 +3871,6 @@ func (m *awsRestjson1_deserializeOpUpdateIncidentRecor
 	output := &UpdateIncidentRecordOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4040,10 +3958,8 @@ func (m *awsRestjson1_deserializeOpUpdateRelatedItems)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4055,7 +3971,6 @@ func (m *awsRestjson1_deserializeOpUpdateRelatedItems)
 	output := &UpdateRelatedItemsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4143,10 +4058,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4158,7 +4071,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationSe
 	output := &UpdateReplicationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4246,10 +4158,8 @@ func (m *awsRestjson1_deserializeOpUpdateResponsePlan)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4261,7 +4171,6 @@ func (m *awsRestjson1_deserializeOpUpdateResponsePlan)
 	output := &UpdateResponsePlanOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4349,10 +4258,8 @@ func (m *awsRestjson1_deserializeOpUpdateTimelineEvent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4364,7 +4271,6 @@ func (m *awsRestjson1_deserializeOpUpdateTimelineEvent
 	output := &UpdateTimelineEventOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
