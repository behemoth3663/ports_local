--- vendor/github.com/aws/aws-sdk-go-v2/service/clouddirectory/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/clouddirectory/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAddFacetToObject) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -61,7 +58,6 @@ func (m *awsRestjson1_deserializeOpAddFacetToObject) H
 	output := &AddFacetToObjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -158,10 +154,8 @@ func (m *awsRestjson1_deserializeOpApplySchema) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -201,7 +195,6 @@ func (m *awsRestjson1_deserializeOpApplySchema) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -347,10 +340,8 @@ func (m *awsRestjson1_deserializeOpAttachObject) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -390,7 +381,6 @@ func (m *awsRestjson1_deserializeOpAttachObject) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -533,10 +523,8 @@ func (m *awsRestjson1_deserializeOpAttachPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -548,7 +536,6 @@ func (m *awsRestjson1_deserializeOpAttachPolicy) Handl
 	output := &AttachPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -645,10 +632,8 @@ func (m *awsRestjson1_deserializeOpAttachToIndex) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -688,7 +673,6 @@ func (m *awsRestjson1_deserializeOpAttachToIndex) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -834,10 +818,8 @@ func (m *awsRestjson1_deserializeOpAttachTypedLink) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -877,7 +859,6 @@ func (m *awsRestjson1_deserializeOpAttachTypedLink) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1013,10 +994,8 @@ func (m *awsRestjson1_deserializeOpBatchRead) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1056,7 +1035,6 @@ func (m *awsRestjson1_deserializeOpBatchRead) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1183,10 +1161,8 @@ func (m *awsRestjson1_deserializeOpBatchWrite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1226,7 +1202,6 @@ func (m *awsRestjson1_deserializeOpBatchWrite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1356,10 +1331,8 @@ func (m *awsRestjson1_deserializeOpCreateDirectory) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1399,7 +1372,6 @@ func (m *awsRestjson1_deserializeOpCreateDirectory) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1560,10 +1532,8 @@ func (m *awsRestjson1_deserializeOpCreateFacet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1575,7 +1545,6 @@ func (m *awsRestjson1_deserializeOpCreateFacet) Handle
 	output := &CreateFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1675,10 +1644,8 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1718,7 +1685,6 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1861,10 +1827,8 @@ func (m *awsRestjson1_deserializeOpCreateObject) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1904,7 +1868,6 @@ func (m *awsRestjson1_deserializeOpCreateObject) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2047,10 +2010,8 @@ func (m *awsRestjson1_deserializeOpCreateSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2090,7 +2051,6 @@ func (m *awsRestjson1_deserializeOpCreateSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2221,10 +2181,8 @@ func (m *awsRestjson1_deserializeOpCreateTypedLinkFace
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2236,7 +2194,6 @@ func (m *awsRestjson1_deserializeOpCreateTypedLinkFace
 	output := &CreateTypedLinkFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2336,10 +2293,8 @@ func (m *awsRestjson1_deserializeOpDeleteDirectory) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2379,7 +2334,6 @@ func (m *awsRestjson1_deserializeOpDeleteDirectory) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2516,10 +2470,8 @@ func (m *awsRestjson1_deserializeOpDeleteFacet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2531,7 +2483,6 @@ func (m *awsRestjson1_deserializeOpDeleteFacet) Handle
 	output := &DeleteFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2628,10 +2579,8 @@ func (m *awsRestjson1_deserializeOpDeleteObject) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2643,7 +2592,6 @@ func (m *awsRestjson1_deserializeOpDeleteObject) Handl
 	output := &DeleteObjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2740,10 +2688,8 @@ func (m *awsRestjson1_deserializeOpDeleteSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2783,7 +2729,6 @@ func (m *awsRestjson1_deserializeOpDeleteSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2917,10 +2862,8 @@ func (m *awsRestjson1_deserializeOpDeleteTypedLinkFace
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2932,7 +2875,6 @@ func (m *awsRestjson1_deserializeOpDeleteTypedLinkFace
 	output := &DeleteTypedLinkFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3026,10 +2968,8 @@ func (m *awsRestjson1_deserializeOpDetachFromIndex) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3069,7 +3009,6 @@ func (m *awsRestjson1_deserializeOpDetachFromIndex) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3209,10 +3148,8 @@ func (m *awsRestjson1_deserializeOpDetachObject) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3252,7 +3189,6 @@ func (m *awsRestjson1_deserializeOpDetachObject) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3389,10 +3325,8 @@ func (m *awsRestjson1_deserializeOpDetachPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3404,7 +3338,6 @@ func (m *awsRestjson1_deserializeOpDetachPolicy) Handl
 	output := &DetachPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3501,10 +3434,8 @@ func (m *awsRestjson1_deserializeOpDetachTypedLink) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3522,7 +3453,6 @@ func (m *awsRestjson1_deserializeOpDetachTypedLink) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3619,10 +3549,8 @@ func (m *awsRestjson1_deserializeOpDisableDirectory) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3662,7 +3590,6 @@ func (m *awsRestjson1_deserializeOpDisableDirectory) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3796,10 +3723,8 @@ func (m *awsRestjson1_deserializeOpEnableDirectory) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3839,7 +3764,6 @@ func (m *awsRestjson1_deserializeOpEnableDirectory) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3973,10 +3897,8 @@ func (m *awsRestjson1_deserializeOpGetAppliedSchemaVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4016,7 +3938,6 @@ func (m *awsRestjson1_deserializeOpGetAppliedSchemaVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4147,10 +4068,8 @@ func (m *awsRestjson1_deserializeOpGetDirectory) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4190,7 +4109,6 @@ func (m *awsRestjson1_deserializeOpGetDirectory) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4314,10 +4232,8 @@ func (m *awsRestjson1_deserializeOpGetFacet) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4357,7 +4273,6 @@ func (m *awsRestjson1_deserializeOpGetFacet) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4487,10 +4402,8 @@ func (m *awsRestjson1_deserializeOpGetLinkAttributes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4530,7 +4443,6 @@ func (m *awsRestjson1_deserializeOpGetLinkAttributes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4663,10 +4575,8 @@ func (m *awsRestjson1_deserializeOpGetObjectAttributes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4706,7 +4616,6 @@ func (m *awsRestjson1_deserializeOpGetObjectAttributes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4839,10 +4748,8 @@ func (m *awsRestjson1_deserializeOpGetObjectInformatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4882,7 +4789,6 @@ func (m *awsRestjson1_deserializeOpGetObjectInformatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5021,10 +4927,8 @@ func (m *awsRestjson1_deserializeOpGetSchemaAsJson) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5064,7 +4968,6 @@ func (m *awsRestjson1_deserializeOpGetSchemaAsJson) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5204,10 +5107,8 @@ func (m *awsRestjson1_deserializeOpGetTypedLinkFacetIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5247,7 +5148,6 @@ func (m *awsRestjson1_deserializeOpGetTypedLinkFacetIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5380,10 +5280,8 @@ func (m *awsRestjson1_deserializeOpListAppliedSchemaAr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5423,7 +5321,6 @@ func (m *awsRestjson1_deserializeOpListAppliedSchemaAr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5562,10 +5459,8 @@ func (m *awsRestjson1_deserializeOpListAttachedIndices
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5605,7 +5500,6 @@ func (m *awsRestjson1_deserializeOpListAttachedIndices
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5744,10 +5638,8 @@ func (m *awsRestjson1_deserializeOpListDevelopmentSche
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5787,7 +5679,6 @@ func (m *awsRestjson1_deserializeOpListDevelopmentSche
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5926,10 +5817,8 @@ func (m *awsRestjson1_deserializeOpListDirectories) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5969,7 +5858,6 @@ func (m *awsRestjson1_deserializeOpListDirectories) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6105,10 +5993,8 @@ func (m *awsRestjson1_deserializeOpListFacetAttributes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6148,7 +6034,6 @@ func (m *awsRestjson1_deserializeOpListFacetAttributes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6290,10 +6175,8 @@ func (m *awsRestjson1_deserializeOpListFacetNames) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6333,7 +6216,6 @@ func (m *awsRestjson1_deserializeOpListFacetNames) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6472,10 +6354,8 @@ func (m *awsRestjson1_deserializeOpListIncomingTypedLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6515,7 +6395,6 @@ func (m *awsRestjson1_deserializeOpListIncomingTypedLi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6660,10 +6539,8 @@ func (m *awsRestjson1_deserializeOpListIndex) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6703,7 +6580,6 @@ func (m *awsRestjson1_deserializeOpListIndex) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6851,10 +6727,8 @@ func (m *awsRestjson1_deserializeOpListManagedSchemaAr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6894,7 +6768,6 @@ func (m *awsRestjson1_deserializeOpListManagedSchemaAr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7027,10 +6900,8 @@ func (m *awsRestjson1_deserializeOpListObjectAttribute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7070,7 +6941,6 @@ func (m *awsRestjson1_deserializeOpListObjectAttribute
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7215,10 +7085,8 @@ func (m *awsRestjson1_deserializeOpListObjectChildren)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7258,7 +7126,6 @@ func (m *awsRestjson1_deserializeOpListObjectChildren)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7403,10 +7270,8 @@ func (m *awsRestjson1_deserializeOpListObjectParentPat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7446,7 +7311,6 @@ func (m *awsRestjson1_deserializeOpListObjectParentPat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7588,10 +7452,8 @@ func (m *awsRestjson1_deserializeOpListObjectParents) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7631,7 +7493,6 @@ func (m *awsRestjson1_deserializeOpListObjectParents) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7781,10 +7642,8 @@ func (m *awsRestjson1_deserializeOpListObjectPolicies)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7824,7 +7683,6 @@ func (m *awsRestjson1_deserializeOpListObjectPolicies)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7966,10 +7824,8 @@ func (m *awsRestjson1_deserializeOpListOutgoingTypedLi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8009,7 +7865,6 @@ func (m *awsRestjson1_deserializeOpListOutgoingTypedLi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8154,10 +8009,8 @@ func (m *awsRestjson1_deserializeOpListPolicyAttachmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8197,7 +8050,6 @@ func (m *awsRestjson1_deserializeOpListPolicyAttachmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8342,10 +8194,8 @@ func (m *awsRestjson1_deserializeOpListPublishedSchema
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8385,7 +8235,6 @@ func (m *awsRestjson1_deserializeOpListPublishedSchema
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8524,10 +8373,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8567,7 +8414,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8706,10 +8552,8 @@ func (m *awsRestjson1_deserializeOpListTypedLinkFacetA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8749,7 +8593,6 @@ func (m *awsRestjson1_deserializeOpListTypedLinkFacetA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8891,10 +8734,8 @@ func (m *awsRestjson1_deserializeOpListTypedLinkFacetN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8934,7 +8775,6 @@ func (m *awsRestjson1_deserializeOpListTypedLinkFacetN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9073,10 +8913,8 @@ func (m *awsRestjson1_deserializeOpLookupPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9116,7 +8954,6 @@ func (m *awsRestjson1_deserializeOpLookupPolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9258,10 +9095,8 @@ func (m *awsRestjson1_deserializeOpPublishSchema) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9301,7 +9136,6 @@ func (m *awsRestjson1_deserializeOpPublishSchema) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9435,10 +9269,8 @@ func (m *awsRestjson1_deserializeOpPutSchemaFromJson) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9478,7 +9310,6 @@ func (m *awsRestjson1_deserializeOpPutSchemaFromJson) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9612,10 +9443,8 @@ func (m *awsRestjson1_deserializeOpRemoveFacetFromObje
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9627,7 +9456,6 @@ func (m *awsRestjson1_deserializeOpRemoveFacetFromObje
 	output := &RemoveFacetFromObjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9724,10 +9552,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9739,7 +9565,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9833,10 +9658,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9848,7 +9671,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9942,10 +9764,8 @@ func (m *awsRestjson1_deserializeOpUpdateFacet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9957,7 +9777,6 @@ func (m *awsRestjson1_deserializeOpUpdateFacet) Handle
 	output := &UpdateFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10060,10 +9879,8 @@ func (m *awsRestjson1_deserializeOpUpdateLinkAttribute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10075,7 +9892,6 @@ func (m *awsRestjson1_deserializeOpUpdateLinkAttribute
 	output := &UpdateLinkAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10172,10 +9988,8 @@ func (m *awsRestjson1_deserializeOpUpdateObjectAttribu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10215,7 +10029,6 @@ func (m *awsRestjson1_deserializeOpUpdateObjectAttribu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10355,10 +10168,8 @@ func (m *awsRestjson1_deserializeOpUpdateSchema) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10398,7 +10209,6 @@ func (m *awsRestjson1_deserializeOpUpdateSchema) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10529,10 +10339,8 @@ func (m *awsRestjson1_deserializeOpUpdateTypedLinkFace
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10544,7 +10352,6 @@ func (m *awsRestjson1_deserializeOpUpdateTypedLinkFace
 	output := &UpdateTypedLinkFacetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10647,10 +10454,8 @@ func (m *awsRestjson1_deserializeOpUpgradeAppliedSchem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10690,7 +10495,6 @@ func (m *awsRestjson1_deserializeOpUpgradeAppliedSchem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10836,10 +10640,8 @@ func (m *awsRestjson1_deserializeOpUpgradePublishedSch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10879,7 +10681,6 @@ func (m *awsRestjson1_deserializeOpUpgradePublishedSch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
