--- vendor/github.com/aws/aws-sdk-go-v2/service/detective/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/detective/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -152,10 +148,8 @@ func (m *awsRestjson1_deserializeOpBatchGetGraphMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -195,7 +189,6 @@ func (m *awsRestjson1_deserializeOpBatchGetGraphMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -318,10 +311,8 @@ func (m *awsRestjson1_deserializeOpBatchGetMembershipD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -361,7 +352,6 @@ func (m *awsRestjson1_deserializeOpBatchGetMembershipD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -484,10 +474,8 @@ func (m *awsRestjson1_deserializeOpCreateGraph) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -527,7 +515,6 @@ func (m *awsRestjson1_deserializeOpCreateGraph) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -649,10 +636,8 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -692,7 +677,6 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -818,10 +802,8 @@ func (m *awsRestjson1_deserializeOpDeleteGraph) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -839,7 +821,6 @@ func (m *awsRestjson1_deserializeOpDeleteGraph) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -921,10 +902,8 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -964,7 +943,6 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1090,10 +1068,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1133,7 +1109,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1255,10 +1230,8 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1276,7 +1249,6 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1358,10 +1330,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1379,7 +1349,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1464,10 +1433,8 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1485,7 +1452,6 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1567,10 +1533,8 @@ func (m *awsRestjson1_deserializeOpGetInvestigation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1610,7 +1574,6 @@ func (m *awsRestjson1_deserializeOpGetInvestigation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1828,10 +1791,8 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1871,7 +1832,6 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1994,10 +1954,8 @@ func (m *awsRestjson1_deserializeOpListDatasourcePacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2037,7 +1995,6 @@ func (m *awsRestjson1_deserializeOpListDatasourcePacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2164,10 +2121,8 @@ func (m *awsRestjson1_deserializeOpListGraphs) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2207,7 +2162,6 @@ func (m *awsRestjson1_deserializeOpListGraphs) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2331,10 +2285,8 @@ func (m *awsRestjson1_deserializeOpListIndicators) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2374,7 +2326,6 @@ func (m *awsRestjson1_deserializeOpListIndicators) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2522,10 +2473,8 @@ func (m *awsRestjson1_deserializeOpListInvestigations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2565,7 +2514,6 @@ func (m *awsRestjson1_deserializeOpListInvestigations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2695,10 +2643,8 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2738,7 +2684,6 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2862,10 +2807,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2905,7 +2848,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3032,10 +2974,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3075,7 +3015,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3202,10 +3141,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3245,7 +3182,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3363,10 +3299,8 @@ func (m *awsRestjson1_deserializeOpRejectInvitation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3384,7 +3318,6 @@ func (m *awsRestjson1_deserializeOpRejectInvitation) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3469,10 +3402,8 @@ func (m *awsRestjson1_deserializeOpStartInvestigation)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3512,7 +3443,6 @@ func (m *awsRestjson1_deserializeOpStartInvestigation)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3637,10 +3567,8 @@ func (m *awsRestjson1_deserializeOpStartMonitoringMemb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3658,7 +3586,6 @@ func (m *awsRestjson1_deserializeOpStartMonitoringMemb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3746,10 +3673,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3761,7 +3686,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3843,10 +3767,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3858,7 +3780,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3940,10 +3861,8 @@ func (m *awsRestjson1_deserializeOpUpdateDatasourcePac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3961,7 +3880,6 @@ func (m *awsRestjson1_deserializeOpUpdateDatasourcePac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4046,10 +3964,8 @@ func (m *awsRestjson1_deserializeOpUpdateInvestigation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4067,7 +3983,6 @@ func (m *awsRestjson1_deserializeOpUpdateInvestigation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4152,10 +4067,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4173,7 +4086,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
