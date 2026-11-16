--- vendor/github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -313,10 +309,8 @@ func (m *awsRestjson1_deserializeOpCreateApplicationVe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -356,7 +350,6 @@ func (m *awsRestjson1_deserializeOpCreateApplicationVe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -545,10 +538,8 @@ func (m *awsRestjson1_deserializeOpCreateCloudFormatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -588,7 +579,6 @@ func (m *awsRestjson1_deserializeOpCreateCloudFormatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -737,10 +727,8 @@ func (m *awsRestjson1_deserializeOpCreateCloudFormatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -780,7 +768,6 @@ func (m *awsRestjson1_deserializeOpCreateCloudFormatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -959,10 +946,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -980,7 +965,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1068,10 +1052,8 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1111,7 +1093,6 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1336,10 +1317,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1379,7 +1358,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1500,10 +1478,8 @@ func (m *awsRestjson1_deserializeOpGetCloudFormationTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1543,7 +1519,6 @@ func (m *awsRestjson1_deserializeOpGetCloudFormationTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1722,10 +1697,8 @@ func (m *awsRestjson1_deserializeOpListApplicationDepe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1765,7 +1738,6 @@ func (m *awsRestjson1_deserializeOpListApplicationDepe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1895,10 +1867,8 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1938,7 +1908,6 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2065,10 +2034,8 @@ func (m *awsRestjson1_deserializeOpListApplicationVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2108,7 +2075,6 @@ func (m *awsRestjson1_deserializeOpListApplicationVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2238,10 +2204,8 @@ func (m *awsRestjson1_deserializeOpPutApplicationPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2281,7 +2245,6 @@ func (m *awsRestjson1_deserializeOpPutApplicationPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2402,10 +2365,8 @@ func (m *awsRestjson1_deserializeOpUnshareApplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2423,7 +2384,6 @@ func (m *awsRestjson1_deserializeOpUnshareApplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2508,10 +2468,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2551,7 +2509,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
