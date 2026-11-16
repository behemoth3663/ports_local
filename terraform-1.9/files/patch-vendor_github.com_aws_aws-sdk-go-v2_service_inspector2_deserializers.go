--- vendor/github.com/aws/aws-sdk-go-v2/service/inspector2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/inspector2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociateMember) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAssociateMember) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -214,10 +210,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateCodeS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -257,7 +251,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateCodeS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -386,10 +379,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -429,7 +420,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -558,10 +548,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAccountStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -601,7 +589,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAccountStat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -727,10 +714,8 @@ func (m *awsRestjson1_deserializeOpBatchGetCodeSnippet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -770,7 +755,6 @@ func (m *awsRestjson1_deserializeOpBatchGetCodeSnippet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -893,10 +877,8 @@ func (m *awsRestjson1_deserializeOpBatchGetFindingDeta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -936,7 +918,6 @@ func (m *awsRestjson1_deserializeOpBatchGetFindingDeta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1059,10 +1040,8 @@ func (m *awsRestjson1_deserializeOpBatchGetFreeTrialIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1102,7 +1081,6 @@ func (m *awsRestjson1_deserializeOpBatchGetFreeTrialIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1225,10 +1203,8 @@ func (m *awsRestjson1_deserializeOpBatchGetMemberEc2De
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1268,7 +1244,6 @@ func (m *awsRestjson1_deserializeOpBatchGetMemberEc2De
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1391,10 +1366,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateMemberEc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1434,7 +1407,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateMemberEc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1557,10 +1529,8 @@ func (m *awsRestjson1_deserializeOpCancelFindingsRepor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1600,7 +1570,6 @@ func (m *awsRestjson1_deserializeOpCancelFindingsRepor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1725,10 +1694,8 @@ func (m *awsRestjson1_deserializeOpCancelSbomExport) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1768,7 +1735,6 @@ func (m *awsRestjson1_deserializeOpCancelSbomExport) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1893,10 +1859,8 @@ func (m *awsRestjson1_deserializeOpCreateCisScanConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1936,7 +1900,6 @@ func (m *awsRestjson1_deserializeOpCreateCisScanConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2058,10 +2021,8 @@ func (m *awsRestjson1_deserializeOpCreateCodeSecurityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2101,7 +2062,6 @@ func (m *awsRestjson1_deserializeOpCreateCodeSecurityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2247,10 +2207,8 @@ func (m *awsRestjson1_deserializeOpCreateCodeSecurityS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2290,7 +2248,6 @@ func (m *awsRestjson1_deserializeOpCreateCodeSecurityS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2418,10 +2375,8 @@ func (m *awsRestjson1_deserializeOpCreateFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2461,7 +2416,6 @@ func (m *awsRestjson1_deserializeOpCreateFilter) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2589,10 +2543,8 @@ func (m *awsRestjson1_deserializeOpCreateFindingsRepor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2632,7 +2584,6 @@ func (m *awsRestjson1_deserializeOpCreateFindingsRepor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2757,10 +2708,8 @@ func (m *awsRestjson1_deserializeOpCreateSbomExport) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2800,7 +2749,6 @@ func (m *awsRestjson1_deserializeOpCreateSbomExport) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2925,10 +2873,8 @@ func (m *awsRestjson1_deserializeOpDeleteCisScanConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2968,7 +2914,6 @@ func (m *awsRestjson1_deserializeOpDeleteCisScanConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3093,10 +3038,8 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSecurityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3136,7 +3079,6 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSecurityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3261,10 +3203,8 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSecurityS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3304,7 +3244,6 @@ func (m *awsRestjson1_deserializeOpDeleteCodeSecurityS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3429,10 +3368,8 @@ func (m *awsRestjson1_deserializeOpDeleteFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3472,7 +3409,6 @@ func (m *awsRestjson1_deserializeOpDeleteFilter) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3597,10 +3533,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3640,7 +3574,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3767,10 +3700,8 @@ func (m *awsRestjson1_deserializeOpDisable) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3810,7 +3741,6 @@ func (m *awsRestjson1_deserializeOpDisable) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3936,10 +3866,8 @@ func (m *awsRestjson1_deserializeOpDisableDelegatedAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3979,7 +3907,6 @@ func (m *awsRestjson1_deserializeOpDisableDelegatedAdm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4107,10 +4034,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMember)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4150,7 +4075,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMember)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4272,10 +4196,8 @@ func (m *awsRestjson1_deserializeOpEnable) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4315,7 +4237,6 @@ func (m *awsRestjson1_deserializeOpEnable) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4441,10 +4362,8 @@ func (m *awsRestjson1_deserializeOpEnableDelegatedAdmi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4484,7 +4403,6 @@ func (m *awsRestjson1_deserializeOpEnableDelegatedAdmi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4612,10 +4530,8 @@ func (m *awsRestjson1_deserializeOpGetCisScanReport) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4655,7 +4571,6 @@ func (m *awsRestjson1_deserializeOpGetCisScanReport) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4789,10 +4704,8 @@ func (m *awsRestjson1_deserializeOpGetCisScanResultDet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4832,7 +4745,6 @@ func (m *awsRestjson1_deserializeOpGetCisScanResultDet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4959,10 +4871,8 @@ func (m *awsRestjson1_deserializeOpGetClustersForImage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5002,7 +4912,6 @@ func (m *awsRestjson1_deserializeOpGetClustersForImage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5129,10 +5038,8 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityInte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5172,7 +5079,6 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityInte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5379,10 +5285,8 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityScan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5422,7 +5326,6 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityScan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5623,10 +5526,8 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityScan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5666,7 +5567,6 @@ func (m *awsRestjson1_deserializeOpGetCodeSecurityScan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5856,10 +5756,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguration) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5899,7 +5797,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguration) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6019,10 +5916,8 @@ func (m *awsRestjson1_deserializeOpGetDelegatedAdminAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6062,7 +5957,6 @@ func (m *awsRestjson1_deserializeOpGetDelegatedAdminAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6183,10 +6077,8 @@ func (m *awsRestjson1_deserializeOpGetEc2DeepInspectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6226,7 +6118,6 @@ func (m *awsRestjson1_deserializeOpGetEc2DeepInspectio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6367,10 +6258,8 @@ func (m *awsRestjson1_deserializeOpGetEncryptionKey) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6410,7 +6299,6 @@ func (m *awsRestjson1_deserializeOpGetEncryptionKey) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6535,10 +6423,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsReportSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6578,7 +6464,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsReportSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6740,10 +6625,8 @@ func (m *awsRestjson1_deserializeOpGetMember) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6783,7 +6666,6 @@ func (m *awsRestjson1_deserializeOpGetMember) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6904,10 +6786,8 @@ func (m *awsRestjson1_deserializeOpGetSbomExport) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6947,7 +6827,6 @@ func (m *awsRestjson1_deserializeOpGetSbomExport) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7118,10 +6997,8 @@ func (m *awsRestjson1_deserializeOpListAccountPermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7161,7 +7038,6 @@ func (m *awsRestjson1_deserializeOpListAccountPermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7288,10 +7164,8 @@ func (m *awsRestjson1_deserializeOpListCisScanConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7331,7 +7205,6 @@ func (m *awsRestjson1_deserializeOpListCisScanConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7458,10 +7331,8 @@ func (m *awsRestjson1_deserializeOpListCisScanResultsA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7501,7 +7372,6 @@ func (m *awsRestjson1_deserializeOpListCisScanResultsA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7628,10 +7498,8 @@ func (m *awsRestjson1_deserializeOpListCisScanResultsA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7671,7 +7539,6 @@ func (m *awsRestjson1_deserializeOpListCisScanResultsA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7798,10 +7665,8 @@ func (m *awsRestjson1_deserializeOpListCisScans) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7841,7 +7706,6 @@ func (m *awsRestjson1_deserializeOpListCisScans) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7968,10 +7832,8 @@ func (m *awsRestjson1_deserializeOpListCodeSecurityInt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8011,7 +7873,6 @@ func (m *awsRestjson1_deserializeOpListCodeSecurityInt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8138,10 +7999,8 @@ func (m *awsRestjson1_deserializeOpListCodeSecuritySca
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8181,7 +8040,6 @@ func (m *awsRestjson1_deserializeOpListCodeSecuritySca
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8311,10 +8169,8 @@ func (m *awsRestjson1_deserializeOpListCodeSecuritySca
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8354,7 +8210,6 @@ func (m *awsRestjson1_deserializeOpListCodeSecuritySca
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8484,10 +8339,8 @@ func (m *awsRestjson1_deserializeOpListCoverage) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8527,7 +8380,6 @@ func (m *awsRestjson1_deserializeOpListCoverage) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8651,10 +8503,8 @@ func (m *awsRestjson1_deserializeOpListCoverageStatist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8694,7 +8544,6 @@ func (m *awsRestjson1_deserializeOpListCoverageStatist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8831,10 +8680,8 @@ func (m *awsRestjson1_deserializeOpListDelegatedAdminA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8874,7 +8721,6 @@ func (m *awsRestjson1_deserializeOpListDelegatedAdminA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9001,10 +8847,8 @@ func (m *awsRestjson1_deserializeOpListFilters) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9044,7 +8888,6 @@ func (m *awsRestjson1_deserializeOpListFilters) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9171,10 +9014,8 @@ func (m *awsRestjson1_deserializeOpListFindingAggregat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9214,7 +9055,6 @@ func (m *awsRestjson1_deserializeOpListFindingAggregat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9347,10 +9187,8 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9390,7 +9228,6 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9514,10 +9351,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9557,7 +9392,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9684,10 +9518,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9727,7 +9559,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9845,10 +9676,8 @@ func (m *awsRestjson1_deserializeOpListUsageTotals) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9888,7 +9717,6 @@ func (m *awsRestjson1_deserializeOpListUsageTotals) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10015,10 +9843,8 @@ func (m *awsRestjson1_deserializeOpResetEncryptionKey)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10030,7 +9856,6 @@ func (m *awsRestjson1_deserializeOpResetEncryptionKey)
 	output := &ResetEncryptionKeyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10115,10 +9940,8 @@ func (m *awsRestjson1_deserializeOpSearchVulnerabiliti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10158,7 +9981,6 @@ func (m *awsRestjson1_deserializeOpSearchVulnerabiliti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10285,10 +10107,8 @@ func (m *awsRestjson1_deserializeOpSendCisSessionHealt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10300,7 +10120,6 @@ func (m *awsRestjson1_deserializeOpSendCisSessionHealt
 	output := &SendCisSessionHealthOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10385,10 +10204,8 @@ func (m *awsRestjson1_deserializeOpSendCisSessionTelem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10400,7 +10217,6 @@ func (m *awsRestjson1_deserializeOpSendCisSessionTelem
 	output := &SendCisSessionTelemetryOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10485,10 +10301,8 @@ func (m *awsRestjson1_deserializeOpStartCisSession) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10500,7 +10314,6 @@ func (m *awsRestjson1_deserializeOpStartCisSession) Ha
 	output := &StartCisSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10585,10 +10398,8 @@ func (m *awsRestjson1_deserializeOpStartCodeSecuritySc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10628,7 +10439,6 @@ func (m *awsRestjson1_deserializeOpStartCodeSecuritySc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10765,10 +10575,8 @@ func (m *awsRestjson1_deserializeOpStopCisSession) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10780,7 +10588,6 @@ func (m *awsRestjson1_deserializeOpStopCisSession) Han
 	output := &StopCisSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10865,10 +10672,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10880,7 +10685,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10965,10 +10769,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10980,7 +10782,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11062,10 +10863,8 @@ func (m *awsRestjson1_deserializeOpUpdateCisScanConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11105,7 +10904,6 @@ func (m *awsRestjson1_deserializeOpUpdateCisScanConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11230,10 +11028,8 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSecurityI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11273,7 +11069,6 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSecurityI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11410,10 +11205,8 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSecurityS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11453,7 +11246,6 @@ func (m *awsRestjson1_deserializeOpUpdateCodeSecurityS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11581,10 +11373,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11596,7 +11386,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 	output := &UpdateConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11678,10 +11467,8 @@ func (m *awsRestjson1_deserializeOpUpdateEc2DeepInspec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11721,7 +11508,6 @@ func (m *awsRestjson1_deserializeOpUpdateEc2DeepInspec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11862,10 +11648,8 @@ func (m *awsRestjson1_deserializeOpUpdateEncryptionKey
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11877,7 +11661,6 @@ func (m *awsRestjson1_deserializeOpUpdateEncryptionKey
 	output := &UpdateEncryptionKeyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11962,10 +11745,8 @@ func (m *awsRestjson1_deserializeOpUpdateFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12005,7 +11786,6 @@ func (m *awsRestjson1_deserializeOpUpdateFilter) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12130,10 +11910,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12173,7 +11951,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12291,10 +12068,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrgEc2DeepIns
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12306,7 +12081,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrgEc2DeepIns
 	output := &UpdateOrgEc2DeepInspectionConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
