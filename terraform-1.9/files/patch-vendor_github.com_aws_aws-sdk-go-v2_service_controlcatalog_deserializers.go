--- vendor/github.com/aws/aws-sdk-go-v2/service/controlcatalog/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/controlcatalog/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpGetControl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpGetControl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -289,10 +285,8 @@ func (m *awsRestjson1_deserializeOpListCommonControls)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -332,7 +326,6 @@ func (m *awsRestjson1_deserializeOpListCommonControls)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -459,10 +452,8 @@ func (m *awsRestjson1_deserializeOpListControlMappings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -502,7 +493,6 @@ func (m *awsRestjson1_deserializeOpListControlMappings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -629,10 +619,8 @@ func (m *awsRestjson1_deserializeOpListControls) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -672,7 +660,6 @@ func (m *awsRestjson1_deserializeOpListControls) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -799,10 +786,8 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -842,7 +827,6 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -969,10 +953,8 @@ func (m *awsRestjson1_deserializeOpListObjectives) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1012,7 +994,6 @@ func (m *awsRestjson1_deserializeOpListObjectives) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
