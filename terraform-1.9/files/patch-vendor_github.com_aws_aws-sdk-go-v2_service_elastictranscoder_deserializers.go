--- vendor/github.com/aws/aws-sdk-go-v2/service/elastictranscoder/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/elastictranscoder/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 	output := &CancelJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -147,10 +143,8 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -190,7 +184,6 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -314,10 +307,8 @@ func (m *awsRestjson1_deserializeOpCreatePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -357,7 +348,6 @@ func (m *awsRestjson1_deserializeOpCreatePipeline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -486,10 +476,8 @@ func (m *awsRestjson1_deserializeOpCreatePreset) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -529,7 +517,6 @@ func (m *awsRestjson1_deserializeOpCreatePreset) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -659,10 +646,8 @@ func (m *awsRestjson1_deserializeOpDeletePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -674,7 +659,6 @@ func (m *awsRestjson1_deserializeOpDeletePipeline) Han
 	output := &DeletePipelineOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -762,10 +746,8 @@ func (m *awsRestjson1_deserializeOpDeletePreset) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -777,7 +759,6 @@ func (m *awsRestjson1_deserializeOpDeletePreset) Handl
 	output := &DeletePresetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -862,10 +843,8 @@ func (m *awsRestjson1_deserializeOpListJobsByPipeline)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -905,7 +884,6 @@ func (m *awsRestjson1_deserializeOpListJobsByPipeline)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1035,10 +1013,8 @@ func (m *awsRestjson1_deserializeOpListJobsByStatus) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1078,7 +1054,6 @@ func (m *awsRestjson1_deserializeOpListJobsByStatus) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1208,10 +1183,8 @@ func (m *awsRestjson1_deserializeOpListPipelines) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1251,7 +1224,6 @@ func (m *awsRestjson1_deserializeOpListPipelines) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1378,10 +1350,8 @@ func (m *awsRestjson1_deserializeOpListPresets) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1421,7 +1391,6 @@ func (m *awsRestjson1_deserializeOpListPresets) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1548,10 +1517,8 @@ func (m *awsRestjson1_deserializeOpReadJob) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1591,7 +1558,6 @@ func (m *awsRestjson1_deserializeOpReadJob) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1712,10 +1678,8 @@ func (m *awsRestjson1_deserializeOpReadPipeline) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1755,7 +1719,6 @@ func (m *awsRestjson1_deserializeOpReadPipeline) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1881,10 +1844,8 @@ func (m *awsRestjson1_deserializeOpReadPreset) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1924,7 +1885,6 @@ func (m *awsRestjson1_deserializeOpReadPreset) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2045,10 +2005,8 @@ func (m *awsRestjson1_deserializeOpTestRole) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2088,7 +2046,6 @@ func (m *awsRestjson1_deserializeOpTestRole) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2218,10 +2175,8 @@ func (m *awsRestjson1_deserializeOpUpdatePipeline) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2261,7 +2216,6 @@ func (m *awsRestjson1_deserializeOpUpdatePipeline) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2390,10 +2344,8 @@ func (m *awsRestjson1_deserializeOpUpdatePipelineNotif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2433,7 +2385,6 @@ func (m *awsRestjson1_deserializeOpUpdatePipelineNotif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2557,10 +2508,8 @@ func (m *awsRestjson1_deserializeOpUpdatePipelineStatu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2600,7 +2549,6 @@ func (m *awsRestjson1_deserializeOpUpdatePipelineStatu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
