--- vendor/github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpDeleteEarthObservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpDeleteEarthObservat
 	output := &DeleteEarthObservationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestjson1_deserializeOpDeleteVectorEnrichm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -163,7 +157,6 @@ func (m *awsRestjson1_deserializeOpDeleteVectorEnrichm
 	output := &DeleteVectorEnrichmentJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -251,10 +244,8 @@ func (m *awsRestjson1_deserializeOpExportEarthObservat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -294,7 +285,6 @@ func (m *awsRestjson1_deserializeOpExportEarthObservat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -470,10 +460,8 @@ func (m *awsRestjson1_deserializeOpExportVectorEnrichm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -513,7 +501,6 @@ func (m *awsRestjson1_deserializeOpExportVectorEnrichm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -680,10 +667,8 @@ func (m *awsRestjson1_deserializeOpGetEarthObservation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -723,7 +708,6 @@ func (m *awsRestjson1_deserializeOpGetEarthObservation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -949,10 +933,8 @@ func (m *awsRestjson1_deserializeOpGetRasterDataCollec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -992,7 +974,6 @@ func (m *awsRestjson1_deserializeOpGetRasterDataCollec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1168,10 +1149,8 @@ func (m *awsRestjson1_deserializeOpGetTile) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1188,7 +1167,6 @@ func (m *awsRestjson1_deserializeOpGetTile) HandleDese
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1282,10 +1260,8 @@ func (m *awsRestjson1_deserializeOpGetVectorEnrichment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1325,7 +1301,6 @@ func (m *awsRestjson1_deserializeOpGetVectorEnrichment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1555,10 +1530,8 @@ func (m *awsRestjson1_deserializeOpListEarthObservatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1598,7 +1571,6 @@ func (m *awsRestjson1_deserializeOpListEarthObservatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1728,10 +1700,8 @@ func (m *awsRestjson1_deserializeOpListRasterDataColle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1771,7 +1741,6 @@ func (m *awsRestjson1_deserializeOpListRasterDataColle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1901,10 +1870,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1944,7 +1911,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2065,10 +2031,8 @@ func (m *awsRestjson1_deserializeOpListVectorEnrichmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2108,7 +2072,6 @@ func (m *awsRestjson1_deserializeOpListVectorEnrichmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2238,10 +2201,8 @@ func (m *awsRestjson1_deserializeOpSearchRasterDataCol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2281,7 +2242,6 @@ func (m *awsRestjson1_deserializeOpSearchRasterDataCol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2424,10 +2384,8 @@ func (m *awsRestjson1_deserializeOpStartEarthObservati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2467,7 +2425,6 @@ func (m *awsRestjson1_deserializeOpStartEarthObservati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2675,10 +2632,8 @@ func (m *awsRestjson1_deserializeOpStartVectorEnrichme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2718,7 +2673,6 @@ func (m *awsRestjson1_deserializeOpStartVectorEnrichme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2935,10 +2889,8 @@ func (m *awsRestjson1_deserializeOpStopEarthObservatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2950,7 +2902,6 @@ func (m *awsRestjson1_deserializeOpStopEarthObservatio
 	output := &StopEarthObservationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3038,10 +2989,8 @@ func (m *awsRestjson1_deserializeOpStopVectorEnrichmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3053,7 +3002,6 @@ func (m *awsRestjson1_deserializeOpStopVectorEnrichmen
 	output := &StopVectorEnrichmentJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3141,10 +3089,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3156,7 +3102,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3241,10 +3186,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3256,7 +3199,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
