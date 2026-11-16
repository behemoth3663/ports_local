--- vendor/github.com/aws/aws-sdk-go-v2/service/fis/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/fis/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateExperimentTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateExperimentTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -205,10 +201,8 @@ func (m *awsRestjson1_deserializeOpCreateTargetAccount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -248,7 +242,6 @@ func (m *awsRestjson1_deserializeOpCreateTargetAccount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -366,10 +359,8 @@ func (m *awsRestjson1_deserializeOpDeleteExperimentTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -409,7 +400,6 @@ func (m *awsRestjson1_deserializeOpDeleteExperimentTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -521,10 +511,8 @@ func (m *awsRestjson1_deserializeOpDeleteTargetAccount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -564,7 +552,6 @@ func (m *awsRestjson1_deserializeOpDeleteTargetAccount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -676,10 +663,8 @@ func (m *awsRestjson1_deserializeOpGetAction) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -719,7 +704,6 @@ func (m *awsRestjson1_deserializeOpGetAction) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -831,10 +815,8 @@ func (m *awsRestjson1_deserializeOpGetExperiment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -874,7 +856,6 @@ func (m *awsRestjson1_deserializeOpGetExperiment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -986,10 +967,8 @@ func (m *awsRestjson1_deserializeOpGetExperimentTarget
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1029,7 +1008,6 @@ func (m *awsRestjson1_deserializeOpGetExperimentTarget
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1141,10 +1119,8 @@ func (m *awsRestjson1_deserializeOpGetExperimentTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1184,7 +1160,6 @@ func (m *awsRestjson1_deserializeOpGetExperimentTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1296,10 +1271,8 @@ func (m *awsRestjson1_deserializeOpGetSafetyLever) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1339,7 +1312,6 @@ func (m *awsRestjson1_deserializeOpGetSafetyLever) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1448,10 +1420,8 @@ func (m *awsRestjson1_deserializeOpGetTargetAccountCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1491,7 +1461,6 @@ func (m *awsRestjson1_deserializeOpGetTargetAccountCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1603,10 +1572,8 @@ func (m *awsRestjson1_deserializeOpGetTargetResourceTy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1646,7 +1613,6 @@ func (m *awsRestjson1_deserializeOpGetTargetResourceTy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1758,10 +1724,8 @@ func (m *awsRestjson1_deserializeOpListActions) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1801,7 +1765,6 @@ func (m *awsRestjson1_deserializeOpListActions) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1919,10 +1882,8 @@ func (m *awsRestjson1_deserializeOpListExperimentResol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1962,7 +1923,6 @@ func (m *awsRestjson1_deserializeOpListExperimentResol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2083,10 +2043,8 @@ func (m *awsRestjson1_deserializeOpListExperiments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2126,7 +2084,6 @@ func (m *awsRestjson1_deserializeOpListExperiments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2244,10 +2201,8 @@ func (m *awsRestjson1_deserializeOpListExperimentTarge
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2287,7 +2242,6 @@ func (m *awsRestjson1_deserializeOpListExperimentTarge
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2408,10 +2362,8 @@ func (m *awsRestjson1_deserializeOpListExperimentTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2451,7 +2403,6 @@ func (m *awsRestjson1_deserializeOpListExperimentTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2569,10 +2520,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2612,7 +2561,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2718,10 +2666,8 @@ func (m *awsRestjson1_deserializeOpListTargetAccountCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2761,7 +2707,6 @@ func (m *awsRestjson1_deserializeOpListTargetAccountCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2882,10 +2827,8 @@ func (m *awsRestjson1_deserializeOpListTargetResourceT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2925,7 +2868,6 @@ func (m *awsRestjson1_deserializeOpListTargetResourceT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3043,10 +2985,8 @@ func (m *awsRestjson1_deserializeOpStartExperiment) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3086,7 +3026,6 @@ func (m *awsRestjson1_deserializeOpStartExperiment) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3204,10 +3143,8 @@ func (m *awsRestjson1_deserializeOpStopExperiment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3247,7 +3184,6 @@ func (m *awsRestjson1_deserializeOpStopExperiment) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3359,10 +3295,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3374,7 +3308,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3444,10 +3377,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3459,7 +3390,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3529,10 +3459,8 @@ func (m *awsRestjson1_deserializeOpUpdateExperimentTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3572,7 +3500,6 @@ func (m *awsRestjson1_deserializeOpUpdateExperimentTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3687,10 +3614,8 @@ func (m *awsRestjson1_deserializeOpUpdateSafetyLeverSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3730,7 +3655,6 @@ func (m *awsRestjson1_deserializeOpUpdateSafetyLeverSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3845,10 +3769,8 @@ func (m *awsRestjson1_deserializeOpUpdateTargetAccount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3888,7 +3810,6 @@ func (m *awsRestjson1_deserializeOpUpdateTargetAccount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
