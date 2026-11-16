--- vendor/github.com/aws/aws-sdk-go-v2/service/cognitosync/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cognitosync/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBulkPublish) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBulkPublish) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -259,7 +253,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -383,10 +376,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -426,7 +417,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -547,10 +537,8 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityPoo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -590,7 +578,6 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityPoo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -711,10 +698,8 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityUsa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -754,7 +739,6 @@ func (m *awsRestjson1_deserializeOpDescribeIdentityUsa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -875,10 +859,8 @@ func (m *awsRestjson1_deserializeOpGetBulkPublishDetai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -918,7 +900,6 @@ func (m *awsRestjson1_deserializeOpGetBulkPublishDetai
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1090,10 +1071,8 @@ func (m *awsRestjson1_deserializeOpGetCognitoEvents) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1133,7 +1112,6 @@ func (m *awsRestjson1_deserializeOpGetCognitoEvents) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1254,10 +1232,8 @@ func (m *awsRestjson1_deserializeOpGetIdentityPoolConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1297,7 +1273,6 @@ func (m *awsRestjson1_deserializeOpGetIdentityPoolConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1432,10 +1407,8 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1475,7 +1448,6 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1615,10 +1587,8 @@ func (m *awsRestjson1_deserializeOpListIdentityPoolUsa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1658,7 +1628,6 @@ func (m *awsRestjson1_deserializeOpListIdentityPoolUsa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1811,10 +1780,8 @@ func (m *awsRestjson1_deserializeOpListRecords) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1854,7 +1821,6 @@ func (m *awsRestjson1_deserializeOpListRecords) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2048,10 +2014,8 @@ func (m *awsRestjson1_deserializeOpRegisterDevice) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2091,7 +2055,6 @@ func (m *awsRestjson1_deserializeOpRegisterDevice) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2219,10 +2182,8 @@ func (m *awsRestjson1_deserializeOpSetCognitoEvents) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2240,7 +2201,6 @@ func (m *awsRestjson1_deserializeOpSetCognitoEvents) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2325,10 +2285,8 @@ func (m *awsRestjson1_deserializeOpSetIdentityPoolConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2368,7 +2326,6 @@ func (m *awsRestjson1_deserializeOpSetIdentityPoolConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2506,10 +2463,8 @@ func (m *awsRestjson1_deserializeOpSubscribeToDataset)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2521,7 +2476,6 @@ func (m *awsRestjson1_deserializeOpSubscribeToDataset)
 	output := &SubscribeToDatasetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2609,10 +2563,8 @@ func (m *awsRestjson1_deserializeOpUnsubscribeFromData
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2624,7 +2576,6 @@ func (m *awsRestjson1_deserializeOpUnsubscribeFromData
 	output := &UnsubscribeFromDatasetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2712,10 +2663,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecords) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2755,7 +2704,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecords) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
