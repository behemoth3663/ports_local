--- vendor/github.com/aws/aws-sdk-go-v2/service/lakeformation/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lakeformation/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAddLFTagsToResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpAddLFTagsToResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -211,10 +207,8 @@ func (m *awsRestjson1_deserializeOpAssumeDecoratedRole
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -254,7 +248,6 @@ func (m *awsRestjson1_deserializeOpAssumeDecoratedRole
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -413,10 +406,8 @@ func (m *awsRestjson1_deserializeOpBatchGrantPermissio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -456,7 +447,6 @@ func (m *awsRestjson1_deserializeOpBatchGrantPermissio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -568,10 +558,8 @@ func (m *awsRestjson1_deserializeOpBatchRevokePermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -611,7 +599,6 @@ func (m *awsRestjson1_deserializeOpBatchRevokePermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -723,10 +710,8 @@ func (m *awsRestjson1_deserializeOpCancelTransaction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -738,7 +723,6 @@ func (m *awsRestjson1_deserializeOpCancelTransaction) 
 	output := &CancelTransactionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -829,10 +813,8 @@ func (m *awsRestjson1_deserializeOpCommitTransaction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -872,7 +854,6 @@ func (m *awsRestjson1_deserializeOpCommitTransaction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1000,10 +981,8 @@ func (m *awsRestjson1_deserializeOpCreateDataCellsFilt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1015,7 +994,6 @@ func (m *awsRestjson1_deserializeOpCreateDataCellsFilt
 	output := &CreateDataCellsFilterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1106,10 +1084,8 @@ func (m *awsRestjson1_deserializeOpCreateLakeFormation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1149,7 +1125,6 @@ func (m *awsRestjson1_deserializeOpCreateLakeFormation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1277,10 +1252,8 @@ func (m *awsRestjson1_deserializeOpCreateLakeFormation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1292,7 +1265,6 @@ func (m *awsRestjson1_deserializeOpCreateLakeFormation
 	output := &CreateLakeFormationOptInOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1383,10 +1355,8 @@ func (m *awsRestjson1_deserializeOpCreateLFTag) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1398,7 +1368,6 @@ func (m *awsRestjson1_deserializeOpCreateLFTag) Handle
 	output := &CreateLFTagOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1486,10 +1455,8 @@ func (m *awsRestjson1_deserializeOpCreateLFTagExpressi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1501,7 +1468,6 @@ func (m *awsRestjson1_deserializeOpCreateLFTagExpressi
 	output := &CreateLFTagExpressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1589,10 +1555,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataCellsFilt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1604,7 +1568,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataCellsFilt
 	output := &DeleteDataCellsFilterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1689,10 +1652,8 @@ func (m *awsRestjson1_deserializeOpDeleteLakeFormation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1704,7 +1665,6 @@ func (m *awsRestjson1_deserializeOpDeleteLakeFormation
 	output := &DeleteLakeFormationIdentityCenterConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1792,10 +1752,8 @@ func (m *awsRestjson1_deserializeOpDeleteLakeFormation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1807,7 +1765,6 @@ func (m *awsRestjson1_deserializeOpDeleteLakeFormation
 	output := &DeleteLakeFormationOptInOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1895,10 +1852,8 @@ func (m *awsRestjson1_deserializeOpDeleteLFTag) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1910,7 +1865,6 @@ func (m *awsRestjson1_deserializeOpDeleteLFTag) Handle
 	output := &DeleteLFTagOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1995,10 +1949,8 @@ func (m *awsRestjson1_deserializeOpDeleteLFTagExpressi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2010,7 +1962,6 @@ func (m *awsRestjson1_deserializeOpDeleteLFTagExpressi
 	output := &DeleteLFTagExpressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2095,10 +2046,8 @@ func (m *awsRestjson1_deserializeOpDeleteObjectsOnCanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2110,7 +2059,6 @@ func (m *awsRestjson1_deserializeOpDeleteObjectsOnCanc
 	output := &DeleteObjectsOnCancelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2204,10 +2152,8 @@ func (m *awsRestjson1_deserializeOpDeregisterResource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2219,7 +2165,6 @@ func (m *awsRestjson1_deserializeOpDeregisterResource)
 	output := &DeregisterResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2301,10 +2246,8 @@ func (m *awsRestjson1_deserializeOpDescribeLakeFormati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2344,7 +2287,6 @@ func (m *awsRestjson1_deserializeOpDescribeLakeFormati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2506,10 +2448,8 @@ func (m *awsRestjson1_deserializeOpDescribeResource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2549,7 +2489,6 @@ func (m *awsRestjson1_deserializeOpDescribeResource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2667,10 +2606,8 @@ func (m *awsRestjson1_deserializeOpDescribeTransaction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2710,7 +2647,6 @@ func (m *awsRestjson1_deserializeOpDescribeTransaction
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2828,10 +2764,8 @@ func (m *awsRestjson1_deserializeOpExtendTransaction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2843,7 +2777,6 @@ func (m *awsRestjson1_deserializeOpExtendTransaction) 
 	output := &ExtendTransactionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2934,10 +2867,8 @@ func (m *awsRestjson1_deserializeOpGetDataCellsFilter)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2977,7 +2908,6 @@ func (m *awsRestjson1_deserializeOpGetDataCellsFilter)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3098,10 +3028,8 @@ func (m *awsRestjson1_deserializeOpGetDataLakePrincipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3141,7 +3069,6 @@ func (m *awsRestjson1_deserializeOpGetDataLakePrincipa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3260,10 +3187,8 @@ func (m *awsRestjson1_deserializeOpGetDataLakeSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3303,7 +3228,6 @@ func (m *awsRestjson1_deserializeOpGetDataLakeSettings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3418,10 +3342,8 @@ func (m *awsRestjson1_deserializeOpGetEffectivePermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3461,7 +3383,6 @@ func (m *awsRestjson1_deserializeOpGetEffectivePermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3588,10 +3509,8 @@ func (m *awsRestjson1_deserializeOpGetLFTag) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3631,7 +3550,6 @@ func (m *awsRestjson1_deserializeOpGetLFTag) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3770,10 +3688,8 @@ func (m *awsRestjson1_deserializeOpGetLFTagExpression)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3813,7 +3729,6 @@ func (m *awsRestjson1_deserializeOpGetLFTagExpression)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3961,10 +3876,8 @@ func (m *awsRestjson1_deserializeOpGetQueryState) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4004,7 +3917,6 @@ func (m *awsRestjson1_deserializeOpGetQueryState) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4132,10 +4044,8 @@ func (m *awsRestjson1_deserializeOpGetQueryStatistics)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4175,7 +4085,6 @@ func (m *awsRestjson1_deserializeOpGetQueryStatistics)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4317,10 +4226,8 @@ func (m *awsRestjson1_deserializeOpGetResourceLFTags) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4360,7 +4267,6 @@ func (m *awsRestjson1_deserializeOpGetResourceLFTags) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4494,10 +4400,8 @@ func (m *awsRestjson1_deserializeOpGetTableObjects) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4537,7 +4441,6 @@ func (m *awsRestjson1_deserializeOpGetTableObjects) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4673,10 +4576,8 @@ func (m *awsRestjson1_deserializeOpGetTemporaryGluePar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4716,7 +4617,6 @@ func (m *awsRestjson1_deserializeOpGetTemporaryGluePar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4878,10 +4778,8 @@ func (m *awsRestjson1_deserializeOpGetTemporaryGlueTab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4921,7 +4819,6 @@ func (m *awsRestjson1_deserializeOpGetTemporaryGlueTab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5088,10 +4985,8 @@ func (m *awsRestjson1_deserializeOpGetWorkUnitResults)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5108,7 +5003,6 @@ func (m *awsRestjson1_deserializeOpGetWorkUnitResults)
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5202,10 +5096,8 @@ func (m *awsRestjson1_deserializeOpGetWorkUnits) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5245,7 +5137,6 @@ func (m *awsRestjson1_deserializeOpGetWorkUnits) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5384,10 +5275,8 @@ func (m *awsRestjson1_deserializeOpGrantPermissions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5399,7 +5288,6 @@ func (m *awsRestjson1_deserializeOpGrantPermissions) H
 	output := &GrantPermissionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5478,10 +5366,8 @@ func (m *awsRestjson1_deserializeOpListDataCellsFilter
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5521,7 +5407,6 @@ func (m *awsRestjson1_deserializeOpListDataCellsFilter
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5648,10 +5533,8 @@ func (m *awsRestjson1_deserializeOpListLakeFormationOp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5691,7 +5574,6 @@ func (m *awsRestjson1_deserializeOpListLakeFormationOp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5818,10 +5700,8 @@ func (m *awsRestjson1_deserializeOpListLFTagExpression
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5861,7 +5741,6 @@ func (m *awsRestjson1_deserializeOpListLFTagExpression
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5991,10 +5870,8 @@ func (m *awsRestjson1_deserializeOpListLFTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6034,7 +5911,6 @@ func (m *awsRestjson1_deserializeOpListLFTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6164,10 +6040,8 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6207,7 +6081,6 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6331,10 +6204,8 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6374,7 +6245,6 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6498,10 +6368,8 @@ func (m *awsRestjson1_deserializeOpListTableStorageOpt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6541,7 +6409,6 @@ func (m *awsRestjson1_deserializeOpListTableStorageOpt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6668,10 +6535,8 @@ func (m *awsRestjson1_deserializeOpListTransactions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6711,7 +6576,6 @@ func (m *awsRestjson1_deserializeOpListTransactions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6835,10 +6699,8 @@ func (m *awsRestjson1_deserializeOpPutDataLakeSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6850,7 +6712,6 @@ func (m *awsRestjson1_deserializeOpPutDataLakeSettings
 	output := &PutDataLakeSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6926,10 +6787,8 @@ func (m *awsRestjson1_deserializeOpRegisterResource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6941,7 +6800,6 @@ func (m *awsRestjson1_deserializeOpRegisterResource) H
 	output := &RegisterResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7032,10 +6890,8 @@ func (m *awsRestjson1_deserializeOpRemoveLFTagsFromRes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7075,7 +6931,6 @@ func (m *awsRestjson1_deserializeOpRemoveLFTagsFromRes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7202,10 +7057,8 @@ func (m *awsRestjson1_deserializeOpRevokePermissions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7217,7 +7070,6 @@ func (m *awsRestjson1_deserializeOpRevokePermissions) 
 	output := &RevokePermissionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7296,10 +7148,8 @@ func (m *awsRestjson1_deserializeOpSearchDatabasesByLF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7339,7 +7189,6 @@ func (m *awsRestjson1_deserializeOpSearchDatabasesByLF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7472,10 +7321,8 @@ func (m *awsRestjson1_deserializeOpSearchTablesByLFTag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7515,7 +7362,6 @@ func (m *awsRestjson1_deserializeOpSearchTablesByLFTag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7648,10 +7494,8 @@ func (m *awsRestjson1_deserializeOpStartQueryPlanning)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7691,7 +7535,6 @@ func (m *awsRestjson1_deserializeOpStartQueryPlanning)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7813,10 +7656,8 @@ func (m *awsRestjson1_deserializeOpStartTransaction) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7856,7 +7697,6 @@ func (m *awsRestjson1_deserializeOpStartTransaction) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7972,10 +7812,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataCellsFilt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7987,7 +7825,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataCellsFilt
 	output := &UpdateDataCellsFilterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8075,10 +7912,8 @@ func (m *awsRestjson1_deserializeOpUpdateLakeFormation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8090,7 +7925,6 @@ func (m *awsRestjson1_deserializeOpUpdateLakeFormation
 	output := &UpdateLakeFormationIdentityCenterConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8178,10 +8012,8 @@ func (m *awsRestjson1_deserializeOpUpdateLFTag) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8193,7 +8025,6 @@ func (m *awsRestjson1_deserializeOpUpdateLFTag) Handle
 	output := &UpdateLFTagOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8281,10 +8112,8 @@ func (m *awsRestjson1_deserializeOpUpdateLFTagExpressi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8296,7 +8125,6 @@ func (m *awsRestjson1_deserializeOpUpdateLFTagExpressi
 	output := &UpdateLFTagExpressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8384,10 +8212,8 @@ func (m *awsRestjson1_deserializeOpUpdateResource) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8399,7 +8225,6 @@ func (m *awsRestjson1_deserializeOpUpdateResource) Han
 	output := &UpdateResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8481,10 +8306,8 @@ func (m *awsRestjson1_deserializeOpUpdateTableObjects)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8496,7 +8319,6 @@ func (m *awsRestjson1_deserializeOpUpdateTableObjects)
 	output := &UpdateTableObjectsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8593,10 +8415,8 @@ func (m *awsRestjson1_deserializeOpUpdateTableStorageO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8636,7 +8456,6 @@ func (m *awsRestjson1_deserializeOpUpdateTableStorageO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
