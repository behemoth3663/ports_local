--- vendor/github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchGetTokenBalanc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchGetTokenBalanc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpGetAssetContract) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpGetAssetContract) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -407,10 +400,8 @@ func (m *awsRestjson1_deserializeOpGetTokenBalance) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -450,7 +441,6 @@ func (m *awsRestjson1_deserializeOpGetTokenBalance) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -598,10 +588,8 @@ func (m *awsRestjson1_deserializeOpGetTransaction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -641,7 +629,6 @@ func (m *awsRestjson1_deserializeOpGetTransaction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -765,10 +752,8 @@ func (m *awsRestjson1_deserializeOpListAssetContracts)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -808,7 +793,6 @@ func (m *awsRestjson1_deserializeOpListAssetContracts)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -938,10 +922,8 @@ func (m *awsRestjson1_deserializeOpListFilteredTransac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -981,7 +963,6 @@ func (m *awsRestjson1_deserializeOpListFilteredTransac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1111,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpListTokenBalances) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1154,7 +1133,6 @@ func (m *awsRestjson1_deserializeOpListTokenBalances) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1284,10 +1262,8 @@ func (m *awsRestjson1_deserializeOpListTransactionEven
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1327,7 +1303,6 @@ func (m *awsRestjson1_deserializeOpListTransactionEven
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1457,10 +1432,8 @@ func (m *awsRestjson1_deserializeOpListTransactions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1500,7 +1473,6 @@ func (m *awsRestjson1_deserializeOpListTransactions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
