--- vendor/github.com/aws/aws-sdk-go-v2/service/notifications/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/notifications/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateChannel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpAssociateChannel) H
 	output := &AssociateChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -151,10 +147,8 @@ func (m *awsRestjson1_deserializeOpAssociateManagedNot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -166,7 +160,6 @@ func (m *awsRestjson1_deserializeOpAssociateManagedNot
 	output := &AssociateManagedNotificationAccountContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -257,10 +250,8 @@ func (m *awsRestjson1_deserializeOpAssociateManagedNot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -272,7 +263,6 @@ func (m *awsRestjson1_deserializeOpAssociateManagedNot
 	output := &AssociateManagedNotificationAdditionalChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -363,10 +353,8 @@ func (m *awsRestjson1_deserializeOpCreateEventRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -406,7 +394,6 @@ func (m *awsRestjson1_deserializeOpCreateEventRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -551,10 +538,8 @@ func (m *awsRestjson1_deserializeOpCreateNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -594,7 +579,6 @@ func (m *awsRestjson1_deserializeOpCreateNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -731,10 +715,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -746,7 +728,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventRule) Ha
 	output := &DeleteEventRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -834,10 +815,8 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -849,7 +828,6 @@ func (m *awsRestjson1_deserializeOpDeleteNotificationC
 	output := &DeleteNotificationConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -937,10 +915,8 @@ func (m *awsRestjson1_deserializeOpDeregisterNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -980,7 +956,6 @@ func (m *awsRestjson1_deserializeOpDeregisterNotificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1113,10 +1088,8 @@ func (m *awsRestjson1_deserializeOpDisableNotification
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1128,7 +1101,6 @@ func (m *awsRestjson1_deserializeOpDisableNotification
 	output := &DisableNotificationsAccessForOrganizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1219,10 +1191,8 @@ func (m *awsRestjson1_deserializeOpDisassociateChannel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1234,7 +1204,6 @@ func (m *awsRestjson1_deserializeOpDisassociateChannel
 	output := &DisassociateChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1319,10 +1288,8 @@ func (m *awsRestjson1_deserializeOpDisassociateManaged
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1334,7 +1301,6 @@ func (m *awsRestjson1_deserializeOpDisassociateManaged
 	output := &DisassociateManagedNotificationAccountContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1422,10 +1388,8 @@ func (m *awsRestjson1_deserializeOpDisassociateManaged
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1437,7 +1401,6 @@ func (m *awsRestjson1_deserializeOpDisassociateManaged
 	output := &DisassociateManagedNotificationAdditionalChannelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1522,10 +1485,8 @@ func (m *awsRestjson1_deserializeOpEnableNotifications
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1537,7 +1498,6 @@ func (m *awsRestjson1_deserializeOpEnableNotifications
 	output := &EnableNotificationsAccessForOrganizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1628,10 +1588,8 @@ func (m *awsRestjson1_deserializeOpGetEventRule) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1671,7 +1629,6 @@ func (m *awsRestjson1_deserializeOpGetEventRule) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1860,10 +1817,8 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1903,7 +1858,6 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2055,10 +2009,8 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2098,7 +2050,6 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2259,10 +2210,8 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2302,7 +2251,6 @@ func (m *awsRestjson1_deserializeOpGetManagedNotificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2454,10 +2402,8 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2497,7 +2443,6 @@ func (m *awsRestjson1_deserializeOpGetNotificationConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2671,10 +2616,8 @@ func (m *awsRestjson1_deserializeOpGetNotificationEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2714,7 +2657,6 @@ func (m *awsRestjson1_deserializeOpGetNotificationEven
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2866,10 +2808,8 @@ func (m *awsRestjson1_deserializeOpGetNotificationsAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2909,7 +2849,6 @@ func (m *awsRestjson1_deserializeOpGetNotificationsAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3027,10 +2966,8 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3070,7 +3007,6 @@ func (m *awsRestjson1_deserializeOpListChannels) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3200,10 +3136,8 @@ func (m *awsRestjson1_deserializeOpListEventRules) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3243,7 +3177,6 @@ func (m *awsRestjson1_deserializeOpListEventRules) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3373,10 +3306,8 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3416,7 +3347,6 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3546,10 +3476,8 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3589,7 +3517,6 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3716,10 +3643,8 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3759,7 +3684,6 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3886,10 +3810,8 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3929,7 +3851,6 @@ func (m *awsRestjson1_deserializeOpListManagedNotifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4056,10 +3977,8 @@ func (m *awsRestjson1_deserializeOpListNotificationCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4099,7 +4018,6 @@ func (m *awsRestjson1_deserializeOpListNotificationCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4226,10 +4144,8 @@ func (m *awsRestjson1_deserializeOpListNotificationEve
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4269,7 +4185,6 @@ func (m *awsRestjson1_deserializeOpListNotificationEve
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4396,10 +4311,8 @@ func (m *awsRestjson1_deserializeOpListNotificationHub
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4439,7 +4352,6 @@ func (m *awsRestjson1_deserializeOpListNotificationHub
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4566,10 +4478,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4609,7 +4519,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4730,10 +4639,8 @@ func (m *awsRestjson1_deserializeOpRegisterNotificatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4773,7 +4680,6 @@ func (m *awsRestjson1_deserializeOpRegisterNotificatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4932,10 +4838,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4947,7 +4851,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5032,10 +4935,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5047,7 +4948,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5132,10 +5032,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5175,7 +5073,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5317,10 +5214,8 @@ func (m *awsRestjson1_deserializeOpUpdateNotificationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5360,7 +5255,6 @@ func (m *awsRestjson1_deserializeOpUpdateNotificationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
