--- vendor/github.com/aws/aws-sdk-go-v2/service/wellarchitected/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/wellarchitected/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateLenses) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpAssociateLenses) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -154,10 +150,8 @@ func (m *awsRestjson1_deserializeOpAssociateProfiles) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -175,7 +169,6 @@ func (m *awsRestjson1_deserializeOpAssociateProfiles) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -263,10 +256,8 @@ func (m *awsRestjson1_deserializeOpCreateLensShare) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -306,7 +297,6 @@ func (m *awsRestjson1_deserializeOpCreateLensShare) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -437,10 +427,8 @@ func (m *awsRestjson1_deserializeOpCreateLensVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -480,7 +468,6 @@ func (m *awsRestjson1_deserializeOpCreateLensVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -620,10 +607,8 @@ func (m *awsRestjson1_deserializeOpCreateMilestone) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -663,7 +648,6 @@ func (m *awsRestjson1_deserializeOpCreateMilestone) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -807,10 +791,8 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -850,7 +832,6 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -987,10 +968,8 @@ func (m *awsRestjson1_deserializeOpCreateProfileShare)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1030,7 +1009,6 @@ func (m *awsRestjson1_deserializeOpCreateProfileShare)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1170,10 +1148,8 @@ func (m *awsRestjson1_deserializeOpCreateReviewTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1213,7 +1189,6 @@ func (m *awsRestjson1_deserializeOpCreateReviewTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1344,10 +1319,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplateShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1387,7 +1360,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplateShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1527,10 +1499,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkload) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1570,7 +1540,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkload) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1710,10 +1679,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkloadShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1753,7 +1720,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkloadShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1893,10 +1859,8 @@ func (m *awsRestjson1_deserializeOpDeleteLens) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1914,7 +1878,6 @@ func (m *awsRestjson1_deserializeOpDeleteLens) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2002,10 +1965,8 @@ func (m *awsRestjson1_deserializeOpDeleteLensShare) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2023,7 +1984,6 @@ func (m *awsRestjson1_deserializeOpDeleteLensShare) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2111,10 +2071,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2132,7 +2090,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2220,10 +2177,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfileShare)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2241,7 +2196,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfileShare)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2329,10 +2283,8 @@ func (m *awsRestjson1_deserializeOpDeleteReviewTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2350,7 +2302,6 @@ func (m *awsRestjson1_deserializeOpDeleteReviewTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2438,10 +2389,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2459,7 +2408,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2547,10 +2495,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkload) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2568,7 +2514,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkload) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2656,10 +2601,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkloadShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2677,7 +2620,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkloadShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2765,10 +2707,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLenses)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2786,7 +2726,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLenses)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2874,10 +2813,8 @@ func (m *awsRestjson1_deserializeOpDisassociateProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2895,7 +2832,6 @@ func (m *awsRestjson1_deserializeOpDisassociateProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2983,10 +2919,8 @@ func (m *awsRestjson1_deserializeOpExportLens) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3026,7 +2960,6 @@ func (m *awsRestjson1_deserializeOpExportLens) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3151,10 +3084,8 @@ func (m *awsRestjson1_deserializeOpGetAnswer) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3194,7 +3125,6 @@ func (m *awsRestjson1_deserializeOpGetAnswer) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3355,10 +3285,8 @@ func (m *awsRestjson1_deserializeOpGetConsolidatedRepo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3398,7 +3326,6 @@ func (m *awsRestjson1_deserializeOpGetConsolidatedRepo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3537,10 +3464,8 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3580,7 +3505,6 @@ func (m *awsRestjson1_deserializeOpGetGlobalSettings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3716,10 +3640,8 @@ func (m *awsRestjson1_deserializeOpGetLens) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3759,7 +3681,6 @@ func (m *awsRestjson1_deserializeOpGetLens) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3880,10 +3801,8 @@ func (m *awsRestjson1_deserializeOpGetLensReview) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3923,7 +3842,6 @@ func (m *awsRestjson1_deserializeOpGetLensReview) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4066,10 +3984,8 @@ func (m *awsRestjson1_deserializeOpGetLensReviewReport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4109,7 +4025,6 @@ func (m *awsRestjson1_deserializeOpGetLensReviewReport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4252,10 +4167,8 @@ func (m *awsRestjson1_deserializeOpGetLensVersionDiffe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4295,7 +4208,6 @@ func (m *awsRestjson1_deserializeOpGetLensVersionDiffe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4461,10 +4373,8 @@ func (m *awsRestjson1_deserializeOpGetMilestone) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4504,7 +4414,6 @@ func (m *awsRestjson1_deserializeOpGetMilestone) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4634,10 +4543,8 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4677,7 +4584,6 @@ func (m *awsRestjson1_deserializeOpGetProfile) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4798,10 +4704,8 @@ func (m *awsRestjson1_deserializeOpGetProfileTemplate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4841,7 +4745,6 @@ func (m *awsRestjson1_deserializeOpGetProfileTemplate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4962,10 +4865,8 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5005,7 +4906,6 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5126,10 +5026,8 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplateAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5169,7 +5067,6 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplateAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5308,10 +5205,8 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplateLe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5351,7 +5246,6 @@ func (m *awsRestjson1_deserializeOpGetReviewTemplateLe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5481,10 +5375,8 @@ func (m *awsRestjson1_deserializeOpGetWorkload) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5524,7 +5416,6 @@ func (m *awsRestjson1_deserializeOpGetWorkload) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5645,10 +5536,8 @@ func (m *awsRestjson1_deserializeOpImportLens) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5688,7 +5577,6 @@ func (m *awsRestjson1_deserializeOpImportLens) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5828,10 +5716,8 @@ func (m *awsRestjson1_deserializeOpListAnswers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5871,7 +5757,6 @@ func (m *awsRestjson1_deserializeOpListAnswers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6041,10 +5926,8 @@ func (m *awsRestjson1_deserializeOpListCheckDetails) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6084,7 +5967,6 @@ func (m *awsRestjson1_deserializeOpListCheckDetails) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6214,10 +6096,8 @@ func (m *awsRestjson1_deserializeOpListCheckSummaries)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6257,7 +6137,6 @@ func (m *awsRestjson1_deserializeOpListCheckSummaries)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6387,10 +6266,8 @@ func (m *awsRestjson1_deserializeOpListLenses) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6430,7 +6307,6 @@ func (m *awsRestjson1_deserializeOpListLenses) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6557,10 +6433,8 @@ func (m *awsRestjson1_deserializeOpListLensReviewImpro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6600,7 +6474,6 @@ func (m *awsRestjson1_deserializeOpListLensReviewImpro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6770,10 +6643,8 @@ func (m *awsRestjson1_deserializeOpListLensReviews) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6813,7 +6684,6 @@ func (m *awsRestjson1_deserializeOpListLensReviews) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6965,10 +6835,8 @@ func (m *awsRestjson1_deserializeOpListLensShares) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7008,7 +6876,6 @@ func (m *awsRestjson1_deserializeOpListLensShares) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7138,10 +7005,8 @@ func (m *awsRestjson1_deserializeOpListMilestones) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7181,7 +7046,6 @@ func (m *awsRestjson1_deserializeOpListMilestones) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7320,10 +7184,8 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7363,7 +7225,6 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7490,10 +7351,8 @@ func (m *awsRestjson1_deserializeOpListProfileNotifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7533,7 +7392,6 @@ func (m *awsRestjson1_deserializeOpListProfileNotifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7660,10 +7518,8 @@ func (m *awsRestjson1_deserializeOpListProfiles) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7703,7 +7559,6 @@ func (m *awsRestjson1_deserializeOpListProfiles) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7830,10 +7685,8 @@ func (m *awsRestjson1_deserializeOpListProfileShares) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7873,7 +7726,6 @@ func (m *awsRestjson1_deserializeOpListProfileShares) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8003,10 +7855,8 @@ func (m *awsRestjson1_deserializeOpListReviewTemplateA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8046,7 +7896,6 @@ func (m *awsRestjson1_deserializeOpListReviewTemplateA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8194,10 +8043,8 @@ func (m *awsRestjson1_deserializeOpListReviewTemplates
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8237,7 +8084,6 @@ func (m *awsRestjson1_deserializeOpListReviewTemplates
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8364,10 +8210,8 @@ func (m *awsRestjson1_deserializeOpListShareInvitation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8407,7 +8251,6 @@ func (m *awsRestjson1_deserializeOpListShareInvitation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8534,10 +8377,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8577,7 +8418,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8689,10 +8529,8 @@ func (m *awsRestjson1_deserializeOpListTemplateShares)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8732,7 +8570,6 @@ func (m *awsRestjson1_deserializeOpListTemplateShares)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8871,10 +8708,8 @@ func (m *awsRestjson1_deserializeOpListWorkloads) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8914,7 +8749,6 @@ func (m *awsRestjson1_deserializeOpListWorkloads) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9041,10 +8875,8 @@ func (m *awsRestjson1_deserializeOpListWorkloadShares)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9084,7 +8916,6 @@ func (m *awsRestjson1_deserializeOpListWorkloadShares)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9223,10 +9054,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9238,7 +9067,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9314,10 +9142,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9329,7 +9155,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9405,10 +9230,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnswer) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9448,7 +9271,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnswer) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9599,10 +9421,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9620,7 +9440,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9705,10 +9524,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9726,7 +9543,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9814,10 +9630,8 @@ func (m *awsRestjson1_deserializeOpUpdateLensReview) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9857,7 +9671,6 @@ func (m *awsRestjson1_deserializeOpUpdateLensReview) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9990,10 +9803,8 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10033,7 +9844,6 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10157,10 +9967,8 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10200,7 +10008,6 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10324,10 +10131,8 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10367,7 +10172,6 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10509,10 +10313,8 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10552,7 +10354,6 @@ func (m *awsRestjson1_deserializeOpUpdateReviewTemplat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10685,10 +10486,8 @@ func (m *awsRestjson1_deserializeOpUpdateShareInvitati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10728,7 +10527,6 @@ func (m *awsRestjson1_deserializeOpUpdateShareInvitati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10852,10 +10650,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkload) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10895,7 +10691,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkload) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11019,10 +10814,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkloadShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11062,7 +10855,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkloadShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11195,10 +10987,8 @@ func (m *awsRestjson1_deserializeOpUpgradeLensReview) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11216,7 +11006,6 @@ func (m *awsRestjson1_deserializeOpUpgradeLensReview) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11307,10 +11096,8 @@ func (m *awsRestjson1_deserializeOpUpgradeProfileVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11328,7 +11115,6 @@ func (m *awsRestjson1_deserializeOpUpgradeProfileVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11419,10 +11205,8 @@ func (m *awsRestjson1_deserializeOpUpgradeReviewTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11440,7 +11224,6 @@ func (m *awsRestjson1_deserializeOpUpgradeReviewTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
