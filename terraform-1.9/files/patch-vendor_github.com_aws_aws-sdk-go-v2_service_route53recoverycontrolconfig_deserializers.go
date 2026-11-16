--- vendor/github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -214,10 +210,8 @@ func (m *awsRestjson1_deserializeOpCreateControlPanel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -257,7 +251,6 @@ func (m *awsRestjson1_deserializeOpCreateControlPanel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -384,10 +377,8 @@ func (m *awsRestjson1_deserializeOpCreateRoutingContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -427,7 +418,6 @@ func (m *awsRestjson1_deserializeOpCreateRoutingContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -554,10 +544,8 @@ func (m *awsRestjson1_deserializeOpCreateSafetyRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -597,7 +585,6 @@ func (m *awsRestjson1_deserializeOpCreateSafetyRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -714,10 +701,8 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -729,7 +714,6 @@ func (m *awsRestjson1_deserializeOpDeleteCluster) Hand
 	output := &DeleteClusterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -817,10 +801,8 @@ func (m *awsRestjson1_deserializeOpDeleteControlPanel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -832,7 +814,6 @@ func (m *awsRestjson1_deserializeOpDeleteControlPanel)
 	output := &DeleteControlPanelOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -920,10 +901,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -935,7 +914,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingContro
 	output := &DeleteRoutingControlOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1023,10 +1001,8 @@ func (m *awsRestjson1_deserializeOpDeleteSafetyRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1038,7 +1014,6 @@ func (m *awsRestjson1_deserializeOpDeleteSafetyRule) H
 	output := &DeleteSafetyRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1117,10 +1092,8 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1160,7 +1133,6 @@ func (m *awsRestjson1_deserializeOpDescribeCluster) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1284,10 +1256,8 @@ func (m *awsRestjson1_deserializeOpDescribeControlPane
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1327,7 +1297,6 @@ func (m *awsRestjson1_deserializeOpDescribeControlPane
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1451,10 +1420,8 @@ func (m *awsRestjson1_deserializeOpDescribeRoutingCont
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1494,7 +1461,6 @@ func (m *awsRestjson1_deserializeOpDescribeRoutingCont
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1618,10 +1584,8 @@ func (m *awsRestjson1_deserializeOpDescribeSafetyRule)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1661,7 +1625,6 @@ func (m *awsRestjson1_deserializeOpDescribeSafetyRule)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1778,10 +1741,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1821,7 +1782,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1937,10 +1897,8 @@ func (m *awsRestjson1_deserializeOpListAssociatedRoute
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1980,7 +1938,6 @@ func (m *awsRestjson1_deserializeOpListAssociatedRoute
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2104,10 +2061,8 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2147,7 +2102,6 @@ func (m *awsRestjson1_deserializeOpListClusters) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2277,10 +2231,8 @@ func (m *awsRestjson1_deserializeOpListControlPanels) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2320,7 +2272,6 @@ func (m *awsRestjson1_deserializeOpListControlPanels) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2450,10 +2401,8 @@ func (m *awsRestjson1_deserializeOpListRoutingControls
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2493,7 +2442,6 @@ func (m *awsRestjson1_deserializeOpListRoutingControls
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2623,10 +2571,8 @@ func (m *awsRestjson1_deserializeOpListSafetyRules) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2666,7 +2612,6 @@ func (m *awsRestjson1_deserializeOpListSafetyRules) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2796,10 +2741,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2839,7 +2782,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2954,10 +2896,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2969,7 +2909,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3048,10 +2987,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3063,7 +3000,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3142,10 +3078,8 @@ func (m *awsRestjson1_deserializeOpUpdateCluster) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3185,7 +3119,6 @@ func (m *awsRestjson1_deserializeOpUpdateCluster) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3309,10 +3242,8 @@ func (m *awsRestjson1_deserializeOpUpdateControlPanel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3352,7 +3283,6 @@ func (m *awsRestjson1_deserializeOpUpdateControlPanel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3476,10 +3406,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingContro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3519,7 +3447,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoutingContro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3643,10 +3570,8 @@ func (m *awsRestjson1_deserializeOpUpdateSafetyRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3686,7 +3611,6 @@ func (m *awsRestjson1_deserializeOpUpdateSafetyRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
