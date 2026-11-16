--- vendor/github.com/aws/aws-sdk-go-v2/service/customerprofiles/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/customerprofiles/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAddProfileKey) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAddProfileKey) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -218,10 +214,8 @@ func (m *awsRestjson1_deserializeOpBatchGetCalculatedA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -261,7 +255,6 @@ func (m *awsRestjson1_deserializeOpBatchGetCalculatedA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -392,10 +385,8 @@ func (m *awsRestjson1_deserializeOpBatchGetProfile) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -435,7 +426,6 @@ func (m *awsRestjson1_deserializeOpBatchGetProfile) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -561,10 +551,8 @@ func (m *awsRestjson1_deserializeOpCreateCalculatedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -604,7 +592,6 @@ func (m *awsRestjson1_deserializeOpCreateCalculatedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -831,10 +818,8 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -874,7 +859,6 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1077,10 +1061,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainLayout)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1120,7 +1102,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainLayout)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1336,10 +1317,8 @@ func (m *awsRestjson1_deserializeOpCreateEventStream) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1379,7 +1358,6 @@ func (m *awsRestjson1_deserializeOpCreateEventStream) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1509,10 +1487,8 @@ func (m *awsRestjson1_deserializeOpCreateEventTrigger)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1552,7 +1528,6 @@ func (m *awsRestjson1_deserializeOpCreateEventTrigger)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1751,10 +1726,8 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationWo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1794,7 +1767,6 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationWo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1928,10 +1900,8 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1971,7 +1941,6 @@ func (m *awsRestjson1_deserializeOpCreateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2096,10 +2065,8 @@ func (m *awsRestjson1_deserializeOpCreateSegmentDefini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2139,7 +2106,6 @@ func (m *awsRestjson1_deserializeOpCreateSegmentDefini
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2312,10 +2278,8 @@ func (m *awsRestjson1_deserializeOpCreateSegmentEstima
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2360,7 +2324,6 @@ func (m *awsRestjson1_deserializeOpCreateSegmentEstima
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2503,10 +2466,8 @@ func (m *awsRestjson1_deserializeOpCreateSegmentSnapsh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2546,7 +2507,6 @@ func (m *awsRestjson1_deserializeOpCreateSegmentSnapsh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2671,10 +2631,8 @@ func (m *awsRestjson1_deserializeOpCreateUploadJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2714,7 +2672,6 @@ func (m *awsRestjson1_deserializeOpCreateUploadJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2839,10 +2796,8 @@ func (m *awsRestjson1_deserializeOpDeleteCalculatedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2854,7 +2809,6 @@ func (m *awsRestjson1_deserializeOpDeleteCalculatedAtt
 	output := &DeleteCalculatedAttributeDefinitionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2939,10 +2893,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2982,7 +2934,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3107,10 +3058,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainLayout)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3150,7 +3099,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainLayout)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3275,10 +3223,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventStream) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3290,7 +3236,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventStream) 
 	output := &DeleteEventStreamOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3375,10 +3320,8 @@ func (m *awsRestjson1_deserializeOpDeleteEventTrigger)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3418,7 +3361,6 @@ func (m *awsRestjson1_deserializeOpDeleteEventTrigger)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3543,10 +3485,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3586,7 +3526,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3711,10 +3650,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3754,7 +3691,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3879,10 +3815,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfileKey) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3922,7 +3856,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfileKey) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4047,10 +3980,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfileObject
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4090,7 +4021,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfileObject
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4215,10 +4145,8 @@ func (m *awsRestjson1_deserializeOpDeleteProfileObject
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4258,7 +4186,6 @@ func (m *awsRestjson1_deserializeOpDeleteProfileObject
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4383,10 +4310,8 @@ func (m *awsRestjson1_deserializeOpDeleteSegmentDefini
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4426,7 +4351,6 @@ func (m *awsRestjson1_deserializeOpDeleteSegmentDefini
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4551,10 +4475,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4566,7 +4488,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 	output := &DeleteWorkflowOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4651,10 +4572,8 @@ func (m *awsRestjson1_deserializeOpDetectProfileObject
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4694,7 +4613,6 @@ func (m *awsRestjson1_deserializeOpDetectProfileObject
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4815,10 +4733,8 @@ func (m *awsRestjson1_deserializeOpGetAutoMergingPrevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4858,7 +4774,6 @@ func (m *awsRestjson1_deserializeOpGetAutoMergingPrevi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5022,10 +4937,8 @@ func (m *awsRestjson1_deserializeOpGetCalculatedAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5065,7 +4978,6 @@ func (m *awsRestjson1_deserializeOpGetCalculatedAttrib
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5292,10 +5204,8 @@ func (m *awsRestjson1_deserializeOpGetCalculatedAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5335,7 +5245,6 @@ func (m *awsRestjson1_deserializeOpGetCalculatedAttrib
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5503,10 +5412,8 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5546,7 +5453,6 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5754,10 +5660,8 @@ func (m *awsRestjson1_deserializeOpGetDomainLayout) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5797,7 +5701,6 @@ func (m *awsRestjson1_deserializeOpGetDomainLayout) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6013,10 +5916,8 @@ func (m *awsRestjson1_deserializeOpGetEventStream) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6056,7 +5957,6 @@ func (m *awsRestjson1_deserializeOpGetEventStream) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6241,10 +6141,8 @@ func (m *awsRestjson1_deserializeOpGetEventTrigger) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6284,7 +6182,6 @@ func (m *awsRestjson1_deserializeOpGetEventTrigger) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6483,10 +6380,8 @@ func (m *awsRestjson1_deserializeOpGetIdentityResoluti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6526,7 +6421,6 @@ func (m *awsRestjson1_deserializeOpGetIdentityResoluti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6757,10 +6651,8 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6800,7 +6692,6 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7017,10 +6908,8 @@ func (m *awsRestjson1_deserializeOpGetMatches) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7060,7 +6949,6 @@ func (m *awsRestjson1_deserializeOpGetMatches) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7219,10 +7107,8 @@ func (m *awsRestjson1_deserializeOpGetProfileObjectTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7262,7 +7148,6 @@ func (m *awsRestjson1_deserializeOpGetProfileObjectTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7518,10 +7403,8 @@ func (m *awsRestjson1_deserializeOpGetProfileObjectTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7561,7 +7444,6 @@ func (m *awsRestjson1_deserializeOpGetProfileObjectTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7732,10 +7614,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7775,7 +7655,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7953,10 +7832,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentEstimate)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8001,7 +7878,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentEstimate)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8171,10 +8047,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8214,7 +8088,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8349,10 +8222,8 @@ func (m *awsRestjson1_deserializeOpGetSegmentSnapshot)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8392,7 +8263,6 @@ func (m *awsRestjson1_deserializeOpGetSegmentSnapshot)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8571,10 +8441,8 @@ func (m *awsRestjson1_deserializeOpGetSimilarProfiles)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8614,7 +8482,6 @@ func (m *awsRestjson1_deserializeOpGetSimilarProfiles)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8809,10 +8676,8 @@ func (m *awsRestjson1_deserializeOpGetUploadJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8852,7 +8717,6 @@ func (m *awsRestjson1_deserializeOpGetUploadJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9068,10 +8932,8 @@ func (m *awsRestjson1_deserializeOpGetUploadJobPath) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9111,7 +8973,6 @@ func (m *awsRestjson1_deserializeOpGetUploadJobPath) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9261,10 +9122,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9304,7 +9163,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9498,10 +9356,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowSteps) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9541,7 +9397,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowSteps) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9689,10 +9544,8 @@ func (m *awsRestjson1_deserializeOpListAccountIntegrat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9732,7 +9585,6 @@ func (m *awsRestjson1_deserializeOpListAccountIntegrat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9862,10 +9714,8 @@ func (m *awsRestjson1_deserializeOpListCalculatedAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9905,7 +9755,6 @@ func (m *awsRestjson1_deserializeOpListCalculatedAttri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10035,10 +9884,8 @@ func (m *awsRestjson1_deserializeOpListCalculatedAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10078,7 +9925,6 @@ func (m *awsRestjson1_deserializeOpListCalculatedAttri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10208,10 +10054,8 @@ func (m *awsRestjson1_deserializeOpListDomainLayouts) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10251,7 +10095,6 @@ func (m *awsRestjson1_deserializeOpListDomainLayouts) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10381,10 +10224,8 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10424,7 +10265,6 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10554,10 +10394,8 @@ func (m *awsRestjson1_deserializeOpListEventStreams) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10597,7 +10435,6 @@ func (m *awsRestjson1_deserializeOpListEventStreams) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10727,10 +10564,8 @@ func (m *awsRestjson1_deserializeOpListEventTriggers) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10770,7 +10605,6 @@ func (m *awsRestjson1_deserializeOpListEventTriggers) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10900,10 +10734,8 @@ func (m *awsRestjson1_deserializeOpListIdentityResolut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10943,7 +10775,6 @@ func (m *awsRestjson1_deserializeOpListIdentityResolut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11073,10 +10904,8 @@ func (m *awsRestjson1_deserializeOpListIntegrations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11116,7 +10945,6 @@ func (m *awsRestjson1_deserializeOpListIntegrations) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11246,10 +11074,8 @@ func (m *awsRestjson1_deserializeOpListObjectTypeAttri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11289,7 +11115,6 @@ func (m *awsRestjson1_deserializeOpListObjectTypeAttri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11419,10 +11244,8 @@ func (m *awsRestjson1_deserializeOpListProfileAttribut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11467,7 +11290,6 @@ func (m *awsRestjson1_deserializeOpListProfileAttribut
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11615,10 +11437,8 @@ func (m *awsRestjson1_deserializeOpListProfileObjects)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11658,7 +11478,6 @@ func (m *awsRestjson1_deserializeOpListProfileObjects)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11788,10 +11607,8 @@ func (m *awsRestjson1_deserializeOpListProfileObjectTy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11831,7 +11648,6 @@ func (m *awsRestjson1_deserializeOpListProfileObjectTy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11961,10 +11777,8 @@ func (m *awsRestjson1_deserializeOpListProfileObjectTy
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12004,7 +11818,6 @@ func (m *awsRestjson1_deserializeOpListProfileObjectTy
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12134,10 +11947,8 @@ func (m *awsRestjson1_deserializeOpListRuleBasedMatche
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12177,7 +11988,6 @@ func (m *awsRestjson1_deserializeOpListRuleBasedMatche
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12307,10 +12117,8 @@ func (m *awsRestjson1_deserializeOpListSegmentDefiniti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12350,7 +12158,6 @@ func (m *awsRestjson1_deserializeOpListSegmentDefiniti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12480,10 +12287,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12523,7 +12328,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12638,10 +12442,8 @@ func (m *awsRestjson1_deserializeOpListUploadJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12681,7 +12483,6 @@ func (m *awsRestjson1_deserializeOpListUploadJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12811,10 +12612,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12854,7 +12653,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12984,10 +12782,8 @@ func (m *awsRestjson1_deserializeOpMergeProfiles) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13027,7 +12823,6 @@ func (m *awsRestjson1_deserializeOpMergeProfiles) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13149,10 +12944,8 @@ func (m *awsRestjson1_deserializeOpPutIntegration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13192,7 +12985,6 @@ func (m *awsRestjson1_deserializeOpPutIntegration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13409,10 +13201,8 @@ func (m *awsRestjson1_deserializeOpPutProfileObject) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13452,7 +13242,6 @@ func (m *awsRestjson1_deserializeOpPutProfileObject) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13577,10 +13366,8 @@ func (m *awsRestjson1_deserializeOpPutProfileObjectTyp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13620,7 +13407,6 @@ func (m *awsRestjson1_deserializeOpPutProfileObjectTyp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13876,10 +13662,8 @@ func (m *awsRestjson1_deserializeOpSearchProfiles) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13919,7 +13703,6 @@ func (m *awsRestjson1_deserializeOpSearchProfiles) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14049,10 +13832,8 @@ func (m *awsRestjson1_deserializeOpStartUploadJob) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14064,7 +13845,6 @@ func (m *awsRestjson1_deserializeOpStartUploadJob) Han
 	output := &StartUploadJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14149,10 +13929,8 @@ func (m *awsRestjson1_deserializeOpStopUploadJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14164,7 +13942,6 @@ func (m *awsRestjson1_deserializeOpStopUploadJob) Hand
 	output := &StopUploadJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14249,10 +14026,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14264,7 +14039,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14343,10 +14117,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14358,7 +14130,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14437,10 +14208,8 @@ func (m *awsRestjson1_deserializeOpUpdateCalculatedAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14480,7 +14249,6 @@ func (m *awsRestjson1_deserializeOpUpdateCalculatedAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14702,10 +14470,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14745,7 +14511,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14948,10 +14713,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainLayout)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14991,7 +14754,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainLayout)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15207,10 +14969,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventTrigger)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15250,7 +15010,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventTrigger)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15449,10 +15208,8 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15492,7 +15249,6 @@ func (m *awsRestjson1_deserializeOpUpdateProfile) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
