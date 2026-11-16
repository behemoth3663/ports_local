--- vendor/github.com/aws/aws-sdk-go-v2/service/neptunegraph/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/neptunegraph/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelExportTask) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCancelExportTask) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -285,10 +281,8 @@ func (m *awsRestjson1_deserializeOpCancelImportTask) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -328,7 +322,6 @@ func (m *awsRestjson1_deserializeOpCancelImportTask) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -507,10 +500,8 @@ func (m *awsRestjson1_deserializeOpCancelQuery) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -528,7 +519,6 @@ func (m *awsRestjson1_deserializeOpCancelQuery) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -613,10 +603,8 @@ func (m *awsRestjson1_deserializeOpCreateGraph) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -656,7 +644,6 @@ func (m *awsRestjson1_deserializeOpCreateGraph) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -918,10 +905,8 @@ func (m *awsRestjson1_deserializeOpCreateGraphSnapshot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -961,7 +946,6 @@ func (m *awsRestjson1_deserializeOpCreateGraphSnapshot
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1150,10 +1134,8 @@ func (m *awsRestjson1_deserializeOpCreateGraphUsingImp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1193,7 +1175,6 @@ func (m *awsRestjson1_deserializeOpCreateGraphUsingImp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1377,10 +1358,8 @@ func (m *awsRestjson1_deserializeOpCreatePrivateGraphE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1420,7 +1399,6 @@ func (m *awsRestjson1_deserializeOpCreatePrivateGraphE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1571,10 +1549,8 @@ func (m *awsRestjson1_deserializeOpDeleteGraph) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1614,7 +1590,6 @@ func (m *awsRestjson1_deserializeOpDeleteGraph) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1876,10 +1851,8 @@ func (m *awsRestjson1_deserializeOpDeleteGraphSnapshot
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1919,7 +1892,6 @@ func (m *awsRestjson1_deserializeOpDeleteGraphSnapshot
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2105,10 +2077,8 @@ func (m *awsRestjson1_deserializeOpDeletePrivateGraphE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2148,7 +2118,6 @@ func (m *awsRestjson1_deserializeOpDeletePrivateGraphE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2296,10 +2265,8 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2316,7 +2283,6 @@ func (m *awsRestjson1_deserializeOpExecuteQuery) Handl
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2413,10 +2379,8 @@ func (m *awsRestjson1_deserializeOpGetExportTask) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2456,7 +2420,6 @@ func (m *awsRestjson1_deserializeOpGetExportTask) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2660,10 +2623,8 @@ func (m *awsRestjson1_deserializeOpGetGraph) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2703,7 +2664,6 @@ func (m *awsRestjson1_deserializeOpGetGraph) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2962,10 +2922,8 @@ func (m *awsRestjson1_deserializeOpGetGraphSnapshot) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3005,7 +2963,6 @@ func (m *awsRestjson1_deserializeOpGetGraphSnapshot) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3188,10 +3145,8 @@ func (m *awsRestjson1_deserializeOpGetGraphSummary) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3231,7 +3186,6 @@ func (m *awsRestjson1_deserializeOpGetGraphSummary) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3374,10 +3328,8 @@ func (m *awsRestjson1_deserializeOpGetImportTask) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3417,7 +3369,6 @@ func (m *awsRestjson1_deserializeOpGetImportTask) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3625,10 +3576,8 @@ func (m *awsRestjson1_deserializeOpGetPrivateGraphEndp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3668,7 +3617,6 @@ func (m *awsRestjson1_deserializeOpGetPrivateGraphEndp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3813,10 +3761,8 @@ func (m *awsRestjson1_deserializeOpGetQuery) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3856,7 +3802,6 @@ func (m *awsRestjson1_deserializeOpGetQuery) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4025,10 +3970,8 @@ func (m *awsRestjson1_deserializeOpListExportTasks) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4068,7 +4011,6 @@ func (m *awsRestjson1_deserializeOpListExportTasks) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4195,10 +4137,8 @@ func (m *awsRestjson1_deserializeOpListGraphs) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4238,7 +4178,6 @@ func (m *awsRestjson1_deserializeOpListGraphs) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4362,10 +4301,8 @@ func (m *awsRestjson1_deserializeOpListGraphSnapshots)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4405,7 +4342,6 @@ func (m *awsRestjson1_deserializeOpListGraphSnapshots)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4532,10 +4468,8 @@ func (m *awsRestjson1_deserializeOpListImportTasks) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4575,7 +4509,6 @@ func (m *awsRestjson1_deserializeOpListImportTasks) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4702,10 +4635,8 @@ func (m *awsRestjson1_deserializeOpListPrivateGraphEnd
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4745,7 +4676,6 @@ func (m *awsRestjson1_deserializeOpListPrivateGraphEnd
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4872,10 +4802,8 @@ func (m *awsRestjson1_deserializeOpListQueries) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4915,7 +4843,6 @@ func (m *awsRestjson1_deserializeOpListQueries) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5033,10 +4960,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5076,7 +5001,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5194,10 +5118,8 @@ func (m *awsRestjson1_deserializeOpResetGraph) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5237,7 +5159,6 @@ func (m *awsRestjson1_deserializeOpResetGraph) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5499,10 +5420,8 @@ func (m *awsRestjson1_deserializeOpRestoreGraphFromSna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5542,7 +5461,6 @@ func (m *awsRestjson1_deserializeOpRestoreGraphFromSna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5807,10 +5725,8 @@ func (m *awsRestjson1_deserializeOpStartExportTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5850,7 +5766,6 @@ func (m *awsRestjson1_deserializeOpStartExportTask) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6052,10 +5967,8 @@ func (m *awsRestjson1_deserializeOpStartImportTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6095,7 +6008,6 @@ func (m *awsRestjson1_deserializeOpStartImportTask) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6279,10 +6191,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6294,7 +6204,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6376,10 +6285,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6391,7 +6298,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6473,10 +6379,8 @@ func (m *awsRestjson1_deserializeOpUpdateGraph) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6516,7 +6420,6 @@ func (m *awsRestjson1_deserializeOpUpdateGraph) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
