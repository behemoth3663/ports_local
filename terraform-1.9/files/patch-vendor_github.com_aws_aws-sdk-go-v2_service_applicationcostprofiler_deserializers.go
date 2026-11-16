--- vendor/github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpDeleteReportDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpDeleteReportDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -209,10 +205,8 @@ func (m *awsRestjson1_deserializeOpGetReportDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -252,7 +246,6 @@ func (m *awsRestjson1_deserializeOpGetReportDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -438,10 +431,8 @@ func (m *awsRestjson1_deserializeOpImportApplicationUs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -481,7 +472,6 @@ func (m *awsRestjson1_deserializeOpImportApplicationUs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -603,10 +593,8 @@ func (m *awsRestjson1_deserializeOpListReportDefinitio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -646,7 +634,6 @@ func (m *awsRestjson1_deserializeOpListReportDefinitio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -773,10 +760,8 @@ func (m *awsRestjson1_deserializeOpPutReportDefinition
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -816,7 +801,6 @@ func (m *awsRestjson1_deserializeOpPutReportDefinition
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -941,10 +925,8 @@ func (m *awsRestjson1_deserializeOpUpdateReportDefinit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -984,7 +966,6 @@ func (m *awsRestjson1_deserializeOpUpdateReportDefinit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
