--- vendor/github.com/aws/aws-sdk-go-v2/service/amp/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/amp/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpCreateAlertManagerD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpCreateAlertManagerD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -217,10 +213,8 @@ func (m *awsRestjson1_deserializeOpCreateLoggingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -260,7 +254,6 @@ func (m *awsRestjson1_deserializeOpCreateLoggingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -378,10 +371,8 @@ func (m *awsRestjson1_deserializeOpCreateQueryLoggingC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -421,7 +412,6 @@ func (m *awsRestjson1_deserializeOpCreateQueryLoggingC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -539,10 +529,8 @@ func (m *awsRestjson1_deserializeOpCreateRuleGroupsNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -582,7 +570,6 @@ func (m *awsRestjson1_deserializeOpCreateRuleGroupsNam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -732,10 +719,8 @@ func (m *awsRestjson1_deserializeOpCreateScraper) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -775,7 +760,6 @@ func (m *awsRestjson1_deserializeOpCreateScraper) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -925,10 +909,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -968,7 +950,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1124,10 +1105,8 @@ func (m *awsRestjson1_deserializeOpDeleteAlertManagerD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1145,7 +1124,6 @@ func (m *awsRestjson1_deserializeOpDeleteAlertManagerD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1233,10 +1211,8 @@ func (m *awsRestjson1_deserializeOpDeleteLoggingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1254,7 +1230,6 @@ func (m *awsRestjson1_deserializeOpDeleteLoggingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1339,10 +1314,8 @@ func (m *awsRestjson1_deserializeOpDeleteQueryLoggingC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1360,7 +1333,6 @@ func (m *awsRestjson1_deserializeOpDeleteQueryLoggingC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1445,10 +1417,8 @@ func (m *awsRestjson1_deserializeOpDeleteRuleGroupsNam
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1466,7 +1436,6 @@ func (m *awsRestjson1_deserializeOpDeleteRuleGroupsNam
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1554,10 +1523,8 @@ func (m *awsRestjson1_deserializeOpDeleteScraper) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1597,7 +1564,6 @@ func (m *awsRestjson1_deserializeOpDeleteScraper) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1730,10 +1696,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1751,7 +1715,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1839,10 +1802,8 @@ func (m *awsRestjson1_deserializeOpDescribeAlertManage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1882,7 +1843,6 @@ func (m *awsRestjson1_deserializeOpDescribeAlertManage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2003,10 +1963,8 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2046,7 +2004,6 @@ func (m *awsRestjson1_deserializeOpDescribeLoggingConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2164,10 +2121,8 @@ func (m *awsRestjson1_deserializeOpDescribeQueryLoggin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2207,7 +2162,6 @@ func (m *awsRestjson1_deserializeOpDescribeQueryLoggin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2325,10 +2279,8 @@ func (m *awsRestjson1_deserializeOpDescribeRuleGroupsN
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2368,7 +2320,6 @@ func (m *awsRestjson1_deserializeOpDescribeRuleGroupsN
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2489,10 +2440,8 @@ func (m *awsRestjson1_deserializeOpDescribeScraper) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2532,7 +2481,6 @@ func (m *awsRestjson1_deserializeOpDescribeScraper) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2653,10 +2601,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2696,7 +2642,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2817,10 +2762,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2860,7 +2803,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2981,10 +2923,8 @@ func (m *awsRestjson1_deserializeOpGetDefaultScraperCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3024,7 +2964,6 @@ func (m *awsRestjson1_deserializeOpGetDefaultScraperCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3147,10 +3086,8 @@ func (m *awsRestjson1_deserializeOpListRuleGroupsNames
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3190,7 +3127,6 @@ func (m *awsRestjson1_deserializeOpListRuleGroupsNames
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3320,10 +3256,8 @@ func (m *awsRestjson1_deserializeOpListScrapers) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3363,7 +3297,6 @@ func (m *awsRestjson1_deserializeOpListScrapers) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3490,10 +3423,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3533,7 +3464,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3654,10 +3584,8 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3697,7 +3625,6 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3824,10 +3751,8 @@ func (m *awsRestjson1_deserializeOpPutAlertManagerDefi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3867,7 +3792,6 @@ func (m *awsRestjson1_deserializeOpPutAlertManagerDefi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3994,10 +3918,8 @@ func (m *awsRestjson1_deserializeOpPutRuleGroupsNamesp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4037,7 +3959,6 @@ func (m *awsRestjson1_deserializeOpPutRuleGroupsNamesp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4187,10 +4108,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4202,7 +4121,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4287,10 +4205,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4302,7 +4218,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4387,10 +4302,8 @@ func (m *awsRestjson1_deserializeOpUpdateLoggingConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4430,7 +4343,6 @@ func (m *awsRestjson1_deserializeOpUpdateLoggingConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4551,10 +4463,8 @@ func (m *awsRestjson1_deserializeOpUpdateQueryLoggingC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4594,7 +4504,6 @@ func (m *awsRestjson1_deserializeOpUpdateQueryLoggingC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4715,10 +4624,8 @@ func (m *awsRestjson1_deserializeOpUpdateScraper) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4758,7 +4665,6 @@ func (m *awsRestjson1_deserializeOpUpdateScraper) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4908,10 +4814,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceAlia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4929,7 +4833,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceAlia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5020,10 +4923,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5063,7 +4964,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
