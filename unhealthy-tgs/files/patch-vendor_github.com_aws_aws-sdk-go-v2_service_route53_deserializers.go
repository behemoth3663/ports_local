--- vendor/github.com/aws/aws-sdk-go-v2/service/route53/deserializers.go.orig	2026-11-21 14:46:19 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -50,10 +49,8 @@ func (m *awsRestxml_deserializeOpActivateKeySigningKey
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorActivateKeySigningKey(response, &metadata)
@@ -89,7 +86,6 @@ func (m *awsRestxml_deserializeOpActivateKeySigningKey
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -215,10 +211,8 @@ func (m *awsRestxml_deserializeOpAssociateVPCWithHoste
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAssociateVPCWithHostedZone(response, &metadata)
@@ -254,7 +248,6 @@ func (m *awsRestxml_deserializeOpAssociateVPCWithHoste
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -386,10 +379,8 @@ func (m *awsRestxml_deserializeOpChangeCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorChangeCidrCollection(response, &metadata)
@@ -425,7 +416,6 @@ func (m *awsRestxml_deserializeOpChangeCidrCollection)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -558,10 +548,8 @@ func (m *awsRestxml_deserializeOpChangeResourceRecordS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorChangeResourceRecordSets(response, &metadata)
@@ -597,7 +585,6 @@ func (m *awsRestxml_deserializeOpChangeResourceRecordS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -720,10 +707,8 @@ func (m *awsRestxml_deserializeOpChangeTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorChangeTagsForResource(response, &metadata)
@@ -731,7 +716,6 @@ func (m *awsRestxml_deserializeOpChangeTagsForResource
 	output := &ChangeTagsForResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -812,10 +796,8 @@ func (m *awsRestxml_deserializeOpCreateCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateCidrCollection(response, &metadata)
@@ -856,7 +838,6 @@ func (m *awsRestxml_deserializeOpCreateCidrCollection)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -988,10 +969,8 @@ func (m *awsRestxml_deserializeOpCreateHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateHealthCheck(response, &metadata)
@@ -1032,7 +1011,6 @@ func (m *awsRestxml_deserializeOpCreateHealthCheck) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1161,10 +1139,8 @@ func (m *awsRestxml_deserializeOpCreateHostedZone) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateHostedZone(response, &metadata)
@@ -1205,7 +1181,6 @@ func (m *awsRestxml_deserializeOpCreateHostedZone) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1370,10 +1345,8 @@ func (m *awsRestxml_deserializeOpCreateKeySigningKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateKeySigningKey(response, &metadata)
@@ -1414,7 +1387,6 @@ func (m *awsRestxml_deserializeOpCreateKeySigningKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1570,10 +1542,8 @@ func (m *awsRestxml_deserializeOpCreateQueryLoggingCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateQueryLoggingConfig(response, &metadata)
@@ -1614,7 +1584,6 @@ func (m *awsRestxml_deserializeOpCreateQueryLoggingCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1752,10 +1721,8 @@ func (m *awsRestxml_deserializeOpCreateReusableDelegat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateReusableDelegationSet(response, &metadata)
@@ -1796,7 +1763,6 @@ func (m *awsRestxml_deserializeOpCreateReusableDelegat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1937,10 +1903,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateTrafficPolicy(response, &metadata)
@@ -1981,7 +1945,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2113,10 +2076,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateTrafficPolicyInstance(response, &metadata)
@@ -2157,7 +2118,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2292,10 +2252,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateTrafficPolicyVersion(response, &metadata)
@@ -2336,7 +2294,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2471,10 +2428,8 @@ func (m *awsRestxml_deserializeOpCreateVPCAssociationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateVPCAssociationAuthorization(response, &metadata)
@@ -2510,7 +2465,6 @@ func (m *awsRestxml_deserializeOpCreateVPCAssociationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2646,10 +2600,8 @@ func (m *awsRestxml_deserializeOpDeactivateKeySigningK
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeactivateKeySigningKey(response, &metadata)
@@ -2685,7 +2637,6 @@ func (m *awsRestxml_deserializeOpDeactivateKeySigningK
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2814,10 +2765,8 @@ func (m *awsRestxml_deserializeOpDeleteCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteCidrCollection(response, &metadata)
@@ -2825,7 +2774,6 @@ func (m *awsRestxml_deserializeOpDeleteCidrCollection)
 	output := &DeleteCidrCollectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2903,10 +2851,8 @@ func (m *awsRestxml_deserializeOpDeleteHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteHealthCheck(response, &metadata)
@@ -2914,7 +2860,6 @@ func (m *awsRestxml_deserializeOpDeleteHealthCheck) Ha
 	output := &DeleteHealthCheckOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2989,10 +2934,8 @@ func (m *awsRestxml_deserializeOpDeleteHostedZone) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteHostedZone(response, &metadata)
@@ -3028,7 +2971,6 @@ func (m *awsRestxml_deserializeOpDeleteHostedZone) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3151,10 +3093,8 @@ func (m *awsRestxml_deserializeOpDeleteKeySigningKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteKeySigningKey(response, &metadata)
@@ -3190,7 +3130,6 @@ func (m *awsRestxml_deserializeOpDeleteKeySigningKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3316,10 +3255,8 @@ func (m *awsRestxml_deserializeOpDeleteQueryLoggingCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteQueryLoggingConfig(response, &metadata)
@@ -3327,7 +3264,6 @@ func (m *awsRestxml_deserializeOpDeleteQueryLoggingCon
 	output := &DeleteQueryLoggingConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3402,10 +3338,8 @@ func (m *awsRestxml_deserializeOpDeleteReusableDelegat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteReusableDelegationSet(response, &metadata)
@@ -3413,7 +3347,6 @@ func (m *awsRestxml_deserializeOpDeleteReusableDelegat
 	output := &DeleteReusableDelegationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3491,10 +3424,8 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteTrafficPolicy(response, &metadata)
@@ -3502,7 +3433,6 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicy) 
 	output := &DeleteTrafficPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3580,10 +3510,8 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteTrafficPolicyInstance(response, &metadata)
@@ -3591,7 +3519,6 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicyIn
 	output := &DeleteTrafficPolicyInstanceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3666,10 +3593,8 @@ func (m *awsRestxml_deserializeOpDeleteVPCAssociationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteVPCAssociationAuthorization(response, &metadata)
@@ -3677,7 +3602,6 @@ func (m *awsRestxml_deserializeOpDeleteVPCAssociationA
 	output := &DeleteVPCAssociationAuthorizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3758,10 +3682,8 @@ func (m *awsRestxml_deserializeOpDisableHostedZoneDNSS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDisableHostedZoneDNSSEC(response, &metadata)
@@ -3797,7 +3719,6 @@ func (m *awsRestxml_deserializeOpDisableHostedZoneDNSS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3929,10 +3850,8 @@ func (m *awsRestxml_deserializeOpDisassociateVPCFromHo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDisassociateVPCFromHostedZone(response, &metadata)
@@ -3968,7 +3887,6 @@ func (m *awsRestxml_deserializeOpDisassociateVPCFromHo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4091,10 +4009,8 @@ func (m *awsRestxml_deserializeOpEnableHostedZoneDNSSE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorEnableHostedZoneDNSSEC(response, &metadata)
@@ -4130,7 +4046,6 @@ func (m *awsRestxml_deserializeOpEnableHostedZoneDNSSE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4265,10 +4180,8 @@ func (m *awsRestxml_deserializeOpGetAccountLimit) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAccountLimit(response, &metadata)
@@ -4304,7 +4217,6 @@ func (m *awsRestxml_deserializeOpGetAccountLimit) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4432,10 +4344,8 @@ func (m *awsRestxml_deserializeOpGetChange) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetChange(response, &metadata)
@@ -4471,7 +4381,6 @@ func (m *awsRestxml_deserializeOpGetChange) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4585,10 +4494,8 @@ func (m *awsRestxml_deserializeOpGetCheckerIpRanges) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetCheckerIpRanges(response, &metadata)
@@ -4624,7 +4531,6 @@ func (m *awsRestxml_deserializeOpGetCheckerIpRanges) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4732,10 +4638,8 @@ func (m *awsRestxml_deserializeOpGetDNSSEC) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDNSSEC(response, &metadata)
@@ -4771,7 +4675,6 @@ func (m *awsRestxml_deserializeOpGetDNSSEC) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4894,10 +4797,8 @@ func (m *awsRestxml_deserializeOpGetGeoLocation) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetGeoLocation(response, &metadata)
@@ -4933,7 +4834,6 @@ func (m *awsRestxml_deserializeOpGetGeoLocation) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5047,10 +4947,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheck) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHealthCheck(response, &metadata)
@@ -5086,7 +4984,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheck) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5203,10 +5100,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckCount) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHealthCheckCount(response, &metadata)
@@ -5242,7 +5137,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckCount) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5361,10 +5255,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckLastFai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHealthCheckLastFailureReason(response, &metadata)
@@ -5400,7 +5292,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckLastFai
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5514,10 +5405,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckStatus)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHealthCheckStatus(response, &metadata)
@@ -5553,7 +5442,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckStatus)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5667,10 +5555,8 @@ func (m *awsRestxml_deserializeOpGetHostedZone) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHostedZone(response, &metadata)
@@ -5706,7 +5592,6 @@ func (m *awsRestxml_deserializeOpGetHostedZone) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5832,10 +5717,8 @@ func (m *awsRestxml_deserializeOpGetHostedZoneCount) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHostedZoneCount(response, &metadata)
@@ -5871,7 +5754,6 @@ func (m *awsRestxml_deserializeOpGetHostedZoneCount) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5993,10 +5875,8 @@ func (m *awsRestxml_deserializeOpGetHostedZoneLimit) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetHostedZoneLimit(response, &metadata)
@@ -6032,7 +5912,6 @@ func (m *awsRestxml_deserializeOpGetHostedZoneLimit) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6166,10 +6045,8 @@ func (m *awsRestxml_deserializeOpGetQueryLoggingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetQueryLoggingConfig(response, &metadata)
@@ -6205,7 +6082,6 @@ func (m *awsRestxml_deserializeOpGetQueryLoggingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6319,10 +6195,8 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetReusableDelegationSet(response, &metadata)
@@ -6358,7 +6232,6 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6475,10 +6348,8 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetReusableDelegationSetLimit(response, &metadata)
@@ -6514,7 +6385,6 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6645,10 +6515,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetTrafficPolicy(response, &metadata)
@@ -6684,7 +6552,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6798,10 +6665,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetTrafficPolicyInstance(response, &metadata)
@@ -6837,7 +6702,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6951,10 +6815,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetTrafficPolicyInstanceCount(response, &metadata)
@@ -6990,7 +6852,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7109,10 +6970,8 @@ func (m *awsRestxml_deserializeOpListCidrBlocks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCidrBlocks(response, &metadata)
@@ -7148,7 +7007,6 @@ func (m *awsRestxml_deserializeOpListCidrBlocks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7278,10 +7136,8 @@ func (m *awsRestxml_deserializeOpListCidrCollections) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCidrCollections(response, &metadata)
@@ -7317,7 +7173,6 @@ func (m *awsRestxml_deserializeOpListCidrCollections) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7441,10 +7296,8 @@ func (m *awsRestxml_deserializeOpListCidrLocations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCidrLocations(response, &metadata)
@@ -7480,7 +7333,6 @@ func (m *awsRestxml_deserializeOpListCidrLocations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7607,10 +7459,8 @@ func (m *awsRestxml_deserializeOpListGeoLocations) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListGeoLocations(response, &metadata)
@@ -7646,7 +7496,6 @@ func (m *awsRestxml_deserializeOpListGeoLocations) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7829,10 +7678,8 @@ func (m *awsRestxml_deserializeOpListHealthChecks) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListHealthChecks(response, &metadata)
@@ -7868,7 +7715,6 @@ func (m *awsRestxml_deserializeOpListHealthChecks) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8041,10 +7887,8 @@ func (m *awsRestxml_deserializeOpListHostedZones) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListHostedZones(response, &metadata)
@@ -8080,7 +7924,6 @@ func (m *awsRestxml_deserializeOpListHostedZones) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8256,10 +8099,8 @@ func (m *awsRestxml_deserializeOpListHostedZonesByName
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListHostedZonesByName(response, &metadata)
@@ -8295,7 +8136,6 @@ func (m *awsRestxml_deserializeOpListHostedZonesByName
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8494,10 +8334,8 @@ func (m *awsRestxml_deserializeOpListHostedZonesByVPC)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListHostedZonesByVPC(response, &metadata)
@@ -8533,7 +8371,6 @@ func (m *awsRestxml_deserializeOpListHostedZonesByVPC)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8677,10 +8514,8 @@ func (m *awsRestxml_deserializeOpListQueryLoggingConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListQueryLoggingConfigs(response, &metadata)
@@ -8716,7 +8551,6 @@ func (m *awsRestxml_deserializeOpListQueryLoggingConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8846,10 +8680,8 @@ func (m *awsRestxml_deserializeOpListResourceRecordSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListResourceRecordSets(response, &metadata)
@@ -8885,7 +8717,6 @@ func (m *awsRestxml_deserializeOpListResourceRecordSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9071,10 +8902,8 @@ func (m *awsRestxml_deserializeOpListReusableDelegatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListReusableDelegationSets(response, &metadata)
@@ -9110,7 +8939,6 @@ func (m *awsRestxml_deserializeOpListReusableDelegatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9280,10 +9108,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -9319,7 +9145,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9442,10 +9267,8 @@ func (m *awsRestxml_deserializeOpListTagsForResources)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTagsForResources(response, &metadata)
@@ -9481,7 +9304,6 @@ func (m *awsRestxml_deserializeOpListTagsForResources)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9604,10 +9426,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicies) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrafficPolicies(response, &metadata)
@@ -9643,7 +9463,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicies) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9800,10 +9619,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrafficPolicyInstances(response, &metadata)
@@ -9839,7 +9656,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10025,10 +9841,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrafficPolicyInstancesByHostedZone(response, &metadata)
@@ -10064,7 +9878,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10240,10 +10053,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrafficPolicyInstancesByPolicy(response, &metadata)
@@ -10279,7 +10090,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10468,10 +10278,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrafficPolicyVersions(response, &metadata)
@@ -10507,7 +10315,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10667,10 +10474,8 @@ func (m *awsRestxml_deserializeOpListVPCAssociationAut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListVPCAssociationAuthorizations(response, &metadata)
@@ -10706,7 +10511,6 @@ func (m *awsRestxml_deserializeOpListVPCAssociationAut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10849,10 +10653,8 @@ func (m *awsRestxml_deserializeOpTestDNSAnswer) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorTestDNSAnswer(response, &metadata)
@@ -10888,7 +10690,6 @@ func (m *awsRestxml_deserializeOpTestDNSAnswer) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11067,10 +10868,8 @@ func (m *awsRestxml_deserializeOpUpdateHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateHealthCheck(response, &metadata)
@@ -11106,7 +10905,6 @@ func (m *awsRestxml_deserializeOpUpdateHealthCheck) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11223,10 +11021,8 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneComme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateHostedZoneComment(response, &metadata)
@@ -11262,7 +11058,6 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneComme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11379,10 +11174,8 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneFeatu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateHostedZoneFeatures(response, &metadata)
@@ -11390,7 +11183,6 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneFeatu
 	output := &UpdateHostedZoneFeaturesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11468,10 +11260,8 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateTrafficPolicyComment(response, &metadata)
@@ -11507,7 +11297,6 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11624,10 +11413,8 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateTrafficPolicyInstance(response, &metadata)
@@ -11663,7 +11450,6 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
