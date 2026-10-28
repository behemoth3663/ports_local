--- vendor/github.com/aws/aws-sdk-go-v2/service/sns/deserializers.go.orig	2026-10-28 15:34:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sns/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -44,11 +43,8 @@ func (m *awsAwsquery_deserializeOpAddPermission) Handl
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorAddPermission(response, &metadata)
@@ -133,11 +129,8 @@ func (m *awsAwsquery_deserializeOpCheckIfPhoneNumberIs
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCheckIfPhoneNumberIsOptedOut(response, &metadata)
@@ -257,11 +250,8 @@ func (m *awsAwsquery_deserializeOpConfirmSubscription)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorConfirmSubscription(response, &metadata)
@@ -390,11 +380,8 @@ func (m *awsAwsquery_deserializeOpCreatePlatformApplic
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreatePlatformApplication(response, &metadata)
@@ -511,11 +498,8 @@ func (m *awsAwsquery_deserializeOpCreatePlatformEndpoi
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreatePlatformEndpoint(response, &metadata)
@@ -635,11 +619,8 @@ func (m *awsAwsquery_deserializeOpCreateSMSSandboxPhon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateSMSSandboxPhoneNumber(response, &metadata)
@@ -765,11 +746,8 @@ func (m *awsAwsquery_deserializeOpCreateTopic) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorCreateTopic(response, &metadata)
@@ -904,11 +882,8 @@ func (m *awsAwsquery_deserializeOpDeleteEndpoint) Hand
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteEndpoint(response, &metadata)
@@ -990,11 +965,8 @@ func (m *awsAwsquery_deserializeOpDeletePlatformApplic
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeletePlatformApplication(response, &metadata)
@@ -1076,11 +1048,8 @@ func (m *awsAwsquery_deserializeOpDeleteSMSSandboxPhon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteSMSSandboxPhoneNumber(response, &metadata)
@@ -1206,11 +1175,8 @@ func (m *awsAwsquery_deserializeOpDeleteTopic) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorDeleteTopic(response, &metadata)
@@ -1307,11 +1273,8 @@ func (m *awsAwsquery_deserializeOpGetDataProtectionPol
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetDataProtectionPolicy(response, &metadata)
@@ -1434,11 +1397,8 @@ func (m *awsAwsquery_deserializeOpGetEndpointAttribute
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetEndpointAttributes(response, &metadata)
@@ -1558,11 +1518,8 @@ func (m *awsAwsquery_deserializeOpGetPlatformApplicati
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetPlatformApplicationAttributes(response, &metadata)
@@ -1682,11 +1639,8 @@ func (m *awsAwsquery_deserializeOpGetSMSAttributes) Ha
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSMSAttributes(response, &metadata)
@@ -1806,11 +1760,8 @@ func (m *awsAwsquery_deserializeOpGetSMSSandboxAccount
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSMSSandboxAccountStatus(response, &metadata)
@@ -1927,11 +1878,8 @@ func (m *awsAwsquery_deserializeOpGetSubscriptionAttri
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetSubscriptionAttributes(response, &metadata)
@@ -2051,11 +1999,8 @@ func (m *awsAwsquery_deserializeOpGetTopicAttributes) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorGetTopicAttributes(response, &metadata)
@@ -2178,11 +2123,8 @@ func (m *awsAwsquery_deserializeOpListEndpointsByPlatf
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListEndpointsByPlatformApplication(response, &metadata)
@@ -2302,11 +2244,8 @@ func (m *awsAwsquery_deserializeOpListOriginationNumbe
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListOriginationNumbers(response, &metadata)
@@ -2429,11 +2368,8 @@ func (m *awsAwsquery_deserializeOpListPhoneNumbersOpte
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPhoneNumbersOptedOut(response, &metadata)
@@ -2553,11 +2489,8 @@ func (m *awsAwsquery_deserializeOpListPlatformApplicat
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListPlatformApplications(response, &metadata)
@@ -2674,11 +2607,8 @@ func (m *awsAwsquery_deserializeOpListSMSSandboxPhoneN
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSMSSandboxPhoneNumbers(response, &metadata)
@@ -2801,11 +2731,8 @@ func (m *awsAwsquery_deserializeOpListSubscriptions) H
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSubscriptions(response, &metadata)
@@ -2922,11 +2849,8 @@ func (m *awsAwsquery_deserializeOpListSubscriptionsByT
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListSubscriptionsByTopic(response, &metadata)
@@ -3046,11 +2970,8 @@ func (m *awsAwsquery_deserializeOpListTagsForResource)
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -3173,11 +3094,8 @@ func (m *awsAwsquery_deserializeOpListTopics) HandleDe
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorListTopics(response, &metadata)
@@ -3294,11 +3212,8 @@ func (m *awsAwsquery_deserializeOpOptInPhoneNumber) Ha
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorOptInPhoneNumber(response, &metadata)
@@ -3418,11 +3333,8 @@ func (m *awsAwsquery_deserializeOpPublish) HandleDeser
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPublish(response, &metadata)
@@ -3575,11 +3487,8 @@ func (m *awsAwsquery_deserializeOpPublishBatch) Handle
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPublishBatch(response, &metadata)
@@ -3747,11 +3656,8 @@ func (m *awsAwsquery_deserializeOpPutDataProtectionPol
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorPutDataProtectionPolicy(response, &metadata)
@@ -3839,11 +3745,8 @@ func (m *awsAwsquery_deserializeOpRemovePermission) Ha
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorRemovePermission(response, &metadata)
@@ -3928,11 +3831,8 @@ func (m *awsAwsquery_deserializeOpSetEndpointAttribute
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetEndpointAttributes(response, &metadata)
@@ -4017,11 +3917,8 @@ func (m *awsAwsquery_deserializeOpSetPlatformApplicati
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetPlatformApplicationAttributes(response, &metadata)
@@ -4106,11 +4003,8 @@ func (m *awsAwsquery_deserializeOpSetSMSAttributes) Ha
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetSMSAttributes(response, &metadata)
@@ -4230,11 +4124,8 @@ func (m *awsAwsquery_deserializeOpSetSubscriptionAttri
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetSubscriptionAttributes(response, &metadata)
@@ -4325,11 +4216,8 @@ func (m *awsAwsquery_deserializeOpSetTopicAttributes) 
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSetTopicAttributes(response, &metadata)
@@ -4417,11 +4305,8 @@ func (m *awsAwsquery_deserializeOpSubscribe) HandleDes
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorSubscribe(response, &metadata)
@@ -4553,11 +4438,8 @@ func (m *awsAwsquery_deserializeOpTagResource) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorTagResource(response, &metadata)
@@ -4686,11 +4568,8 @@ func (m *awsAwsquery_deserializeOpUnsubscribe) HandleD
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUnsubscribe(response, &metadata)
@@ -4778,11 +4657,8 @@ func (m *awsAwsquery_deserializeOpUntagResource) Handl
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorUntagResource(response, &metadata)
@@ -4911,11 +4787,8 @@ func (m *awsAwsquery_deserializeOpVerifySMSSandboxPhon
 	}
 
 	defer func() { smithyhttp.CloseResponseBody(ctx, response, false, err) }()
-
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsAwsquery_deserializeOpErrorVerifySMSSandboxPhoneNumber(response, &metadata)
