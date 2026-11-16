--- vendor/github.com/aws/aws-sdk-go-v2/service/workdocs/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/workdocs/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAbortDocumentVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpAbortDocumentVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -157,10 +153,8 @@ func (m *awsRestjson1_deserializeOpActivateUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -200,7 +194,6 @@ func (m *awsRestjson1_deserializeOpActivateUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -321,10 +314,8 @@ func (m *awsRestjson1_deserializeOpAddResourcePermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -364,7 +355,6 @@ func (m *awsRestjson1_deserializeOpAddResourcePermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -485,10 +475,8 @@ func (m *awsRestjson1_deserializeOpCreateComment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -528,7 +516,6 @@ func (m *awsRestjson1_deserializeOpCreateComment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -658,10 +645,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomMetadat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -673,7 +658,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomMetadat
 	output := &CreateCustomMetadataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -764,10 +748,8 @@ func (m *awsRestjson1_deserializeOpCreateFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -807,7 +789,6 @@ func (m *awsRestjson1_deserializeOpCreateFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -943,10 +924,8 @@ func (m *awsRestjson1_deserializeOpCreateLabels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -958,7 +937,6 @@ func (m *awsRestjson1_deserializeOpCreateLabels) Handl
 	output := &CreateLabelsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1046,10 +1024,8 @@ func (m *awsRestjson1_deserializeOpCreateNotificationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1089,7 +1065,6 @@ func (m *awsRestjson1_deserializeOpCreateNotificationS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1207,10 +1182,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1250,7 +1223,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1371,10 +1343,8 @@ func (m *awsRestjson1_deserializeOpDeactivateUser) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1392,7 +1362,6 @@ func (m *awsRestjson1_deserializeOpDeactivateUser) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1477,10 +1446,8 @@ func (m *awsRestjson1_deserializeOpDeleteComment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1498,7 +1465,6 @@ func (m *awsRestjson1_deserializeOpDeleteComment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1589,10 +1555,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomMetadat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1604,7 +1568,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomMetadat
 	output := &DeleteCustomMetadataOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1692,10 +1655,8 @@ func (m *awsRestjson1_deserializeOpDeleteDocument) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1713,7 +1674,6 @@ func (m *awsRestjson1_deserializeOpDeleteDocument) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1810,10 +1770,8 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1831,7 +1789,6 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1925,10 +1882,8 @@ func (m *awsRestjson1_deserializeOpDeleteFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1946,7 +1901,6 @@ func (m *awsRestjson1_deserializeOpDeleteFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2043,10 +1997,8 @@ func (m *awsRestjson1_deserializeOpDeleteFolderContent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2064,7 +2016,6 @@ func (m *awsRestjson1_deserializeOpDeleteFolderContent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2155,10 +2106,8 @@ func (m *awsRestjson1_deserializeOpDeleteLabels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2170,7 +2119,6 @@ func (m *awsRestjson1_deserializeOpDeleteLabels) Handl
 	output := &DeleteLabelsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2258,10 +2206,8 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2279,7 +2225,6 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2361,10 +2306,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2382,7 +2325,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2467,10 +2409,8 @@ func (m *awsRestjson1_deserializeOpDescribeActivities)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2510,7 +2450,6 @@ func (m *awsRestjson1_deserializeOpDescribeActivities)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2640,10 +2579,8 @@ func (m *awsRestjson1_deserializeOpDescribeComments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2683,7 +2620,6 @@ func (m *awsRestjson1_deserializeOpDescribeComments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2816,10 +2752,8 @@ func (m *awsRestjson1_deserializeOpDescribeDocumentVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2859,7 +2793,6 @@ func (m *awsRestjson1_deserializeOpDescribeDocumentVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2998,10 +2931,8 @@ func (m *awsRestjson1_deserializeOpDescribeFolderConte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3041,7 +2972,6 @@ func (m *awsRestjson1_deserializeOpDescribeFolderConte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3179,10 +3109,8 @@ func (m *awsRestjson1_deserializeOpDescribeGroups) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3222,7 +3150,6 @@ func (m *awsRestjson1_deserializeOpDescribeGroups) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3349,10 +3276,8 @@ func (m *awsRestjson1_deserializeOpDescribeNotificatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3392,7 +3317,6 @@ func (m *awsRestjson1_deserializeOpDescribeNotificatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3516,10 +3440,8 @@ func (m *awsRestjson1_deserializeOpDescribeResourcePer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3559,7 +3481,6 @@ func (m *awsRestjson1_deserializeOpDescribeResourcePer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3689,10 +3610,8 @@ func (m *awsRestjson1_deserializeOpDescribeRootFolders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3732,7 +3651,6 @@ func (m *awsRestjson1_deserializeOpDescribeRootFolders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3862,10 +3780,8 @@ func (m *awsRestjson1_deserializeOpDescribeUsers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3905,7 +3821,6 @@ func (m *awsRestjson1_deserializeOpDescribeUsers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3969,8 @@ func (m *awsRestjson1_deserializeOpGetCurrentUser) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4097,7 +4010,6 @@ func (m *awsRestjson1_deserializeOpGetCurrentUser) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4218,10 +4130,8 @@ func (m *awsRestjson1_deserializeOpGetDocument) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4261,7 +4171,6 @@ func (m *awsRestjson1_deserializeOpGetDocument) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4393,10 +4302,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentPath) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4436,7 +4343,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentPath) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4557,10 +4463,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4600,7 +4504,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4732,10 +4635,8 @@ func (m *awsRestjson1_deserializeOpGetFolder) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4775,7 +4676,6 @@ func (m *awsRestjson1_deserializeOpGetFolder) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4907,10 +4807,8 @@ func (m *awsRestjson1_deserializeOpGetFolderPath) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4950,7 +4848,6 @@ func (m *awsRestjson1_deserializeOpGetFolderPath) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5071,10 +4968,8 @@ func (m *awsRestjson1_deserializeOpGetResources) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5114,7 +5009,6 @@ func (m *awsRestjson1_deserializeOpGetResources) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5249,10 +5143,8 @@ func (m *awsRestjson1_deserializeOpInitiateDocumentVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5292,7 +5184,6 @@ func (m *awsRestjson1_deserializeOpInitiateDocumentVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5445,10 +5336,8 @@ func (m *awsRestjson1_deserializeOpRemoveAllResourcePe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5466,7 +5355,6 @@ func (m *awsRestjson1_deserializeOpRemoveAllResourcePe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5548,10 +5436,8 @@ func (m *awsRestjson1_deserializeOpRemoveResourcePermi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5569,7 +5455,6 @@ func (m *awsRestjson1_deserializeOpRemoveResourcePermi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5651,10 +5536,8 @@ func (m *awsRestjson1_deserializeOpRestoreDocumentVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5672,7 +5555,6 @@ func (m *awsRestjson1_deserializeOpRestoreDocumentVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5766,10 +5648,8 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5809,7 +5689,6 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5936,10 +5815,8 @@ func (m *awsRestjson1_deserializeOpUpdateDocument) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5957,7 +5834,6 @@ func (m *awsRestjson1_deserializeOpUpdateDocument) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6057,10 +5933,8 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6078,7 +5952,6 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6172,10 +6045,8 @@ func (m *awsRestjson1_deserializeOpUpdateFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6193,7 +6064,6 @@ func (m *awsRestjson1_deserializeOpUpdateFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6293,10 +6163,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6336,7 +6204,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
