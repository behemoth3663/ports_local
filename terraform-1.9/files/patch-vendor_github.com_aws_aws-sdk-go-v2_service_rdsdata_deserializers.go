--- vendor/github.com/aws/aws-sdk-go-v2/service/rdsdata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/rdsdata/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpBatchExecuteStateme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpBatchExecuteStateme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -240,10 +236,8 @@ func (m *awsRestjson1_deserializeOpBeginTransaction) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -283,7 +277,6 @@ func (m *awsRestjson1_deserializeOpBeginTransaction) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -438,10 +431,8 @@ func (m *awsRestjson1_deserializeOpCommitTransaction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -481,7 +472,6 @@ func (m *awsRestjson1_deserializeOpCommitTransaction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -636,10 +626,8 @@ func (m *awsRestjson1_deserializeOpExecuteSql) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -679,7 +667,6 @@ func (m *awsRestjson1_deserializeOpExecuteSql) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -800,10 +787,8 @@ func (m *awsRestjson1_deserializeOpExecuteStatement) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -843,7 +828,6 @@ func (m *awsRestjson1_deserializeOpExecuteStatement) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1029,10 +1013,8 @@ func (m *awsRestjson1_deserializeOpRollbackTransaction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1072,7 +1054,6 @@ func (m *awsRestjson1_deserializeOpRollbackTransaction
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
