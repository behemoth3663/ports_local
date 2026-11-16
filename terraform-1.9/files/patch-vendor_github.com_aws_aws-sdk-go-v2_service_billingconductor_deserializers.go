--- vendor/github.com/aws/aws-sdk-go-v2/service/billingconductor/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/billingconductor/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociateAccounts) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAssociateAccounts) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -220,10 +216,8 @@ func (m *awsRestjson1_deserializeOpAssociatePricingRul
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -263,7 +257,6 @@ func (m *awsRestjson1_deserializeOpAssociatePricingRul
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -394,10 +387,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -437,7 +428,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -569,10 +559,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -612,7 +600,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -741,10 +728,8 @@ func (m *awsRestjson1_deserializeOpCreateBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -784,7 +769,6 @@ func (m *awsRestjson1_deserializeOpCreateBillingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -912,10 +896,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomLineIte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -955,7 +937,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomLineIte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1083,10 +1064,8 @@ func (m *awsRestjson1_deserializeOpCreatePricingPlan) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1126,7 +1105,6 @@ func (m *awsRestjson1_deserializeOpCreatePricingPlan) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1257,10 +1235,8 @@ func (m *awsRestjson1_deserializeOpCreatePricingRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1300,7 +1276,6 @@ func (m *awsRestjson1_deserializeOpCreatePricingRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1428,10 +1403,8 @@ func (m *awsRestjson1_deserializeOpDeleteBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1471,7 +1444,6 @@ func (m *awsRestjson1_deserializeOpDeleteBillingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1593,10 +1565,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomLineIte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1636,7 +1606,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomLineIte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1761,10 +1730,8 @@ func (m *awsRestjson1_deserializeOpDeletePricingPlan) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1804,7 +1771,6 @@ func (m *awsRestjson1_deserializeOpDeletePricingPlan) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1929,10 +1895,8 @@ func (m *awsRestjson1_deserializeOpDeletePricingRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1972,7 +1936,6 @@ func (m *awsRestjson1_deserializeOpDeletePricingRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2097,10 +2060,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAccount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2140,7 +2101,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAccount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2268,10 +2228,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePricing
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2311,7 +2269,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePricing
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2439,10 +2396,8 @@ func (m *awsRestjson1_deserializeOpGetBillingGroupCost
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2482,7 +2437,6 @@ func (m *awsRestjson1_deserializeOpGetBillingGroupCost
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2612,10 +2566,8 @@ func (m *awsRestjson1_deserializeOpListAccountAssociat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2655,7 +2607,6 @@ func (m *awsRestjson1_deserializeOpListAccountAssociat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2785,10 +2736,8 @@ func (m *awsRestjson1_deserializeOpListBillingGroupCos
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2828,7 +2777,6 @@ func (m *awsRestjson1_deserializeOpListBillingGroupCos
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2958,10 +2906,8 @@ func (m *awsRestjson1_deserializeOpListBillingGroups) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3001,7 +2947,6 @@ func (m *awsRestjson1_deserializeOpListBillingGroups) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3131,10 +3076,8 @@ func (m *awsRestjson1_deserializeOpListCustomLineItems
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3174,7 +3117,6 @@ func (m *awsRestjson1_deserializeOpListCustomLineItems
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3304,10 +3246,8 @@ func (m *awsRestjson1_deserializeOpListCustomLineItemV
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3347,7 +3287,6 @@ func (m *awsRestjson1_deserializeOpListCustomLineItemV
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3474,10 +3413,8 @@ func (m *awsRestjson1_deserializeOpListPricingPlans) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3517,7 +3454,6 @@ func (m *awsRestjson1_deserializeOpListPricingPlans) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3653,10 +3589,8 @@ func (m *awsRestjson1_deserializeOpListPricingPlansAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3696,7 +3630,6 @@ func (m *awsRestjson1_deserializeOpListPricingPlansAss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3844,10 +3777,8 @@ func (m *awsRestjson1_deserializeOpListPricingRules) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3887,7 +3818,6 @@ func (m *awsRestjson1_deserializeOpListPricingRules) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4023,10 +3953,8 @@ func (m *awsRestjson1_deserializeOpListPricingRulesAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4066,7 +3994,6 @@ func (m *awsRestjson1_deserializeOpListPricingRulesAss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4214,10 +4141,8 @@ func (m *awsRestjson1_deserializeOpListResourcesAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4257,7 +4182,6 @@ func (m *awsRestjson1_deserializeOpListResourcesAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4396,10 +4320,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4439,7 +4361,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4560,10 +4481,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4575,7 +4494,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4660,10 +4578,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4675,7 +4591,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4760,10 +4675,8 @@ func (m *awsRestjson1_deserializeOpUpdateBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4803,7 +4716,6 @@ func (m *awsRestjson1_deserializeOpUpdateBillingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5016,10 +4928,8 @@ func (m *awsRestjson1_deserializeOpUpdateCustomLineIte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5059,7 +4969,6 @@ func (m *awsRestjson1_deserializeOpUpdateCustomLineIte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5242,10 +5151,8 @@ func (m *awsRestjson1_deserializeOpUpdatePricingPlan) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5285,7 +5192,6 @@ func (m *awsRestjson1_deserializeOpUpdatePricingPlan) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5457,10 +5363,8 @@ func (m *awsRestjson1_deserializeOpUpdatePricingRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5500,7 +5404,6 @@ func (m *awsRestjson1_deserializeOpUpdatePricingRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
