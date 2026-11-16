--- vendor/github.com/aws/aws-sdk-go-v2/service/connectcampaigns/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connectcampaigns/deserializers.go
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
 
@@ -337,10 +330,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -358,7 +349,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectInstan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -446,10 +436,8 @@ func (m *awsRestjson1_deserializeOpDeleteInstanceOnboa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -467,7 +455,6 @@ func (m *awsRestjson1_deserializeOpDeleteInstanceOnboa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +539,8 @@ func (m *awsRestjson1_deserializeOpDescribeCampaign) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -595,7 +580,6 @@ func (m *awsRestjson1_deserializeOpDescribeCampaign) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -713,10 +697,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignState) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -756,7 +738,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignState) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -881,10 +862,8 @@ func (m *awsRestjson1_deserializeOpGetCampaignStateBat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -924,7 +903,6 @@ func (m *awsRestjson1_deserializeOpGetCampaignStateBat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1047,10 +1025,8 @@ func (m *awsRestjson1_deserializeOpGetConnectInstanceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1090,7 +1066,6 @@ func (m *awsRestjson1_deserializeOpGetConnectInstanceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1208,10 +1183,8 @@ func (m *awsRestjson1_deserializeOpGetInstanceOnboardi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1251,7 +1224,6 @@ func (m *awsRestjson1_deserializeOpGetInstanceOnboardi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1369,10 +1341,8 @@ func (m *awsRestjson1_deserializeOpListCampaigns) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1412,7 +1382,6 @@ func (m *awsRestjson1_deserializeOpListCampaigns) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1536,10 +1505,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1579,7 +1546,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1700,10 +1666,8 @@ func (m *awsRestjson1_deserializeOpPauseCampaign) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1721,7 +1685,6 @@ func (m *awsRestjson1_deserializeOpPauseCampaign) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1812,10 +1775,8 @@ func (m *awsRestjson1_deserializeOpPutDialRequestBatch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1855,7 +1816,6 @@ func (m *awsRestjson1_deserializeOpPutDialRequestBatch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1987,10 +1947,8 @@ func (m *awsRestjson1_deserializeOpResumeCampaign) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2008,7 +1966,6 @@ func (m *awsRestjson1_deserializeOpResumeCampaign) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2099,10 +2056,8 @@ func (m *awsRestjson1_deserializeOpStartCampaign) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2120,7 +2075,6 @@ func (m *awsRestjson1_deserializeOpStartCampaign) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2211,10 +2165,8 @@ func (m *awsRestjson1_deserializeOpStartInstanceOnboar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2254,7 +2206,6 @@ func (m *awsRestjson1_deserializeOpStartInstanceOnboar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2378,10 +2329,8 @@ func (m *awsRestjson1_deserializeOpStopCampaign) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2399,7 +2348,6 @@ func (m *awsRestjson1_deserializeOpStopCampaign) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2490,10 +2438,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2511,7 +2457,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2596,10 +2541,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2617,7 +2560,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2702,10 +2644,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignDiale
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2723,7 +2663,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignDiale
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2808,10 +2747,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignName)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2829,7 +2766,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignName)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2914,10 +2850,8 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignOutbo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2935,7 +2869,6 @@ func (m *awsRestjson1_deserializeOpUpdateCampaignOutbo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
