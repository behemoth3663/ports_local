--- vendor/github.com/aws/aws-sdk-go-v2/service/rolesanywhere/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/rolesanywhere/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -200,10 +196,8 @@ func (m *awsRestjson1_deserializeOpCreateTrustAnchor) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -243,7 +237,6 @@ func (m *awsRestjson1_deserializeOpCreateTrustAnchor) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -355,10 +348,8 @@ func (m *awsRestjson1_deserializeOpDeleteAttributeMapp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -398,7 +389,6 @@ func (m *awsRestjson1_deserializeOpDeleteAttributeMapp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -513,10 +503,8 @@ func (m *awsRestjson1_deserializeOpDeleteCrl) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -556,7 +544,6 @@ func (m *awsRestjson1_deserializeOpDeleteCrl) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -668,10 +655,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -711,7 +696,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -823,10 +807,8 @@ func (m *awsRestjson1_deserializeOpDeleteTrustAnchor) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -866,7 +848,6 @@ func (m *awsRestjson1_deserializeOpDeleteTrustAnchor) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -978,10 +959,8 @@ func (m *awsRestjson1_deserializeOpDisableCrl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1021,7 +1000,6 @@ func (m *awsRestjson1_deserializeOpDisableCrl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1133,10 +1111,8 @@ func (m *awsRestjson1_deserializeOpDisableProfile) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1176,7 +1152,6 @@ func (m *awsRestjson1_deserializeOpDisableProfile) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1288,10 +1263,8 @@ func (m *awsRestjson1_deserializeOpDisableTrustAnchor)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1331,7 +1304,6 @@ func (m *awsRestjson1_deserializeOpDisableTrustAnchor)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1443,10 +1415,8 @@ func (m *awsRestjson1_deserializeOpEnableCrl) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1486,7 +1456,6 @@ func (m *awsRestjson1_deserializeOpEnableCrl) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1598,10 +1567,8 @@ func (m *awsRestjson1_deserializeOpEnableProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1641,7 +1608,6 @@ func (m *awsRestjson1_deserializeOpEnableProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1753,10 +1719,8 @@ func (m *awsRestjson1_deserializeOpEnableTrustAnchor) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1796,7 +1760,6 @@ func (m *awsRestjson1_deserializeOpEnableTrustAnchor) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1908,10 +1871,8 @@ func (m *awsRestjson1_deserializeOpGetCrl) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1951,7 +1912,6 @@ func (m *awsRestjson1_deserializeOpGetCrl) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2060,10 +2020,8 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2103,7 +2061,6 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2215,10 +2172,8 @@ func (m *awsRestjson1_deserializeOpGetSubject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2258,7 +2213,6 @@ func (m *awsRestjson1_deserializeOpGetSubject) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2370,10 +2324,8 @@ func (m *awsRestjson1_deserializeOpGetTrustAnchor) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2413,7 +2365,6 @@ func (m *awsRestjson1_deserializeOpGetTrustAnchor) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2528,10 +2479,8 @@ func (m *awsRestjson1_deserializeOpImportCrl) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2571,7 +2520,6 @@ func (m *awsRestjson1_deserializeOpImportCrl) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2683,10 +2631,8 @@ func (m *awsRestjson1_deserializeOpListCrls) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2726,7 +2672,6 @@ func (m *awsRestjson1_deserializeOpListCrls) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2847,10 +2792,8 @@ func (m *awsRestjson1_deserializeOpListProfiles) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2890,7 +2833,6 @@ func (m *awsRestjson1_deserializeOpListProfiles) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3011,10 +2953,8 @@ func (m *awsRestjson1_deserializeOpListSubjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3054,7 +2994,6 @@ func (m *awsRestjson1_deserializeOpListSubjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3175,10 +3114,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3218,7 +3155,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3333,10 +3269,8 @@ func (m *awsRestjson1_deserializeOpListTrustAnchors) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3376,7 +3310,6 @@ func (m *awsRestjson1_deserializeOpListTrustAnchors) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3497,10 +3430,8 @@ func (m *awsRestjson1_deserializeOpPutAttributeMapping
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3540,7 +3471,6 @@ func (m *awsRestjson1_deserializeOpPutAttributeMapping
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3655,10 +3585,8 @@ func (m *awsRestjson1_deserializeOpPutNotificationSett
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3698,7 +3626,6 @@ func (m *awsRestjson1_deserializeOpPutNotificationSett
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3813,10 +3740,8 @@ func (m *awsRestjson1_deserializeOpResetNotificationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3856,7 +3781,6 @@ func (m *awsRestjson1_deserializeOpResetNotificationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3971,10 +3895,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3986,7 +3908,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4068,10 +3989,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4083,7 +4002,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4162,10 +4080,8 @@ func (m *awsRestjson1_deserializeOpUpdateCrl) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4205,7 +4121,6 @@ func (m *awsRestjson1_deserializeOpUpdateCrl) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4320,10 +4235,8 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4363,7 +4276,6 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4478,10 +4390,8 @@ func (m *awsRestjson1_deserializeOpUpdateTrustAnchor) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4521,7 +4431,6 @@ func (m *awsRestjson1_deserializeOpUpdateTrustAnchor) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
