--- vendor/github.com/aws/aws-sdk-go-v2/service/geoplaces/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/geoplaces/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAutocomplete) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpAutocomplete) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -223,10 +219,8 @@ func (m *awsRestjson1_deserializeOpGeocode) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -271,7 +265,6 @@ func (m *awsRestjson1_deserializeOpGeocode) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -401,10 +394,8 @@ func (m *awsRestjson1_deserializeOpGetPlace) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -449,7 +440,6 @@ func (m *awsRestjson1_deserializeOpGetPlace) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -694,10 +684,8 @@ func (m *awsRestjson1_deserializeOpReverseGeocode) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -742,7 +730,6 @@ func (m *awsRestjson1_deserializeOpReverseGeocode) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -872,10 +859,8 @@ func (m *awsRestjson1_deserializeOpSearchNearby) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -920,7 +905,6 @@ func (m *awsRestjson1_deserializeOpSearchNearby) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1059,10 +1043,8 @@ func (m *awsRestjson1_deserializeOpSearchText) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1107,7 +1089,6 @@ func (m *awsRestjson1_deserializeOpSearchText) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1246,10 +1227,8 @@ func (m *awsRestjson1_deserializeOpSuggest) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1294,7 +1273,6 @@ func (m *awsRestjson1_deserializeOpSuggest) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
