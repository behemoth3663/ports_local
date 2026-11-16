--- vendor/github.com/aws/aws-sdk-go-v2/service/s3tables/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3tables/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateNamespace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateNamespace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -221,10 +217,8 @@ func (m *awsRestjson1_deserializeOpCreateTable) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -264,7 +258,6 @@ func (m *awsRestjson1_deserializeOpCreateTable) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -401,10 +394,8 @@ func (m *awsRestjson1_deserializeOpCreateTableBucket) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -444,7 +435,6 @@ func (m *awsRestjson1_deserializeOpCreateTableBucket) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -572,10 +562,8 @@ func (m *awsRestjson1_deserializeOpDeleteNamespace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -593,7 +581,6 @@ func (m *awsRestjson1_deserializeOpDeleteNamespace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -681,10 +668,8 @@ func (m *awsRestjson1_deserializeOpDeleteTable) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -702,7 +687,6 @@ func (m *awsRestjson1_deserializeOpDeleteTable) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -790,10 +774,8 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucket) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -811,7 +793,6 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucket) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -899,10 +880,8 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucketEn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -920,7 +899,6 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucketEn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1008,10 +986,8 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucketPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1029,7 +1005,6 @@ func (m *awsRestjson1_deserializeOpDeleteTableBucketPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1117,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpDeleteTablePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1138,7 +1111,6 @@ func (m *awsRestjson1_deserializeOpDeleteTablePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1226,10 +1198,8 @@ func (m *awsRestjson1_deserializeOpGetNamespace) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1269,7 +1239,6 @@ func (m *awsRestjson1_deserializeOpGetNamespace) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1445,10 +1414,8 @@ func (m *awsRestjson1_deserializeOpGetTable) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1488,7 +1455,6 @@ func (m *awsRestjson1_deserializeOpGetTable) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1758,10 +1724,8 @@ func (m *awsRestjson1_deserializeOpGetTableBucket) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1801,7 +1765,6 @@ func (m *awsRestjson1_deserializeOpGetTableBucket) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1981,10 +1944,8 @@ func (m *awsRestjson1_deserializeOpGetTableBucketEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2024,7 +1985,6 @@ func (m *awsRestjson1_deserializeOpGetTableBucketEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2148,10 +2108,8 @@ func (m *awsRestjson1_deserializeOpGetTableBucketMaint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2191,7 +2149,6 @@ func (m *awsRestjson1_deserializeOpGetTableBucketMaint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2324,10 +2281,8 @@ func (m *awsRestjson1_deserializeOpGetTableBucketPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2367,7 +2322,6 @@ func (m *awsRestjson1_deserializeOpGetTableBucketPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2495,10 +2449,8 @@ func (m *awsRestjson1_deserializeOpGetTableEncryption)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2538,7 +2490,6 @@ func (m *awsRestjson1_deserializeOpGetTableEncryption)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2662,10 +2613,8 @@ func (m *awsRestjson1_deserializeOpGetTableMaintenance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2705,7 +2654,6 @@ func (m *awsRestjson1_deserializeOpGetTableMaintenance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2838,10 +2786,8 @@ func (m *awsRestjson1_deserializeOpGetTableMaintenance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2881,7 +2827,6 @@ func (m *awsRestjson1_deserializeOpGetTableMaintenance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3014,10 +2959,8 @@ func (m *awsRestjson1_deserializeOpGetTableMetadataLoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3057,7 +3000,6 @@ func (m *awsRestjson1_deserializeOpGetTableMetadataLoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3203,10 +3145,8 @@ func (m *awsRestjson1_deserializeOpGetTablePolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3246,7 +3186,6 @@ func (m *awsRestjson1_deserializeOpGetTablePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3374,10 +3313,8 @@ func (m *awsRestjson1_deserializeOpListNamespaces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3417,7 +3354,6 @@ func (m *awsRestjson1_deserializeOpListNamespaces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3553,10 +3489,8 @@ func (m *awsRestjson1_deserializeOpListTableBuckets) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3596,7 +3530,6 @@ func (m *awsRestjson1_deserializeOpListTableBuckets) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3732,10 +3665,8 @@ func (m *awsRestjson1_deserializeOpListTables) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3775,7 +3706,6 @@ func (m *awsRestjson1_deserializeOpListTables) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3908,10 +3838,8 @@ func (m *awsRestjson1_deserializeOpPutTableBucketEncry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3929,7 +3857,6 @@ func (m *awsRestjson1_deserializeOpPutTableBucketEncry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4017,10 +3944,8 @@ func (m *awsRestjson1_deserializeOpPutTableBucketMaint
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4038,7 +3963,6 @@ func (m *awsRestjson1_deserializeOpPutTableBucketMaint
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4126,10 +4050,8 @@ func (m *awsRestjson1_deserializeOpPutTableBucketPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4147,7 +4069,6 @@ func (m *awsRestjson1_deserializeOpPutTableBucketPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4235,10 +4156,8 @@ func (m *awsRestjson1_deserializeOpPutTableMaintenance
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4256,7 +4175,6 @@ func (m *awsRestjson1_deserializeOpPutTableMaintenance
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4344,10 +4262,8 @@ func (m *awsRestjson1_deserializeOpPutTablePolicy) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4365,7 +4281,6 @@ func (m *awsRestjson1_deserializeOpPutTablePolicy) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4453,10 +4368,8 @@ func (m *awsRestjson1_deserializeOpRenameTable) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4474,7 +4387,6 @@ func (m *awsRestjson1_deserializeOpRenameTable) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4562,10 +4474,8 @@ func (m *awsRestjson1_deserializeOpUpdateTableMetadata
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4605,7 +4515,6 @@ func (m *awsRestjson1_deserializeOpUpdateTableMetadata
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
