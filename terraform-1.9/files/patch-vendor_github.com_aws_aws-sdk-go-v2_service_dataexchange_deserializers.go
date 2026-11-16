--- vendor/github.com/aws/aws-sdk-go-v2/service/dataexchange/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/dataexchange/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAcceptDataGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAcceptDataGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -341,10 +337,8 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -362,7 +356,6 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -447,10 +440,8 @@ func (m *awsRestjson1_deserializeOpCreateDataGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -490,7 +481,6 @@ func (m *awsRestjson1_deserializeOpCreateDataGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -756,10 +746,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -799,7 +787,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1014,10 +1001,8 @@ func (m *awsRestjson1_deserializeOpCreateEventAction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1057,7 +1042,6 @@ func (m *awsRestjson1_deserializeOpCreateEventAction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1232,10 +1216,8 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1275,7 +1257,6 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1466,10 +1447,8 @@ func (m *awsRestjson1_deserializeOpCreateRevision) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1509,7 +1488,6 @@ func (m *awsRestjson1_deserializeOpCreateRevision) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1741,10 +1719,8 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1762,7 +1738,6 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1850,10 +1825,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataGrant) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1871,7 +1844,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataGrant) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1956,10 +1928,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1977,7 +1947,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2065,10 +2034,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventAction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2086,7 +2053,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventAction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2168,10 +2134,8 @@ func (m *awsRestjson1_deserializeOpDeleteRevision) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2189,7 +2153,6 @@ func (m *awsRestjson1_deserializeOpDeleteRevision) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2277,10 +2240,8 @@ func (m *awsRestjson1_deserializeOpGetAsset) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2320,7 +2281,6 @@ func (m *awsRestjson1_deserializeOpGetAsset) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2527,10 +2487,8 @@ func (m *awsRestjson1_deserializeOpGetDataGrant) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2570,7 +2528,6 @@ func (m *awsRestjson1_deserializeOpGetDataGrant) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2833,10 +2790,8 @@ func (m *awsRestjson1_deserializeOpGetDataSet) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2876,7 +2831,6 @@ func (m *awsRestjson1_deserializeOpGetDataSet) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3088,10 +3042,8 @@ func (m *awsRestjson1_deserializeOpGetEventAction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3131,7 +3083,6 @@ func (m *awsRestjson1_deserializeOpGetEventAction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3303,10 +3254,8 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3346,7 +3295,6 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3531,10 +3479,8 @@ func (m *awsRestjson1_deserializeOpGetReceivedDataGran
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3574,7 +3520,6 @@ func (m *awsRestjson1_deserializeOpGetReceivedDataGran
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3823,10 +3768,8 @@ func (m *awsRestjson1_deserializeOpGetRevision) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3866,7 +3809,6 @@ func (m *awsRestjson1_deserializeOpGetRevision) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4095,10 +4037,8 @@ func (m *awsRestjson1_deserializeOpListDataGrants) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4138,7 +4078,6 @@ func (m *awsRestjson1_deserializeOpListDataGrants) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4268,10 +4207,8 @@ func (m *awsRestjson1_deserializeOpListDataSetRevision
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4311,7 +4248,6 @@ func (m *awsRestjson1_deserializeOpListDataSetRevision
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4438,10 +4374,8 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4481,7 +4415,6 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4608,10 +4541,8 @@ func (m *awsRestjson1_deserializeOpListEventActions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4651,7 +4582,6 @@ func (m *awsRestjson1_deserializeOpListEventActions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4778,10 +4708,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4821,7 +4749,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4948,10 +4875,8 @@ func (m *awsRestjson1_deserializeOpListReceivedDataGra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4991,7 +4916,6 @@ func (m *awsRestjson1_deserializeOpListReceivedDataGra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5121,10 +5045,8 @@ func (m *awsRestjson1_deserializeOpListRevisionAssets)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5164,7 +5086,6 @@ func (m *awsRestjson1_deserializeOpListRevisionAssets)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5291,10 +5212,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5334,7 +5253,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5440,10 +5358,8 @@ func (m *awsRestjson1_deserializeOpRevokeRevision) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5483,7 +5399,6 @@ func (m *awsRestjson1_deserializeOpRevokeRevision) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5713,10 +5628,8 @@ func (m *awsRestjson1_deserializeOpSendApiAsset) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5738,7 +5651,6 @@ func (m *awsRestjson1_deserializeOpSendApiAsset) Handl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5862,10 +5774,8 @@ func (m *awsRestjson1_deserializeOpSendDataSetNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5877,7 +5787,6 @@ func (m *awsRestjson1_deserializeOpSendDataSetNotifica
 	output := &SendDataSetNotificationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5965,10 +5874,8 @@ func (m *awsRestjson1_deserializeOpStartJob) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5980,7 +5887,6 @@ func (m *awsRestjson1_deserializeOpStartJob) HandleDes
 	output := &StartJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6068,10 +5974,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6089,7 +5993,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6159,10 +6062,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6180,7 +6081,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6250,10 +6150,8 @@ func (m *awsRestjson1_deserializeOpUpdateAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6293,7 +6191,6 @@ func (m *awsRestjson1_deserializeOpUpdateAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6506,10 +6403,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6549,7 +6444,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6759,10 +6653,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventAction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6802,7 +6694,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventAction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6972,10 +6863,8 @@ func (m *awsRestjson1_deserializeOpUpdateRevision) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7015,7 +6904,6 @@ func (m *awsRestjson1_deserializeOpUpdateRevision) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
