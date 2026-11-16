--- vendor/github.com/aws/aws-sdk-go-v2/service/lookoutvision/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/lookoutvision/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -387,10 +380,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -430,7 +421,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -557,10 +547,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -572,7 +560,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 	output := &DeleteDatasetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -660,10 +647,8 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -703,7 +688,6 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -831,10 +815,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -874,7 +856,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1002,10 +983,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1045,7 +1024,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1169,10 +1147,8 @@ func (m *awsRestjson1_deserializeOpDescribeModel) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1212,7 +1188,6 @@ func (m *awsRestjson1_deserializeOpDescribeModel) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1336,10 +1311,8 @@ func (m *awsRestjson1_deserializeOpDescribeModelPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1379,7 +1352,6 @@ func (m *awsRestjson1_deserializeOpDescribeModelPackag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1500,10 +1472,8 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1543,7 +1513,6 @@ func (m *awsRestjson1_deserializeOpDescribeProject) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1667,10 +1636,8 @@ func (m *awsRestjson1_deserializeOpDetectAnomalies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1710,7 +1677,6 @@ func (m *awsRestjson1_deserializeOpDetectAnomalies) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1834,10 +1800,8 @@ func (m *awsRestjson1_deserializeOpListDatasetEntries)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1877,7 +1841,6 @@ func (m *awsRestjson1_deserializeOpListDatasetEntries)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2010,10 +1973,8 @@ func (m *awsRestjson1_deserializeOpListModelPackagingJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2053,7 +2014,6 @@ func (m *awsRestjson1_deserializeOpListModelPackagingJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2183,10 +2143,8 @@ func (m *awsRestjson1_deserializeOpListModels) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2226,7 +2184,6 @@ func (m *awsRestjson1_deserializeOpListModels) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2359,10 +2316,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2402,7 +2357,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2535,10 +2489,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2578,7 +2530,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2702,10 +2653,8 @@ func (m *awsRestjson1_deserializeOpStartModel) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2745,7 +2694,6 @@ func (m *awsRestjson1_deserializeOpStartModel) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2876,10 +2824,8 @@ func (m *awsRestjson1_deserializeOpStartModelPackaging
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2919,7 +2865,6 @@ func (m *awsRestjson1_deserializeOpStartModelPackaging
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3050,10 +2995,8 @@ func (m *awsRestjson1_deserializeOpStopModel) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3093,7 +3036,6 @@ func (m *awsRestjson1_deserializeOpStopModel) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3221,10 +3163,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3236,7 +3176,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3327,10 +3266,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3342,7 +3279,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3430,10 +3366,8 @@ func (m *awsRestjson1_deserializeOpUpdateDatasetEntrie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3473,7 +3407,6 @@ func (m *awsRestjson1_deserializeOpUpdateDatasetEntrie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
