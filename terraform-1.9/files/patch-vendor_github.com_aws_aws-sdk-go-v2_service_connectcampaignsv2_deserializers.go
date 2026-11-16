--- vendor/github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connectcampaignsv2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -234,10 +230,8 @@ func (m *awsRestjson1_deserializeOpDeleteCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpDeleteCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -337,10 +330,8 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignChann
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -358,7 +349,6 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignChann
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -443,10 +433,8 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignCommu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -464,7 +452,6 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignCommu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +539,8 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignCommu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -573,7 +558,6 @@ func (m *awsRestjson1_deserializeOpDeleteCampaignCommu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -661,10 +645,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -682,7 +664,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -770,10 +751,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -791,7 +770,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -876,10 +854,8 @@ func (m *awsRestjson1_deserializeOpDeleteInstanceOnboa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -897,7 +873,6 @@ func (m *awsRestjson1_deserializeOpDeleteInstanceOnboa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -982,10 +957,8 @@ func (m *awsRestjson1_deserializeOpDescribeCampaign) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1025,7 +998,6 @@ func (m *awsRestjson1_deserializeOpDescribeCampaign) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1143,10 +1115,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignState) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1186,7 +1156,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignState) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1311,10 +1280,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignStateBat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1354,7 +1321,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignStateBat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1477,10 +1443,8 @@ func (m *awsRestjson1_deserializeOpGetConnectInstanceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1520,7 +1484,6 @@ func (m *awsRestjson1_deserializeOpGetConnectInstanceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1638,10 +1601,8 @@ func (m *awsRestjson1_deserializeOpGetInstanceCommunic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1681,7 +1642,6 @@ func (m *awsRestjson1_deserializeOpGetInstanceCommunic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1799,10 +1759,8 @@ func (m *awsRestjson1_deserializeOpGetInstanceOnboardi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1842,7 +1800,6 @@ func (m *awsRestjson1_deserializeOpGetInstanceOnboardi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1960,10 +1917,8 @@ func (m *awsRestjson1_deserializeOpListCampaigns) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2003,7 +1958,6 @@ func (m *awsRestjson1_deserializeOpListCampaigns) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2127,10 +2081,8 @@ func (m *awsRestjson1_deserializeOpListConnectInstance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2170,7 +2122,6 @@ func (m *awsRestjson1_deserializeOpListConnectInstance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2300,10 +2251,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2343,7 +2292,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2464,10 +2412,8 @@ func (m *awsRestjson1_deserializeOpPauseCampaign) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2485,7 +2431,6 @@ func (m *awsRestjson1_deserializeOpPauseCampaign) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2576,10 +2521,8 @@ func (m *awsRestjson1_deserializeOpPutConnectInstanceI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2597,7 +2540,6 @@ func (m *awsRestjson1_deserializeOpPutConnectInstanceI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2685,10 +2627,8 @@ func (m *awsRestjson1_deserializeOpPutInstanceCommunic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2706,7 +2646,6 @@ func (m *awsRestjson1_deserializeOpPutInstanceCommunic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2791,10 +2730,8 @@ func (m *awsRestjson1_deserializeOpPutOutboundRequestB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2834,7 +2771,6 @@ func (m *awsRestjson1_deserializeOpPutOutboundRequestB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2966,10 +2902,8 @@ func (m *awsRestjson1_deserializeOpPutProfileOutboundR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3009,7 +2943,6 @@ func (m *awsRestjson1_deserializeOpPutProfileOutboundR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3141,10 +3074,8 @@ func (m *awsRestjson1_deserializeOpResumeCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3162,7 +3093,6 @@ func (m *awsRestjson1_deserializeOpResumeCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3253,10 +3183,8 @@ func (m *awsRestjson1_deserializeOpStartCampaign) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3274,7 +3202,6 @@ func (m *awsRestjson1_deserializeOpStartCampaign) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3365,10 +3292,8 @@ func (m *awsRestjson1_deserializeOpStartInstanceOnboar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3408,7 +3333,6 @@ func (m *awsRestjson1_deserializeOpStartInstanceOnboar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3532,10 +3456,8 @@ func (m *awsRestjson1_deserializeOpStopCampaign) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3553,7 +3475,6 @@ func (m *awsRestjson1_deserializeOpStopCampaign) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3644,10 +3565,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3665,7 +3584,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3750,10 +3668,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3771,7 +3687,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3856,10 +3771,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignChann
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3877,7 +3790,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignChann
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3962,10 +3874,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignCommu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3983,7 +3893,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignCommu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4071,10 +3980,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignCommu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4092,7 +3999,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignCommu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4180,10 +4086,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignFlowA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4201,7 +4105,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignFlowA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4289,10 +4192,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignName)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4310,7 +4211,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignName)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4395,10 +4295,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignSched
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4416,7 +4314,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignSched
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4504,10 +4401,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignSourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4525,7 +4420,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignSourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
