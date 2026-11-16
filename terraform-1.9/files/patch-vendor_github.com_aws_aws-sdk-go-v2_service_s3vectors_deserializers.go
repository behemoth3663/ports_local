--- vendor/github.com/aws/aws-sdk-go-v2/service/s3vectors/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3vectors/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -62,7 +59,6 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 	output := &CreateIndexOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -156,10 +152,8 @@ func (m *awsRestjson1_deserializeOpCreateVectorBucket)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -171,7 +165,6 @@ func (m *awsRestjson1_deserializeOpCreateVectorBucket)
 	output := &CreateVectorBucketOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -262,10 +255,8 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -277,7 +268,6 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 	output := &DeleteIndexOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -365,10 +355,8 @@ func (m *awsRestjson1_deserializeOpDeleteVectorBucket)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -380,7 +368,6 @@ func (m *awsRestjson1_deserializeOpDeleteVectorBucket)
 	output := &DeleteVectorBucketOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -471,10 +458,8 @@ func (m *awsRestjson1_deserializeOpDeleteVectorBucketP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -486,7 +471,6 @@ func (m *awsRestjson1_deserializeOpDeleteVectorBucketP
 	output := &DeleteVectorBucketPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -577,10 +561,8 @@ func (m *awsRestjson1_deserializeOpDeleteVectors) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -592,7 +574,6 @@ func (m *awsRestjson1_deserializeOpDeleteVectors) Hand
 	output := &DeleteVectorsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -695,10 +676,8 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -738,7 +717,6 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -865,10 +843,8 @@ func (m *awsRestjson1_deserializeOpGetVectorBucket) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -908,7 +884,6 @@ func (m *awsRestjson1_deserializeOpGetVectorBucket) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1035,10 +1010,8 @@ func (m *awsRestjson1_deserializeOpGetVectorBucketPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1078,7 +1051,6 @@ func (m *awsRestjson1_deserializeOpGetVectorBucketPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1209,10 +1181,8 @@ func (m *awsRestjson1_deserializeOpGetVectors) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1252,7 +1222,6 @@ func (m *awsRestjson1_deserializeOpGetVectors) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1391,10 +1360,8 @@ func (m *awsRestjson1_deserializeOpListIndexes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1434,7 +1401,6 @@ func (m *awsRestjson1_deserializeOpListIndexes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1570,10 +1536,8 @@ func (m *awsRestjson1_deserializeOpListVectorBuckets) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1613,7 +1577,6 @@ func (m *awsRestjson1_deserializeOpListVectorBuckets) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1746,10 +1709,8 @@ func (m *awsRestjson1_deserializeOpListVectors) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1789,7 +1750,6 @@ func (m *awsRestjson1_deserializeOpListVectors) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1925,10 +1885,8 @@ func (m *awsRestjson1_deserializeOpPutVectorBucketPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1940,7 +1898,6 @@ func (m *awsRestjson1_deserializeOpPutVectorBucketPoli
 	output := &PutVectorBucketPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2031,10 +1988,8 @@ func (m *awsRestjson1_deserializeOpPutVectors) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2046,7 +2001,6 @@ func (m *awsRestjson1_deserializeOpPutVectors) HandleD
 	output := &PutVectorsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2149,10 +2103,8 @@ func (m *awsRestjson1_deserializeOpQueryVectors) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2192,7 +2144,6 @@ func (m *awsRestjson1_deserializeOpQueryVectors) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
