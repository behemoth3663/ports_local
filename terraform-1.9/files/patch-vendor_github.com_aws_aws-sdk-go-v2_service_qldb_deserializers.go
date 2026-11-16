--- vendor/github.com/aws/aws-sdk-go-v2/service/qldb/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/qldb/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCancelJournalKinesi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCancelJournalKinesi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -208,10 +204,8 @@ func (m *awsRestjson1_deserializeOpCreateLedger) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -251,7 +245,6 @@ func (m *awsRestjson1_deserializeOpCreateLedger) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -434,10 +427,8 @@ func (m *awsRestjson1_deserializeOpDeleteLedger) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -455,7 +446,6 @@ func (m *awsRestjson1_deserializeOpDeleteLedger) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -537,10 +527,8 @@ func (m *awsRestjson1_deserializeOpDescribeJournalKine
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -580,7 +568,6 @@ func (m *awsRestjson1_deserializeOpDescribeJournalKine
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -695,10 +682,8 @@ func (m *awsRestjson1_deserializeOpDescribeJournalS3Ex
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -738,7 +723,6 @@ func (m *awsRestjson1_deserializeOpDescribeJournalS3Ex
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -847,10 +831,8 @@ func (m *awsRestjson1_deserializeOpDescribeLedger) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -890,7 +872,6 @@ func (m *awsRestjson1_deserializeOpDescribeLedger) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1063,10 +1044,8 @@ func (m *awsRestjson1_deserializeOpExportJournalToS3) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1106,7 +1085,6 @@ func (m *awsRestjson1_deserializeOpExportJournalToS3) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1222,10 +1200,8 @@ func (m *awsRestjson1_deserializeOpGetBlock) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1265,7 +1241,6 @@ func (m *awsRestjson1_deserializeOpGetBlock) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1385,10 +1360,8 @@ func (m *awsRestjson1_deserializeOpGetDigest) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1428,7 +1401,6 @@ func (m *awsRestjson1_deserializeOpGetDigest) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1556,10 +1528,8 @@ func (m *awsRestjson1_deserializeOpGetRevision) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1599,7 +1569,6 @@ func (m *awsRestjson1_deserializeOpGetRevision) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1719,10 +1688,8 @@ func (m *awsRestjson1_deserializeOpListJournalKinesisS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1762,7 +1729,6 @@ func (m *awsRestjson1_deserializeOpListJournalKinesisS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1886,10 +1852,8 @@ func (m *awsRestjson1_deserializeOpListJournalS3Export
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1929,7 +1893,6 @@ func (m *awsRestjson1_deserializeOpListJournalS3Export
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2044,10 +2007,8 @@ func (m *awsRestjson1_deserializeOpListJournalS3Export
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2087,7 +2048,6 @@ func (m *awsRestjson1_deserializeOpListJournalS3Export
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2202,10 +2162,8 @@ func (m *awsRestjson1_deserializeOpListLedgers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2245,7 +2203,6 @@ func (m *awsRestjson1_deserializeOpListLedgers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2360,10 +2317,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2403,7 +2358,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2515,10 +2469,8 @@ func (m *awsRestjson1_deserializeOpStreamJournalToKine
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2558,7 +2510,6 @@ func (m *awsRestjson1_deserializeOpStreamJournalToKine
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2677,10 +2628,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2692,7 +2641,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2768,10 +2716,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2783,7 +2729,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2859,10 +2804,8 @@ func (m *awsRestjson1_deserializeOpUpdateLedger) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2902,7 +2845,6 @@ func (m *awsRestjson1_deserializeOpUpdateLedger) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3066,10 +3008,8 @@ func (m *awsRestjson1_deserializeOpUpdateLedgerPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3109,7 +3049,6 @@ func (m *awsRestjson1_deserializeOpUpdateLedgerPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
