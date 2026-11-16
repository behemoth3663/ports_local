--- vendor/github.com/aws/aws-sdk-go-v2/service/amplifybackend/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/amplifybackend/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCloneBackend) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCloneBackend) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -255,10 +251,8 @@ func (m *awsRestjson1_deserializeOpCreateBackend) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -298,7 +292,6 @@ func (m *awsRestjson1_deserializeOpCreateBackend) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -465,10 +458,8 @@ func (m *awsRestjson1_deserializeOpCreateBackendAPI) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -508,7 +499,6 @@ func (m *awsRestjson1_deserializeOpCreateBackendAPI) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -675,10 +665,8 @@ func (m *awsRestjson1_deserializeOpCreateBackendAuth) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -718,7 +706,6 @@ func (m *awsRestjson1_deserializeOpCreateBackendAuth) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -885,10 +872,8 @@ func (m *awsRestjson1_deserializeOpCreateBackendConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -928,7 +913,6 @@ func (m *awsRestjson1_deserializeOpCreateBackendConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1077,10 +1061,8 @@ func (m *awsRestjson1_deserializeOpCreateBackendStorag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1120,7 +1102,6 @@ func (m *awsRestjson1_deserializeOpCreateBackendStorag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1269,10 +1250,8 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1312,7 +1291,6 @@ func (m *awsRestjson1_deserializeOpCreateToken) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1461,10 +1439,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackend) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1504,7 +1480,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackend) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1671,10 +1646,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackendAPI) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1714,7 +1687,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackendAPI) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1881,10 +1853,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackendAuth) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1924,7 +1894,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackendAuth) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2091,10 +2060,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackendStorag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2134,7 +2101,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackendStorag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2283,10 +2249,8 @@ func (m *awsRestjson1_deserializeOpDeleteToken) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2326,7 +2290,6 @@ func (m *awsRestjson1_deserializeOpDeleteToken) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2448,10 +2411,8 @@ func (m *awsRestjson1_deserializeOpGenerateBackendAPIM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2491,7 +2452,6 @@ func (m *awsRestjson1_deserializeOpGenerateBackendAPIM
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2658,10 +2618,8 @@ func (m *awsRestjson1_deserializeOpGetBackend) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2701,7 +2659,6 @@ func (m *awsRestjson1_deserializeOpGetBackend) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2873,10 +2830,8 @@ func (m *awsRestjson1_deserializeOpGetBackendAPI) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2916,7 +2871,6 @@ func (m *awsRestjson1_deserializeOpGetBackendAPI) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3070,10 +3024,8 @@ func (m *awsRestjson1_deserializeOpGetBackendAPIModels
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3113,7 +3065,6 @@ func (m *awsRestjson1_deserializeOpGetBackendAPIModels
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3253,10 +3204,8 @@ func (m *awsRestjson1_deserializeOpGetBackendAuth) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3296,7 +3245,6 @@ func (m *awsRestjson1_deserializeOpGetBackendAuth) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3450,10 +3398,8 @@ func (m *awsRestjson1_deserializeOpGetBackendJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3493,7 +3439,6 @@ func (m *awsRestjson1_deserializeOpGetBackendJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3678,10 +3623,8 @@ func (m *awsRestjson1_deserializeOpGetBackendStorage) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3721,7 +3664,6 @@ func (m *awsRestjson1_deserializeOpGetBackendStorage) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3866,10 +3808,8 @@ func (m *awsRestjson1_deserializeOpGetToken) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3909,7 +3849,6 @@ func (m *awsRestjson1_deserializeOpGetToken) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4058,10 +3997,8 @@ func (m *awsRestjson1_deserializeOpImportBackendAuth) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4101,7 +4038,6 @@ func (m *awsRestjson1_deserializeOpImportBackendAuth) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4268,10 +4204,8 @@ func (m *awsRestjson1_deserializeOpImportBackendStorag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4311,7 +4245,6 @@ func (m *awsRestjson1_deserializeOpImportBackendStorag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4460,10 +4393,8 @@ func (m *awsRestjson1_deserializeOpListBackendJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4503,7 +4434,6 @@ func (m *awsRestjson1_deserializeOpListBackendJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4630,10 +4560,8 @@ func (m *awsRestjson1_deserializeOpListS3Buckets) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4673,7 +4601,6 @@ func (m *awsRestjson1_deserializeOpListS3Buckets) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4800,10 +4727,8 @@ func (m *awsRestjson1_deserializeOpRemoveAllBackends) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4843,7 +4768,6 @@ func (m *awsRestjson1_deserializeOpRemoveAllBackends) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5001,10 +4925,8 @@ func (m *awsRestjson1_deserializeOpRemoveBackendConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5044,7 +4966,6 @@ func (m *awsRestjson1_deserializeOpRemoveBackendConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5166,10 +5087,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackendAPI) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5209,7 +5128,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackendAPI) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5376,10 +5294,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackendAuth) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5419,7 +5335,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackendAuth) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5586,10 +5501,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackendConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5629,7 +5542,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackendConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5774,10 +5686,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackendJob) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5817,7 +5727,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackendJob) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6002,10 +5911,8 @@ func (m *awsRestjson1_deserializeOpUpdateBackendStorag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6045,7 +5952,6 @@ func (m *awsRestjson1_deserializeOpUpdateBackendStorag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
