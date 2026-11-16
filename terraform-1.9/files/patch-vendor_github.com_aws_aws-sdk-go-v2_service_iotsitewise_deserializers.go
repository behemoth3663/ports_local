--- vendor/github.com/aws/aws-sdk-go-v2/service/iotsitewise/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iotsitewise/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpAssociateAssets) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -69,7 +66,6 @@ func (m *awsRestjson1_deserializeOpAssociateAssets) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -160,10 +156,8 @@ func (m *awsRestjson1_deserializeOpAssociateTimeSeries
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -181,7 +175,6 @@ func (m *awsRestjson1_deserializeOpAssociateTimeSeries
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -266,10 +259,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateProje
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -309,7 +300,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateProje
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -430,10 +420,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociatePr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -473,7 +461,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociatePr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -591,10 +578,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -634,7 +619,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -771,10 +755,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -814,7 +796,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -951,10 +932,8 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -994,7 +973,6 @@ func (m *awsRestjson1_deserializeOpBatchGetAssetProper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1131,10 +1109,8 @@ func (m *awsRestjson1_deserializeOpBatchPutAssetProper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1174,7 +1150,6 @@ func (m *awsRestjson1_deserializeOpBatchPutAssetProper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1301,10 +1276,8 @@ func (m *awsRestjson1_deserializeOpCreateAccessPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1344,7 +1317,6 @@ func (m *awsRestjson1_deserializeOpCreateAccessPolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1478,10 +1450,8 @@ func (m *awsRestjson1_deserializeOpCreateAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1521,7 +1491,6 @@ func (m *awsRestjson1_deserializeOpCreateAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1666,10 +1635,8 @@ func (m *awsRestjson1_deserializeOpCreateAssetModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1709,7 +1676,6 @@ func (m *awsRestjson1_deserializeOpCreateAssetModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1854,10 +1820,8 @@ func (m *awsRestjson1_deserializeOpCreateAssetModelCom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1897,7 +1861,6 @@ func (m *awsRestjson1_deserializeOpCreateAssetModelCom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2041,10 +2004,8 @@ func (m *awsRestjson1_deserializeOpCreateBulkImportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2084,7 +2045,6 @@ func (m *awsRestjson1_deserializeOpCreateBulkImportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2233,10 +2193,8 @@ func (m *awsRestjson1_deserializeOpCreateDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2276,7 +2234,6 @@ func (m *awsRestjson1_deserializeOpCreateDashboard) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2410,10 +2367,8 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2453,7 +2408,6 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2598,10 +2552,8 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2641,7 +2593,6 @@ func (m *awsRestjson1_deserializeOpCreateGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2775,10 +2726,8 @@ func (m *awsRestjson1_deserializeOpCreatePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2818,7 +2767,6 @@ func (m *awsRestjson1_deserializeOpCreatePortal) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2975,10 +2923,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3018,7 +2964,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3152,10 +3097,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3167,7 +3110,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessPolicy)
 	output := &DeleteAccessPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3249,10 +3191,8 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3292,7 +3232,6 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3413,10 +3352,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssetModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3456,7 +3393,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssetModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3580,10 +3516,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssetModelCom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3623,7 +3557,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssetModelCom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3747,10 +3680,8 @@ func (m *awsRestjson1_deserializeOpDeleteDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3762,7 +3693,6 @@ func (m *awsRestjson1_deserializeOpDeleteDashboard) Ha
 	output := &DeleteDashboardOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3844,10 +3774,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3887,7 +3815,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4008,10 +3935,8 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4029,7 +3954,6 @@ func (m *awsRestjson1_deserializeOpDeleteGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4114,10 +4038,8 @@ func (m *awsRestjson1_deserializeOpDeletePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4157,7 +4079,6 @@ func (m *awsRestjson1_deserializeOpDeletePortal) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4278,10 +4199,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4293,7 +4212,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 	output := &DeleteProjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4375,10 +4293,8 @@ func (m *awsRestjson1_deserializeOpDeleteTimeSeries) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4396,7 +4312,6 @@ func (m *awsRestjson1_deserializeOpDeleteTimeSeries) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4481,10 +4396,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccessPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4524,7 +4437,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccessPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4706,10 +4618,8 @@ func (m *awsRestjson1_deserializeOpDescribeAction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4749,7 +4659,6 @@ func (m *awsRestjson1_deserializeOpDescribeAction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4906,10 +4815,8 @@ func (m *awsRestjson1_deserializeOpDescribeAsset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4949,7 +4856,6 @@ func (m *awsRestjson1_deserializeOpDescribeAsset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5173,10 +5079,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetCompos
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5216,7 +5120,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetCompos
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5403,10 +5306,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetModel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5451,7 +5352,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetModel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5696,10 +5596,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetModelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5739,7 +5637,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetModelC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5931,10 +5828,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetProper
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5974,7 +5869,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetProper
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6133,10 +6027,8 @@ func (m *awsRestjson1_deserializeOpDescribeBulkImportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6176,7 +6068,6 @@ func (m *awsRestjson1_deserializeOpDescribeBulkImportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6390,10 +6281,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboard) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6433,7 +6322,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboard) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6632,10 +6520,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6675,7 +6561,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6875,10 +6760,8 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultEncr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6918,7 +6801,6 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultEncr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7051,10 +6933,8 @@ func (m *awsRestjson1_deserializeOpDescribeGateway) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7094,7 +6974,6 @@ func (m *awsRestjson1_deserializeOpDescribeGateway) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7285,10 +7164,8 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayCapa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7328,7 +7205,6 @@ func (m *awsRestjson1_deserializeOpDescribeGatewayCapa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7477,10 +7353,8 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7520,7 +7394,6 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingOpti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7638,10 +7511,8 @@ func (m *awsRestjson1_deserializeOpDescribePortal) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7681,7 +7552,6 @@ func (m *awsRestjson1_deserializeOpDescribePortal) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7945,10 +7815,8 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7988,7 +7856,6 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8178,10 +8045,8 @@ func (m *awsRestjson1_deserializeOpDescribeStorageConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8221,7 +8086,6 @@ func (m *awsRestjson1_deserializeOpDescribeStorageConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8412,10 +8276,8 @@ func (m *awsRestjson1_deserializeOpDescribeTimeSeries)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8455,7 +8317,6 @@ func (m *awsRestjson1_deserializeOpDescribeTimeSeries)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8663,10 +8524,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAssets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8684,7 +8543,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAssets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8769,10 +8627,8 @@ func (m *awsRestjson1_deserializeOpDisassociateTimeSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8790,7 +8646,6 @@ func (m *awsRestjson1_deserializeOpDisassociateTimeSer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8875,10 +8730,8 @@ func (m *awsRestjson1_deserializeOpExecuteAction) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8918,7 +8771,6 @@ func (m *awsRestjson1_deserializeOpExecuteAction) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9046,10 +8898,8 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9089,7 +8939,6 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9230,10 +9079,8 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyAgg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9273,7 +9120,6 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyAgg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9403,10 +9249,8 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyVal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9446,7 +9290,6 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyVal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9567,10 +9410,8 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyVal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9610,7 +9451,6 @@ func (m *awsRestjson1_deserializeOpGetAssetPropertyVal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9740,10 +9580,8 @@ func (m *awsRestjson1_deserializeOpGetInterpolatedAsse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9783,7 +9621,6 @@ func (m *awsRestjson1_deserializeOpGetInterpolatedAsse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9913,10 +9750,8 @@ func (m *awsRestjson1_deserializeOpInvokeAssistant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9933,7 +9768,6 @@ func (m *awsRestjson1_deserializeOpInvokeAssistant) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10037,10 +9871,8 @@ func (m *awsRestjson1_deserializeOpListAccessPolicies)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10080,7 +9912,6 @@ func (m *awsRestjson1_deserializeOpListAccessPolicies)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10204,10 +10035,8 @@ func (m *awsRestjson1_deserializeOpListActions) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10247,7 +10076,6 @@ func (m *awsRestjson1_deserializeOpListActions) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10374,10 +10202,8 @@ func (m *awsRestjson1_deserializeOpListAssetModelCompo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10417,7 +10243,6 @@ func (m *awsRestjson1_deserializeOpListAssetModelCompo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10544,10 +10369,8 @@ func (m *awsRestjson1_deserializeOpListAssetModelPrope
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10587,7 +10410,6 @@ func (m *awsRestjson1_deserializeOpListAssetModelPrope
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10714,10 +10536,8 @@ func (m *awsRestjson1_deserializeOpListAssetModels) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10757,7 +10577,6 @@ func (m *awsRestjson1_deserializeOpListAssetModels) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10881,10 +10700,8 @@ func (m *awsRestjson1_deserializeOpListAssetProperties
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10924,7 +10741,6 @@ func (m *awsRestjson1_deserializeOpListAssetProperties
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11051,10 +10867,8 @@ func (m *awsRestjson1_deserializeOpListAssetRelationsh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11094,7 +10908,6 @@ func (m *awsRestjson1_deserializeOpListAssetRelationsh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11221,10 +11034,8 @@ func (m *awsRestjson1_deserializeOpListAssets) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11264,7 +11075,6 @@ func (m *awsRestjson1_deserializeOpListAssets) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11391,10 +11201,8 @@ func (m *awsRestjson1_deserializeOpListAssociatedAsset
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11434,7 +11242,6 @@ func (m *awsRestjson1_deserializeOpListAssociatedAsset
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11561,10 +11368,8 @@ func (m *awsRestjson1_deserializeOpListBulkImportJobs)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11604,7 +11409,6 @@ func (m *awsRestjson1_deserializeOpListBulkImportJobs)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11731,10 +11535,8 @@ func (m *awsRestjson1_deserializeOpListCompositionRela
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11774,7 +11576,6 @@ func (m *awsRestjson1_deserializeOpListCompositionRela
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11901,10 +11702,8 @@ func (m *awsRestjson1_deserializeOpListDashboards) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11944,7 +11743,6 @@ func (m *awsRestjson1_deserializeOpListDashboards) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12068,10 +11866,8 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12111,7 +11907,6 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12235,10 +12030,8 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12278,7 +12071,6 @@ func (m *awsRestjson1_deserializeOpListGateways) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12402,10 +12194,8 @@ func (m *awsRestjson1_deserializeOpListPortals) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12445,7 +12235,6 @@ func (m *awsRestjson1_deserializeOpListPortals) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12569,10 +12358,8 @@ func (m *awsRestjson1_deserializeOpListProjectAssets) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12612,7 +12399,6 @@ func (m *awsRestjson1_deserializeOpListProjectAssets) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12736,10 +12522,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12779,7 +12563,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12903,10 +12686,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12946,7 +12727,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13073,10 +12853,8 @@ func (m *awsRestjson1_deserializeOpListTimeSeries) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13116,7 +12894,6 @@ func (m *awsRestjson1_deserializeOpListTimeSeries) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13243,10 +13020,8 @@ func (m *awsRestjson1_deserializeOpPutDefaultEncryptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13286,7 +13061,6 @@ func (m *awsRestjson1_deserializeOpPutDefaultEncryptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13425,10 +13199,8 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13440,7 +13212,6 @@ func (m *awsRestjson1_deserializeOpPutLoggingOptions) 
 	output := &PutLoggingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13525,10 +13296,8 @@ func (m *awsRestjson1_deserializeOpPutStorageConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13568,7 +13337,6 @@ func (m *awsRestjson1_deserializeOpPutStorageConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13746,10 +13514,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13761,7 +13527,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13855,10 +13620,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13870,7 +13633,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13961,10 +13723,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccessPolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13976,7 +13736,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccessPolicy)
 	output := &UpdateAccessPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14058,10 +13817,8 @@ func (m *awsRestjson1_deserializeOpUpdateAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14101,7 +13858,6 @@ func (m *awsRestjson1_deserializeOpUpdateAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14225,10 +13981,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssetModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14268,7 +14022,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssetModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14398,10 +14151,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssetModelCom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14441,7 +14192,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssetModelCom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14576,10 +14326,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssetProperty
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14597,7 +14345,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssetProperty
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14682,10 +14429,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14697,7 +14442,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboard) Ha
 	output := &UpdateDashboardOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14779,10 +14523,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14822,7 +14564,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14964,10 +14705,8 @@ func (m *awsRestjson1_deserializeOpUpdateGateway) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14985,7 +14724,6 @@ func (m *awsRestjson1_deserializeOpUpdateGateway) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15070,10 +14808,8 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayCapabi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15113,7 +14849,6 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayCapabi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15250,10 +14985,8 @@ func (m *awsRestjson1_deserializeOpUpdatePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15293,7 +15026,6 @@ func (m *awsRestjson1_deserializeOpUpdatePortal) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15414,10 +15146,8 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15429,7 +15159,6 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 	output := &UpdateProjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
