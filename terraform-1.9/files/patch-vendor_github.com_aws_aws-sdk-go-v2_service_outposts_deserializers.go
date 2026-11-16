--- vendor/github.com/aws/aws-sdk-go-v2/service/outposts/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/outposts/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelCapacityTask)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpCancelCapacityTask)
 	output := &CancelCapacityTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -145,10 +141,8 @@ func (m *awsRestjson1_deserializeOpCancelOrder) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -160,7 +154,6 @@ func (m *awsRestjson1_deserializeOpCancelOrder) Handle
 	output := &CancelOrderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -245,10 +238,8 @@ func (m *awsRestjson1_deserializeOpCreateOrder) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -288,7 +279,6 @@ func (m *awsRestjson1_deserializeOpCreateOrder) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -412,10 +402,8 @@ func (m *awsRestjson1_deserializeOpCreateOutpost) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -455,7 +443,6 @@ func (m *awsRestjson1_deserializeOpCreateOutpost) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -579,10 +566,8 @@ func (m *awsRestjson1_deserializeOpCreateSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -622,7 +607,6 @@ func (m *awsRestjson1_deserializeOpCreateSite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -743,10 +727,8 @@ func (m *awsRestjson1_deserializeOpDeleteOutpost) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -758,7 +740,6 @@ func (m *awsRestjson1_deserializeOpDeleteOutpost) Hand
 	output := &DeleteOutpostOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -843,10 +824,8 @@ func (m *awsRestjson1_deserializeOpDeleteSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -858,7 +837,6 @@ func (m *awsRestjson1_deserializeOpDeleteSite) HandleD
 	output := &DeleteSiteOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -943,10 +921,8 @@ func (m *awsRestjson1_deserializeOpGetCapacityTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -986,7 +962,6 @@ func (m *awsRestjson1_deserializeOpGetCapacityTask) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1225,10 +1200,8 @@ func (m *awsRestjson1_deserializeOpGetCatalogItem) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1268,7 +1241,6 @@ func (m *awsRestjson1_deserializeOpGetCatalogItem) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1386,10 +1358,8 @@ func (m *awsRestjson1_deserializeOpGetConnection) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1429,7 +1399,6 @@ func (m *awsRestjson1_deserializeOpGetConnection) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1556,10 +1525,8 @@ func (m *awsRestjson1_deserializeOpGetOrder) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1599,7 +1566,6 @@ func (m *awsRestjson1_deserializeOpGetOrder) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1714,10 +1680,8 @@ func (m *awsRestjson1_deserializeOpGetOutpost) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1757,7 +1721,6 @@ func (m *awsRestjson1_deserializeOpGetOutpost) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1875,10 +1838,8 @@ func (m *awsRestjson1_deserializeOpGetOutpostBillingIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1918,7 +1879,6 @@ func (m *awsRestjson1_deserializeOpGetOutpostBillingIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2051,10 +2011,8 @@ func (m *awsRestjson1_deserializeOpGetOutpostInstanceT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2094,7 +2052,6 @@ func (m *awsRestjson1_deserializeOpGetOutpostInstanceT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2239,10 +2196,8 @@ func (m *awsRestjson1_deserializeOpGetOutpostSupported
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2282,7 +2237,6 @@ func (m *awsRestjson1_deserializeOpGetOutpostSupported
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2409,10 +2363,8 @@ func (m *awsRestjson1_deserializeOpGetSite) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2452,7 +2404,6 @@ func (m *awsRestjson1_deserializeOpGetSite) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2570,10 +2521,8 @@ func (m *awsRestjson1_deserializeOpGetSiteAddress) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2613,7 +2562,6 @@ func (m *awsRestjson1_deserializeOpGetSiteAddress) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2749,10 +2697,8 @@ func (m *awsRestjson1_deserializeOpListAssetInstances)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2792,7 +2738,6 @@ func (m *awsRestjson1_deserializeOpListAssetInstances)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2919,10 +2864,8 @@ func (m *awsRestjson1_deserializeOpListAssets) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2962,7 +2905,6 @@ func (m *awsRestjson1_deserializeOpListAssets) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3089,10 +3031,8 @@ func (m *awsRestjson1_deserializeOpListBlockingInstanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3132,7 +3072,6 @@ func (m *awsRestjson1_deserializeOpListBlockingInstanc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3259,10 +3198,8 @@ func (m *awsRestjson1_deserializeOpListCapacityTasks) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3302,7 +3239,6 @@ func (m *awsRestjson1_deserializeOpListCapacityTasks) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3429,10 +3365,8 @@ func (m *awsRestjson1_deserializeOpListCatalogItems) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3472,7 +3406,6 @@ func (m *awsRestjson1_deserializeOpListCatalogItems) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3599,10 +3532,8 @@ func (m *awsRestjson1_deserializeOpListOrders) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3642,7 +3573,6 @@ func (m *awsRestjson1_deserializeOpListOrders) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3769,10 +3699,8 @@ func (m *awsRestjson1_deserializeOpListOutposts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3812,7 +3740,6 @@ func (m *awsRestjson1_deserializeOpListOutposts) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3936,10 +3863,8 @@ func (m *awsRestjson1_deserializeOpListSites) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3979,7 +3904,6 @@ func (m *awsRestjson1_deserializeOpListSites) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4103,10 +4027,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4146,7 +4068,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4261,10 +4182,8 @@ func (m *awsRestjson1_deserializeOpStartCapacityTask) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4304,7 +4223,6 @@ func (m *awsRestjson1_deserializeOpStartCapacityTask) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4546,10 +4464,8 @@ func (m *awsRestjson1_deserializeOpStartConnection) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4589,7 +4505,6 @@ func (m *awsRestjson1_deserializeOpStartConnection) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4720,10 +4635,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4735,7 +4648,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4814,10 +4726,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4829,7 +4739,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4908,10 +4817,8 @@ func (m *awsRestjson1_deserializeOpUpdateOutpost) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4951,7 +4858,6 @@ func (m *awsRestjson1_deserializeOpUpdateOutpost) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5072,10 +4978,8 @@ func (m *awsRestjson1_deserializeOpUpdateSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5115,7 +5019,6 @@ func (m *awsRestjson1_deserializeOpUpdateSite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5236,10 +5139,8 @@ func (m *awsRestjson1_deserializeOpUpdateSiteAddress) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5279,7 +5180,6 @@ func (m *awsRestjson1_deserializeOpUpdateSiteAddress) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5409,10 +5309,8 @@ func (m *awsRestjson1_deserializeOpUpdateSiteRackPhysi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5452,7 +5350,6 @@ func (m *awsRestjson1_deserializeOpUpdateSiteRackPhysi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
