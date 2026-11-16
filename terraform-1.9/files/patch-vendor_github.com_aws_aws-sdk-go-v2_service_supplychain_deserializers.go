--- vendor/github.com/aws/aws-sdk-go-v2/service/supplychain/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/supplychain/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateBillOfMateria
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateBillOfMateria
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -219,10 +215,8 @@ func (m *awsRestjson1_deserializeOpCreateDataIntegrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -262,7 +256,6 @@ func (m *awsRestjson1_deserializeOpCreateDataIntegrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -402,10 +395,8 @@ func (m *awsRestjson1_deserializeOpCreateDataLakeDatas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -445,7 +436,6 @@ func (m *awsRestjson1_deserializeOpCreateDataLakeDatas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -572,10 +562,8 @@ func (m *awsRestjson1_deserializeOpCreateDataLakeNames
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -615,7 +603,6 @@ func (m *awsRestjson1_deserializeOpCreateDataLakeNames
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -742,10 +729,8 @@ func (m *awsRestjson1_deserializeOpCreateInstance) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -785,7 +770,6 @@ func (m *awsRestjson1_deserializeOpCreateInstance) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -912,10 +896,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataIntegrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -955,7 +937,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataIntegrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1095,10 +1076,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataLakeDatas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1138,7 +1117,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataLakeDatas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1287,10 +1265,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataLakeNames
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1330,7 +1306,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataLakeNames
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1470,10 +1445,8 @@ func (m *awsRestjson1_deserializeOpDeleteInstance) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1513,7 +1486,6 @@ func (m *awsRestjson1_deserializeOpDeleteInstance) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1640,10 +1612,8 @@ func (m *awsRestjson1_deserializeOpGetBillOfMaterialsI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1683,7 +1653,6 @@ func (m *awsRestjson1_deserializeOpGetBillOfMaterialsI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1810,10 +1779,8 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1853,7 +1820,6 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1980,10 +1946,8 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2023,7 +1987,6 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2150,10 +2113,8 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationF
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2193,7 +2154,6 @@ func (m *awsRestjson1_deserializeOpGetDataIntegrationF
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2320,10 +2280,8 @@ func (m *awsRestjson1_deserializeOpGetDataLakeDataset)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2363,7 +2321,6 @@ func (m *awsRestjson1_deserializeOpGetDataLakeDataset)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2490,10 +2447,8 @@ func (m *awsRestjson1_deserializeOpGetDataLakeNamespac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2533,7 +2488,6 @@ func (m *awsRestjson1_deserializeOpGetDataLakeNamespac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2660,10 +2614,8 @@ func (m *awsRestjson1_deserializeOpGetInstance) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2703,7 +2655,6 @@ func (m *awsRestjson1_deserializeOpGetInstance) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2830,10 +2781,8 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2873,7 +2822,6 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3009,10 +2957,8 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3052,7 +2998,6 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3188,10 +3133,8 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3231,7 +3174,6 @@ func (m *awsRestjson1_deserializeOpListDataIntegration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3367,10 +3309,8 @@ func (m *awsRestjson1_deserializeOpListDataLakeDataset
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3410,7 +3350,6 @@ func (m *awsRestjson1_deserializeOpListDataLakeDataset
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3546,10 +3485,8 @@ func (m *awsRestjson1_deserializeOpListDataLakeNamespa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3589,7 +3526,6 @@ func (m *awsRestjson1_deserializeOpListDataLakeNamespa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3725,10 +3661,8 @@ func (m *awsRestjson1_deserializeOpListInstances) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3768,7 +3702,6 @@ func (m *awsRestjson1_deserializeOpListInstances) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3904,10 +3837,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3947,7 +3878,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4074,10 +4004,8 @@ func (m *awsRestjson1_deserializeOpSendDataIntegration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4117,7 +4045,6 @@ func (m *awsRestjson1_deserializeOpSendDataIntegration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4248,10 +4175,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4263,7 +4188,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4354,10 +4278,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4369,7 +4291,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4460,10 +4381,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataIntegrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4503,7 +4422,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataIntegrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4630,10 +4548,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataLakeDatas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4673,7 +4589,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataLakeDatas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4800,10 +4715,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataLakeNames
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4843,7 +4756,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataLakeNames
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4970,10 +4882,8 @@ func (m *awsRestjson1_deserializeOpUpdateInstance) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5013,7 +4923,6 @@ func (m *awsRestjson1_deserializeOpUpdateInstance) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
