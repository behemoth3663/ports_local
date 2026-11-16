--- vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateChannelFlo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpAssociateChannelFlo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -160,10 +156,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateChannelM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -203,7 +197,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateChannelM
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -338,10 +331,8 @@ func (m *awsRestjson1_deserializeOpChannelFlowCallback
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -381,7 +372,6 @@ func (m *awsRestjson1_deserializeOpChannelFlowCallback
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -521,10 +511,8 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -564,7 +552,6 @@ func (m *awsRestjson1_deserializeOpCreateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -698,10 +685,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelBan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -741,7 +726,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelBan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -880,10 +864,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelFlow) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -923,7 +905,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelFlow) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1057,10 +1038,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1100,7 +1079,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1251,10 +1229,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelModera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1294,7 +1270,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelModera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1433,10 +1408,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1454,7 +1427,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1545,10 +1517,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelBan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1566,7 +1536,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelBan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1654,10 +1623,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelFlow) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1675,7 +1642,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelFlow) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1766,10 +1732,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1787,7 +1751,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1878,10 +1841,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelMessag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1899,7 +1860,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelMessag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1987,10 +1947,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelModera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2008,7 +1966,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelModera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2096,10 +2053,8 @@ func (m *awsRestjson1_deserializeOpDeleteMessagingStre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2117,7 +2072,6 @@ func (m *awsRestjson1_deserializeOpDeleteMessagingStre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2205,10 +2159,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2248,7 +2200,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2372,10 +2323,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelBan)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2415,7 +2364,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelBan)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2542,10 +2490,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelFlow
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2585,7 +2531,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelFlow
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2709,10 +2654,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMemb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2752,7 +2695,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMemb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2879,10 +2821,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMemb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2922,7 +2862,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMemb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3046,10 +2985,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3089,7 +3026,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3213,10 +3149,8 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3256,7 +3190,6 @@ func (m *awsRestjson1_deserializeOpDescribeChannelMode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3383,10 +3316,8 @@ func (m *awsRestjson1_deserializeOpDisassociateChannel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3404,7 +3335,6 @@ func (m *awsRestjson1_deserializeOpDisassociateChannel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3498,10 +3428,8 @@ func (m *awsRestjson1_deserializeOpGetChannelMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3541,7 +3469,6 @@ func (m *awsRestjson1_deserializeOpGetChannelMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3679,10 +3606,8 @@ func (m *awsRestjson1_deserializeOpGetChannelMessage) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3722,7 +3647,6 @@ func (m *awsRestjson1_deserializeOpGetChannelMessage) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3849,10 +3773,8 @@ func (m *awsRestjson1_deserializeOpGetChannelMessageSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3892,7 +3814,6 @@ func (m *awsRestjson1_deserializeOpGetChannelMessageSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4016,10 +3937,8 @@ func (m *awsRestjson1_deserializeOpGetMessagingSession
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4059,7 +3978,6 @@ func (m *awsRestjson1_deserializeOpGetMessagingSession
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4180,10 +4098,8 @@ func (m *awsRestjson1_deserializeOpGetMessagingStreami
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4223,7 +4139,6 @@ func (m *awsRestjson1_deserializeOpGetMessagingStreami
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4350,10 +4265,8 @@ func (m *awsRestjson1_deserializeOpListChannelBans) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4393,7 +4306,6 @@ func (m *awsRestjson1_deserializeOpListChannelBans) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4535,10 +4447,8 @@ func (m *awsRestjson1_deserializeOpListChannelFlows) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4578,7 +4488,6 @@ func (m *awsRestjson1_deserializeOpListChannelFlows) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4711,10 +4620,8 @@ func (m *awsRestjson1_deserializeOpListChannelMembersh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4754,7 +4661,6 @@ func (m *awsRestjson1_deserializeOpListChannelMembersh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4896,10 +4802,8 @@ func (m *awsRestjson1_deserializeOpListChannelMembersh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4939,7 +4843,6 @@ func (m *awsRestjson1_deserializeOpListChannelMembersh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5072,10 +4975,8 @@ func (m *awsRestjson1_deserializeOpListChannelMessages
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5115,7 +5016,6 @@ func (m *awsRestjson1_deserializeOpListChannelMessages
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5266,10 +5166,8 @@ func (m *awsRestjson1_deserializeOpListChannelModerato
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5309,7 +5207,6 @@ func (m *awsRestjson1_deserializeOpListChannelModerato
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5451,10 +5348,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5494,7 +5389,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5627,10 +5521,8 @@ func (m *awsRestjson1_deserializeOpListChannelsAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5670,7 +5562,6 @@ func (m *awsRestjson1_deserializeOpListChannelsAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5803,10 +5694,8 @@ func (m *awsRestjson1_deserializeOpListChannelsModerat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5846,7 +5735,6 @@ func (m *awsRestjson1_deserializeOpListChannelsModerat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5979,10 +5867,8 @@ func (m *awsRestjson1_deserializeOpListSubChannels) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6022,7 +5908,6 @@ func (m *awsRestjson1_deserializeOpListSubChannels) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6164,10 +6049,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6207,7 +6090,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6331,10 +6213,8 @@ func (m *awsRestjson1_deserializeOpPutChannelExpiratio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6374,7 +6254,6 @@ func (m *awsRestjson1_deserializeOpPutChannelExpiratio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6510,10 +6389,8 @@ func (m *awsRestjson1_deserializeOpPutChannelMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6553,7 +6430,6 @@ func (m *awsRestjson1_deserializeOpPutChannelMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6694,10 +6570,8 @@ func (m *awsRestjson1_deserializeOpPutMessagingStreami
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6737,7 +6611,6 @@ func (m *awsRestjson1_deserializeOpPutMessagingStreami
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6867,10 +6740,8 @@ func (m *awsRestjson1_deserializeOpRedactChannelMessag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6910,7 +6781,6 @@ func (m *awsRestjson1_deserializeOpRedactChannelMessag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7059,10 +6929,8 @@ func (m *awsRestjson1_deserializeOpSearchChannels) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7102,7 +6970,6 @@ func (m *awsRestjson1_deserializeOpSearchChannels) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7235,10 +7102,8 @@ func (m *awsRestjson1_deserializeOpSendChannelMessage)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7278,7 +7143,6 @@ func (m *awsRestjson1_deserializeOpSendChannelMessage)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7432,10 +7296,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7453,7 +7315,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7544,10 +7405,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7565,7 +7424,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7653,10 +7511,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7696,7 +7552,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7827,10 +7682,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelFlow) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7870,7 +7723,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelFlow) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8001,10 +7853,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelMessag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8044,7 +7894,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelMessag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8198,10 +8047,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelReadMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8241,7 +8088,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelReadMa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
