--- vendor/github.com/aws/aws-sdk-go-v2/service/s3outposts/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3outposts/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateEndpoint) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateEndpoint) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -219,10 +215,8 @@ func (m *awsRestjson1_deserializeOpDeleteEndpoint) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -240,7 +234,6 @@ func (m *awsRestjson1_deserializeOpDeleteEndpoint) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -328,10 +321,8 @@ func (m *awsRestjson1_deserializeOpListEndpoints) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -371,7 +362,6 @@ func (m *awsRestjson1_deserializeOpListEndpoints) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -501,10 +491,8 @@ func (m *awsRestjson1_deserializeOpListOutpostsWithS3)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -544,7 +532,6 @@ func (m *awsRestjson1_deserializeOpListOutpostsWithS3)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -671,10 +658,8 @@ func (m *awsRestjson1_deserializeOpListSharedEndpoints
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -714,7 +699,6 @@ func (m *awsRestjson1_deserializeOpListSharedEndpoints
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
