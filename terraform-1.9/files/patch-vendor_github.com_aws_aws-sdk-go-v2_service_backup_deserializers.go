--- vendor/github.com/aws/aws-sdk-go-v2/service/backup/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/backup/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateBackupVaul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpAssociateBackupVaul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -151,10 +147,8 @@ func (m *awsRestjson1_deserializeOpCancelLegalHold) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -166,7 +160,6 @@ func (m *awsRestjson1_deserializeOpCancelLegalHold) Ha
 	output := &CancelLegalHoldOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -251,10 +244,8 @@ func (m *awsRestjson1_deserializeOpCreateBackupPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -294,7 +285,6 @@ func (m *awsRestjson1_deserializeOpCreateBackupPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -458,10 +448,8 @@ func (m *awsRestjson1_deserializeOpCreateBackupSelecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -501,7 +489,6 @@ func (m *awsRestjson1_deserializeOpCreateBackupSelecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -651,10 +638,8 @@ func (m *awsRestjson1_deserializeOpCreateBackupVault) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -694,7 +679,6 @@ func (m *awsRestjson1_deserializeOpCreateBackupVault) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -844,10 +828,8 @@ func (m *awsRestjson1_deserializeOpCreateFramework) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -887,7 +869,6 @@ func (m *awsRestjson1_deserializeOpCreateFramework) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1021,10 +1002,8 @@ func (m *awsRestjson1_deserializeOpCreateLegalHold) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1064,7 +1043,6 @@ func (m *awsRestjson1_deserializeOpCreateLegalHold) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1243,10 +1221,8 @@ func (m *awsRestjson1_deserializeOpCreateLogicallyAirG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1286,7 +1262,6 @@ func (m *awsRestjson1_deserializeOpCreateLogicallyAirG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1448,10 +1423,8 @@ func (m *awsRestjson1_deserializeOpCreateReportPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1491,7 +1464,6 @@ func (m *awsRestjson1_deserializeOpCreateReportPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1641,10 +1613,8 @@ func (m *awsRestjson1_deserializeOpCreateRestoreAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1684,7 +1654,6 @@ func (m *awsRestjson1_deserializeOpCreateRestoreAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1849,10 +1818,8 @@ func (m *awsRestjson1_deserializeOpCreateRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1892,7 +1859,6 @@ func (m *awsRestjson1_deserializeOpCreateRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2045,10 +2011,8 @@ func (m *awsRestjson1_deserializeOpCreateRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2088,7 +2052,6 @@ func (m *awsRestjson1_deserializeOpCreateRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2250,10 +2213,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2293,7 +2254,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2452,10 +2412,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupSelecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2473,7 +2431,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupSelecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2555,10 +2512,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVault) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2576,7 +2531,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVault) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2661,10 +2615,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2682,7 +2634,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2764,10 +2715,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultLo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2785,7 +2734,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultLo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2870,10 +2818,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultNo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2891,7 +2837,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackupVaultNo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2973,10 +2918,8 @@ func (m *awsRestjson1_deserializeOpDeleteFramework) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2994,7 +2937,6 @@ func (m *awsRestjson1_deserializeOpDeleteFramework) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3079,10 +3021,8 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryPoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3100,7 +3040,6 @@ func (m *awsRestjson1_deserializeOpDeleteRecoveryPoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3188,10 +3127,8 @@ func (m *awsRestjson1_deserializeOpDeleteReportPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3209,7 +3146,6 @@ func (m *awsRestjson1_deserializeOpDeleteReportPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3294,10 +3230,8 @@ func (m *awsRestjson1_deserializeOpDeleteRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3315,7 +3249,6 @@ func (m *awsRestjson1_deserializeOpDeleteRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3391,10 +3324,8 @@ func (m *awsRestjson1_deserializeOpDeleteRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3412,7 +3343,6 @@ func (m *awsRestjson1_deserializeOpDeleteRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3488,10 +3418,8 @@ func (m *awsRestjson1_deserializeOpDescribeBackupJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3531,7 +3459,6 @@ func (m *awsRestjson1_deserializeOpDescribeBackupJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3925,10 +3852,8 @@ func (m *awsRestjson1_deserializeOpDescribeBackupVault
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3968,7 +3893,6 @@ func (m *awsRestjson1_deserializeOpDescribeBackupVault
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4247,10 +4171,8 @@ func (m *awsRestjson1_deserializeOpDescribeCopyJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4290,7 +4212,6 @@ func (m *awsRestjson1_deserializeOpDescribeCopyJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4408,10 +4329,8 @@ func (m *awsRestjson1_deserializeOpDescribeFramework) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4451,7 +4370,6 @@ func (m *awsRestjson1_deserializeOpDescribeFramework) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4639,10 +4557,8 @@ func (m *awsRestjson1_deserializeOpDescribeGlobalSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4682,7 +4598,6 @@ func (m *awsRestjson1_deserializeOpDescribeGlobalSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4810,10 +4725,8 @@ func (m *awsRestjson1_deserializeOpDescribeProtectedRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4853,7 +4766,6 @@ func (m *awsRestjson1_deserializeOpDescribeProtectedRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5072,10 +4984,8 @@ func (m *awsRestjson1_deserializeOpDescribeRecoveryPoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5115,7 +5025,6 @@ func (m *awsRestjson1_deserializeOpDescribeRecoveryPoi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5491,10 +5400,8 @@ func (m *awsRestjson1_deserializeOpDescribeRegionSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5534,7 +5441,6 @@ func (m *awsRestjson1_deserializeOpDescribeRegionSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5648,10 +5554,8 @@ func (m *awsRestjson1_deserializeOpDescribeReportJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5691,7 +5595,6 @@ func (m *awsRestjson1_deserializeOpDescribeReportJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5806,10 +5709,8 @@ func (m *awsRestjson1_deserializeOpDescribeReportPlan)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5849,7 +5750,6 @@ func (m *awsRestjson1_deserializeOpDescribeReportPlan)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5967,10 +5867,8 @@ func (m *awsRestjson1_deserializeOpDescribeRestoreJob)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6010,7 +5908,6 @@ func (m *awsRestjson1_deserializeOpDescribeRestoreJob)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6322,10 +6219,8 @@ func (m *awsRestjson1_deserializeOpDisassociateBackupV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6343,7 +6238,6 @@ func (m *awsRestjson1_deserializeOpDisassociateBackupV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6428,10 +6322,8 @@ func (m *awsRestjson1_deserializeOpDisassociateRecover
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6449,7 +6341,6 @@ func (m *awsRestjson1_deserializeOpDisassociateRecover
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6537,10 +6428,8 @@ func (m *awsRestjson1_deserializeOpDisassociateRecover
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6558,7 +6447,6 @@ func (m *awsRestjson1_deserializeOpDisassociateRecover
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6643,10 +6531,8 @@ func (m *awsRestjson1_deserializeOpExportBackupPlanTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6686,7 +6572,6 @@ func (m *awsRestjson1_deserializeOpExportBackupPlanTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6808,10 +6693,8 @@ func (m *awsRestjson1_deserializeOpGetBackupPlan) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6851,7 +6734,6 @@ func (m *awsRestjson1_deserializeOpGetBackupPlan) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7058,10 +6940,8 @@ func (m *awsRestjson1_deserializeOpGetBackupPlanFromJS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7101,7 +6981,6 @@ func (m *awsRestjson1_deserializeOpGetBackupPlanFromJS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7222,10 +7101,8 @@ func (m *awsRestjson1_deserializeOpGetBackupPlanFromTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7265,7 +7142,6 @@ func (m *awsRestjson1_deserializeOpGetBackupPlanFromTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7383,10 +7259,8 @@ func (m *awsRestjson1_deserializeOpGetBackupSelection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7426,7 +7300,6 @@ func (m *awsRestjson1_deserializeOpGetBackupSelection)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7587,10 +7460,8 @@ func (m *awsRestjson1_deserializeOpGetBackupVaultAcces
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7630,7 +7501,6 @@ func (m *awsRestjson1_deserializeOpGetBackupVaultAcces
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7770,10 +7640,8 @@ func (m *awsRestjson1_deserializeOpGetBackupVaultNotif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7813,7 +7681,6 @@ func (m *awsRestjson1_deserializeOpGetBackupVaultNotif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7958,10 +7825,8 @@ func (m *awsRestjson1_deserializeOpGetLegalHold) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8001,7 +7866,6 @@ func (m *awsRestjson1_deserializeOpGetLegalHold) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8221,10 +8085,8 @@ func (m *awsRestjson1_deserializeOpGetRecoveryPointInd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8264,7 +8126,6 @@ func (m *awsRestjson1_deserializeOpGetRecoveryPointInd
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8483,10 +8344,8 @@ func (m *awsRestjson1_deserializeOpGetRecoveryPointRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8526,7 +8385,6 @@ func (m *awsRestjson1_deserializeOpGetRecoveryPointRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8671,10 +8529,8 @@ func (m *awsRestjson1_deserializeOpGetRestoreJobMetada
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8714,7 +8570,6 @@ func (m *awsRestjson1_deserializeOpGetRestoreJobMetada
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8841,10 +8696,8 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8884,7 +8737,6 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9002,10 +8854,8 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingPl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9045,7 +8895,6 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingPl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9157,10 +9006,8 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9200,7 +9047,6 @@ func (m *awsRestjson1_deserializeOpGetRestoreTestingSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9312,10 +9158,8 @@ func (m *awsRestjson1_deserializeOpGetSupportedResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9355,7 +9199,6 @@ func (m *awsRestjson1_deserializeOpGetSupportedResourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9464,10 +9307,8 @@ func (m *awsRestjson1_deserializeOpListBackupJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9507,7 +9348,6 @@ func (m *awsRestjson1_deserializeOpListBackupJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9628,10 +9468,8 @@ func (m *awsRestjson1_deserializeOpListBackupJobSummar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9671,7 +9509,6 @@ func (m *awsRestjson1_deserializeOpListBackupJobSummar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9801,10 +9638,8 @@ func (m *awsRestjson1_deserializeOpListBackupPlans) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9844,7 +9679,6 @@ func (m *awsRestjson1_deserializeOpListBackupPlans) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9971,10 +9805,8 @@ func (m *awsRestjson1_deserializeOpListBackupPlanTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10014,7 +9846,6 @@ func (m *awsRestjson1_deserializeOpListBackupPlanTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10141,10 +9972,8 @@ func (m *awsRestjson1_deserializeOpListBackupPlanVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10184,7 +10013,6 @@ func (m *awsRestjson1_deserializeOpListBackupPlanVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10311,10 +10139,8 @@ func (m *awsRestjson1_deserializeOpListBackupSelection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10354,7 +10180,6 @@ func (m *awsRestjson1_deserializeOpListBackupSelection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10481,10 +10306,8 @@ func (m *awsRestjson1_deserializeOpListBackupVaults) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10524,7 +10347,6 @@ func (m *awsRestjson1_deserializeOpListBackupVaults) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10651,10 +10473,8 @@ func (m *awsRestjson1_deserializeOpListCopyJobs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10694,7 +10514,6 @@ func (m *awsRestjson1_deserializeOpListCopyJobs) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10815,10 +10634,8 @@ func (m *awsRestjson1_deserializeOpListCopyJobSummarie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10858,7 +10675,6 @@ func (m *awsRestjson1_deserializeOpListCopyJobSummarie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10988,10 +10804,8 @@ func (m *awsRestjson1_deserializeOpListFrameworks) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11031,7 +10845,6 @@ func (m *awsRestjson1_deserializeOpListFrameworks) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11152,10 +10965,8 @@ func (m *awsRestjson1_deserializeOpListIndexedRecovery
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11195,7 +11006,6 @@ func (m *awsRestjson1_deserializeOpListIndexedRecovery
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11319,10 +11129,8 @@ func (m *awsRestjson1_deserializeOpListLegalHolds) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11362,7 +11170,6 @@ func (m *awsRestjson1_deserializeOpListLegalHolds) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11483,10 +11290,8 @@ func (m *awsRestjson1_deserializeOpListProtectedResour
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11526,7 +11331,6 @@ func (m *awsRestjson1_deserializeOpListProtectedResour
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11647,10 +11451,8 @@ func (m *awsRestjson1_deserializeOpListProtectedResour
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11690,7 +11492,6 @@ func (m *awsRestjson1_deserializeOpListProtectedResour
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11814,10 +11615,8 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11857,7 +11656,6 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11984,10 +11782,8 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12027,7 +11823,6 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12151,10 +11946,8 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12194,7 +11987,6 @@ func (m *awsRestjson1_deserializeOpListRecoveryPointsB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12321,10 +12113,8 @@ func (m *awsRestjson1_deserializeOpListReportJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12364,7 +12154,6 @@ func (m *awsRestjson1_deserializeOpListReportJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12488,10 +12277,8 @@ func (m *awsRestjson1_deserializeOpListReportPlans) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12531,7 +12318,6 @@ func (m *awsRestjson1_deserializeOpListReportPlans) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12652,10 +12438,8 @@ func (m *awsRestjson1_deserializeOpListRestoreAccessBa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12695,7 +12479,6 @@ func (m *awsRestjson1_deserializeOpListRestoreAccessBa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12822,10 +12605,8 @@ func (m *awsRestjson1_deserializeOpListRestoreJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12865,7 +12646,6 @@ func (m *awsRestjson1_deserializeOpListRestoreJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12992,10 +12772,8 @@ func (m *awsRestjson1_deserializeOpListRestoreJobsByPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13035,7 +12813,6 @@ func (m *awsRestjson1_deserializeOpListRestoreJobsByPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13162,10 +12939,8 @@ func (m *awsRestjson1_deserializeOpListRestoreJobSumma
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13205,7 +12980,6 @@ func (m *awsRestjson1_deserializeOpListRestoreJobSumma
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13335,10 +13109,8 @@ func (m *awsRestjson1_deserializeOpListRestoreTestingP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13378,7 +13150,6 @@ func (m *awsRestjson1_deserializeOpListRestoreTestingP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13499,10 +13270,8 @@ func (m *awsRestjson1_deserializeOpListRestoreTestingS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13542,7 +13311,6 @@ func (m *awsRestjson1_deserializeOpListRestoreTestingS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13666,10 +13434,8 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13709,7 +13475,6 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13836,10 +13601,8 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultAcces
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13857,7 +13620,6 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultAcces
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13939,10 +13701,8 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultLockC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13960,7 +13720,6 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultLockC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14045,10 +13804,8 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultNotif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14066,7 +13823,6 @@ func (m *awsRestjson1_deserializeOpPutBackupVaultNotif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14148,10 +13904,8 @@ func (m *awsRestjson1_deserializeOpPutRestoreValidatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14169,7 +13923,6 @@ func (m *awsRestjson1_deserializeOpPutRestoreValidatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14254,10 +14007,8 @@ func (m *awsRestjson1_deserializeOpRevokeRestoreAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14275,7 +14026,6 @@ func (m *awsRestjson1_deserializeOpRevokeRestoreAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14360,10 +14110,8 @@ func (m *awsRestjson1_deserializeOpStartBackupJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14403,7 +14151,6 @@ func (m *awsRestjson1_deserializeOpStartBackupJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14565,10 +14312,8 @@ func (m *awsRestjson1_deserializeOpStartCopyJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14608,7 +14353,6 @@ func (m *awsRestjson1_deserializeOpStartCopyJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14761,10 +14505,8 @@ func (m *awsRestjson1_deserializeOpStartReportJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14804,7 +14546,6 @@ func (m *awsRestjson1_deserializeOpStartReportJob) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14926,10 +14667,8 @@ func (m *awsRestjson1_deserializeOpStartRestoreJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14969,7 +14708,6 @@ func (m *awsRestjson1_deserializeOpStartRestoreJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15094,10 +14832,8 @@ func (m *awsRestjson1_deserializeOpStopBackupJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15115,7 +14851,6 @@ func (m *awsRestjson1_deserializeOpStopBackupJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15200,10 +14935,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15221,7 +14954,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15306,10 +15038,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15327,7 +15057,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15409,10 +15138,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackupPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15452,7 +15179,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackupPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15613,10 +15339,8 @@ func (m *awsRestjson1_deserializeOpUpdateFramework) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15656,7 +15380,6 @@ func (m *awsRestjson1_deserializeOpUpdateFramework) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15812,10 +15535,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15833,7 +15554,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15915,10 +15635,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryPoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15958,7 +15676,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryPoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16110,10 +15827,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryPoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16153,7 +15868,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecoveryPoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16297,10 +16011,8 @@ func (m *awsRestjson1_deserializeOpUpdateRegionSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16318,7 +16030,6 @@ func (m *awsRestjson1_deserializeOpUpdateRegionSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16397,10 +16108,8 @@ func (m *awsRestjson1_deserializeOpUpdateReportPlan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16440,7 +16149,6 @@ func (m *awsRestjson1_deserializeOpUpdateReportPlan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16590,10 +16298,8 @@ func (m *awsRestjson1_deserializeOpUpdateRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16633,7 +16339,6 @@ func (m *awsRestjson1_deserializeOpUpdateRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16799,10 +16504,8 @@ func (m *awsRestjson1_deserializeOpUpdateRestoreTestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16842,7 +16545,6 @@ func (m *awsRestjson1_deserializeOpUpdateRestoreTestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
