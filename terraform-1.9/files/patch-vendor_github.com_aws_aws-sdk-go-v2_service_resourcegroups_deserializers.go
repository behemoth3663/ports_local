--- vendor/github.com/aws/aws-sdk-go-v2/service/resourcegroups/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/resourcegroups/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelTagSyncTask) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -66,7 +63,6 @@ func (m *awsRestjson1_deserializeOpCancelTagSyncTask) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -154,10 +150,8 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -197,7 +191,6 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -333,10 +326,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -376,7 +367,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -500,10 +490,8 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -543,7 +531,6 @@ func (m *awsRestjson1_deserializeOpGetAccountSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -664,10 +651,8 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -707,7 +692,6 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -831,10 +815,8 @@ func (m *awsRestjson1_deserializeOpGetGroupConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -874,7 +856,6 @@ func (m *awsRestjson1_deserializeOpGetGroupConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -998,10 +979,8 @@ func (m *awsRestjson1_deserializeOpGetGroupQuery) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1041,7 +1020,6 @@ func (m *awsRestjson1_deserializeOpGetGroupQuery) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1165,10 +1143,8 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1208,7 +1184,6 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1341,10 +1316,8 @@ func (m *awsRestjson1_deserializeOpGetTagSyncTask) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1384,7 +1357,6 @@ func (m *awsRestjson1_deserializeOpGetTagSyncTask) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1599,10 +1571,8 @@ func (m *awsRestjson1_deserializeOpGroupResources) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1642,7 +1612,6 @@ func (m *awsRestjson1_deserializeOpGroupResources) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1776,10 +1745,8 @@ func (m *awsRestjson1_deserializeOpListGroupingStatuse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1819,7 +1786,6 @@ func (m *awsRestjson1_deserializeOpListGroupingStatuse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1958,10 +1924,8 @@ func (m *awsRestjson1_deserializeOpListGroupResources)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2001,7 +1965,6 @@ func (m *awsRestjson1_deserializeOpListGroupResources)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2147,10 +2110,8 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2190,7 +2151,6 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2325,10 +2285,8 @@ func (m *awsRestjson1_deserializeOpListTagSyncTasks) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2368,7 +2326,6 @@ func (m *awsRestjson1_deserializeOpListTagSyncTasks) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2501,10 +2458,8 @@ func (m *awsRestjson1_deserializeOpPutGroupConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2516,7 +2471,6 @@ func (m *awsRestjson1_deserializeOpPutGroupConfigurati
 	output := &PutGroupConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2604,10 +2558,8 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2647,7 +2599,6 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2785,10 +2736,8 @@ func (m *awsRestjson1_deserializeOpStartTagSyncTask) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2828,7 +2777,6 @@ func (m *awsRestjson1_deserializeOpStartTagSyncTask) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3009,10 +2957,8 @@ func (m *awsRestjson1_deserializeOpTag) HandleDeserial
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3052,7 +2998,6 @@ func (m *awsRestjson1_deserializeOpTag) HandleDeserial
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3185,10 +3130,8 @@ func (m *awsRestjson1_deserializeOpUngroupResources) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3228,7 +3171,6 @@ func (m *awsRestjson1_deserializeOpUngroupResources) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3362,10 +3304,8 @@ func (m *awsRestjson1_deserializeOpUntag) HandleDeseri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3405,7 +3345,6 @@ func (m *awsRestjson1_deserializeOpUntag) HandleDeseri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3538,10 +3477,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3581,7 +3518,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3702,10 +3638,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3745,7 +3679,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3869,10 +3802,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroupQuery) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3912,7 +3843,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroupQuery) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
