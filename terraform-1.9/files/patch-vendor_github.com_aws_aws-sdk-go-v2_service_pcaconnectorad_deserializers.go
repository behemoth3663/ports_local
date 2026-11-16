--- vendor/github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -219,10 +215,8 @@ func (m *awsRestjson1_deserializeOpCreateDirectoryRegi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -262,7 +256,6 @@ func (m *awsRestjson1_deserializeOpCreateDirectoryRegi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -390,10 +383,8 @@ func (m *awsRestjson1_deserializeOpCreateServicePrinci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -411,7 +402,6 @@ func (m *awsRestjson1_deserializeOpCreateServicePrinci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -499,10 +489,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -542,7 +530,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -673,10 +660,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplateGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -694,7 +679,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplateGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -785,10 +769,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -806,7 +788,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -894,10 +875,8 @@ func (m *awsRestjson1_deserializeOpDeleteDirectoryRegi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -915,7 +894,6 @@ func (m *awsRestjson1_deserializeOpDeleteDirectoryRegi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1000,10 +978,8 @@ func (m *awsRestjson1_deserializeOpDeleteServicePrinci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1021,7 +997,6 @@ func (m *awsRestjson1_deserializeOpDeleteServicePrinci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1106,10 +1081,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1127,7 +1100,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1215,10 +1187,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1236,7 +1206,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1324,10 +1293,8 @@ func (m *awsRestjson1_deserializeOpGetConnector) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1367,7 +1334,6 @@ func (m *awsRestjson1_deserializeOpGetConnector) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1488,10 +1454,8 @@ func (m *awsRestjson1_deserializeOpGetDirectoryRegistr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1531,7 +1495,6 @@ func (m *awsRestjson1_deserializeOpGetDirectoryRegistr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1652,10 +1615,8 @@ func (m *awsRestjson1_deserializeOpGetServicePrincipal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1695,7 +1656,6 @@ func (m *awsRestjson1_deserializeOpGetServicePrincipal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1816,10 +1776,8 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1859,7 +1817,6 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1980,10 +1937,8 @@ func (m *awsRestjson1_deserializeOpGetTemplateGroupAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2023,7 +1978,6 @@ func (m *awsRestjson1_deserializeOpGetTemplateGroupAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2144,10 +2098,8 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2187,7 +2139,6 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2314,10 +2265,8 @@ func (m *awsRestjson1_deserializeOpListDirectoryRegist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2357,7 +2306,6 @@ func (m *awsRestjson1_deserializeOpListDirectoryRegist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2484,10 +2432,8 @@ func (m *awsRestjson1_deserializeOpListServicePrincipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2527,7 +2473,6 @@ func (m *awsRestjson1_deserializeOpListServicePrincipa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2657,10 +2602,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2700,7 +2643,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2821,10 +2763,8 @@ func (m *awsRestjson1_deserializeOpListTemplateGroupAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2864,7 +2804,6 @@ func (m *awsRestjson1_deserializeOpListTemplateGroupAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2994,10 +2933,8 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3037,7 +2974,6 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3167,10 +3103,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3188,7 +3122,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3273,10 +3206,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3294,7 +3225,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3379,10 +3309,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3400,7 +3328,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3488,10 +3415,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3509,7 +3434,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
