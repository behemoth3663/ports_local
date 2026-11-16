--- vendor/github.com/aws/aws-sdk-go-v2/service/chatbot/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/chatbot/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateToConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpAssociateToConfigur
 	output := &AssociateToConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -138,10 +134,8 @@ func (m *awsRestjson1_deserializeOpCreateChimeWebhookC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -181,7 +175,6 @@ func (m *awsRestjson1_deserializeOpCreateChimeWebhookC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -302,10 +295,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomAction)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -345,7 +336,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomAction)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -470,10 +460,8 @@ func (m *awsRestjson1_deserializeOpCreateMicrosoftTeam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -513,7 +501,6 @@ func (m *awsRestjson1_deserializeOpCreateMicrosoftTeam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -634,10 +621,8 @@ func (m *awsRestjson1_deserializeOpCreateSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -677,7 +662,6 @@ func (m *awsRestjson1_deserializeOpCreateSlackChannelC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -798,10 +782,8 @@ func (m *awsRestjson1_deserializeOpDeleteChimeWebhookC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -813,7 +795,6 @@ func (m *awsRestjson1_deserializeOpDeleteChimeWebhookC
 	output := &DeleteChimeWebhookConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -895,10 +876,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomAction)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -910,7 +889,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomAction)
 	output := &DeleteCustomActionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -992,10 +970,8 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1007,7 +983,6 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 	output := &DeleteMicrosoftTeamsChannelConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1089,10 +1064,8 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1104,7 +1077,6 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 	output := &DeleteMicrosoftTeamsConfiguredTeamOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1180,10 +1152,8 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1195,7 +1165,6 @@ func (m *awsRestjson1_deserializeOpDeleteMicrosoftTeam
 	output := &DeleteMicrosoftTeamsUserIdentityOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1274,10 +1243,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1289,7 +1256,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlackChannelC
 	output := &DeleteSlackChannelConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1371,10 +1337,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlackUserIden
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1386,7 +1350,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlackUserIden
 	output := &DeleteSlackUserIdentityOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1465,10 +1428,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlackWorkspac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1480,7 +1441,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlackWorkspac
 	output := &DeleteSlackWorkspaceAuthorizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1556,10 +1516,8 @@ func (m *awsRestjson1_deserializeOpDescribeChimeWebhoo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1599,7 +1557,6 @@ func (m *awsRestjson1_deserializeOpDescribeChimeWebhoo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1723,10 +1680,8 @@ func (m *awsRestjson1_deserializeOpDescribeSlackChanne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1766,7 +1721,6 @@ func (m *awsRestjson1_deserializeOpDescribeSlackChanne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1890,10 +1844,8 @@ func (m *awsRestjson1_deserializeOpDescribeSlackUserId
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1933,7 +1885,6 @@ func (m *awsRestjson1_deserializeOpDescribeSlackUserId
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2057,10 +2008,8 @@ func (m *awsRestjson1_deserializeOpDescribeSlackWorksp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2100,7 +2049,6 @@ func (m *awsRestjson1_deserializeOpDescribeSlackWorksp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2224,10 +2172,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2239,7 +2185,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromCon
 	output := &DisassociateFromConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2318,10 +2263,8 @@ func (m *awsRestjson1_deserializeOpGetAccountPreferenc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2361,7 +2304,6 @@ func (m *awsRestjson1_deserializeOpGetAccountPreferenc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2473,10 +2415,8 @@ func (m *awsRestjson1_deserializeOpGetCustomAction) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2516,7 +2456,6 @@ func (m *awsRestjson1_deserializeOpGetCustomAction) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2634,10 +2573,8 @@ func (m *awsRestjson1_deserializeOpGetMicrosoftTeamsCh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2677,7 +2614,6 @@ func (m *awsRestjson1_deserializeOpGetMicrosoftTeamsCh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2792,10 +2728,8 @@ func (m *awsRestjson1_deserializeOpListAssociations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2835,7 +2769,6 @@ func (m *awsRestjson1_deserializeOpListAssociations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2950,10 +2883,8 @@ func (m *awsRestjson1_deserializeOpListCustomActions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2993,7 +2924,6 @@ func (m *awsRestjson1_deserializeOpListCustomActions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3117,10 +3047,8 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3160,7 +3088,6 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3284,10 +3211,8 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3327,7 +3252,6 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3451,10 +3375,8 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsU
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3494,7 +3416,6 @@ func (m *awsRestjson1_deserializeOpListMicrosoftTeamsU
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3618,10 +3539,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3661,7 +3580,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3776,10 +3694,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3791,7 +3707,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3873,10 +3788,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3888,7 +3801,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3967,10 +3879,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountPrefer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4010,7 +3920,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountPrefer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4125,10 +4034,8 @@ func (m *awsRestjson1_deserializeOpUpdateChimeWebhookC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4168,7 +4075,6 @@ func (m *awsRestjson1_deserializeOpUpdateChimeWebhookC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4286,10 +4192,8 @@ func (m *awsRestjson1_deserializeOpUpdateCustomAction)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4329,7 +4233,6 @@ func (m *awsRestjson1_deserializeOpUpdateCustomAction)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4451,10 +4354,8 @@ func (m *awsRestjson1_deserializeOpUpdateMicrosoftTeam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4494,7 +4395,6 @@ func (m *awsRestjson1_deserializeOpUpdateMicrosoftTeam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4612,10 +4512,8 @@ func (m *awsRestjson1_deserializeOpUpdateSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4655,7 +4553,6 @@ func (m *awsRestjson1_deserializeOpUpdateSlackChannelC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
