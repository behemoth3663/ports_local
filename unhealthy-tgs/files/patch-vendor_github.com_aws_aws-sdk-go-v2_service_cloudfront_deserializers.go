--- vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/deserializers.go.orig	2026-11-21 14:46:26 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cloudfront/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -50,10 +49,8 @@ func (m *awsRestxml_deserializeOpAssociateAlias) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAssociateAlias(response, &metadata)
@@ -67,7 +64,6 @@ func (m *awsRestxml_deserializeOpAssociateAlias) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAssociateDistributionTenantWebACL(response, &metadata)
@@ -192,7 +186,6 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -350,10 +343,8 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorAssociateDistributionWebACL(response, &metadata)
@@ -394,7 +385,6 @@ func (m *awsRestxml_deserializeOpAssociateDistribution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +542,8 @@ func (m *awsRestxml_deserializeOpCopyDistribution) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCopyDistribution(response, &metadata)
@@ -596,7 +584,6 @@ func (m *awsRestxml_deserializeOpCopyDistribution) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -913,10 +900,8 @@ func (m *awsRestxml_deserializeOpCreateAnycastIpList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateAnycastIpList(response, &metadata)
@@ -957,7 +942,6 @@ func (m *awsRestxml_deserializeOpCreateAnycastIpList) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1095,10 +1079,8 @@ func (m *awsRestxml_deserializeOpCreateCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateCachePolicy(response, &metadata)
@@ -1139,7 +1121,6 @@ func (m *awsRestxml_deserializeOpCreateCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1288,10 +1269,8 @@ func (m *awsRestxml_deserializeOpCreateCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateCloudFrontOriginAccessIdentity(response, &metadata)
@@ -1332,7 +1311,6 @@ func (m *awsRestxml_deserializeOpCreateCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1472,10 +1450,8 @@ func (m *awsRestxml_deserializeOpCreateConnectionFunct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateConnectionFunction(response, &metadata)
@@ -1516,7 +1492,6 @@ func (m *awsRestxml_deserializeOpCreateConnectionFunct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1662,10 +1637,8 @@ func (m *awsRestxml_deserializeOpCreateConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateConnectionGroup(response, &metadata)
@@ -1706,7 +1679,6 @@ func (m *awsRestxml_deserializeOpCreateConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1844,10 +1816,8 @@ func (m *awsRestxml_deserializeOpCreateContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateContinuousDeploymentPolicy(response, &metadata)
@@ -1888,7 +1858,6 @@ func (m *awsRestxml_deserializeOpCreateContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2031,10 +2000,8 @@ func (m *awsRestxml_deserializeOpCreateDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateDistribution(response, &metadata)
@@ -2075,7 +2042,6 @@ func (m *awsRestxml_deserializeOpCreateDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2401,10 +2367,8 @@ func (m *awsRestxml_deserializeOpCreateDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateDistributionTenant(response, &metadata)
@@ -2445,7 +2409,6 @@ func (m *awsRestxml_deserializeOpCreateDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2589,10 +2552,8 @@ func (m *awsRestxml_deserializeOpCreateDistributionWit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateDistributionWithTags(response, &metadata)
@@ -2633,7 +2594,6 @@ func (m *awsRestxml_deserializeOpCreateDistributionWit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2959,10 +2919,8 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateFieldLevelEncryptionConfig(response, &metadata)
@@ -3003,7 +2961,6 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3152,10 +3109,8 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateFieldLevelEncryptionProfile(response, &metadata)
@@ -3196,7 +3151,6 @@ func (m *awsRestxml_deserializeOpCreateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3345,10 +3299,8 @@ func (m *awsRestxml_deserializeOpCreateFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateFunction(response, &metadata)
@@ -3389,7 +3341,6 @@ func (m *awsRestxml_deserializeOpCreateFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3529,10 +3480,8 @@ func (m *awsRestxml_deserializeOpCreateInvalidation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateInvalidation(response, &metadata)
@@ -3573,7 +3522,6 @@ func (m *awsRestxml_deserializeOpCreateInvalidation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3714,10 +3662,8 @@ func (m *awsRestxml_deserializeOpCreateInvalidationFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateInvalidationForDistributionTenant(response, &metadata)
@@ -3758,7 +3704,6 @@ func (m *awsRestxml_deserializeOpCreateInvalidationFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3899,10 +3844,8 @@ func (m *awsRestxml_deserializeOpCreateKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateKeyGroup(response, &metadata)
@@ -3943,7 +3886,6 @@ func (m *awsRestxml_deserializeOpCreateKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4080,10 +4022,8 @@ func (m *awsRestxml_deserializeOpCreateKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateKeyValueStore(response, &metadata)
@@ -4124,7 +4064,6 @@ func (m *awsRestxml_deserializeOpCreateKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4267,10 +4206,8 @@ func (m *awsRestxml_deserializeOpCreateMonitoringSubsc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateMonitoringSubscription(response, &metadata)
@@ -4306,7 +4243,6 @@ func (m *awsRestxml_deserializeOpCreateMonitoringSubsc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4426,10 +4362,8 @@ func (m *awsRestxml_deserializeOpCreateOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateOriginAccessControl(response, &metadata)
@@ -4470,7 +4404,6 @@ func (m *awsRestxml_deserializeOpCreateOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4604,10 +4537,8 @@ func (m *awsRestxml_deserializeOpCreateOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateOriginRequestPolicy(response, &metadata)
@@ -4648,7 +4579,6 @@ func (m *awsRestxml_deserializeOpCreateOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4797,10 +4727,8 @@ func (m *awsRestxml_deserializeOpCreatePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreatePublicKey(response, &metadata)
@@ -4841,7 +4769,6 @@ func (m *awsRestxml_deserializeOpCreatePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4975,10 +4902,8 @@ func (m *awsRestxml_deserializeOpCreateRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateRealtimeLogConfig(response, &metadata)
@@ -5014,7 +4939,6 @@ func (m *awsRestxml_deserializeOpCreateRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5134,10 +5058,8 @@ func (m *awsRestxml_deserializeOpCreateResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateResponseHeadersPolicy(response, &metadata)
@@ -5178,7 +5100,6 @@ func (m *awsRestxml_deserializeOpCreateResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5327,10 +5248,8 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateStreamingDistribution(response, &metadata)
@@ -5371,7 +5290,6 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5535,10 +5453,8 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateStreamingDistributionWithTags(response, &metadata)
@@ -5579,7 +5495,6 @@ func (m *awsRestxml_deserializeOpCreateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5746,10 +5661,8 @@ func (m *awsRestxml_deserializeOpCreateTrustStore) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateTrustStore(response, &metadata)
@@ -5790,7 +5703,6 @@ func (m *awsRestxml_deserializeOpCreateTrustStore) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5928,10 +5840,8 @@ func (m *awsRestxml_deserializeOpCreateVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorCreateVpcOrigin(response, &metadata)
@@ -5972,7 +5882,6 @@ func (m *awsRestxml_deserializeOpCreateVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6118,10 +6027,8 @@ func (m *awsRestxml_deserializeOpDeleteAnycastIpList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteAnycastIpList(response, &metadata)
@@ -6135,7 +6042,6 @@ func (m *awsRestxml_deserializeOpDeleteAnycastIpList) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6225,10 +6131,8 @@ func (m *awsRestxml_deserializeOpDeleteCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteCachePolicy(response, &metadata)
@@ -6242,7 +6146,6 @@ func (m *awsRestxml_deserializeOpDeleteCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6326,10 +6229,8 @@ func (m *awsRestxml_deserializeOpDeleteCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteCloudFrontOriginAccessIdentity(response, &metadata)
@@ -6343,7 +6244,6 @@ func (m *awsRestxml_deserializeOpDeleteCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6424,10 +6324,8 @@ func (m *awsRestxml_deserializeOpDeleteConnectionFunct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteConnectionFunction(response, &metadata)
@@ -6441,7 +6339,6 @@ func (m *awsRestxml_deserializeOpDeleteConnectionFunct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6528,10 +6425,8 @@ func (m *awsRestxml_deserializeOpDeleteConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteConnectionGroup(response, &metadata)
@@ -6545,7 +6440,6 @@ func (m *awsRestxml_deserializeOpDeleteConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6629,10 +6523,8 @@ func (m *awsRestxml_deserializeOpDeleteContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteContinuousDeploymentPolicy(response, &metadata)
@@ -6646,7 +6538,6 @@ func (m *awsRestxml_deserializeOpDeleteContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6730,10 +6621,8 @@ func (m *awsRestxml_deserializeOpDeleteDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteDistribution(response, &metadata)
@@ -6747,7 +6636,6 @@ func (m *awsRestxml_deserializeOpDeleteDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6831,10 +6719,8 @@ func (m *awsRestxml_deserializeOpDeleteDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteDistributionTenant(response, &metadata)
@@ -6848,7 +6734,6 @@ func (m *awsRestxml_deserializeOpDeleteDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6929,10 +6814,8 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteFieldLevelEncryptionConfig(response, &metadata)
@@ -6946,7 +6829,6 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7027,10 +6909,8 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteFieldLevelEncryptionProfile(response, &metadata)
@@ -7044,7 +6924,6 @@ func (m *awsRestxml_deserializeOpDeleteFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7125,10 +7004,8 @@ func (m *awsRestxml_deserializeOpDeleteFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteFunction(response, &metadata)
@@ -7142,7 +7019,6 @@ func (m *awsRestxml_deserializeOpDeleteFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7223,10 +7099,8 @@ func (m *awsRestxml_deserializeOpDeleteKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteKeyGroup(response, &metadata)
@@ -7240,7 +7114,6 @@ func (m *awsRestxml_deserializeOpDeleteKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7318,10 +7191,8 @@ func (m *awsRestxml_deserializeOpDeleteKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteKeyValueStore(response, &metadata)
@@ -7335,7 +7206,6 @@ func (m *awsRestxml_deserializeOpDeleteKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7419,10 +7289,8 @@ func (m *awsRestxml_deserializeOpDeleteMonitoringSubsc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteMonitoringSubscription(response, &metadata)
@@ -7430,7 +7298,6 @@ func (m *awsRestxml_deserializeOpDeleteMonitoringSubsc
 	output := &DeleteMonitoringSubscriptionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7508,10 +7375,8 @@ func (m *awsRestxml_deserializeOpDeleteOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteOriginAccessControl(response, &metadata)
@@ -7525,7 +7390,6 @@ func (m *awsRestxml_deserializeOpDeleteOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7606,10 +7470,8 @@ func (m *awsRestxml_deserializeOpDeleteOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteOriginRequestPolicy(response, &metadata)
@@ -7623,7 +7485,6 @@ func (m *awsRestxml_deserializeOpDeleteOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7707,10 +7568,8 @@ func (m *awsRestxml_deserializeOpDeletePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeletePublicKey(response, &metadata)
@@ -7724,7 +7583,6 @@ func (m *awsRestxml_deserializeOpDeletePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7805,10 +7663,8 @@ func (m *awsRestxml_deserializeOpDeleteRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteRealtimeLogConfig(response, &metadata)
@@ -7822,7 +7678,6 @@ func (m *awsRestxml_deserializeOpDeleteRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7900,10 +7755,8 @@ func (m *awsRestxml_deserializeOpDeleteResourcePolicy)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteResourcePolicy(response, &metadata)
@@ -7917,7 +7770,6 @@ func (m *awsRestxml_deserializeOpDeleteResourcePolicy)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8001,10 +7853,8 @@ func (m *awsRestxml_deserializeOpDeleteResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteResponseHeadersPolicy(response, &metadata)
@@ -8018,7 +7868,6 @@ func (m *awsRestxml_deserializeOpDeleteResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8102,10 +7951,8 @@ func (m *awsRestxml_deserializeOpDeleteStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteStreamingDistribution(response, &metadata)
@@ -8119,7 +7966,6 @@ func (m *awsRestxml_deserializeOpDeleteStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8200,10 +8046,8 @@ func (m *awsRestxml_deserializeOpDeleteTrustStore) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteTrustStore(response, &metadata)
@@ -8217,7 +8061,6 @@ func (m *awsRestxml_deserializeOpDeleteTrustStore) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8301,10 +8144,8 @@ func (m *awsRestxml_deserializeOpDeleteVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDeleteVpcOrigin(response, &metadata)
@@ -8345,7 +8186,6 @@ func (m *awsRestxml_deserializeOpDeleteVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8489,10 +8329,8 @@ func (m *awsRestxml_deserializeOpDescribeConnectionFun
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDescribeConnectionFunction(response, &metadata)
@@ -8533,7 +8371,6 @@ func (m *awsRestxml_deserializeOpDescribeConnectionFun
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8665,10 +8502,8 @@ func (m *awsRestxml_deserializeOpDescribeFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDescribeFunction(response, &metadata)
@@ -8709,7 +8544,6 @@ func (m *awsRestxml_deserializeOpDescribeFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8835,10 +8669,8 @@ func (m *awsRestxml_deserializeOpDescribeKeyValueStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDescribeKeyValueStore(response, &metadata)
@@ -8879,7 +8711,6 @@ func (m *awsRestxml_deserializeOpDescribeKeyValueStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9011,10 +8842,8 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDisassociateDistributionTenantWebACL(response, &metadata)
@@ -9055,7 +8884,6 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9197,10 +9025,8 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorDisassociateDistributionWebACL(response, &metadata)
@@ -9241,7 +9067,6 @@ func (m *awsRestxml_deserializeOpDisassociateDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9383,10 +9208,8 @@ func (m *awsRestxml_deserializeOpGetAnycastIpList) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetAnycastIpList(response, &metadata)
@@ -9427,7 +9250,6 @@ func (m *awsRestxml_deserializeOpGetAnycastIpList) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9559,10 +9381,8 @@ func (m *awsRestxml_deserializeOpGetCachePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetCachePolicy(response, &metadata)
@@ -9603,7 +9423,6 @@ func (m *awsRestxml_deserializeOpGetCachePolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9729,10 +9548,8 @@ func (m *awsRestxml_deserializeOpGetCachePolicyConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetCachePolicyConfig(response, &metadata)
@@ -9773,7 +9590,6 @@ func (m *awsRestxml_deserializeOpGetCachePolicyConfig)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9899,10 +9715,8 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetCloudFrontOriginAccessIdentity(response, &metadata)
@@ -9943,7 +9757,6 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10069,10 +9882,8 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetCloudFrontOriginAccessIdentityConfig(response, &metadata)
@@ -10113,7 +9924,6 @@ func (m *awsRestxml_deserializeOpGetCloudFrontOriginAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10239,10 +10049,8 @@ func (m *awsRestxml_deserializeOpGetConnectionFunction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetConnectionFunction(response, &metadata)
@@ -10260,7 +10068,6 @@ func (m *awsRestxml_deserializeOpGetConnectionFunction
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10373,10 +10180,8 @@ func (m *awsRestxml_deserializeOpGetConnectionGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetConnectionGroup(response, &metadata)
@@ -10417,7 +10222,6 @@ func (m *awsRestxml_deserializeOpGetConnectionGroup) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10543,10 +10347,8 @@ func (m *awsRestxml_deserializeOpGetConnectionGroupByR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetConnectionGroupByRoutingEndpoint(response, &metadata)
@@ -10587,7 +10389,6 @@ func (m *awsRestxml_deserializeOpGetConnectionGroupByR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10713,10 +10514,8 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetContinuousDeploymentPolicy(response, &metadata)
@@ -10757,7 +10556,6 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10883,10 +10681,8 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetContinuousDeploymentPolicyConfig(response, &metadata)
@@ -10927,7 +10723,6 @@ func (m *awsRestxml_deserializeOpGetContinuousDeployme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11053,10 +10848,8 @@ func (m *awsRestxml_deserializeOpGetDistribution) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDistribution(response, &metadata)
@@ -11097,7 +10890,6 @@ func (m *awsRestxml_deserializeOpGetDistribution) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11223,10 +11015,8 @@ func (m *awsRestxml_deserializeOpGetDistributionConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDistributionConfig(response, &metadata)
@@ -11267,7 +11057,6 @@ func (m *awsRestxml_deserializeOpGetDistributionConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11393,10 +11182,8 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDistributionTenant(response, &metadata)
@@ -11437,7 +11224,6 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11563,10 +11349,8 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetDistributionTenantByDomain(response, &metadata)
@@ -11607,7 +11391,6 @@ func (m *awsRestxml_deserializeOpGetDistributionTenant
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11733,10 +11516,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetFieldLevelEncryption(response, &metadata)
@@ -11777,7 +11558,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11903,10 +11683,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetFieldLevelEncryptionConfig(response, &metadata)
@@ -11947,7 +11725,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12073,10 +11850,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetFieldLevelEncryptionProfile(response, &metadata)
@@ -12117,7 +11892,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12243,10 +12017,8 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetFieldLevelEncryptionProfileConfig(response, &metadata)
@@ -12287,7 +12059,6 @@ func (m *awsRestxml_deserializeOpGetFieldLevelEncrypti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12413,10 +12184,8 @@ func (m *awsRestxml_deserializeOpGetFunction) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetFunction(response, &metadata)
@@ -12434,7 +12203,6 @@ func (m *awsRestxml_deserializeOpGetFunction) HandleDe
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12544,10 +12312,8 @@ func (m *awsRestxml_deserializeOpGetInvalidation) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetInvalidation(response, &metadata)
@@ -12583,7 +12349,6 @@ func (m *awsRestxml_deserializeOpGetInvalidation) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12700,10 +12465,8 @@ func (m *awsRestxml_deserializeOpGetInvalidationForDis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetInvalidationForDistributionTenant(response, &metadata)
@@ -12739,7 +12502,6 @@ func (m *awsRestxml_deserializeOpGetInvalidationForDis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12856,10 +12618,8 @@ func (m *awsRestxml_deserializeOpGetKeyGroup) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetKeyGroup(response, &metadata)
@@ -12900,7 +12660,6 @@ func (m *awsRestxml_deserializeOpGetKeyGroup) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13023,10 +12782,8 @@ func (m *awsRestxml_deserializeOpGetKeyGroupConfig) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetKeyGroupConfig(response, &metadata)
@@ -13067,7 +12824,6 @@ func (m *awsRestxml_deserializeOpGetKeyGroupConfig) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13190,10 +12946,8 @@ func (m *awsRestxml_deserializeOpGetManagedCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetManagedCertificateDetails(response, &metadata)
@@ -13229,7 +12983,6 @@ func (m *awsRestxml_deserializeOpGetManagedCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13343,10 +13096,8 @@ func (m *awsRestxml_deserializeOpGetMonitoringSubscrip
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetMonitoringSubscription(response, &metadata)
@@ -13382,7 +13133,6 @@ func (m *awsRestxml_deserializeOpGetMonitoringSubscrip
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13502,10 +13252,8 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetOriginAccessControl(response, &metadata)
@@ -13546,7 +13294,6 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13672,10 +13419,8 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetOriginAccessControlConfig(response, &metadata)
@@ -13716,7 +13461,6 @@ func (m *awsRestxml_deserializeOpGetOriginAccessContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13842,10 +13586,8 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetOriginRequestPolicy(response, &metadata)
@@ -13886,7 +13628,6 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14012,10 +13753,8 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetOriginRequestPolicyConfig(response, &metadata)
@@ -14056,7 +13795,6 @@ func (m *awsRestxml_deserializeOpGetOriginRequestPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14182,10 +13920,8 @@ func (m *awsRestxml_deserializeOpGetPublicKey) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetPublicKey(response, &metadata)
@@ -14226,7 +13962,6 @@ func (m *awsRestxml_deserializeOpGetPublicKey) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14352,10 +14087,8 @@ func (m *awsRestxml_deserializeOpGetPublicKeyConfig) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetPublicKeyConfig(response, &metadata)
@@ -14396,7 +14129,6 @@ func (m *awsRestxml_deserializeOpGetPublicKeyConfig) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14522,10 +14254,8 @@ func (m *awsRestxml_deserializeOpGetRealtimeLogConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetRealtimeLogConfig(response, &metadata)
@@ -14561,7 +14291,6 @@ func (m *awsRestxml_deserializeOpGetRealtimeLogConfig)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14678,10 +14407,8 @@ func (m *awsRestxml_deserializeOpGetResourcePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetResourcePolicy(response, &metadata)
@@ -14717,7 +14444,6 @@ func (m *awsRestxml_deserializeOpGetResourcePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14857,10 +14583,8 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetResponseHeadersPolicy(response, &metadata)
@@ -14901,7 +14625,6 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15027,10 +14750,8 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetResponseHeadersPolicyConfig(response, &metadata)
@@ -15071,7 +14792,6 @@ func (m *awsRestxml_deserializeOpGetResponseHeadersPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15197,10 +14917,8 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetStreamingDistribution(response, &metadata)
@@ -15241,7 +14959,6 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15367,10 +15084,8 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetStreamingDistributionConfig(response, &metadata)
@@ -15411,7 +15126,6 @@ func (m *awsRestxml_deserializeOpGetStreamingDistribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15537,10 +15251,8 @@ func (m *awsRestxml_deserializeOpGetTrustStore) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetTrustStore(response, &metadata)
@@ -15581,7 +15293,6 @@ func (m *awsRestxml_deserializeOpGetTrustStore) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15710,10 +15421,8 @@ func (m *awsRestxml_deserializeOpGetVpcOrigin) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorGetVpcOrigin(response, &metadata)
@@ -15754,7 +15463,6 @@ func (m *awsRestxml_deserializeOpGetVpcOrigin) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15886,10 +15594,8 @@ func (m *awsRestxml_deserializeOpListAnycastIpLists) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListAnycastIpLists(response, &metadata)
@@ -15925,7 +15631,6 @@ func (m *awsRestxml_deserializeOpListAnycastIpLists) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16045,10 +15750,8 @@ func (m *awsRestxml_deserializeOpListCachePolicies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCachePolicies(response, &metadata)
@@ -16084,7 +15787,6 @@ func (m *awsRestxml_deserializeOpListCachePolicies) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16201,10 +15903,8 @@ func (m *awsRestxml_deserializeOpListCloudFrontOriginA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListCloudFrontOriginAccessIdentities(response, &metadata)
@@ -16240,7 +15940,6 @@ func (m *awsRestxml_deserializeOpListCloudFrontOriginA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16351,10 +16050,8 @@ func (m *awsRestxml_deserializeOpListConflictingAliase
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListConflictingAliases(response, &metadata)
@@ -16390,7 +16087,6 @@ func (m *awsRestxml_deserializeOpListConflictingAliase
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16504,10 +16200,8 @@ func (m *awsRestxml_deserializeOpListConnectionFunctio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListConnectionFunctions(response, &metadata)
@@ -16543,7 +16237,6 @@ func (m *awsRestxml_deserializeOpListConnectionFunctio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16673,10 +16366,8 @@ func (m *awsRestxml_deserializeOpListConnectionGroups)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListConnectionGroups(response, &metadata)
@@ -16712,7 +16403,6 @@ func (m *awsRestxml_deserializeOpListConnectionGroups)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16842,10 +16532,8 @@ func (m *awsRestxml_deserializeOpListContinuousDeploym
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListContinuousDeploymentPolicies(response, &metadata)
@@ -16881,7 +16569,6 @@ func (m *awsRestxml_deserializeOpListContinuousDeploym
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16998,10 +16685,8 @@ func (m *awsRestxml_deserializeOpListDistributions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributions(response, &metadata)
@@ -17037,7 +16722,6 @@ func (m *awsRestxml_deserializeOpListDistributions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17148,10 +16832,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByAnycastIpListId(response, &metadata)
@@ -17187,7 +16869,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17307,10 +16988,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByCa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByCachePolicyId(response, &metadata)
@@ -17346,7 +17025,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByCa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17463,10 +17141,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByConnectionFunction(response, &metadata)
@@ -17502,7 +17178,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17619,10 +17294,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByConnectionMode(response, &metadata)
@@ -17658,7 +17331,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17772,10 +17444,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByKe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByKeyGroup(response, &metadata)
@@ -17811,7 +17481,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByKe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17925,10 +17594,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByOr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByOriginRequestPolicyId(response, &metadata)
@@ -17964,7 +17631,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByOr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18081,10 +17747,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByOw
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByOwnedResource(response, &metadata)
@@ -18120,7 +17784,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByOw
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18240,10 +17903,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByRealtimeLogConfig(response, &metadata)
@@ -18279,7 +17940,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18390,10 +18050,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByResponseHeadersPolicyId(response, &metadata)
@@ -18429,7 +18087,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18546,10 +18203,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByTrustStore(response, &metadata)
@@ -18585,7 +18240,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18702,10 +18356,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByVp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByVpcOriginId(response, &metadata)
@@ -18741,7 +18393,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByVp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18861,10 +18512,8 @@ func (m *awsRestxml_deserializeOpListDistributionsByWe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionsByWebACLId(response, &metadata)
@@ -18900,7 +18549,6 @@ func (m *awsRestxml_deserializeOpListDistributionsByWe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19014,10 +18662,8 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionTenants(response, &metadata)
@@ -19053,7 +18699,6 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19183,10 +18828,8 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDistributionTenantsByCustomization(response, &metadata)
@@ -19222,7 +18865,6 @@ func (m *awsRestxml_deserializeOpListDistributionTenan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19352,10 +18994,8 @@ func (m *awsRestxml_deserializeOpListDomainConflicts) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListDomainConflicts(response, &metadata)
@@ -19391,7 +19031,6 @@ func (m *awsRestxml_deserializeOpListDomainConflicts) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19521,10 +19160,8 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListFieldLevelEncryptionConfigs(response, &metadata)
@@ -19560,7 +19197,6 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19671,10 +19307,8 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListFieldLevelEncryptionProfiles(response, &metadata)
@@ -19710,7 +19344,6 @@ func (m *awsRestxml_deserializeOpListFieldLevelEncrypt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19821,10 +19454,8 @@ func (m *awsRestxml_deserializeOpListFunctions) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListFunctions(response, &metadata)
@@ -19860,7 +19491,6 @@ func (m *awsRestxml_deserializeOpListFunctions) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19974,10 +19604,8 @@ func (m *awsRestxml_deserializeOpListInvalidations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListInvalidations(response, &metadata)
@@ -20013,7 +19641,6 @@ func (m *awsRestxml_deserializeOpListInvalidations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20130,10 +19757,8 @@ func (m *awsRestxml_deserializeOpListInvalidationsForD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListInvalidationsForDistributionTenant(response, &metadata)
@@ -20169,7 +19794,6 @@ func (m *awsRestxml_deserializeOpListInvalidationsForD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20286,10 +19910,8 @@ func (m *awsRestxml_deserializeOpListKeyGroups) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListKeyGroups(response, &metadata)
@@ -20325,7 +19947,6 @@ func (m *awsRestxml_deserializeOpListKeyGroups) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20436,10 +20057,8 @@ func (m *awsRestxml_deserializeOpListKeyValueStores) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListKeyValueStores(response, &metadata)
@@ -20475,7 +20094,6 @@ func (m *awsRestxml_deserializeOpListKeyValueStores) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20592,10 +20210,8 @@ func (m *awsRestxml_deserializeOpListOriginAccessContr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListOriginAccessControls(response, &metadata)
@@ -20631,7 +20247,6 @@ func (m *awsRestxml_deserializeOpListOriginAccessContr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20742,10 +20357,8 @@ func (m *awsRestxml_deserializeOpListOriginRequestPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListOriginRequestPolicies(response, &metadata)
@@ -20781,7 +20394,6 @@ func (m *awsRestxml_deserializeOpListOriginRequestPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20898,10 +20510,8 @@ func (m *awsRestxml_deserializeOpListPublicKeys) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListPublicKeys(response, &metadata)
@@ -20937,7 +20547,6 @@ func (m *awsRestxml_deserializeOpListPublicKeys) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21048,10 +20657,8 @@ func (m *awsRestxml_deserializeOpListRealtimeLogConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListRealtimeLogConfigs(response, &metadata)
@@ -21087,7 +20694,6 @@ func (m *awsRestxml_deserializeOpListRealtimeLogConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21204,10 +20810,8 @@ func (m *awsRestxml_deserializeOpListResponseHeadersPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListResponseHeadersPolicies(response, &metadata)
@@ -21243,7 +20847,6 @@ func (m *awsRestxml_deserializeOpListResponseHeadersPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21360,10 +20963,8 @@ func (m *awsRestxml_deserializeOpListStreamingDistribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListStreamingDistributions(response, &metadata)
@@ -21399,7 +21000,6 @@ func (m *awsRestxml_deserializeOpListStreamingDistribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21510,10 +21110,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTagsForResource(response, &metadata)
@@ -21549,7 +21147,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21669,10 +21266,8 @@ func (m *awsRestxml_deserializeOpListTrustStores) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListTrustStores(response, &metadata)
@@ -21708,7 +21303,6 @@ func (m *awsRestxml_deserializeOpListTrustStores) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21838,10 +21432,8 @@ func (m *awsRestxml_deserializeOpListVpcOrigins) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorListVpcOrigins(response, &metadata)
@@ -21877,7 +21469,6 @@ func (m *awsRestxml_deserializeOpListVpcOrigins) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21997,10 +21588,8 @@ func (m *awsRestxml_deserializeOpPublishConnectionFunc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPublishConnectionFunction(response, &metadata)
@@ -22036,7 +21625,6 @@ func (m *awsRestxml_deserializeOpPublishConnectionFunc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22162,10 +21750,8 @@ func (m *awsRestxml_deserializeOpPublishFunction) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPublishFunction(response, &metadata)
@@ -22201,7 +21787,6 @@ func (m *awsRestxml_deserializeOpPublishFunction) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22324,10 +21909,8 @@ func (m *awsRestxml_deserializeOpPutResourcePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorPutResourcePolicy(response, &metadata)
@@ -22363,7 +21946,6 @@ func (m *awsRestxml_deserializeOpPutResourcePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22496,10 +22078,8 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorTagResource(response, &metadata)
@@ -22513,7 +22093,6 @@ func (m *awsRestxml_deserializeOpTagResource) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22591,10 +22170,8 @@ func (m *awsRestxml_deserializeOpTestConnectionFunctio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorTestConnectionFunction(response, &metadata)
@@ -22630,7 +22207,6 @@ func (m *awsRestxml_deserializeOpTestConnectionFunctio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22756,10 +22332,8 @@ func (m *awsRestxml_deserializeOpTestFunction) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorTestFunction(response, &metadata)
@@ -22795,7 +22369,6 @@ func (m *awsRestxml_deserializeOpTestFunction) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22918,10 +22491,8 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUntagResource(response, &metadata)
@@ -22935,7 +22506,6 @@ func (m *awsRestxml_deserializeOpUntagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23013,10 +22583,8 @@ func (m *awsRestxml_deserializeOpUpdateAnycastIpList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateAnycastIpList(response, &metadata)
@@ -23057,7 +22625,6 @@ func (m *awsRestxml_deserializeOpUpdateAnycastIpList) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23195,10 +22762,8 @@ func (m *awsRestxml_deserializeOpUpdateCachePolicy) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateCachePolicy(response, &metadata)
@@ -23239,7 +22804,6 @@ func (m *awsRestxml_deserializeOpUpdateCachePolicy) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23392,10 +22956,8 @@ func (m *awsRestxml_deserializeOpUpdateCloudFrontOrigi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateCloudFrontOriginAccessIdentity(response, &metadata)
@@ -23436,7 +22998,6 @@ func (m *awsRestxml_deserializeOpUpdateCloudFrontOrigi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23580,10 +23141,8 @@ func (m *awsRestxml_deserializeOpUpdateConnectionFunct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateConnectionFunction(response, &metadata)
@@ -23624,7 +23183,6 @@ func (m *awsRestxml_deserializeOpUpdateConnectionFunct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23765,10 +23323,8 @@ func (m *awsRestxml_deserializeOpUpdateConnectionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateConnectionGroup(response, &metadata)
@@ -23809,7 +23365,6 @@ func (m *awsRestxml_deserializeOpUpdateConnectionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23953,10 +23508,8 @@ func (m *awsRestxml_deserializeOpUpdateContinuousDeplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateContinuousDeploymentPolicy(response, &metadata)
@@ -23997,7 +23550,6 @@ func (m *awsRestxml_deserializeOpUpdateContinuousDeplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24138,10 +23690,8 @@ func (m *awsRestxml_deserializeOpUpdateDistribution) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateDistribution(response, &metadata)
@@ -24182,7 +23732,6 @@ func (m *awsRestxml_deserializeOpUpdateDistribution) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24503,10 +24052,8 @@ func (m *awsRestxml_deserializeOpUpdateDistributionTen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateDistributionTenant(response, &metadata)
@@ -24547,7 +24094,6 @@ func (m *awsRestxml_deserializeOpUpdateDistributionTen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24694,10 +24240,8 @@ func (m *awsRestxml_deserializeOpUpdateDistributionWit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateDistributionWithStagingConfig(response, &metadata)
@@ -24738,7 +24282,6 @@ func (m *awsRestxml_deserializeOpUpdateDistributionWit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25047,10 +24590,8 @@ func (m *awsRestxml_deserializeOpUpdateDomainAssociati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateDomainAssociation(response, &metadata)
@@ -25091,7 +24632,6 @@ func (m *awsRestxml_deserializeOpUpdateDomainAssociati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25249,10 +24789,8 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateFieldLevelEncryptionConfig(response, &metadata)
@@ -25293,7 +24831,6 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25446,10 +24983,8 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateFieldLevelEncryptionProfile(response, &metadata)
@@ -25490,7 +25025,6 @@ func (m *awsRestxml_deserializeOpUpdateFieldLevelEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25646,10 +25180,8 @@ func (m *awsRestxml_deserializeOpUpdateFunction) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateFunction(response, &metadata)
@@ -25690,7 +25222,6 @@ func (m *awsRestxml_deserializeOpUpdateFunction) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25828,10 +25359,8 @@ func (m *awsRestxml_deserializeOpUpdateKeyGroup) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateKeyGroup(response, &metadata)
@@ -25872,7 +25401,6 @@ func (m *awsRestxml_deserializeOpUpdateKeyGroup) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26010,10 +25538,8 @@ func (m *awsRestxml_deserializeOpUpdateKeyValueStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateKeyValueStore(response, &metadata)
@@ -26054,7 +25580,6 @@ func (m *awsRestxml_deserializeOpUpdateKeyValueStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26192,10 +25717,8 @@ func (m *awsRestxml_deserializeOpUpdateOriginAccessCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateOriginAccessControl(response, &metadata)
@@ -26236,7 +25759,6 @@ func (m *awsRestxml_deserializeOpUpdateOriginAccessCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26377,10 +25899,8 @@ func (m *awsRestxml_deserializeOpUpdateOriginRequestPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateOriginRequestPolicy(response, &metadata)
@@ -26421,7 +25941,6 @@ func (m *awsRestxml_deserializeOpUpdateOriginRequestPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26574,10 +26093,8 @@ func (m *awsRestxml_deserializeOpUpdatePublicKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdatePublicKey(response, &metadata)
@@ -26618,7 +26135,6 @@ func (m *awsRestxml_deserializeOpUpdatePublicKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26759,10 +26275,8 @@ func (m *awsRestxml_deserializeOpUpdateRealtimeLogConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateRealtimeLogConfig(response, &metadata)
@@ -26798,7 +26312,6 @@ func (m *awsRestxml_deserializeOpUpdateRealtimeLogConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26915,10 +26428,8 @@ func (m *awsRestxml_deserializeOpUpdateResponseHeaders
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateResponseHeadersPolicy(response, &metadata)
@@ -26959,7 +26470,6 @@ func (m *awsRestxml_deserializeOpUpdateResponseHeaders
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27112,10 +26622,8 @@ func (m *awsRestxml_deserializeOpUpdateStreamingDistri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateStreamingDistribution(response, &metadata)
@@ -27156,7 +26664,6 @@ func (m *awsRestxml_deserializeOpUpdateStreamingDistri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27318,10 +26825,8 @@ func (m *awsRestxml_deserializeOpUpdateTrustStore) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateTrustStore(response, &metadata)
@@ -27362,7 +26867,6 @@ func (m *awsRestxml_deserializeOpUpdateTrustStore) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27497,10 +27001,8 @@ func (m *awsRestxml_deserializeOpUpdateVpcOrigin) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorUpdateVpcOrigin(response, &metadata)
@@ -27541,7 +27043,6 @@ func (m *awsRestxml_deserializeOpUpdateVpcOrigin) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27694,10 +27195,8 @@ func (m *awsRestxml_deserializeOpVerifyDnsConfiguratio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if response.StatusCode < 200 || response.StatusCode >= 300 {
 		return out, metadata, awsRestxml_deserializeOpErrorVerifyDnsConfiguration(response, &metadata)
@@ -27733,7 +27232,6 @@ func (m *awsRestxml_deserializeOpVerifyDnsConfiguratio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
