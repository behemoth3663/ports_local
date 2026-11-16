--- vendor/github.com/aws/aws-sdk-go-v2/service/mediapackagev2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mediapackagev2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCancelHarvestJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -61,7 +58,6 @@ func (m *awsRestjson1_deserializeOpCancelHarvestJob) H
 	output := &CancelHarvestJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -149,10 +145,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -192,7 +186,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -420,10 +413,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -463,7 +454,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -667,10 +657,8 @@ func (m *awsRestjson1_deserializeOpCreateHarvestJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -710,7 +698,6 @@ func (m *awsRestjson1_deserializeOpCreateHarvestJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -965,10 +952,8 @@ func (m *awsRestjson1_deserializeOpCreateOriginEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1008,7 +993,6 @@ func (m *awsRestjson1_deserializeOpCreateOriginEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1273,10 +1257,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1288,7 +1270,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 	output := &DeleteChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1373,10 +1354,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1388,7 +1367,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelGroup)
 	output := &DeleteChannelGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1473,10 +1451,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPolicy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1488,7 +1464,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelPolicy
 	output := &DeleteChannelPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1573,10 +1548,8 @@ func (m *awsRestjson1_deserializeOpDeleteOriginEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1588,7 +1561,6 @@ func (m *awsRestjson1_deserializeOpDeleteOriginEndpoin
 	output := &DeleteOriginEndpointOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1670,10 +1642,8 @@ func (m *awsRestjson1_deserializeOpDeleteOriginEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1685,7 +1655,6 @@ func (m *awsRestjson1_deserializeOpDeleteOriginEndpoin
 	output := &DeleteOriginEndpointPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1770,10 +1739,8 @@ func (m *awsRestjson1_deserializeOpGetChannel) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1813,7 +1780,6 @@ func (m *awsRestjson1_deserializeOpGetChannel) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2051,10 +2017,8 @@ func (m *awsRestjson1_deserializeOpGetChannelGroup) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2094,7 +2058,6 @@ func (m *awsRestjson1_deserializeOpGetChannelGroup) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2292,10 +2255,8 @@ func (m *awsRestjson1_deserializeOpGetChannelPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2335,7 +2296,6 @@ func (m *awsRestjson1_deserializeOpGetChannelPolicy) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2478,10 +2438,8 @@ func (m *awsRestjson1_deserializeOpGetHarvestJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2521,7 +2479,6 @@ func (m *awsRestjson1_deserializeOpGetHarvestJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2770,10 +2727,8 @@ func (m *awsRestjson1_deserializeOpGetOriginEndpoint) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2813,7 +2768,6 @@ func (m *awsRestjson1_deserializeOpGetOriginEndpoint) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3088,10 +3042,8 @@ func (m *awsRestjson1_deserializeOpGetOriginEndpointPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3131,7 +3083,6 @@ func (m *awsRestjson1_deserializeOpGetOriginEndpointPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3288,10 +3239,8 @@ func (m *awsRestjson1_deserializeOpListChannelGroups) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3331,7 +3280,6 @@ func (m *awsRestjson1_deserializeOpListChannelGroups) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3458,10 +3406,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3501,7 +3447,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3631,10 +3576,8 @@ func (m *awsRestjson1_deserializeOpListHarvestJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3674,7 +3617,6 @@ func (m *awsRestjson1_deserializeOpListHarvestJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3804,10 +3746,8 @@ func (m *awsRestjson1_deserializeOpListOriginEndpoints
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3847,7 +3787,6 @@ func (m *awsRestjson1_deserializeOpListOriginEndpoints
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3977,10 +3916,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4020,7 +3957,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4129,10 +4065,8 @@ func (m *awsRestjson1_deserializeOpPutChannelPolicy) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4144,7 +4078,6 @@ func (m *awsRestjson1_deserializeOpPutChannelPolicy) H
 	output := &PutChannelPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4232,10 +4165,8 @@ func (m *awsRestjson1_deserializeOpPutOriginEndpointPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4247,7 +4178,6 @@ func (m *awsRestjson1_deserializeOpPutOriginEndpointPo
 	output := &PutOriginEndpointPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4335,10 +4265,8 @@ func (m *awsRestjson1_deserializeOpResetChannelState) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4378,7 +4306,6 @@ func (m *awsRestjson1_deserializeOpResetChannelState) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4540,10 +4467,8 @@ func (m *awsRestjson1_deserializeOpResetOriginEndpoint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4583,7 +4508,6 @@ func (m *awsRestjson1_deserializeOpResetOriginEndpoint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4754,10 +4678,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4775,7 +4697,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4848,10 +4769,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4869,7 +4788,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4942,10 +4860,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4985,7 +4901,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5210,10 +5125,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5253,7 +5166,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5454,10 +5366,8 @@ func (m *awsRestjson1_deserializeOpUpdateOriginEndpoin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5497,7 +5407,6 @@ func (m *awsRestjson1_deserializeOpUpdateOriginEndpoin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
