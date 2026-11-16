--- vendor/github.com/aws/aws-sdk-go-v2/service/m2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/m2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelBatchJobExecu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpCancelBatchJobExecu
 	output := &CancelBatchJobExecutionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -191,7 +185,6 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -341,10 +334,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSetExport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -384,7 +375,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSetExport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -515,10 +505,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSetImport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -558,7 +546,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSetImport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -689,10 +676,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -732,7 +717,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -863,10 +847,8 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -906,7 +888,6 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1034,10 +1015,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1049,7 +1028,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 	output := &DeleteApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1134,10 +1112,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplicationFr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1149,7 +1125,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplicationFr
 	output := &DeleteApplicationFromEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1237,10 +1212,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1252,7 +1225,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 	output := &DeleteEnvironmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1337,10 +1309,8 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1380,7 +1350,6 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1662,10 +1631,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1705,7 +1672,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1895,10 +1861,8 @@ func (m *awsRestjson1_deserializeOpGetBatchJobExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1938,7 +1902,6 @@ func (m *awsRestjson1_deserializeOpGetBatchJobExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2177,10 +2140,8 @@ func (m *awsRestjson1_deserializeOpGetDataSetDetails) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2220,7 +2181,6 @@ func (m *awsRestjson1_deserializeOpGetDataSetDetails) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2455,10 +2415,8 @@ func (m *awsRestjson1_deserializeOpGetDataSetExportTas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2498,7 +2456,6 @@ func (m *awsRestjson1_deserializeOpGetDataSetExportTas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2655,10 +2612,8 @@ func (m *awsRestjson1_deserializeOpGetDataSetImportTas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2698,7 +2653,6 @@ func (m *awsRestjson1_deserializeOpGetDataSetImportTas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2837,10 +2791,8 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2880,7 +2832,6 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3070,10 +3021,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3113,7 +3062,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3423,10 +3371,8 @@ func (m *awsRestjson1_deserializeOpGetSignedBluinsight
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3466,7 +3412,6 @@ func (m *awsRestjson1_deserializeOpGetSignedBluinsight
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3585,10 +3530,8 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3628,7 +3571,6 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3755,10 +3697,8 @@ func (m *awsRestjson1_deserializeOpListApplicationVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3798,7 +3738,6 @@ func (m *awsRestjson1_deserializeOpListApplicationVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3928,10 +3867,8 @@ func (m *awsRestjson1_deserializeOpListBatchJobDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3971,7 +3908,6 @@ func (m *awsRestjson1_deserializeOpListBatchJobDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4101,10 +4037,8 @@ func (m *awsRestjson1_deserializeOpListBatchJobExecuti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4144,7 +4078,6 @@ func (m *awsRestjson1_deserializeOpListBatchJobExecuti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4274,10 +4207,8 @@ func (m *awsRestjson1_deserializeOpListBatchJobRestart
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4317,7 +4248,6 @@ func (m *awsRestjson1_deserializeOpListBatchJobRestart
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4441,10 +4371,8 @@ func (m *awsRestjson1_deserializeOpListDataSetExportHi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4484,7 +4412,6 @@ func (m *awsRestjson1_deserializeOpListDataSetExportHi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4614,10 +4541,8 @@ func (m *awsRestjson1_deserializeOpListDataSetImportHi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4657,7 +4582,6 @@ func (m *awsRestjson1_deserializeOpListDataSetImportHi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4787,10 +4711,8 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4830,7 +4752,6 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4969,10 +4890,8 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5012,7 +4931,6 @@ func (m *awsRestjson1_deserializeOpListDeployments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5142,10 +5060,8 @@ func (m *awsRestjson1_deserializeOpListEngineVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5185,7 +5101,6 @@ func (m *awsRestjson1_deserializeOpListEngineVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5312,10 +5227,8 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5355,7 +5268,6 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5482,10 +5394,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5525,7 +5435,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5646,10 +5555,8 @@ func (m *awsRestjson1_deserializeOpStartApplication) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5661,7 +5568,6 @@ func (m *awsRestjson1_deserializeOpStartApplication) H
 	output := &StartApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5749,10 +5655,8 @@ func (m *awsRestjson1_deserializeOpStartBatchJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5792,7 +5696,6 @@ func (m *awsRestjson1_deserializeOpStartBatchJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5920,10 +5823,8 @@ func (m *awsRestjson1_deserializeOpStopApplication) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5935,7 +5836,6 @@ func (m *awsRestjson1_deserializeOpStopApplication) Ha
 	output := &StopApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6023,10 +5923,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6038,7 +5936,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6126,10 +6023,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6141,7 +6036,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6226,10 +6120,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6269,7 +6161,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6401,10 +6292,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6444,7 +6333,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
