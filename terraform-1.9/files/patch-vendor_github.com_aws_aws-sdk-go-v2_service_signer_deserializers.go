--- vendor/github.com/aws/aws-sdk-go-v2/service/signer/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/signer/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAddProfilePermissio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAddProfilePermissio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -220,10 +216,8 @@ func (m *awsRestjson1_deserializeOpCancelSigningProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -241,7 +235,6 @@ func (m *awsRestjson1_deserializeOpCancelSigningProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -323,10 +316,8 @@ func (m *awsRestjson1_deserializeOpDescribeSigningJob)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -366,7 +357,6 @@ func (m *awsRestjson1_deserializeOpDescribeSigningJob)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -647,10 +637,8 @@ func (m *awsRestjson1_deserializeOpGetRevocationStatus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -690,7 +678,6 @@ func (m *awsRestjson1_deserializeOpGetRevocationStatus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -808,10 +795,8 @@ func (m *awsRestjson1_deserializeOpGetSigningPlatform)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -851,7 +836,6 @@ func (m *awsRestjson1_deserializeOpGetSigningPlatform)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1041,10 +1025,8 @@ func (m *awsRestjson1_deserializeOpGetSigningProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1084,7 +1066,6 @@ func (m *awsRestjson1_deserializeOpGetSigningProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1299,10 +1280,8 @@ func (m *awsRestjson1_deserializeOpListProfilePermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1342,7 +1321,6 @@ func (m *awsRestjson1_deserializeOpListProfilePermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1494,10 +1472,8 @@ func (m *awsRestjson1_deserializeOpListSigningJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1537,7 +1513,6 @@ func (m *awsRestjson1_deserializeOpListSigningJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1664,10 +1639,8 @@ func (m *awsRestjson1_deserializeOpListSigningPlatform
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1707,7 +1680,6 @@ func (m *awsRestjson1_deserializeOpListSigningPlatform
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1834,10 +1806,8 @@ func (m *awsRestjson1_deserializeOpListSigningProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1877,7 +1847,6 @@ func (m *awsRestjson1_deserializeOpListSigningProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2001,10 +1970,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2044,7 +2011,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2162,10 +2128,8 @@ func (m *awsRestjson1_deserializeOpPutSigningProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2205,7 +2169,6 @@ func (m *awsRestjson1_deserializeOpPutSigningProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2348,10 +2311,8 @@ func (m *awsRestjson1_deserializeOpRemoveProfilePermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2391,7 +2352,6 @@ func (m *awsRestjson1_deserializeOpRemoveProfilePermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2519,10 +2479,8 @@ func (m *awsRestjson1_deserializeOpRevokeSignature) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2540,7 +2498,6 @@ func (m *awsRestjson1_deserializeOpRevokeSignature) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2625,10 +2582,8 @@ func (m *awsRestjson1_deserializeOpRevokeSigningProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2646,7 +2601,6 @@ func (m *awsRestjson1_deserializeOpRevokeSigningProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2731,10 +2685,8 @@ func (m *awsRestjson1_deserializeOpSignPayload) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2774,7 +2726,6 @@ func (m *awsRestjson1_deserializeOpSignPayload) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2926,10 +2877,8 @@ func (m *awsRestjson1_deserializeOpStartSigningJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2969,7 +2918,6 @@ func (m *awsRestjson1_deserializeOpStartSigningJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3106,10 +3054,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3121,7 +3067,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3203,10 +3148,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3218,7 +3161,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
