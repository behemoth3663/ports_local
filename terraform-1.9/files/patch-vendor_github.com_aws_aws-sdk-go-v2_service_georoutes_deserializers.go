--- vendor/github.com/aws/aws-sdk-go-v2/service/georoutes/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/georoutes/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCalculateIsolines) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpCalculateIsolines) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -260,10 +256,8 @@ func (m *awsRestjson1_deserializeOpCalculateRouteMatri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -308,7 +302,6 @@ func (m *awsRestjson1_deserializeOpCalculateRouteMatri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -456,10 +449,8 @@ func (m *awsRestjson1_deserializeOpCalculateRoutes) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -504,7 +495,6 @@ func (m *awsRestjson1_deserializeOpCalculateRoutes) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -648,10 +638,8 @@ func (m *awsRestjson1_deserializeOpOptimizeWaypoints) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -696,7 +684,6 @@ func (m *awsRestjson1_deserializeOpOptimizeWaypoints) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -867,10 +854,8 @@ func (m *awsRestjson1_deserializeOpSnapToRoads) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -915,7 +900,6 @@ func (m *awsRestjson1_deserializeOpSnapToRoads) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
