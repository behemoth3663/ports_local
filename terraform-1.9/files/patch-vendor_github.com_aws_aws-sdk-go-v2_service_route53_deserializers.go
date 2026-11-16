--- vendor/github.com/aws/aws-sdk-go-v2/service/route53/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -47,10 +46,8 @@ func (m *awsRestxml_deserializeOpActivateKeySigningKey
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestxml_deserializeOpActivateKeySigningKey
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -204,10 +200,8 @@ func (m *awsRestxml_deserializeOpAssociateVPCWithHoste
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -247,7 +241,6 @@ func (m *awsRestxml_deserializeOpAssociateVPCWithHoste
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -367,10 +360,8 @@ func (m *awsRestxml_deserializeOpChangeCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -410,7 +401,6 @@ func (m *awsRestxml_deserializeOpChangeCidrCollection)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -531,10 +521,8 @@ func (m *awsRestxml_deserializeOpChangeResourceRecordS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -574,7 +562,6 @@ func (m *awsRestxml_deserializeOpChangeResourceRecordS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -685,10 +672,8 @@ func (m *awsRestxml_deserializeOpChangeTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -700,7 +685,6 @@ func (m *awsRestxml_deserializeOpChangeTagsForResource
 	output := &ChangeTagsForResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -769,10 +753,8 @@ func (m *awsRestxml_deserializeOpCreateCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -817,7 +799,6 @@ func (m *awsRestxml_deserializeOpCreateCidrCollection)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -937,10 +918,8 @@ func (m *awsRestxml_deserializeOpCreateHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -985,7 +964,6 @@ func (m *awsRestxml_deserializeOpCreateHealthCheck) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1102,10 +1080,8 @@ func (m *awsRestxml_deserializeOpCreateHostedZone) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1150,7 +1126,6 @@ func (m *awsRestxml_deserializeOpCreateHostedZone) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1303,10 +1278,8 @@ func (m *awsRestxml_deserializeOpCreateKeySigningKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1351,7 +1324,6 @@ func (m *awsRestxml_deserializeOpCreateKeySigningKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1495,10 +1467,8 @@ func (m *awsRestxml_deserializeOpCreateQueryLoggingCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1543,7 +1513,6 @@ func (m *awsRestxml_deserializeOpCreateQueryLoggingCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1669,10 +1638,8 @@ func (m *awsRestxml_deserializeOpCreateReusableDelegat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1717,7 +1684,6 @@ func (m *awsRestxml_deserializeOpCreateReusableDelegat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1846,10 +1812,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1894,7 +1858,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2014,10 +1977,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2062,7 +2023,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2185,10 +2145,8 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2233,7 +2191,6 @@ func (m *awsRestxml_deserializeOpCreateTrafficPolicyVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2356,10 +2313,8 @@ func (m *awsRestxml_deserializeOpCreateVPCAssociationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2399,7 +2354,6 @@ func (m *awsRestxml_deserializeOpCreateVPCAssociationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2523,10 +2477,8 @@ func (m *awsRestxml_deserializeOpDeactivateKeySigningK
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2566,7 +2518,6 @@ func (m *awsRestxml_deserializeOpDeactivateKeySigningK
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2683,10 +2634,8 @@ func (m *awsRestxml_deserializeOpDeleteCidrCollection)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2698,7 +2647,6 @@ func (m *awsRestxml_deserializeOpDeleteCidrCollection)
 	output := &DeleteCidrCollectionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2764,10 +2712,8 @@ func (m *awsRestxml_deserializeOpDeleteHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2779,7 +2725,6 @@ func (m *awsRestxml_deserializeOpDeleteHealthCheck) Ha
 	output := &DeleteHealthCheckOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2842,10 +2787,8 @@ func (m *awsRestxml_deserializeOpDeleteHostedZone) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2885,7 +2828,6 @@ func (m *awsRestxml_deserializeOpDeleteHostedZone) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2996,10 +2938,8 @@ func (m *awsRestxml_deserializeOpDeleteKeySigningKey) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3039,7 +2979,6 @@ func (m *awsRestxml_deserializeOpDeleteKeySigningKey) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3153,10 +3092,8 @@ func (m *awsRestxml_deserializeOpDeleteQueryLoggingCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3168,7 +3105,6 @@ func (m *awsRestxml_deserializeOpDeleteQueryLoggingCon
 	output := &DeleteQueryLoggingConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3231,10 +3167,8 @@ func (m *awsRestxml_deserializeOpDeleteReusableDelegat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3246,7 +3180,6 @@ func (m *awsRestxml_deserializeOpDeleteReusableDelegat
 	output := &DeleteReusableDelegationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3312,10 +3245,8 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3327,7 +3258,6 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicy) 
 	output := &DeleteTrafficPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3393,10 +3323,8 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3408,7 +3336,6 @@ func (m *awsRestxml_deserializeOpDeleteTrafficPolicyIn
 	output := &DeleteTrafficPolicyInstanceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3471,10 +3398,8 @@ func (m *awsRestxml_deserializeOpDeleteVPCAssociationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3486,7 +3411,6 @@ func (m *awsRestxml_deserializeOpDeleteVPCAssociationA
 	output := &DeleteVPCAssociationAuthorizationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3555,10 +3479,8 @@ func (m *awsRestxml_deserializeOpDisableHostedZoneDNSS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3598,7 +3520,6 @@ func (m *awsRestxml_deserializeOpDisableHostedZoneDNSS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3718,10 +3639,8 @@ func (m *awsRestxml_deserializeOpDisassociateVPCFromHo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3761,7 +3680,6 @@ func (m *awsRestxml_deserializeOpDisassociateVPCFromHo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3872,10 +3790,8 @@ func (m *awsRestxml_deserializeOpEnableHostedZoneDNSSE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3915,7 +3831,6 @@ func (m *awsRestxml_deserializeOpEnableHostedZoneDNSSE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4038,10 +3953,8 @@ func (m *awsRestxml_deserializeOpGetAccountLimit) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4081,7 +3994,6 @@ func (m *awsRestxml_deserializeOpGetAccountLimit) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4197,10 +4109,8 @@ func (m *awsRestxml_deserializeOpGetChange) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4240,7 +4150,6 @@ func (m *awsRestxml_deserializeOpGetChange) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4342,10 +4251,8 @@ func (m *awsRestxml_deserializeOpGetCheckerIpRanges) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4385,7 +4292,6 @@ func (m *awsRestxml_deserializeOpGetCheckerIpRanges) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4481,10 +4387,8 @@ func (m *awsRestxml_deserializeOpGetDNSSEC) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4524,7 +4428,6 @@ func (m *awsRestxml_deserializeOpGetDNSSEC) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4635,10 +4538,8 @@ func (m *awsRestxml_deserializeOpGetGeoLocation) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4678,7 +4579,6 @@ func (m *awsRestxml_deserializeOpGetGeoLocation) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4780,10 +4680,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheck) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4823,7 +4721,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheck) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4928,10 +4825,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckCount) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4971,7 +4866,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckCount) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5078,10 +4972,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckLastFai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5121,7 +5013,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckLastFai
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5223,10 +5114,8 @@ func (m *awsRestxml_deserializeOpGetHealthCheckStatus)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5266,7 +5155,6 @@ func (m *awsRestxml_deserializeOpGetHealthCheckStatus)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5368,10 +5256,8 @@ func (m *awsRestxml_deserializeOpGetHostedZone) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5411,7 +5297,6 @@ func (m *awsRestxml_deserializeOpGetHostedZone) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5525,10 +5410,8 @@ func (m *awsRestxml_deserializeOpGetHostedZoneCount) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5568,7 +5451,6 @@ func (m *awsRestxml_deserializeOpGetHostedZoneCount) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5678,10 +5560,8 @@ func (m *awsRestxml_deserializeOpGetHostedZoneLimit) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5721,7 +5601,6 @@ func (m *awsRestxml_deserializeOpGetHostedZoneLimit) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5843,10 +5722,8 @@ func (m *awsRestxml_deserializeOpGetQueryLoggingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5886,7 +5763,6 @@ func (m *awsRestxml_deserializeOpGetQueryLoggingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5988,10 +5864,8 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6031,7 +5905,6 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6136,10 +6009,8 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6179,7 +6050,6 @@ func (m *awsRestxml_deserializeOpGetReusableDelegation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6298,10 +6168,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6341,7 +6209,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6443,10 +6310,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6486,7 +6351,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6588,10 +6452,8 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6631,7 +6493,6 @@ func (m *awsRestxml_deserializeOpGetTrafficPolicyInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6738,10 +6599,8 @@ func (m *awsRestxml_deserializeOpListCidrBlocks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6781,7 +6640,6 @@ func (m *awsRestxml_deserializeOpListCidrBlocks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6899,10 +6757,8 @@ func (m *awsRestxml_deserializeOpListCidrCollections) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6942,7 +6798,6 @@ func (m *awsRestxml_deserializeOpListCidrCollections) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7054,10 +6909,8 @@ func (m *awsRestxml_deserializeOpListCidrLocations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7097,7 +6950,6 @@ func (m *awsRestxml_deserializeOpListCidrLocations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7212,10 +7064,8 @@ func (m *awsRestxml_deserializeOpListGeoLocations) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7255,7 +7105,6 @@ func (m *awsRestxml_deserializeOpListGeoLocations) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7426,10 +7275,8 @@ func (m *awsRestxml_deserializeOpListHealthChecks) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7469,7 +7316,6 @@ func (m *awsRestxml_deserializeOpListHealthChecks) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7630,10 +7476,8 @@ func (m *awsRestxml_deserializeOpListHostedZones) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7673,7 +7517,6 @@ func (m *awsRestxml_deserializeOpListHostedZones) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7837,10 +7680,8 @@ func (m *awsRestxml_deserializeOpListHostedZonesByName
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7880,7 +7721,6 @@ func (m *awsRestxml_deserializeOpListHostedZonesByName
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8067,10 +7907,8 @@ func (m *awsRestxml_deserializeOpListHostedZonesByVPC)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8110,7 +7948,6 @@ func (m *awsRestxml_deserializeOpListHostedZonesByVPC)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8242,10 +8079,8 @@ func (m *awsRestxml_deserializeOpListQueryLoggingConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8285,7 +8120,6 @@ func (m *awsRestxml_deserializeOpListQueryLoggingConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8403,10 +8237,8 @@ func (m *awsRestxml_deserializeOpListResourceRecordSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8446,7 +8278,6 @@ func (m *awsRestxml_deserializeOpListResourceRecordSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8620,10 +8451,8 @@ func (m *awsRestxml_deserializeOpListReusableDelegatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8663,7 +8492,6 @@ func (m *awsRestxml_deserializeOpListReusableDelegatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8821,10 +8649,8 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8864,7 +8690,6 @@ func (m *awsRestxml_deserializeOpListTagsForResource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8975,10 +8800,8 @@ func (m *awsRestxml_deserializeOpListTagsForResources)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9018,7 +8841,6 @@ func (m *awsRestxml_deserializeOpListTagsForResources)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9129,10 +8951,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicies) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9172,7 +8992,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicies) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9317,10 +9136,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9360,7 +9177,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9534,10 +9350,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9577,7 +9391,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9741,10 +9554,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9784,7 +9595,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9961,10 +9771,8 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10004,7 +9812,6 @@ func (m *awsRestxml_deserializeOpListTrafficPolicyVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10152,10 +9959,8 @@ func (m *awsRestxml_deserializeOpListVPCAssociationAut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10195,7 +10000,6 @@ func (m *awsRestxml_deserializeOpListVPCAssociationAut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10326,10 +10130,8 @@ func (m *awsRestxml_deserializeOpTestDNSAnswer) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10369,7 +10171,6 @@ func (m *awsRestxml_deserializeOpTestDNSAnswer) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10536,10 +10337,8 @@ func (m *awsRestxml_deserializeOpUpdateHealthCheck) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10579,7 +10378,6 @@ func (m *awsRestxml_deserializeOpUpdateHealthCheck) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10684,10 +10482,8 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneComme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10727,7 +10523,6 @@ func (m *awsRestxml_deserializeOpUpdateHostedZoneComme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10832,10 +10627,8 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10875,7 +10668,6 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10980,10 +10772,8 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11023,7 +10813,6 @@ func (m *awsRestxml_deserializeOpUpdateTrafficPolicyIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
