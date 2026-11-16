--- vendor/github.com/aws/aws-sdk-go-v2/service/personalizeevents/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/personalizeevents/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpPutActionInteractio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpPutActionInteractio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -145,10 +141,8 @@ func (m *awsRestjson1_deserializeOpPutActions) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -166,7 +160,6 @@ func (m *awsRestjson1_deserializeOpPutActions) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -245,10 +238,8 @@ func (m *awsRestjson1_deserializeOpPutEvents) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -266,7 +257,6 @@ func (m *awsRestjson1_deserializeOpPutEvents) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -339,10 +329,8 @@ func (m *awsRestjson1_deserializeOpPutItems) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -360,7 +348,6 @@ func (m *awsRestjson1_deserializeOpPutItems) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -439,10 +426,8 @@ func (m *awsRestjson1_deserializeOpPutUsers) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -460,7 +445,6 @@ func (m *awsRestjson1_deserializeOpPutUsers) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
