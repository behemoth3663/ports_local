--- vendor/github.com/aws/aws-sdk-go-v2/service/omics/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/omics/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAbortMultipartReadS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -62,7 +59,6 @@ func (m *awsRestjson1_deserializeOpAbortMultipartReadS
 	output := &AbortMultipartReadSetUploadOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -156,10 +152,8 @@ func (m *awsRestjson1_deserializeOpAcceptShare) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -199,7 +193,6 @@ func (m *awsRestjson1_deserializeOpAcceptShare) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -330,10 +323,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteReadSet)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -373,7 +364,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteReadSet)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -497,10 +487,8 @@ func (m *awsRestjson1_deserializeOpCancelAnnotationImp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -512,7 +500,6 @@ func (m *awsRestjson1_deserializeOpCancelAnnotationImp
 	output := &CancelAnnotationImportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -597,10 +584,8 @@ func (m *awsRestjson1_deserializeOpCancelRun) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -618,7 +603,6 @@ func (m *awsRestjson1_deserializeOpCancelRun) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -712,10 +696,8 @@ func (m *awsRestjson1_deserializeOpCancelVariantImport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -727,7 +709,6 @@ func (m *awsRestjson1_deserializeOpCancelVariantImport
 	output := &CancelVariantImportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -812,10 +793,8 @@ func (m *awsRestjson1_deserializeOpCompleteMultipartRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -855,7 +834,6 @@ func (m *awsRestjson1_deserializeOpCompleteMultipartRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -989,10 +967,8 @@ func (m *awsRestjson1_deserializeOpCreateAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1032,7 +1008,6 @@ func (m *awsRestjson1_deserializeOpCreateAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1222,10 +1197,8 @@ func (m *awsRestjson1_deserializeOpCreateAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1265,7 +1238,6 @@ func (m *awsRestjson1_deserializeOpCreateAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1450,10 +1422,8 @@ func (m *awsRestjson1_deserializeOpCreateMultipartRead
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1493,7 +1463,6 @@ func (m *awsRestjson1_deserializeOpCreateMultipartRead
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1717,10 +1686,8 @@ func (m *awsRestjson1_deserializeOpCreateReferenceStor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1760,7 +1727,6 @@ func (m *awsRestjson1_deserializeOpCreateReferenceStor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1933,10 +1899,8 @@ func (m *awsRestjson1_deserializeOpCreateRunCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1976,7 +1940,6 @@ func (m *awsRestjson1_deserializeOpCreateRunCache) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2133,10 +2096,8 @@ func (m *awsRestjson1_deserializeOpCreateRunGroup) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2176,7 +2137,6 @@ func (m *awsRestjson1_deserializeOpCreateRunGroup) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2324,10 +2284,8 @@ func (m *awsRestjson1_deserializeOpCreateSequenceStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2367,7 +2325,6 @@ func (m *awsRestjson1_deserializeOpCreateSequenceStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2586,10 +2543,8 @@ func (m *awsRestjson1_deserializeOpCreateShare) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2629,7 +2584,6 @@ func (m *awsRestjson1_deserializeOpCreateShare) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2778,10 +2732,8 @@ func (m *awsRestjson1_deserializeOpCreateVariantStore)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2821,7 +2773,6 @@ func (m *awsRestjson1_deserializeOpCreateVariantStore)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2988,10 +2939,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3031,7 +2980,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3197,10 +3145,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3240,7 +3186,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkflowVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3415,10 +3360,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3458,7 +3401,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3586,10 +3528,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3629,7 +3569,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3753,10 +3692,8 @@ func (m *awsRestjson1_deserializeOpDeleteReference) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3768,7 +3705,6 @@ func (m *awsRestjson1_deserializeOpDeleteReference) Ha
 	output := &DeleteReferenceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3859,10 +3795,8 @@ func (m *awsRestjson1_deserializeOpDeleteReferenceStor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3874,7 +3808,6 @@ func (m *awsRestjson1_deserializeOpDeleteReferenceStor
 	output := &DeleteReferenceStoreOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3965,10 +3898,8 @@ func (m *awsRestjson1_deserializeOpDeleteRun) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3986,7 +3917,6 @@ func (m *awsRestjson1_deserializeOpDeleteRun) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4080,10 +4010,8 @@ func (m *awsRestjson1_deserializeOpDeleteRunCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4101,7 +4029,6 @@ func (m *awsRestjson1_deserializeOpDeleteRunCache) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4195,10 +4122,8 @@ func (m *awsRestjson1_deserializeOpDeleteRunGroup) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4216,7 +4141,6 @@ func (m *awsRestjson1_deserializeOpDeleteRunGroup) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4310,10 +4234,8 @@ func (m *awsRestjson1_deserializeOpDeleteS3AccessPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4325,7 +4247,6 @@ func (m *awsRestjson1_deserializeOpDeleteS3AccessPolic
 	output := &DeleteS3AccessPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4416,10 +4337,8 @@ func (m *awsRestjson1_deserializeOpDeleteSequenceStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4431,7 +4350,6 @@ func (m *awsRestjson1_deserializeOpDeleteSequenceStore
 	output := &DeleteSequenceStoreOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4522,10 +4440,8 @@ func (m *awsRestjson1_deserializeOpDeleteShare) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4565,7 +4481,6 @@ func (m *awsRestjson1_deserializeOpDeleteShare) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4696,10 +4611,8 @@ func (m *awsRestjson1_deserializeOpDeleteVariantStore)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4739,7 +4652,6 @@ func (m *awsRestjson1_deserializeOpDeleteVariantStore)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4867,10 +4779,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4888,7 +4798,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4982,10 +4891,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5003,7 +4910,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkflowVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5097,10 +5003,8 @@ func (m *awsRestjson1_deserializeOpGetAnnotationImport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5140,7 +5044,6 @@ func (m *awsRestjson1_deserializeOpGetAnnotationImport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5373,10 +5276,8 @@ func (m *awsRestjson1_deserializeOpGetAnnotationStore)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5416,7 +5317,6 @@ func (m *awsRestjson1_deserializeOpGetAnnotationStore)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5667,10 +5567,8 @@ func (m *awsRestjson1_deserializeOpGetAnnotationStoreV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5710,7 +5608,6 @@ func (m *awsRestjson1_deserializeOpGetAnnotationStoreV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5947,10 +5844,8 @@ func (m *awsRestjson1_deserializeOpGetReadSet) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5967,7 +5862,6 @@ func (m *awsRestjson1_deserializeOpGetReadSet) HandleD
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6070,10 +5964,8 @@ func (m *awsRestjson1_deserializeOpGetReadSetActivatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6113,7 +6005,6 @@ func (m *awsRestjson1_deserializeOpGetReadSetActivatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6299,10 +6190,8 @@ func (m *awsRestjson1_deserializeOpGetReadSetExportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6342,7 +6231,6 @@ func (m *awsRestjson1_deserializeOpGetReadSetExportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6537,10 +6425,8 @@ func (m *awsRestjson1_deserializeOpGetReadSetImportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6580,7 +6466,6 @@ func (m *awsRestjson1_deserializeOpGetReadSetImportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6775,10 +6660,8 @@ func (m *awsRestjson1_deserializeOpGetReadSetMetadata)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6818,7 +6701,6 @@ func (m *awsRestjson1_deserializeOpGetReadSetMetadata)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7082,10 +6964,8 @@ func (m *awsRestjson1_deserializeOpGetReference) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7102,7 +6982,6 @@ func (m *awsRestjson1_deserializeOpGetReference) Handl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7202,10 +7081,8 @@ func (m *awsRestjson1_deserializeOpGetReferenceImportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7245,7 +7122,6 @@ func (m *awsRestjson1_deserializeOpGetReferenceImportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7440,10 +7316,8 @@ func (m *awsRestjson1_deserializeOpGetReferenceMetadat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7483,7 +7357,6 @@ func (m *awsRestjson1_deserializeOpGetReferenceMetadat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7714,10 +7587,8 @@ func (m *awsRestjson1_deserializeOpGetReferenceStore) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7757,7 +7628,6 @@ func (m *awsRestjson1_deserializeOpGetReferenceStore) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7930,10 +7800,8 @@ func (m *awsRestjson1_deserializeOpGetRun) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7973,7 +7841,6 @@ func (m *awsRestjson1_deserializeOpGetRun) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8426,10 +8293,8 @@ func (m *awsRestjson1_deserializeOpGetRunCache) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8469,7 +8334,6 @@ func (m *awsRestjson1_deserializeOpGetRunCache) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8684,10 +8548,8 @@ func (m *awsRestjson1_deserializeOpGetRunGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8727,7 +8589,6 @@ func (m *awsRestjson1_deserializeOpGetRunGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8949,10 +8810,8 @@ func (m *awsRestjson1_deserializeOpGetRunTask) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8992,7 +8851,6 @@ func (m *awsRestjson1_deserializeOpGetRunTask) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9276,10 +9134,8 @@ func (m *awsRestjson1_deserializeOpGetS3AccessPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9319,7 +9175,6 @@ func (m *awsRestjson1_deserializeOpGetS3AccessPolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9493,10 +9348,8 @@ func (m *awsRestjson1_deserializeOpGetSequenceStore) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9536,7 +9389,6 @@ func (m *awsRestjson1_deserializeOpGetSequenceStore) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9768,10 +9620,8 @@ func (m *awsRestjson1_deserializeOpGetShare) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9811,7 +9661,6 @@ func (m *awsRestjson1_deserializeOpGetShare) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9938,10 +9787,8 @@ func (m *awsRestjson1_deserializeOpGetVariantImportJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9981,7 +9828,6 @@ func (m *awsRestjson1_deserializeOpGetVariantImportJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10200,10 +10046,8 @@ func (m *awsRestjson1_deserializeOpGetVariantStore) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10243,7 +10087,6 @@ func (m *awsRestjson1_deserializeOpGetVariantStore) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10467,10 +10310,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10510,7 +10351,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflow) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10802,10 +10642,8 @@ func (m *awsRestjson1_deserializeOpGetWorkflowVersion)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10845,7 +10683,6 @@ func (m *awsRestjson1_deserializeOpGetWorkflowVersion)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11146,10 +10983,8 @@ func (m *awsRestjson1_deserializeOpListAnnotationImpor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11189,7 +11024,6 @@ func (m *awsRestjson1_deserializeOpListAnnotationImpor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11319,10 +11153,8 @@ func (m *awsRestjson1_deserializeOpListAnnotationStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11362,7 +11194,6 @@ func (m *awsRestjson1_deserializeOpListAnnotationStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11492,10 +11323,8 @@ func (m *awsRestjson1_deserializeOpListAnnotationStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11535,7 +11364,6 @@ func (m *awsRestjson1_deserializeOpListAnnotationStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11665,10 +11493,8 @@ func (m *awsRestjson1_deserializeOpListMultipartReadSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11708,7 +11534,6 @@ func (m *awsRestjson1_deserializeOpListMultipartReadSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11847,10 +11672,8 @@ func (m *awsRestjson1_deserializeOpListReadSetActivati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11890,7 +11713,6 @@ func (m *awsRestjson1_deserializeOpListReadSetActivati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12023,10 +11845,8 @@ func (m *awsRestjson1_deserializeOpListReadSetExportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12066,7 +11886,6 @@ func (m *awsRestjson1_deserializeOpListReadSetExportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12199,10 +12018,8 @@ func (m *awsRestjson1_deserializeOpListReadSetImportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12242,7 +12059,6 @@ func (m *awsRestjson1_deserializeOpListReadSetImportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12375,10 +12191,8 @@ func (m *awsRestjson1_deserializeOpListReadSets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12418,7 +12232,6 @@ func (m *awsRestjson1_deserializeOpListReadSets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12551,10 +12364,8 @@ func (m *awsRestjson1_deserializeOpListReadSetUploadPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12594,7 +12405,6 @@ func (m *awsRestjson1_deserializeOpListReadSetUploadPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12733,10 +12543,8 @@ func (m *awsRestjson1_deserializeOpListReferenceImport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12776,7 +12584,6 @@ func (m *awsRestjson1_deserializeOpListReferenceImport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12909,10 +12716,8 @@ func (m *awsRestjson1_deserializeOpListReferences) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12952,7 +12757,6 @@ func (m *awsRestjson1_deserializeOpListReferences) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13085,10 +12889,8 @@ func (m *awsRestjson1_deserializeOpListReferenceStores
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13128,7 +12930,6 @@ func (m *awsRestjson1_deserializeOpListReferenceStores
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13258,10 +13059,8 @@ func (m *awsRestjson1_deserializeOpListRunCaches) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13301,7 +13100,6 @@ func (m *awsRestjson1_deserializeOpListRunCaches) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13440,10 +13238,8 @@ func (m *awsRestjson1_deserializeOpListRunGroups) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13483,7 +13279,6 @@ func (m *awsRestjson1_deserializeOpListRunGroups) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13622,10 +13417,8 @@ func (m *awsRestjson1_deserializeOpListRuns) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13665,7 +13458,6 @@ func (m *awsRestjson1_deserializeOpListRuns) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13804,10 +13596,8 @@ func (m *awsRestjson1_deserializeOpListRunTasks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13847,7 +13637,6 @@ func (m *awsRestjson1_deserializeOpListRunTasks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13986,10 +13775,8 @@ func (m *awsRestjson1_deserializeOpListSequenceStores)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14029,7 +13816,6 @@ func (m *awsRestjson1_deserializeOpListSequenceStores)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14159,10 +13945,8 @@ func (m *awsRestjson1_deserializeOpListShares) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14202,7 +13986,6 @@ func (m *awsRestjson1_deserializeOpListShares) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14338,10 +14121,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14381,7 +14162,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14511,10 +14291,8 @@ func (m *awsRestjson1_deserializeOpListVariantImportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14554,7 +14332,6 @@ func (m *awsRestjson1_deserializeOpListVariantImportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14684,10 +14461,8 @@ func (m *awsRestjson1_deserializeOpListVariantStores) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14727,7 +14502,6 @@ func (m *awsRestjson1_deserializeOpListVariantStores) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14857,10 +14631,8 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14900,7 +14672,6 @@ func (m *awsRestjson1_deserializeOpListWorkflows) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15039,10 +14810,8 @@ func (m *awsRestjson1_deserializeOpListWorkflowVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15082,7 +14851,6 @@ func (m *awsRestjson1_deserializeOpListWorkflowVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15221,10 +14989,8 @@ func (m *awsRestjson1_deserializeOpPutS3AccessPolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15264,7 +15030,6 @@ func (m *awsRestjson1_deserializeOpPutS3AccessPolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15413,10 +15178,8 @@ func (m *awsRestjson1_deserializeOpStartAnnotationImpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15456,7 +15219,6 @@ func (m *awsRestjson1_deserializeOpStartAnnotationImpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15584,10 +15346,8 @@ func (m *awsRestjson1_deserializeOpStartReadSetActivat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15627,7 +15387,6 @@ func (m *awsRestjson1_deserializeOpStartReadSetActivat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15789,10 +15548,8 @@ func (m *awsRestjson1_deserializeOpStartReadSetExportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15832,7 +15589,6 @@ func (m *awsRestjson1_deserializeOpStartReadSetExportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16003,10 +15759,8 @@ func (m *awsRestjson1_deserializeOpStartReadSetImportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16046,7 +15800,6 @@ func (m *awsRestjson1_deserializeOpStartReadSetImportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16217,10 +15970,8 @@ func (m *awsRestjson1_deserializeOpStartReferenceImpor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16260,7 +16011,6 @@ func (m *awsRestjson1_deserializeOpStartReferenceImpor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16431,10 +16181,8 @@ func (m *awsRestjson1_deserializeOpStartRun) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16474,7 +16222,6 @@ func (m *awsRestjson1_deserializeOpStartRun) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16649,10 +16396,8 @@ func (m *awsRestjson1_deserializeOpStartVariantImportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16692,7 +16437,6 @@ func (m *awsRestjson1_deserializeOpStartVariantImportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16820,10 +16564,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16835,7 +16577,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16929,10 +16670,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16944,7 +16683,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17038,10 +16776,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17081,7 +16817,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17278,10 +17013,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnnotationSto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17321,7 +17054,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnnotationSto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17517,10 +17249,8 @@ func (m *awsRestjson1_deserializeOpUpdateRunCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17538,7 +17268,6 @@ func (m *awsRestjson1_deserializeOpUpdateRunCache) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17632,10 +17361,8 @@ func (m *awsRestjson1_deserializeOpUpdateRunGroup) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17653,7 +17380,6 @@ func (m *awsRestjson1_deserializeOpUpdateRunGroup) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17747,10 +17473,8 @@ func (m *awsRestjson1_deserializeOpUpdateSequenceStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17790,7 +17514,6 @@ func (m *awsRestjson1_deserializeOpUpdateSequenceStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18025,10 +17748,8 @@ func (m *awsRestjson1_deserializeOpUpdateVariantStore)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18068,7 +17789,6 @@ func (m *awsRestjson1_deserializeOpUpdateVariantStore)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18251,10 +17971,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflow) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18272,7 +17990,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflow) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18366,10 +18083,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18387,7 +18102,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkflowVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18481,10 +18195,8 @@ func (m *awsRestjson1_deserializeOpUploadReadSetPart) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18524,7 +18236,6 @@ func (m *awsRestjson1_deserializeOpUploadReadSetPart) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
