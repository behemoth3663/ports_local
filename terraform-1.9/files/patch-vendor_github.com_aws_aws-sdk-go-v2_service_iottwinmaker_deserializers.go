--- vendor/github.com/aws/aws-sdk-go-v2/service/iottwinmaker/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iottwinmaker/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchPutPropertyVal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchPutPropertyVal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -208,10 +204,8 @@ func (m *awsRestjson1_deserializeOpCancelMetadataTrans
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -251,7 +245,6 @@ func (m *awsRestjson1_deserializeOpCancelMetadataTrans
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -414,10 +407,8 @@ func (m *awsRestjson1_deserializeOpCreateComponentType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -457,7 +448,6 @@ func (m *awsRestjson1_deserializeOpCreateComponentType
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -610,10 +600,8 @@ func (m *awsRestjson1_deserializeOpCreateEntity) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -653,7 +641,6 @@ func (m *awsRestjson1_deserializeOpCreateEntity) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -815,10 +802,8 @@ func (m *awsRestjson1_deserializeOpCreateMetadataTrans
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -858,7 +843,6 @@ func (m *awsRestjson1_deserializeOpCreateMetadataTrans
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1019,10 +1003,8 @@ func (m *awsRestjson1_deserializeOpCreateScene) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1062,7 +1044,6 @@ func (m *awsRestjson1_deserializeOpCreateScene) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1206,10 +1187,8 @@ func (m *awsRestjson1_deserializeOpCreateSyncJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1249,7 +1228,6 @@ func (m *awsRestjson1_deserializeOpCreateSyncJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1402,10 +1380,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1445,7 +1421,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1589,10 +1564,8 @@ func (m *awsRestjson1_deserializeOpDeleteComponentType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1632,7 +1605,6 @@ func (m *awsRestjson1_deserializeOpDeleteComponentType
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1757,10 +1729,8 @@ func (m *awsRestjson1_deserializeOpDeleteEntity) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1800,7 +1770,6 @@ func (m *awsRestjson1_deserializeOpDeleteEntity) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1925,10 +1894,8 @@ func (m *awsRestjson1_deserializeOpDeleteScene) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1940,7 +1907,6 @@ func (m *awsRestjson1_deserializeOpDeleteScene) Handle
 	output := &DeleteSceneOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2025,10 +1991,8 @@ func (m *awsRestjson1_deserializeOpDeleteSyncJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2068,7 +2032,6 @@ func (m *awsRestjson1_deserializeOpDeleteSyncJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2196,10 +2159,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2239,7 +2200,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2364,10 +2324,8 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2407,7 +2365,6 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2545,10 +2502,8 @@ func (m *awsRestjson1_deserializeOpGetComponentType) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2588,7 +2543,6 @@ func (m *awsRestjson1_deserializeOpGetComponentType) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2847,10 +2801,8 @@ func (m *awsRestjson1_deserializeOpGetEntity) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2890,7 +2842,6 @@ func (m *awsRestjson1_deserializeOpGetEntity) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3129,10 +3080,8 @@ func (m *awsRestjson1_deserializeOpGetMetadataTransfer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3172,7 +3121,6 @@ func (m *awsRestjson1_deserializeOpGetMetadataTransfer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3385,10 +3333,8 @@ func (m *awsRestjson1_deserializeOpGetPricingPlan) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3428,7 +3374,6 @@ func (m *awsRestjson1_deserializeOpGetPricingPlan) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3551,10 +3496,8 @@ func (m *awsRestjson1_deserializeOpGetPropertyValue) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3594,7 +3537,6 @@ func (m *awsRestjson1_deserializeOpGetPropertyValue) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3735,10 +3677,8 @@ func (m *awsRestjson1_deserializeOpGetPropertyValueHis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3778,7 +3718,6 @@ func (m *awsRestjson1_deserializeOpGetPropertyValueHis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3914,10 +3853,8 @@ func (m *awsRestjson1_deserializeOpGetScene) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3957,7 +3894,6 @@ func (m *awsRestjson1_deserializeOpGetScene) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4170,10 +4106,8 @@ func (m *awsRestjson1_deserializeOpGetSyncJob) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4213,7 +4147,6 @@ func (m *awsRestjson1_deserializeOpGetSyncJob) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4405,10 +4338,8 @@ func (m *awsRestjson1_deserializeOpGetWorkspace) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4448,7 +4379,6 @@ func (m *awsRestjson1_deserializeOpGetWorkspace) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4646,10 +4576,8 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4689,7 +4617,6 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4819,10 +4746,8 @@ func (m *awsRestjson1_deserializeOpListComponentTypes)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4862,7 +4787,6 @@ func (m *awsRestjson1_deserializeOpListComponentTypes)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5011,10 +4935,8 @@ func (m *awsRestjson1_deserializeOpListEntities) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5054,7 +4976,6 @@ func (m *awsRestjson1_deserializeOpListEntities) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5181,10 +5102,8 @@ func (m *awsRestjson1_deserializeOpListMetadataTransfe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5224,7 +5143,6 @@ func (m *awsRestjson1_deserializeOpListMetadataTransfe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5351,10 +5269,8 @@ func (m *awsRestjson1_deserializeOpListProperties) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5394,7 +5310,6 @@ func (m *awsRestjson1_deserializeOpListProperties) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5524,10 +5439,8 @@ func (m *awsRestjson1_deserializeOpListScenes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5567,7 +5480,6 @@ func (m *awsRestjson1_deserializeOpListScenes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5694,10 +5606,8 @@ func (m *awsRestjson1_deserializeOpListSyncJobs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5737,7 +5647,6 @@ func (m *awsRestjson1_deserializeOpListSyncJobs) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5867,10 +5776,8 @@ func (m *awsRestjson1_deserializeOpListSyncResources) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5910,7 +5817,6 @@ func (m *awsRestjson1_deserializeOpListSyncResources) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6040,10 +5946,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6083,7 +5987,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6204,10 +6107,8 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6247,7 +6148,6 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6374,10 +6274,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6389,7 +6287,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6468,10 +6365,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6483,7 +6378,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6559,10 +6453,8 @@ func (m *awsRestjson1_deserializeOpUpdateComponentType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6602,7 +6494,6 @@ func (m *awsRestjson1_deserializeOpUpdateComponentType
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6757,10 +6648,8 @@ func (m *awsRestjson1_deserializeOpUpdateEntity) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6800,7 +6689,6 @@ func (m *awsRestjson1_deserializeOpUpdateEntity) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6947,10 +6835,8 @@ func (m *awsRestjson1_deserializeOpUpdatePricingPlan) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6990,7 +6876,6 @@ func (m *awsRestjson1_deserializeOpUpdatePricingPlan) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7113,10 +6998,8 @@ func (m *awsRestjson1_deserializeOpUpdateScene) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7156,7 +7039,6 @@ func (m *awsRestjson1_deserializeOpUpdateScene) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7288,10 +7170,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7331,7 +7211,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
