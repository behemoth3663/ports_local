--- vendor/github.com/aws/aws-sdk-go-v2/service/amplify/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/amplify/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpCreateApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -208,10 +204,8 @@ func (m *awsRestjson1_deserializeOpCreateBackendEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -251,7 +245,6 @@ func (m *awsRestjson1_deserializeOpCreateBackendEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -372,10 +365,8 @@ func (m *awsRestjson1_deserializeOpCreateBranch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -415,7 +406,6 @@ func (m *awsRestjson1_deserializeOpCreateBranch) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -539,10 +529,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -582,7 +570,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -718,10 +705,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -761,7 +746,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -885,10 +869,8 @@ func (m *awsRestjson1_deserializeOpCreateWebhook) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -928,7 +910,6 @@ func (m *awsRestjson1_deserializeOpCreateWebhook) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1052,10 +1033,8 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1095,7 +1074,6 @@ func (m *awsRestjson1_deserializeOpDeleteApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1216,10 +1194,8 @@ func (m *awsRestjson1_deserializeOpDeleteBackendEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1259,7 +1235,6 @@ func (m *awsRestjson1_deserializeOpDeleteBackendEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1380,10 +1355,8 @@ func (m *awsRestjson1_deserializeOpDeleteBranch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1423,7 +1396,6 @@ func (m *awsRestjson1_deserializeOpDeleteBranch) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1544,10 +1516,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1587,7 +1557,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1708,10 +1677,8 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1751,7 +1718,6 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1872,10 +1838,8 @@ func (m *awsRestjson1_deserializeOpDeleteWebhook) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1915,7 +1879,6 @@ func (m *awsRestjson1_deserializeOpDeleteWebhook) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2036,10 +1999,8 @@ func (m *awsRestjson1_deserializeOpGenerateAccessLogs)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2079,7 +2040,6 @@ func (m *awsRestjson1_deserializeOpGenerateAccessLogs)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2201,10 +2161,8 @@ func (m *awsRestjson1_deserializeOpGetApp) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2244,7 +2202,6 @@ func (m *awsRestjson1_deserializeOpGetApp) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2362,10 +2319,8 @@ func (m *awsRestjson1_deserializeOpGetArtifactUrl) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2405,7 +2360,6 @@ func (m *awsRestjson1_deserializeOpGetArtifactUrl) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2539,10 +2493,8 @@ func (m *awsRestjson1_deserializeOpGetBackendEnvironme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2582,7 +2534,6 @@ func (m *awsRestjson1_deserializeOpGetBackendEnvironme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2700,10 +2651,8 @@ func (m *awsRestjson1_deserializeOpGetBranch) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2743,7 +2692,6 @@ func (m *awsRestjson1_deserializeOpGetBranch) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2861,10 +2809,8 @@ func (m *awsRestjson1_deserializeOpGetDomainAssociatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2904,7 +2850,6 @@ func (m *awsRestjson1_deserializeOpGetDomainAssociatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3022,10 +2967,8 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3065,7 +3008,6 @@ func (m *awsRestjson1_deserializeOpGetJob) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3186,10 +3128,8 @@ func (m *awsRestjson1_deserializeOpGetWebhook) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3229,7 +3169,6 @@ func (m *awsRestjson1_deserializeOpGetWebhook) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3350,10 +3289,8 @@ func (m *awsRestjson1_deserializeOpListApps) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3393,7 +3330,6 @@ func (m *awsRestjson1_deserializeOpListApps) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3517,10 +3453,8 @@ func (m *awsRestjson1_deserializeOpListArtifacts) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3560,7 +3494,6 @@ func (m *awsRestjson1_deserializeOpListArtifacts) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3687,10 +3620,8 @@ func (m *awsRestjson1_deserializeOpListBackendEnvironm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3730,7 +3661,6 @@ func (m *awsRestjson1_deserializeOpListBackendEnvironm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3854,10 +3784,8 @@ func (m *awsRestjson1_deserializeOpListBranches) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3897,7 +3825,6 @@ func (m *awsRestjson1_deserializeOpListBranches) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4021,10 +3948,8 @@ func (m *awsRestjson1_deserializeOpListDomainAssociati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4064,7 +3989,6 @@ func (m *awsRestjson1_deserializeOpListDomainAssociati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4188,10 +4112,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4231,7 +4153,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4358,10 +4279,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4401,7 +4320,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4516,10 +4434,8 @@ func (m *awsRestjson1_deserializeOpListWebhooks) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4559,7 +4475,6 @@ func (m *awsRestjson1_deserializeOpListWebhooks) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4686,10 +4601,8 @@ func (m *awsRestjson1_deserializeOpStartDeployment) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4729,7 +4642,6 @@ func (m *awsRestjson1_deserializeOpStartDeployment) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4850,10 +4762,8 @@ func (m *awsRestjson1_deserializeOpStartJob) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4893,7 +4803,6 @@ func (m *awsRestjson1_deserializeOpStartJob) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5014,10 +4923,8 @@ func (m *awsRestjson1_deserializeOpStopJob) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5057,7 +4964,6 @@ func (m *awsRestjson1_deserializeOpStopJob) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5178,10 +5084,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5193,7 +5097,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5272,10 +5175,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5287,7 +5188,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5366,10 +5266,8 @@ func (m *awsRestjson1_deserializeOpUpdateApp) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5409,7 +5307,6 @@ func (m *awsRestjson1_deserializeOpUpdateApp) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5527,10 +5424,8 @@ func (m *awsRestjson1_deserializeOpUpdateBranch) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5570,7 +5465,6 @@ func (m *awsRestjson1_deserializeOpUpdateBranch) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5691,10 +5585,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5734,7 +5626,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5855,10 +5746,8 @@ func (m *awsRestjson1_deserializeOpUpdateWebhook) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5898,7 +5787,6 @@ func (m *awsRestjson1_deserializeOpUpdateWebhook) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
