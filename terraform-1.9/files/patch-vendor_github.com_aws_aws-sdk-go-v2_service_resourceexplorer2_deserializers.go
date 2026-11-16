--- vendor/github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/deserializers.go
@@ -16,7 +16,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAssociateDefaultVie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpAssociateDefaultVie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -215,10 +211,8 @@ func (m *awsRestjson1_deserializeOpBatchGetView) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -258,7 +252,6 @@ func (m *awsRestjson1_deserializeOpBatchGetView) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -384,10 +377,8 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -574,10 +564,8 @@ func (m *awsRestjson1_deserializeOpCreateView) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -617,7 +605,6 @@ func (m *awsRestjson1_deserializeOpCreateView) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -744,10 +731,8 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -787,7 +772,6 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -934,10 +918,8 @@ func (m *awsRestjson1_deserializeOpDeleteView) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -977,7 +959,6 @@ func (m *awsRestjson1_deserializeOpDeleteView) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1105,10 +1086,8 @@ func (m *awsRestjson1_deserializeOpDisassociateDefault
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1126,7 +1105,6 @@ func (m *awsRestjson1_deserializeOpDisassociateDefault
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1211,10 +1189,8 @@ func (m *awsRestjson1_deserializeOpGetAccountLevelServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1254,7 +1230,6 @@ func (m *awsRestjson1_deserializeOpGetAccountLevelServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1372,10 +1347,8 @@ func (m *awsRestjson1_deserializeOpGetDefaultView) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1415,7 +1388,6 @@ func (m *awsRestjson1_deserializeOpGetDefaultView) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1540,10 +1512,8 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1583,7 +1553,6 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1767,10 +1736,8 @@ func (m *awsRestjson1_deserializeOpGetManagedView) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1810,7 +1777,6 @@ func (m *awsRestjson1_deserializeOpGetManagedView) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1934,10 +1900,8 @@ func (m *awsRestjson1_deserializeOpGetView) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1977,7 +1941,6 @@ func (m *awsRestjson1_deserializeOpGetView) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2106,10 +2069,8 @@ func (m *awsRestjson1_deserializeOpListIndexes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2149,7 +2110,6 @@ func (m *awsRestjson1_deserializeOpListIndexes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2276,10 +2236,8 @@ func (m *awsRestjson1_deserializeOpListIndexesForMembe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2319,7 +2277,6 @@ func (m *awsRestjson1_deserializeOpListIndexesForMembe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2446,10 +2403,8 @@ func (m *awsRestjson1_deserializeOpListManagedViews) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2489,7 +2444,6 @@ func (m *awsRestjson1_deserializeOpListManagedViews) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2619,10 +2573,8 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2662,7 +2614,6 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2804,10 +2755,8 @@ func (m *awsRestjson1_deserializeOpListSupportedResour
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2847,7 +2796,6 @@ func (m *awsRestjson1_deserializeOpListSupportedResour
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2974,10 +2922,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3017,7 +2963,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3141,10 +3086,8 @@ func (m *awsRestjson1_deserializeOpListViews) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3184,7 +3127,6 @@ func (m *awsRestjson1_deserializeOpListViews) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3311,10 +3253,8 @@ func (m *awsRestjson1_deserializeOpSearch) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3354,7 +3294,6 @@ func (m *awsRestjson1_deserializeOpSearch) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3501,10 +3440,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3516,7 +3453,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3604,10 +3540,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3619,7 +3553,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3707,10 +3640,8 @@ func (m *awsRestjson1_deserializeOpUpdateIndexType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3750,7 +3681,6 @@ func (m *awsRestjson1_deserializeOpUpdateIndexType) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3912,10 +3842,8 @@ func (m *awsRestjson1_deserializeOpUpdateView) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3955,7 +3883,6 @@ func (m *awsRestjson1_deserializeOpUpdateView) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
