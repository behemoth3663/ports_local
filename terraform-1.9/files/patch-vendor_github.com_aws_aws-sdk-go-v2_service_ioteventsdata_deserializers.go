--- vendor/github.com/aws/aws-sdk-go-v2/service/ioteventsdata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ioteventsdata/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpBatchAcknowledgeAla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpBatchAcknowledgeAla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -205,10 +201,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDetector
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -248,7 +242,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDetector
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -366,10 +359,8 @@ func (m *awsRestjson1_deserializeOpBatchDisableAlarm) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -409,7 +400,6 @@ func (m *awsRestjson1_deserializeOpBatchDisableAlarm) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -527,10 +517,8 @@ func (m *awsRestjson1_deserializeOpBatchEnableAlarm) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -570,7 +558,6 @@ func (m *awsRestjson1_deserializeOpBatchEnableAlarm) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -688,10 +675,8 @@ func (m *awsRestjson1_deserializeOpBatchPutMessage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -731,7 +716,6 @@ func (m *awsRestjson1_deserializeOpBatchPutMessage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -849,10 +833,8 @@ func (m *awsRestjson1_deserializeOpBatchResetAlarm) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -892,7 +874,6 @@ func (m *awsRestjson1_deserializeOpBatchResetAlarm) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1010,10 +991,8 @@ func (m *awsRestjson1_deserializeOpBatchSnoozeAlarm) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1053,7 +1032,6 @@ func (m *awsRestjson1_deserializeOpBatchSnoozeAlarm) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1171,10 +1149,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateDetector
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1214,7 +1190,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateDetector
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1332,10 +1307,8 @@ func (m *awsRestjson1_deserializeOpDescribeAlarm) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1375,7 +1348,6 @@ func (m *awsRestjson1_deserializeOpDescribeAlarm) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1496,10 +1468,8 @@ func (m *awsRestjson1_deserializeOpDescribeDetector) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1539,7 +1509,6 @@ func (m *awsRestjson1_deserializeOpDescribeDetector) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1660,10 +1629,8 @@ func (m *awsRestjson1_deserializeOpListAlarms) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1703,7 +1670,6 @@ func (m *awsRestjson1_deserializeOpListAlarms) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1833,10 +1799,8 @@ func (m *awsRestjson1_deserializeOpListDetectors) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1876,7 +1840,6 @@ func (m *awsRestjson1_deserializeOpListDetectors) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
