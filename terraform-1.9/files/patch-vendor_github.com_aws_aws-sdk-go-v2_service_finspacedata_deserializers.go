--- vendor/github.com/aws/aws-sdk-go-v2/service/finspacedata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/finspacedata/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateUserToPerm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -64,7 +61,6 @@ func (m *awsRestjson1_deserializeOpAssociateUserToPerm
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -162,10 +158,8 @@ func (m *awsRestjson1_deserializeOpCreateChangeset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -205,7 +199,6 @@ func (m *awsRestjson1_deserializeOpCreateChangeset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -345,10 +338,8 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -388,7 +379,6 @@ func (m *awsRestjson1_deserializeOpCreateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -519,10 +509,8 @@ func (m *awsRestjson1_deserializeOpCreateDataView) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -562,7 +550,6 @@ func (m *awsRestjson1_deserializeOpCreateDataView) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -699,10 +686,8 @@ func (m *awsRestjson1_deserializeOpCreatePermissionGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -742,7 +727,6 @@ func (m *awsRestjson1_deserializeOpCreatePermissionGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -870,10 +854,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -913,7 +895,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1041,10 +1022,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1084,7 +1063,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1215,10 +1193,8 @@ func (m *awsRestjson1_deserializeOpDeletePermissionGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1258,7 +1234,6 @@ func (m *awsRestjson1_deserializeOpDeletePermissionGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1389,10 +1364,8 @@ func (m *awsRestjson1_deserializeOpDisableUser) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1432,7 +1405,6 @@ func (m *awsRestjson1_deserializeOpDisableUser) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1560,10 +1532,8 @@ func (m *awsRestjson1_deserializeOpDisassociateUserFro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1580,7 +1550,6 @@ func (m *awsRestjson1_deserializeOpDisassociateUserFro
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1678,10 +1647,8 @@ func (m *awsRestjson1_deserializeOpEnableUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1721,7 +1688,6 @@ func (m *awsRestjson1_deserializeOpEnableUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1852,10 +1818,8 @@ func (m *awsRestjson1_deserializeOpGetChangeset) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1895,7 +1859,6 @@ func (m *awsRestjson1_deserializeOpGetChangeset) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2131,10 +2094,8 @@ func (m *awsRestjson1_deserializeOpGetDataset) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2174,7 +2135,6 @@ func (m *awsRestjson1_deserializeOpGetDataset) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2387,10 +2347,8 @@ func (m *awsRestjson1_deserializeOpGetDataView) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2430,7 +2388,6 @@ func (m *awsRestjson1_deserializeOpGetDataView) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2650,10 +2607,8 @@ func (m *awsRestjson1_deserializeOpGetExternalDataView
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2693,7 +2648,6 @@ func (m *awsRestjson1_deserializeOpGetExternalDataView
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2819,10 +2773,8 @@ func (m *awsRestjson1_deserializeOpGetPermissionGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2862,7 +2814,6 @@ func (m *awsRestjson1_deserializeOpGetPermissionGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2983,10 +2934,8 @@ func (m *awsRestjson1_deserializeOpGetProgrammaticAcce
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3026,7 +2975,6 @@ func (m *awsRestjson1_deserializeOpGetProgrammaticAcce
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3157,10 +3105,8 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3200,7 +3146,6 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3453,10 +3398,8 @@ func (m *awsRestjson1_deserializeOpGetWorkingLocation)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3496,7 +3439,6 @@ func (m *awsRestjson1_deserializeOpGetWorkingLocation)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3636,10 +3578,8 @@ func (m *awsRestjson1_deserializeOpListChangesets) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3679,7 +3619,6 @@ func (m *awsRestjson1_deserializeOpListChangesets) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3812,10 +3751,8 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3855,7 +3792,6 @@ func (m *awsRestjson1_deserializeOpListDatasets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3985,10 +3921,8 @@ func (m *awsRestjson1_deserializeOpListDataViews) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4028,7 +3962,6 @@ func (m *awsRestjson1_deserializeOpListDataViews) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4158,10 +4091,8 @@ func (m *awsRestjson1_deserializeOpListPermissionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4201,7 +4132,6 @@ func (m *awsRestjson1_deserializeOpListPermissionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4328,10 +4258,8 @@ func (m *awsRestjson1_deserializeOpListPermissionGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4371,7 +4299,6 @@ func (m *awsRestjson1_deserializeOpListPermissionGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4501,10 +4428,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4544,7 +4469,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4671,10 +4595,8 @@ func (m *awsRestjson1_deserializeOpListUsersByPermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4714,7 +4636,6 @@ func (m *awsRestjson1_deserializeOpListUsersByPermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4844,10 +4765,8 @@ func (m *awsRestjson1_deserializeOpResetUserPassword) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4887,7 +4806,6 @@ func (m *awsRestjson1_deserializeOpResetUserPassword) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5024,10 +4942,8 @@ func (m *awsRestjson1_deserializeOpUpdateChangeset) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5067,7 +4983,6 @@ func (m *awsRestjson1_deserializeOpUpdateChangeset) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5204,10 +5119,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5247,7 +5160,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataset) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5375,10 +5287,8 @@ func (m *awsRestjson1_deserializeOpUpdatePermissionGro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5418,7 +5328,6 @@ func (m *awsRestjson1_deserializeOpUpdatePermissionGro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5546,10 +5455,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5589,7 +5496,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
