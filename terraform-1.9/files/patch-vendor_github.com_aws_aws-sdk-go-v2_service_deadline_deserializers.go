--- vendor/github.com/aws/aws-sdk-go-v2/service/deadline/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/deadline/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToFa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -63,7 +60,6 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToFa
 	output := &AssociateMemberToFarmOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -151,10 +147,8 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToFl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -166,7 +160,6 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToFl
 	output := &AssociateMemberToFleetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -254,10 +247,8 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -269,7 +260,6 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToJo
 	output := &AssociateMemberToJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -357,10 +347,8 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToQu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -372,7 +360,6 @@ func (m *awsRestjson1_deserializeOpAssociateMemberToQu
 	output := &AssociateMemberToQueueOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -460,10 +447,8 @@ func (m *awsRestjson1_deserializeOpAssumeFleetRoleForR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -503,7 +488,6 @@ func (m *awsRestjson1_deserializeOpAssumeFleetRoleForR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -624,10 +608,8 @@ func (m *awsRestjson1_deserializeOpAssumeFleetRoleForW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -667,7 +649,6 @@ func (m *awsRestjson1_deserializeOpAssumeFleetRoleForW
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -791,10 +772,8 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -834,7 +813,6 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -955,10 +933,8 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForU
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -998,7 +974,6 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForU
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1119,10 +1094,8 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForW
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1162,7 +1135,6 @@ func (m *awsRestjson1_deserializeOpAssumeQueueRoleForW
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1286,10 +1258,8 @@ func (m *awsRestjson1_deserializeOpBatchGetJobEntity) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1329,7 +1299,6 @@ func (m *awsRestjson1_deserializeOpBatchGetJobEntity) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1455,10 +1424,8 @@ func (m *awsRestjson1_deserializeOpCopyJobTemplate) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1498,7 +1465,6 @@ func (m *awsRestjson1_deserializeOpCopyJobTemplate) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1623,10 +1589,8 @@ func (m *awsRestjson1_deserializeOpCreateBudget) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1666,7 +1630,6 @@ func (m *awsRestjson1_deserializeOpCreateBudget) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1794,10 +1757,8 @@ func (m *awsRestjson1_deserializeOpCreateFarm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1837,7 +1798,6 @@ func (m *awsRestjson1_deserializeOpCreateFarm) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1965,10 +1925,8 @@ func (m *awsRestjson1_deserializeOpCreateFleet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2008,7 +1966,6 @@ func (m *awsRestjson1_deserializeOpCreateFleet) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2136,10 +2093,8 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2179,7 +2134,6 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2307,10 +2261,8 @@ func (m *awsRestjson1_deserializeOpCreateLicenseEndpoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2350,7 +2302,6 @@ func (m *awsRestjson1_deserializeOpCreateLicenseEndpoi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2478,10 +2429,8 @@ func (m *awsRestjson1_deserializeOpCreateLimit) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2521,7 +2470,6 @@ func (m *awsRestjson1_deserializeOpCreateLimit) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2649,10 +2597,8 @@ func (m *awsRestjson1_deserializeOpCreateMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2692,7 +2638,6 @@ func (m *awsRestjson1_deserializeOpCreateMonitor) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2826,10 +2771,8 @@ func (m *awsRestjson1_deserializeOpCreateQueue) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2869,7 +2812,6 @@ func (m *awsRestjson1_deserializeOpCreateQueue) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2997,10 +2939,8 @@ func (m *awsRestjson1_deserializeOpCreateQueueEnvironm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3040,7 +2980,6 @@ func (m *awsRestjson1_deserializeOpCreateQueueEnvironm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3168,10 +3107,8 @@ func (m *awsRestjson1_deserializeOpCreateQueueFleetAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3183,7 +3120,6 @@ func (m *awsRestjson1_deserializeOpCreateQueueFleetAss
 	output := &CreateQueueFleetAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3268,10 +3204,8 @@ func (m *awsRestjson1_deserializeOpCreateQueueLimitAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3283,7 +3217,6 @@ func (m *awsRestjson1_deserializeOpCreateQueueLimitAss
 	output := &CreateQueueLimitAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3368,10 +3301,8 @@ func (m *awsRestjson1_deserializeOpCreateStorageProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3411,7 +3342,6 @@ func (m *awsRestjson1_deserializeOpCreateStorageProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3539,10 +3469,8 @@ func (m *awsRestjson1_deserializeOpCreateWorker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3582,7 +3510,6 @@ func (m *awsRestjson1_deserializeOpCreateWorker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3710,10 +3637,8 @@ func (m *awsRestjson1_deserializeOpDeleteBudget) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3725,7 +3650,6 @@ func (m *awsRestjson1_deserializeOpDeleteBudget) Handl
 	output := &DeleteBudgetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3810,10 +3734,8 @@ func (m *awsRestjson1_deserializeOpDeleteFarm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3825,7 +3747,6 @@ func (m *awsRestjson1_deserializeOpDeleteFarm) HandleD
 	output := &DeleteFarmOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3910,10 +3831,8 @@ func (m *awsRestjson1_deserializeOpDeleteFleet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3925,7 +3844,6 @@ func (m *awsRestjson1_deserializeOpDeleteFleet) Handle
 	output := &DeleteFleetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4013,10 +3931,8 @@ func (m *awsRestjson1_deserializeOpDeleteLicenseEndpoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4028,7 +3944,6 @@ func (m *awsRestjson1_deserializeOpDeleteLicenseEndpoi
 	output := &DeleteLicenseEndpointOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4116,10 +4031,8 @@ func (m *awsRestjson1_deserializeOpDeleteLimit) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4131,7 +4044,6 @@ func (m *awsRestjson1_deserializeOpDeleteLimit) Handle
 	output := &DeleteLimitOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4213,10 +4125,8 @@ func (m *awsRestjson1_deserializeOpDeleteMeteredProduc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4228,7 +4138,6 @@ func (m *awsRestjson1_deserializeOpDeleteMeteredProduc
 	output := &DeleteMeteredProductOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4313,10 +4222,8 @@ func (m *awsRestjson1_deserializeOpDeleteMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4328,7 +4235,6 @@ func (m *awsRestjson1_deserializeOpDeleteMonitor) Hand
 	output := &DeleteMonitorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4413,10 +4319,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueue) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4428,7 +4332,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueue) Handle
 	output := &DeleteQueueOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4516,10 +4419,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueueEnvironm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4531,7 +4432,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueueEnvironm
 	output := &DeleteQueueEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4613,10 +4513,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueueFleetAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4628,7 +4526,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueueFleetAss
 	output := &DeleteQueueFleetAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4716,10 +4613,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueueLimitAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4731,7 +4626,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueueLimitAss
 	output := &DeleteQueueLimitAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4819,10 +4713,8 @@ func (m *awsRestjson1_deserializeOpDeleteStorageProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4834,7 +4726,6 @@ func (m *awsRestjson1_deserializeOpDeleteStorageProfil
 	output := &DeleteStorageProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4916,10 +4807,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4931,7 +4820,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorker) Handl
 	output := &DeleteWorkerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5019,10 +4907,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5034,7 +4920,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 	output := &DisassociateMemberFromFarmOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5119,10 +5004,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5134,7 +5017,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 	output := &DisassociateMemberFromFleetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5222,10 +5104,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5237,7 +5117,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 	output := &DisassociateMemberFromJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5322,10 +5201,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5337,7 +5214,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMemberF
 	output := &DisassociateMemberFromQueueOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5425,10 +5301,8 @@ func (m *awsRestjson1_deserializeOpGetBudget) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5468,7 +5342,6 @@ func (m *awsRestjson1_deserializeOpGetBudget) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5731,10 +5604,8 @@ func (m *awsRestjson1_deserializeOpGetFarm) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5774,7 +5645,6 @@ func (m *awsRestjson1_deserializeOpGetFarm) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5970,10 +5840,8 @@ func (m *awsRestjson1_deserializeOpGetFleet) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6013,7 +5881,6 @@ func (m *awsRestjson1_deserializeOpGetFleet) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6303,10 +6170,8 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6346,7 +6211,6 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6693,10 +6557,8 @@ func (m *awsRestjson1_deserializeOpGetLicenseEndpoint)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6736,7 +6598,6 @@ func (m *awsRestjson1_deserializeOpGetLicenseEndpoint)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6907,10 +6768,8 @@ func (m *awsRestjson1_deserializeOpGetLimit) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6950,7 +6809,6 @@ func (m *awsRestjson1_deserializeOpGetLimit) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7181,10 +7039,8 @@ func (m *awsRestjson1_deserializeOpGetMonitor) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7224,7 +7080,6 @@ func (m *awsRestjson1_deserializeOpGetMonitor) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7447,10 +7302,8 @@ func (m *awsRestjson1_deserializeOpGetQueue) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7490,7 +7343,6 @@ func (m *awsRestjson1_deserializeOpGetQueue) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7742,10 +7594,8 @@ func (m *awsRestjson1_deserializeOpGetQueueEnvironment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7785,7 +7635,6 @@ func (m *awsRestjson1_deserializeOpGetQueueEnvironment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7994,10 +7843,8 @@ func (m *awsRestjson1_deserializeOpGetQueueFleetAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8037,7 +7884,6 @@ func (m *awsRestjson1_deserializeOpGetQueueFleetAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8224,10 +8070,8 @@ func (m *awsRestjson1_deserializeOpGetQueueLimitAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8267,7 +8111,6 @@ func (m *awsRestjson1_deserializeOpGetQueueLimitAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8454,10 +8297,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8497,7 +8338,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8721,10 +8561,8 @@ func (m *awsRestjson1_deserializeOpGetSessionAction) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8764,7 +8602,6 @@ func (m *awsRestjson1_deserializeOpGetSessionAction) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9017,10 +8854,8 @@ func (m *awsRestjson1_deserializeOpGetSessionsStatisti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9060,7 +8895,6 @@ func (m *awsRestjson1_deserializeOpGetSessionsStatisti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9208,10 +9042,8 @@ func (m *awsRestjson1_deserializeOpGetStep) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9251,7 +9083,6 @@ func (m *awsRestjson1_deserializeOpGetStep) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9533,10 +9364,8 @@ func (m *awsRestjson1_deserializeOpGetStorageProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9576,7 +9405,6 @@ func (m *awsRestjson1_deserializeOpGetStorageProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9768,10 +9596,8 @@ func (m *awsRestjson1_deserializeOpGetStorageProfileFo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9811,7 +9637,6 @@ func (m *awsRestjson1_deserializeOpGetStorageProfileFo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9959,10 +9784,8 @@ func (m *awsRestjson1_deserializeOpGetTask) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10002,7 +9825,6 @@ func (m *awsRestjson1_deserializeOpGetTask) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10242,10 +10064,8 @@ func (m *awsRestjson1_deserializeOpGetWorker) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10285,7 +10105,6 @@ func (m *awsRestjson1_deserializeOpGetWorker) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10491,10 +10310,8 @@ func (m *awsRestjson1_deserializeOpListAvailableMetere
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10534,7 +10351,6 @@ func (m *awsRestjson1_deserializeOpListAvailableMetere
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10655,10 +10471,8 @@ func (m *awsRestjson1_deserializeOpListBudgets) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10698,7 +10512,6 @@ func (m *awsRestjson1_deserializeOpListBudgets) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10828,10 +10641,8 @@ func (m *awsRestjson1_deserializeOpListFarmMembers) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10871,7 +10682,6 @@ func (m *awsRestjson1_deserializeOpListFarmMembers) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11001,10 +10811,8 @@ func (m *awsRestjson1_deserializeOpListFarms) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11044,7 +10852,6 @@ func (m *awsRestjson1_deserializeOpListFarms) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11171,10 +10978,8 @@ func (m *awsRestjson1_deserializeOpListFleetMembers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11214,7 +11019,6 @@ func (m *awsRestjson1_deserializeOpListFleetMembers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11344,10 +11148,8 @@ func (m *awsRestjson1_deserializeOpListFleets) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11387,7 +11189,6 @@ func (m *awsRestjson1_deserializeOpListFleets) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11517,10 +11318,8 @@ func (m *awsRestjson1_deserializeOpListJobMembers) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11560,7 +11359,6 @@ func (m *awsRestjson1_deserializeOpListJobMembers) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11690,10 +11488,8 @@ func (m *awsRestjson1_deserializeOpListJobParameterDef
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11733,7 +11529,6 @@ func (m *awsRestjson1_deserializeOpListJobParameterDef
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11863,10 +11658,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11906,7 +11699,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12036,10 +11828,8 @@ func (m *awsRestjson1_deserializeOpListLicenseEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12079,7 +11869,6 @@ func (m *awsRestjson1_deserializeOpListLicenseEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12209,10 +11998,8 @@ func (m *awsRestjson1_deserializeOpListLimits) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12252,7 +12039,6 @@ func (m *awsRestjson1_deserializeOpListLimits) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12382,10 +12168,8 @@ func (m *awsRestjson1_deserializeOpListMeteredProducts
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12425,7 +12209,6 @@ func (m *awsRestjson1_deserializeOpListMeteredProducts
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12555,10 +12338,8 @@ func (m *awsRestjson1_deserializeOpListMonitors) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12598,7 +12379,6 @@ func (m *awsRestjson1_deserializeOpListMonitors) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12725,10 +12505,8 @@ func (m *awsRestjson1_deserializeOpListQueueEnvironmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12768,7 +12546,6 @@ func (m *awsRestjson1_deserializeOpListQueueEnvironmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12898,10 +12675,8 @@ func (m *awsRestjson1_deserializeOpListQueueFleetAssoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12941,7 +12716,6 @@ func (m *awsRestjson1_deserializeOpListQueueFleetAssoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13068,10 +12842,8 @@ func (m *awsRestjson1_deserializeOpListQueueLimitAssoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13111,7 +12883,6 @@ func (m *awsRestjson1_deserializeOpListQueueLimitAssoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13238,10 +13009,8 @@ func (m *awsRestjson1_deserializeOpListQueueMembers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13281,7 +13050,6 @@ func (m *awsRestjson1_deserializeOpListQueueMembers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13411,10 +13179,8 @@ func (m *awsRestjson1_deserializeOpListQueues) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13454,7 +13220,6 @@ func (m *awsRestjson1_deserializeOpListQueues) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13584,10 +13349,8 @@ func (m *awsRestjson1_deserializeOpListSessionActions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13627,7 +13390,6 @@ func (m *awsRestjson1_deserializeOpListSessionActions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13757,10 +13519,8 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13800,7 +13560,6 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13930,10 +13689,8 @@ func (m *awsRestjson1_deserializeOpListSessionsForWork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13973,7 +13730,6 @@ func (m *awsRestjson1_deserializeOpListSessionsForWork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14103,10 +13859,8 @@ func (m *awsRestjson1_deserializeOpListStepConsumers) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14146,7 +13900,6 @@ func (m *awsRestjson1_deserializeOpListStepConsumers) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14276,10 +14029,8 @@ func (m *awsRestjson1_deserializeOpListStepDependencie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14319,7 +14070,6 @@ func (m *awsRestjson1_deserializeOpListStepDependencie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14449,10 +14199,8 @@ func (m *awsRestjson1_deserializeOpListSteps) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14492,7 +14240,6 @@ func (m *awsRestjson1_deserializeOpListSteps) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14622,10 +14369,8 @@ func (m *awsRestjson1_deserializeOpListStorageProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14665,7 +14410,6 @@ func (m *awsRestjson1_deserializeOpListStorageProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14795,10 +14539,8 @@ func (m *awsRestjson1_deserializeOpListStorageProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14838,7 +14580,6 @@ func (m *awsRestjson1_deserializeOpListStorageProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14968,10 +14709,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15011,7 +14750,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15132,10 +14870,8 @@ func (m *awsRestjson1_deserializeOpListTasks) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15175,7 +14911,6 @@ func (m *awsRestjson1_deserializeOpListTasks) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15305,10 +15040,8 @@ func (m *awsRestjson1_deserializeOpListWorkers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15348,7 +15081,6 @@ func (m *awsRestjson1_deserializeOpListWorkers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15478,10 +15210,8 @@ func (m *awsRestjson1_deserializeOpPutMeteredProduct) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15493,7 +15223,6 @@ func (m *awsRestjson1_deserializeOpPutMeteredProduct) 
 	output := &PutMeteredProductOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15578,10 +15307,8 @@ func (m *awsRestjson1_deserializeOpSearchJobs) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15621,7 +15348,6 @@ func (m *awsRestjson1_deserializeOpSearchJobs) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15768,10 +15494,8 @@ func (m *awsRestjson1_deserializeOpSearchSteps) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15811,7 +15535,6 @@ func (m *awsRestjson1_deserializeOpSearchSteps) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15958,10 +15681,8 @@ func (m *awsRestjson1_deserializeOpSearchTasks) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16001,7 +15722,6 @@ func (m *awsRestjson1_deserializeOpSearchTasks) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16148,10 +15868,8 @@ func (m *awsRestjson1_deserializeOpSearchWorkers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16191,7 +15909,6 @@ func (m *awsRestjson1_deserializeOpSearchWorkers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16338,10 +16055,8 @@ func (m *awsRestjson1_deserializeOpStartSessionsStatis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16381,7 +16096,6 @@ func (m *awsRestjson1_deserializeOpStartSessionsStatis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16506,10 +16220,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16521,7 +16233,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16609,10 +16320,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16624,7 +16333,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16712,10 +16420,8 @@ func (m *awsRestjson1_deserializeOpUpdateBudget) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16727,7 +16433,6 @@ func (m *awsRestjson1_deserializeOpUpdateBudget) Handl
 	output := &UpdateBudgetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16812,10 +16517,8 @@ func (m *awsRestjson1_deserializeOpUpdateFarm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16827,7 +16530,6 @@ func (m *awsRestjson1_deserializeOpUpdateFarm) HandleD
 	output := &UpdateFarmOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16912,10 +16614,8 @@ func (m *awsRestjson1_deserializeOpUpdateFleet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16927,7 +16627,6 @@ func (m *awsRestjson1_deserializeOpUpdateFleet) Handle
 	output := &UpdateFleetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17015,10 +16714,8 @@ func (m *awsRestjson1_deserializeOpUpdateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17030,7 +16727,6 @@ func (m *awsRestjson1_deserializeOpUpdateJob) HandleDe
 	output := &UpdateJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17118,10 +16814,8 @@ func (m *awsRestjson1_deserializeOpUpdateLimit) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17133,7 +16827,6 @@ func (m *awsRestjson1_deserializeOpUpdateLimit) Handle
 	output := &UpdateLimitOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17218,10 +16911,8 @@ func (m *awsRestjson1_deserializeOpUpdateMonitor) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17233,7 +16924,6 @@ func (m *awsRestjson1_deserializeOpUpdateMonitor) Hand
 	output := &UpdateMonitorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17318,10 +17008,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueue) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17333,7 +17021,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueue) Handle
 	output := &UpdateQueueOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17418,10 +17105,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueEnvironm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17433,7 +17118,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueEnvironm
 	output := &UpdateQueueEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17518,10 +17202,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueFleetAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17533,7 +17215,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueFleetAss
 	output := &UpdateQueueFleetAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17618,10 +17299,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueueLimitAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17633,7 +17312,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueueLimitAss
 	output := &UpdateQueueLimitAssociationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17718,10 +17396,8 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17733,7 +17409,6 @@ func (m *awsRestjson1_deserializeOpUpdateSession) Hand
 	output := &UpdateSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17821,10 +17496,8 @@ func (m *awsRestjson1_deserializeOpUpdateStep) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17836,7 +17509,6 @@ func (m *awsRestjson1_deserializeOpUpdateStep) HandleD
 	output := &UpdateStepOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17924,10 +17596,8 @@ func (m *awsRestjson1_deserializeOpUpdateStorageProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17939,7 +17609,6 @@ func (m *awsRestjson1_deserializeOpUpdateStorageProfil
 	output := &UpdateStorageProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18024,10 +17693,8 @@ func (m *awsRestjson1_deserializeOpUpdateTask) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18039,7 +17706,6 @@ func (m *awsRestjson1_deserializeOpUpdateTask) HandleD
 	output := &UpdateTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18127,10 +17793,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18170,7 +17834,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18299,10 +17962,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkerSchedul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18342,7 +18003,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkerSchedul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
