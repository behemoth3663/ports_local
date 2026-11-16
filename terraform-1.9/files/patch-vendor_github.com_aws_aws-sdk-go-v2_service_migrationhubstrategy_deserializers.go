--- vendor/github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationCompo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationCompo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -221,10 +217,8 @@ func (m *awsRestjson1_deserializeOpGetApplicationCompo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -264,7 +258,6 @@ func (m *awsRestjson1_deserializeOpGetApplicationCompo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -379,10 +372,8 @@ func (m *awsRestjson1_deserializeOpGetAssessment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -422,7 +413,6 @@ func (m *awsRestjson1_deserializeOpGetAssessment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -554,10 +544,8 @@ func (m *awsRestjson1_deserializeOpGetImportFileTask) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -597,7 +585,6 @@ func (m *awsRestjson1_deserializeOpGetImportFileTask) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -834,10 +821,8 @@ func (m *awsRestjson1_deserializeOpGetLatestAssessment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -877,7 +862,6 @@ func (m *awsRestjson1_deserializeOpGetLatestAssessment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -999,10 +983,8 @@ func (m *awsRestjson1_deserializeOpGetPortfolioPrefere
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1042,7 +1024,6 @@ func (m *awsRestjson1_deserializeOpGetPortfolioPrefere
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1179,10 +1160,8 @@ func (m *awsRestjson1_deserializeOpGetPortfolioSummary
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1222,7 +1201,6 @@ func (m *awsRestjson1_deserializeOpGetPortfolioSummary
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1337,10 +1315,8 @@ func (m *awsRestjson1_deserializeOpGetRecommendationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1380,7 +1356,6 @@ func (m *awsRestjson1_deserializeOpGetRecommendationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1510,10 +1485,8 @@ func (m *awsRestjson1_deserializeOpGetServerDetails) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1553,7 +1526,6 @@ func (m *awsRestjson1_deserializeOpGetServerDetails) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1688,10 +1660,8 @@ func (m *awsRestjson1_deserializeOpGetServerStrategies
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1731,7 +1701,6 @@ func (m *awsRestjson1_deserializeOpGetServerStrategies
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1852,10 +1821,8 @@ func (m *awsRestjson1_deserializeOpListAnalyzableServe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1895,7 +1862,6 @@ func (m *awsRestjson1_deserializeOpListAnalyzableServe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2022,10 +1988,8 @@ func (m *awsRestjson1_deserializeOpListApplicationComp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2065,7 +2029,6 @@ func (m *awsRestjson1_deserializeOpListApplicationComp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2192,10 +2155,8 @@ func (m *awsRestjson1_deserializeOpListCollectors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2235,7 +2196,6 @@ func (m *awsRestjson1_deserializeOpListCollectors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2362,10 +2322,8 @@ func (m *awsRestjson1_deserializeOpListImportFileTask)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2405,7 +2363,6 @@ func (m *awsRestjson1_deserializeOpListImportFileTask)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2532,10 +2489,8 @@ func (m *awsRestjson1_deserializeOpListServers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2575,7 +2530,6 @@ func (m *awsRestjson1_deserializeOpListServers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2702,10 +2656,8 @@ func (m *awsRestjson1_deserializeOpPutPortfolioPrefere
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2717,7 +2669,6 @@ func (m *awsRestjson1_deserializeOpPutPortfolioPrefere
 	output := &PutPortfolioPreferencesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2802,10 +2753,8 @@ func (m *awsRestjson1_deserializeOpStartAssessment) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2845,7 +2794,6 @@ func (m *awsRestjson1_deserializeOpStartAssessment) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2967,10 +2915,8 @@ func (m *awsRestjson1_deserializeOpStartImportFileTask
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3010,7 +2956,6 @@ func (m *awsRestjson1_deserializeOpStartImportFileTask
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3135,10 +3080,8 @@ func (m *awsRestjson1_deserializeOpStartRecommendation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3178,7 +3121,6 @@ func (m *awsRestjson1_deserializeOpStartRecommendation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3303,10 +3245,8 @@ func (m *awsRestjson1_deserializeOpStopAssessment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3318,7 +3258,6 @@ func (m *awsRestjson1_deserializeOpStopAssessment) Han
 	output := &StopAssessmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3400,10 +3339,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3415,7 +3352,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationCo
 	output := &UpdateApplicationComponentConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3497,10 +3433,8 @@ func (m *awsRestjson1_deserializeOpUpdateServerConfig)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3512,7 +3446,6 @@ func (m *awsRestjson1_deserializeOpUpdateServerConfig)
 	output := &UpdateServerConfigOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
