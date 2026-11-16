--- vendor/github.com/aws/aws-sdk-go-v2/service/qapps/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/qapps/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -48,10 +47,8 @@ func (m *awsRestjson1_deserializeOpAssociateLibraryIte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -69,7 +66,6 @@ func (m *awsRestjson1_deserializeOpAssociateLibraryIte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -163,10 +159,8 @@ func (m *awsRestjson1_deserializeOpAssociateQAppWithUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -184,7 +178,6 @@ func (m *awsRestjson1_deserializeOpAssociateQAppWithUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -275,10 +268,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateCategory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -296,7 +287,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateCategory
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -387,10 +377,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteCategory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -408,7 +396,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteCategory
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -499,10 +486,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateCategory
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -520,7 +505,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateCategory
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -611,10 +595,8 @@ func (m *awsRestjson1_deserializeOpCreateLibraryItem) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -654,7 +636,6 @@ func (m *awsRestjson1_deserializeOpCreateLibraryItem) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -860,10 +841,8 @@ func (m *awsRestjson1_deserializeOpCreatePresignedUrl)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -903,7 +882,6 @@ func (m *awsRestjson1_deserializeOpCreatePresignedUrl)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1055,10 +1033,8 @@ func (m *awsRestjson1_deserializeOpCreateQApp) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1098,7 +1074,6 @@ func (m *awsRestjson1_deserializeOpCreateQApp) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1339,10 +1314,8 @@ func (m *awsRestjson1_deserializeOpDeleteLibraryItem) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1360,7 +1333,6 @@ func (m *awsRestjson1_deserializeOpDeleteLibraryItem) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1451,10 +1423,8 @@ func (m *awsRestjson1_deserializeOpDeleteQApp) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1472,7 +1442,6 @@ func (m *awsRestjson1_deserializeOpDeleteQApp) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1560,10 +1529,8 @@ func (m *awsRestjson1_deserializeOpDescribeQAppPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1603,7 +1570,6 @@ func (m *awsRestjson1_deserializeOpDescribeQAppPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1745,10 +1711,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLibrary
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1766,7 +1730,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLibrary
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1860,10 +1823,8 @@ func (m *awsRestjson1_deserializeOpDisassociateQAppFro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1881,7 +1842,6 @@ func (m *awsRestjson1_deserializeOpDisassociateQAppFro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1969,10 +1929,8 @@ func (m *awsRestjson1_deserializeOpExportQAppSessionDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2012,7 +1970,6 @@ func (m *awsRestjson1_deserializeOpExportQAppSessionDa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2168,10 +2125,8 @@ func (m *awsRestjson1_deserializeOpGetLibraryItem) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2211,7 +2166,6 @@ func (m *awsRestjson1_deserializeOpGetLibraryItem) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2463,10 +2417,8 @@ func (m *awsRestjson1_deserializeOpGetQApp) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2506,7 +2458,6 @@ func (m *awsRestjson1_deserializeOpGetQApp) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2746,10 +2697,8 @@ func (m *awsRestjson1_deserializeOpGetQAppSession) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2789,7 +2738,6 @@ func (m *awsRestjson1_deserializeOpGetQAppSession) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2987,10 +2935,8 @@ func (m *awsRestjson1_deserializeOpGetQAppSessionMetad
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3030,7 +2976,6 @@ func (m *awsRestjson1_deserializeOpGetQAppSessionMetad
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3193,10 +3138,8 @@ func (m *awsRestjson1_deserializeOpImportDocument) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3236,7 +3179,6 @@ func (m *awsRestjson1_deserializeOpImportDocument) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3370,10 +3312,8 @@ func (m *awsRestjson1_deserializeOpListCategories) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3413,7 +3353,6 @@ func (m *awsRestjson1_deserializeOpListCategories) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3537,10 +3476,8 @@ func (m *awsRestjson1_deserializeOpListLibraryItems) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3580,7 +3517,6 @@ func (m *awsRestjson1_deserializeOpListLibraryItems) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3713,10 +3649,8 @@ func (m *awsRestjson1_deserializeOpListQApps) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3756,7 +3690,6 @@ func (m *awsRestjson1_deserializeOpListQApps) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3886,10 +3819,8 @@ func (m *awsRestjson1_deserializeOpListQAppSessionData
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3929,7 +3860,6 @@ func (m *awsRestjson1_deserializeOpListQAppSessionData
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4083,10 +4013,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4126,7 +4054,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4247,10 +4174,8 @@ func (m *awsRestjson1_deserializeOpPredictQApp) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4290,7 +4215,6 @@ func (m *awsRestjson1_deserializeOpPredictQApp) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4420,10 +4344,8 @@ func (m *awsRestjson1_deserializeOpStartQAppSession) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4463,7 +4385,6 @@ func (m *awsRestjson1_deserializeOpStartQAppSession) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4603,10 +4524,8 @@ func (m *awsRestjson1_deserializeOpStopQAppSession) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4624,7 +4543,6 @@ func (m *awsRestjson1_deserializeOpStopQAppSession) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4715,10 +4633,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4730,7 +4646,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4818,10 +4733,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4833,7 +4746,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4918,10 +4830,8 @@ func (m *awsRestjson1_deserializeOpUpdateLibraryItem) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4961,7 +4871,6 @@ func (m *awsRestjson1_deserializeOpUpdateLibraryItem) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5216,10 +5125,8 @@ func (m *awsRestjson1_deserializeOpUpdateLibraryItemMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5237,7 +5144,6 @@ func (m *awsRestjson1_deserializeOpUpdateLibraryItemMe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5328,10 +5234,8 @@ func (m *awsRestjson1_deserializeOpUpdateQApp) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5371,7 +5275,6 @@ func (m *awsRestjson1_deserializeOpUpdateQApp) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5609,10 +5512,8 @@ func (m *awsRestjson1_deserializeOpUpdateQAppPermissio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5652,7 +5553,6 @@ func (m *awsRestjson1_deserializeOpUpdateQAppPermissio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5794,10 +5694,8 @@ func (m *awsRestjson1_deserializeOpUpdateQAppSession) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5837,7 +5735,6 @@ func (m *awsRestjson1_deserializeOpUpdateQAppSession) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5977,10 +5874,8 @@ func (m *awsRestjson1_deserializeOpUpdateQAppSessionMe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6020,7 +5915,6 @@ func (m *awsRestjson1_deserializeOpUpdateQAppSessionMe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
