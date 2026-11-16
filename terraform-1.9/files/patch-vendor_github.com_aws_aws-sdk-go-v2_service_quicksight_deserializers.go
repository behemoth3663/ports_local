--- vendor/github.com/aws/aws-sdk-go-v2/service/quicksight/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/quicksight/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateTopicRev
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateTopicRev
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -255,10 +251,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteTopicRev
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -303,7 +297,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteTopicRev
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -468,10 +461,8 @@ func (m *awsRestjson1_deserializeOpCancelIngestion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -516,7 +507,6 @@ func (m *awsRestjson1_deserializeOpCancelIngestion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -671,10 +661,8 @@ func (m *awsRestjson1_deserializeOpCreateAccountCustom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -719,7 +707,6 @@ func (m *awsRestjson1_deserializeOpCreateAccountCustom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -894,10 +881,8 @@ func (m *awsRestjson1_deserializeOpCreateAccountSubscr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -942,7 +927,6 @@ func (m *awsRestjson1_deserializeOpCreateAccountSubscr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1093,10 +1077,8 @@ func (m *awsRestjson1_deserializeOpCreateAnalysis) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1141,7 +1123,6 @@ func (m *awsRestjson1_deserializeOpCreateAnalysis) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1311,10 +1292,8 @@ func (m *awsRestjson1_deserializeOpCreateBrand) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1354,7 +1333,6 @@ func (m *awsRestjson1_deserializeOpCreateBrand) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1492,10 +1470,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1535,7 +1511,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1697,10 +1672,8 @@ func (m *awsRestjson1_deserializeOpCreateDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1745,7 +1718,6 @@ func (m *awsRestjson1_deserializeOpCreateDashboard) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1924,10 +1896,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1972,7 +1942,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2154,10 +2123,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2202,7 +2169,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2375,10 +2341,8 @@ func (m *awsRestjson1_deserializeOpCreateFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2423,7 +2387,6 @@ func (m *awsRestjson1_deserializeOpCreateFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2587,10 +2550,8 @@ func (m *awsRestjson1_deserializeOpCreateFolderMembers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2630,7 +2591,6 @@ func (m *awsRestjson1_deserializeOpCreateFolderMembers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2782,10 +2742,8 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2830,7 +2788,6 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2981,10 +2938,8 @@ func (m *awsRestjson1_deserializeOpCreateGroupMembersh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3029,7 +2984,6 @@ func (m *awsRestjson1_deserializeOpCreateGroupMembersh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3174,10 +3128,8 @@ func (m *awsRestjson1_deserializeOpCreateIAMPolicyAssi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3222,7 +3174,6 @@ func (m *awsRestjson1_deserializeOpCreateIAMPolicyAssi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3403,10 +3354,8 @@ func (m *awsRestjson1_deserializeOpCreateIngestion) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3451,7 +3400,6 @@ func (m *awsRestjson1_deserializeOpCreateIngestion) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3618,10 +3566,8 @@ func (m *awsRestjson1_deserializeOpCreateNamespace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3666,7 +3612,6 @@ func (m *awsRestjson1_deserializeOpCreateNamespace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3860,10 +3805,8 @@ func (m *awsRestjson1_deserializeOpCreateRefreshSchedu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3908,7 +3851,6 @@ func (m *awsRestjson1_deserializeOpCreateRefreshSchedu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4069,10 +4011,8 @@ func (m *awsRestjson1_deserializeOpCreateRoleMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4117,7 +4057,6 @@ func (m *awsRestjson1_deserializeOpCreateRoleMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4257,10 +4196,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4305,7 +4242,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4487,10 +4423,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplateAlias
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4535,7 +4469,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplateAlias
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4680,10 +4613,8 @@ func (m *awsRestjson1_deserializeOpCreateTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4728,7 +4659,6 @@ func (m *awsRestjson1_deserializeOpCreateTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4907,10 +4837,8 @@ func (m *awsRestjson1_deserializeOpCreateThemeAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4955,7 +4883,6 @@ func (m *awsRestjson1_deserializeOpCreateThemeAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5103,10 +5030,8 @@ func (m *awsRestjson1_deserializeOpCreateTopic) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5151,7 +5076,6 @@ func (m *awsRestjson1_deserializeOpCreateTopic) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5321,10 +5245,8 @@ func (m *awsRestjson1_deserializeOpCreateTopicRefreshS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5369,7 +5291,6 @@ func (m *awsRestjson1_deserializeOpCreateTopicRefreshS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5539,10 +5460,8 @@ func (m *awsRestjson1_deserializeOpCreateVPCConnection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5587,7 +5506,6 @@ func (m *awsRestjson1_deserializeOpCreateVPCConnection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5766,10 +5684,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccountCustom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5814,7 +5730,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccountCustom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5960,10 +5875,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccountSubscr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6008,7 +5921,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccountSubscr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6148,10 +6060,8 @@ func (m *awsRestjson1_deserializeOpDeleteAnalysis) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6196,7 +6106,6 @@ func (m *awsRestjson1_deserializeOpDeleteAnalysis) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6367,10 +6276,8 @@ func (m *awsRestjson1_deserializeOpDeleteBrand) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6410,7 +6317,6 @@ func (m *awsRestjson1_deserializeOpDeleteBrand) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6538,10 +6444,8 @@ func (m *awsRestjson1_deserializeOpDeleteBrandAssignme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6581,7 +6485,6 @@ func (m *awsRestjson1_deserializeOpDeleteBrandAssignme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6709,10 +6612,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6752,7 +6653,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6911,10 +6811,8 @@ func (m *awsRestjson1_deserializeOpDeleteDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6959,7 +6857,6 @@ func (m *awsRestjson1_deserializeOpDeleteDashboard) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7114,10 +7011,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7162,7 +7057,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7314,10 +7208,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSetRefres
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7362,7 +7254,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSetRefres
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7502,10 +7393,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7550,7 +7439,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7702,10 +7590,8 @@ func (m *awsRestjson1_deserializeOpDeleteDefaultQBusin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7750,7 +7636,6 @@ func (m *awsRestjson1_deserializeOpDeleteDefaultQBusin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7884,10 +7769,8 @@ func (m *awsRestjson1_deserializeOpDeleteFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7932,7 +7815,6 @@ func (m *awsRestjson1_deserializeOpDeleteFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8093,10 +7975,8 @@ func (m *awsRestjson1_deserializeOpDeleteFolderMembers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8136,7 +8016,6 @@ func (m *awsRestjson1_deserializeOpDeleteFolderMembers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8277,10 +8156,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8325,7 +8202,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8465,10 +8341,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroupMembersh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8513,7 +8387,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroupMembersh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8653,10 +8526,8 @@ func (m *awsRestjson1_deserializeOpDeleteIAMPolicyAssi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8701,7 +8572,6 @@ func (m *awsRestjson1_deserializeOpDeleteIAMPolicyAssi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8850,10 +8720,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdentityPropa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8898,7 +8766,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdentityPropa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9032,10 +8899,8 @@ func (m *awsRestjson1_deserializeOpDeleteNamespace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9080,7 +8945,6 @@ func (m *awsRestjson1_deserializeOpDeleteNamespace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9220,10 +9084,8 @@ func (m *awsRestjson1_deserializeOpDeleteRefreshSchedu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9268,7 +9130,6 @@ func (m *awsRestjson1_deserializeOpDeleteRefreshSchedu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9423,10 +9284,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoleCustomPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9466,7 +9325,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoleCustomPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9610,10 +9468,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoleMembershi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9658,7 +9514,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoleMembershi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9798,10 +9653,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9846,7 +9699,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10004,10 +9856,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateAlias
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10052,7 +9902,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplateAlias
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10213,10 +10062,8 @@ func (m *awsRestjson1_deserializeOpDeleteTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10261,7 +10108,6 @@ func (m *awsRestjson1_deserializeOpDeleteTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10419,10 +10265,8 @@ func (m *awsRestjson1_deserializeOpDeleteThemeAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10467,7 +10311,6 @@ func (m *awsRestjson1_deserializeOpDeleteThemeAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10631,10 +10474,8 @@ func (m *awsRestjson1_deserializeOpDeleteTopic) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10679,7 +10520,6 @@ func (m *awsRestjson1_deserializeOpDeleteTopic) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10834,10 +10674,8 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRefreshS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10882,7 +10720,6 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRefreshS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11052,10 +10889,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11100,7 +10935,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11240,10 +11074,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserByPrincip
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11288,7 +11120,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserByPrincip
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11428,10 +11259,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserCustomPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11476,7 +11305,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserCustomPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11619,10 +11447,8 @@ func (m *awsRestjson1_deserializeOpDeleteVPCConnection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11667,7 +11493,6 @@ func (m *awsRestjson1_deserializeOpDeleteVPCConnection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11843,10 +11668,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountCust
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11891,7 +11714,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountCust
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12060,10 +11882,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountSett
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12108,7 +11928,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountSett
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12250,10 +12069,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountSubs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12298,7 +12115,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountSubs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12440,10 +12256,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysis) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12488,7 +12302,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysis) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12630,10 +12443,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysisDef
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12678,7 +12489,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysisDef
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12867,10 +12677,8 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysisPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12915,7 +12723,6 @@ func (m *awsRestjson1_deserializeOpDescribeAnalysisPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13072,10 +12879,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetBundle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13120,7 +12925,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetBundle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13388,10 +13192,8 @@ func (m *awsRestjson1_deserializeOpDescribeAssetBundle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13436,7 +13238,6 @@ func (m *awsRestjson1_deserializeOpDescribeAssetBundle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13665,10 +13466,8 @@ func (m *awsRestjson1_deserializeOpDescribeBrand) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13708,7 +13507,6 @@ func (m *awsRestjson1_deserializeOpDescribeBrand) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13846,10 +13644,8 @@ func (m *awsRestjson1_deserializeOpDescribeBrandAssign
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13889,7 +13685,6 @@ func (m *awsRestjson1_deserializeOpDescribeBrandAssign
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14026,10 +13821,8 @@ func (m *awsRestjson1_deserializeOpDescribeBrandPublis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14069,7 +13862,6 @@ func (m *awsRestjson1_deserializeOpDescribeBrandPublis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14207,10 +13999,8 @@ func (m *awsRestjson1_deserializeOpDescribeCustomPermi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14250,7 +14040,6 @@ func (m *awsRestjson1_deserializeOpDescribeCustomPermi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14399,10 +14188,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboard) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14447,7 +14234,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboard) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14589,10 +14375,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14637,7 +14421,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14831,10 +14614,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardPe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14879,7 +14660,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardPe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15041,10 +14821,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardSn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15084,7 +14862,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardSn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15309,10 +15086,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardSn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15357,7 +15132,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardSn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15557,10 +15331,8 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardsQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15605,7 +15377,6 @@ func (m *awsRestjson1_deserializeOpDescribeDashboardsQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15751,10 +15522,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataSet) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15799,7 +15568,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataSet) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15938,10 +15706,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataSetPerm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15986,7 +15752,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataSetPerm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16143,10 +15908,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataSetRefr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16191,7 +15954,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataSetRefr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16336,10 +16098,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataSource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16384,7 +16144,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataSource)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16523,10 +16282,8 @@ func (m *awsRestjson1_deserializeOpDescribeDataSourceP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16571,7 +16328,6 @@ func (m *awsRestjson1_deserializeOpDescribeDataSourceP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16728,10 +16484,8 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultQBus
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16776,7 +16530,6 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultQBus
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16919,10 +16672,8 @@ func (m *awsRestjson1_deserializeOpDescribeFolder) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16967,7 +16718,6 @@ func (m *awsRestjson1_deserializeOpDescribeFolder) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17109,10 +16859,8 @@ func (m *awsRestjson1_deserializeOpDescribeFolderPermi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17157,7 +16905,6 @@ func (m *awsRestjson1_deserializeOpDescribeFolderPermi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17329,10 +17076,8 @@ func (m *awsRestjson1_deserializeOpDescribeFolderResol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17377,7 +17122,6 @@ func (m *awsRestjson1_deserializeOpDescribeFolderResol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17549,10 +17293,8 @@ func (m *awsRestjson1_deserializeOpDescribeGroup) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17597,7 +17339,6 @@ func (m *awsRestjson1_deserializeOpDescribeGroup) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17742,10 +17483,8 @@ func (m *awsRestjson1_deserializeOpDescribeGroupMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17790,7 +17529,6 @@ func (m *awsRestjson1_deserializeOpDescribeGroupMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17935,10 +17673,8 @@ func (m *awsRestjson1_deserializeOpDescribeIAMPolicyAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17983,7 +17719,6 @@ func (m *awsRestjson1_deserializeOpDescribeIAMPolicyAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18125,10 +17860,8 @@ func (m *awsRestjson1_deserializeOpDescribeIngestion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18173,7 +17906,6 @@ func (m *awsRestjson1_deserializeOpDescribeIngestion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18315,10 +18047,8 @@ func (m *awsRestjson1_deserializeOpDescribeIpRestricti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18363,7 +18093,6 @@ func (m *awsRestjson1_deserializeOpDescribeIpRestricti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18530,10 +18259,8 @@ func (m *awsRestjson1_deserializeOpDescribeKeyRegistra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18573,7 +18300,6 @@ func (m *awsRestjson1_deserializeOpDescribeKeyRegistra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18722,10 +18448,8 @@ func (m *awsRestjson1_deserializeOpDescribeNamespace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18770,7 +18494,6 @@ func (m *awsRestjson1_deserializeOpDescribeNamespace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18912,10 +18635,8 @@ func (m *awsRestjson1_deserializeOpDescribeQPersonaliz
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18960,7 +18681,6 @@ func (m *awsRestjson1_deserializeOpDescribeQPersonaliz
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19106,10 +18826,8 @@ func (m *awsRestjson1_deserializeOpDescribeQuickSightQ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19154,7 +18872,6 @@ func (m *awsRestjson1_deserializeOpDescribeQuickSightQ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19300,10 +19017,8 @@ func (m *awsRestjson1_deserializeOpDescribeRefreshSche
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19348,7 +19063,6 @@ func (m *awsRestjson1_deserializeOpDescribeRefreshSche
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19499,10 +19213,8 @@ func (m *awsRestjson1_deserializeOpDescribeRoleCustomP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19542,7 +19254,6 @@ func (m *awsRestjson1_deserializeOpDescribeRoleCustomP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19695,10 +19406,8 @@ func (m *awsRestjson1_deserializeOpDescribeTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19743,7 +19452,6 @@ func (m *awsRestjson1_deserializeOpDescribeTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19891,10 +19599,8 @@ func (m *awsRestjson1_deserializeOpDescribeTemplateAli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19939,7 +19645,6 @@ func (m *awsRestjson1_deserializeOpDescribeTemplateAli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20075,10 +19780,8 @@ func (m *awsRestjson1_deserializeOpDescribeTemplateDef
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20123,7 +19826,6 @@ func (m *awsRestjson1_deserializeOpDescribeTemplateDef
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20312,10 +20014,8 @@ func (m *awsRestjson1_deserializeOpDescribeTemplatePer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20360,7 +20060,6 @@ func (m *awsRestjson1_deserializeOpDescribeTemplatePer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20520,10 +20219,8 @@ func (m *awsRestjson1_deserializeOpDescribeTheme) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20568,7 +20265,6 @@ func (m *awsRestjson1_deserializeOpDescribeTheme) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20713,10 +20409,8 @@ func (m *awsRestjson1_deserializeOpDescribeThemeAlias)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20761,7 +20455,6 @@ func (m *awsRestjson1_deserializeOpDescribeThemeAlias)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20903,10 +20596,8 @@ func (m *awsRestjson1_deserializeOpDescribeThemePermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20951,7 +20642,6 @@ func (m *awsRestjson1_deserializeOpDescribeThemePermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21111,10 +20801,8 @@ func (m *awsRestjson1_deserializeOpDescribeTopic) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21159,7 +20847,6 @@ func (m *awsRestjson1_deserializeOpDescribeTopic) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21321,10 +21008,8 @@ func (m *awsRestjson1_deserializeOpDescribeTopicPermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21369,7 +21054,6 @@ func (m *awsRestjson1_deserializeOpDescribeTopicPermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21526,10 +21210,8 @@ func (m *awsRestjson1_deserializeOpDescribeTopicRefres
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21574,7 +21256,6 @@ func (m *awsRestjson1_deserializeOpDescribeTopicRefres
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21713,10 +21394,8 @@ func (m *awsRestjson1_deserializeOpDescribeTopicRefres
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21761,7 +21440,6 @@ func (m *awsRestjson1_deserializeOpDescribeTopicRefres
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21936,10 +21614,8 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21984,7 +21660,6 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22129,10 +21804,8 @@ func (m *awsRestjson1_deserializeOpDescribeVPCConnecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22172,7 +21845,6 @@ func (m *awsRestjson1_deserializeOpDescribeVPCConnecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22318,10 +21990,8 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22366,7 +22036,6 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22527,10 +22196,8 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22575,7 +22242,6 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22730,10 +22396,8 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22778,7 +22442,6 @@ func (m *awsRestjson1_deserializeOpGenerateEmbedUrlFor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22933,10 +22596,8 @@ func (m *awsRestjson1_deserializeOpGetDashboardEmbedUr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22981,7 +22642,6 @@ func (m *awsRestjson1_deserializeOpGetDashboardEmbedUr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23145,10 +22805,8 @@ func (m *awsRestjson1_deserializeOpGetSessionEmbedUrl)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23193,7 +22851,6 @@ func (m *awsRestjson1_deserializeOpGetSessionEmbedUrl)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23348,10 +23005,8 @@ func (m *awsRestjson1_deserializeOpListAnalyses) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23396,7 +23051,6 @@ func (m *awsRestjson1_deserializeOpListAnalyses) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23541,10 +23195,8 @@ func (m *awsRestjson1_deserializeOpListAssetBundleExpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23589,7 +23241,6 @@ func (m *awsRestjson1_deserializeOpListAssetBundleExpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23737,10 +23388,8 @@ func (m *awsRestjson1_deserializeOpListAssetBundleImpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23785,7 +23434,6 @@ func (m *awsRestjson1_deserializeOpListAssetBundleImpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23933,10 +23581,8 @@ func (m *awsRestjson1_deserializeOpListBrands) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23976,7 +23622,6 @@ func (m *awsRestjson1_deserializeOpListBrands) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24103,10 +23748,8 @@ func (m *awsRestjson1_deserializeOpListCustomPermissio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24151,7 +23794,6 @@ func (m *awsRestjson1_deserializeOpListCustomPermissio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24305,10 +23947,8 @@ func (m *awsRestjson1_deserializeOpListDashboards) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24353,7 +23993,6 @@ func (m *awsRestjson1_deserializeOpListDashboards) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24498,10 +24137,8 @@ func (m *awsRestjson1_deserializeOpListDashboardVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24546,7 +24183,6 @@ func (m *awsRestjson1_deserializeOpListDashboardVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24697,10 +24333,8 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24745,7 +24379,6 @@ func (m *awsRestjson1_deserializeOpListDataSets) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24893,10 +24526,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24941,7 +24572,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25089,10 +24719,8 @@ func (m *awsRestjson1_deserializeOpListFolderMembers) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25137,7 +24765,6 @@ func (m *awsRestjson1_deserializeOpListFolderMembers) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25291,10 +24918,8 @@ func (m *awsRestjson1_deserializeOpListFolders) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25339,7 +24964,6 @@ func (m *awsRestjson1_deserializeOpListFolders) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25493,10 +25117,8 @@ func (m *awsRestjson1_deserializeOpListFoldersForResou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25541,7 +25163,6 @@ func (m *awsRestjson1_deserializeOpListFoldersForResou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25695,10 +25316,8 @@ func (m *awsRestjson1_deserializeOpListGroupMembership
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25743,7 +25362,6 @@ func (m *awsRestjson1_deserializeOpListGroupMembership
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25900,10 +25518,8 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25948,7 +25564,6 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26105,10 +25720,8 @@ func (m *awsRestjson1_deserializeOpListIAMPolicyAssign
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26153,7 +25766,6 @@ func (m *awsRestjson1_deserializeOpListIAMPolicyAssign
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26304,10 +25916,8 @@ func (m *awsRestjson1_deserializeOpListIAMPolicyAssign
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26352,7 +25962,6 @@ func (m *awsRestjson1_deserializeOpListIAMPolicyAssign
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26506,10 +26115,8 @@ func (m *awsRestjson1_deserializeOpListIdentityPropaga
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26554,7 +26161,6 @@ func (m *awsRestjson1_deserializeOpListIdentityPropaga
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26702,10 +26308,8 @@ func (m *awsRestjson1_deserializeOpListIngestions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26750,7 +26354,6 @@ func (m *awsRestjson1_deserializeOpListIngestions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26904,10 +26507,8 @@ func (m *awsRestjson1_deserializeOpListNamespaces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26952,7 +26553,6 @@ func (m *awsRestjson1_deserializeOpListNamespaces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27109,10 +26709,8 @@ func (m *awsRestjson1_deserializeOpListRefreshSchedule
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27157,7 +26755,6 @@ func (m *awsRestjson1_deserializeOpListRefreshSchedule
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27299,10 +26896,8 @@ func (m *awsRestjson1_deserializeOpListRoleMemberships
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27347,7 +26942,6 @@ func (m *awsRestjson1_deserializeOpListRoleMemberships
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27507,10 +27101,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27555,7 +27147,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27694,10 +27285,8 @@ func (m *awsRestjson1_deserializeOpListTemplateAliases
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27742,7 +27331,6 @@ func (m *awsRestjson1_deserializeOpListTemplateAliases
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27890,10 +27478,8 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27938,7 +27524,6 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28089,10 +27674,8 @@ func (m *awsRestjson1_deserializeOpListTemplateVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28137,7 +27720,6 @@ func (m *awsRestjson1_deserializeOpListTemplateVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28288,10 +27870,8 @@ func (m *awsRestjson1_deserializeOpListThemeAliases) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28336,7 +27916,6 @@ func (m *awsRestjson1_deserializeOpListThemeAliases) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28490,10 +28069,8 @@ func (m *awsRestjson1_deserializeOpListThemes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28538,7 +28115,6 @@ func (m *awsRestjson1_deserializeOpListThemes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28692,10 +28268,8 @@ func (m *awsRestjson1_deserializeOpListThemeVersions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28740,7 +28314,6 @@ func (m *awsRestjson1_deserializeOpListThemeVersions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28894,10 +28467,8 @@ func (m *awsRestjson1_deserializeOpListTopicRefreshSch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28942,7 +28513,6 @@ func (m *awsRestjson1_deserializeOpListTopicRefreshSch
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29108,10 +28678,8 @@ func (m *awsRestjson1_deserializeOpListTopicReviewedAn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29156,7 +28724,6 @@ func (m *awsRestjson1_deserializeOpListTopicReviewedAn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29313,10 +28880,8 @@ func (m *awsRestjson1_deserializeOpListTopics) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29361,7 +28926,6 @@ func (m *awsRestjson1_deserializeOpListTopics) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29509,10 +29073,8 @@ func (m *awsRestjson1_deserializeOpListUserGroups) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29557,7 +29119,6 @@ func (m *awsRestjson1_deserializeOpListUserGroups) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29711,10 +29272,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29759,7 +29318,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29916,10 +29474,8 @@ func (m *awsRestjson1_deserializeOpListVPCConnections)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29964,7 +29520,6 @@ func (m *awsRestjson1_deserializeOpListVPCConnections)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30115,10 +29670,8 @@ func (m *awsRestjson1_deserializeOpPredictQAResults) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30163,7 +29716,6 @@ func (m *awsRestjson1_deserializeOpPredictQAResults) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30304,10 +29856,8 @@ func (m *awsRestjson1_deserializeOpPutDataSetRefreshPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30352,7 +29902,6 @@ func (m *awsRestjson1_deserializeOpPutDataSetRefreshPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30495,10 +30044,8 @@ func (m *awsRestjson1_deserializeOpRegisterUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30543,7 +30090,6 @@ func (m *awsRestjson1_deserializeOpRegisterUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30703,10 +30249,8 @@ func (m *awsRestjson1_deserializeOpRestoreAnalysis) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30751,7 +30295,6 @@ func (m *awsRestjson1_deserializeOpRestoreAnalysis) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30917,10 +30460,8 @@ func (m *awsRestjson1_deserializeOpSearchAnalyses) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30965,7 +30506,6 @@ func (m *awsRestjson1_deserializeOpSearchAnalyses) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31116,10 +30656,8 @@ func (m *awsRestjson1_deserializeOpSearchDashboards) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31164,7 +30702,6 @@ func (m *awsRestjson1_deserializeOpSearchDashboards) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31315,10 +30852,8 @@ func (m *awsRestjson1_deserializeOpSearchDataSets) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31363,7 +30898,6 @@ func (m *awsRestjson1_deserializeOpSearchDataSets) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31514,10 +31048,8 @@ func (m *awsRestjson1_deserializeOpSearchDataSources) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31562,7 +31094,6 @@ func (m *awsRestjson1_deserializeOpSearchDataSources) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31713,10 +31244,8 @@ func (m *awsRestjson1_deserializeOpSearchFolders) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31761,7 +31290,6 @@ func (m *awsRestjson1_deserializeOpSearchFolders) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31918,10 +31446,8 @@ func (m *awsRestjson1_deserializeOpSearchGroups) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31966,7 +31492,6 @@ func (m *awsRestjson1_deserializeOpSearchGroups) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32123,10 +31648,8 @@ func (m *awsRestjson1_deserializeOpSearchTopics) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32171,7 +31694,6 @@ func (m *awsRestjson1_deserializeOpSearchTopics) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32322,10 +31844,8 @@ func (m *awsRestjson1_deserializeOpStartAssetBundleExp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32370,7 +31890,6 @@ func (m *awsRestjson1_deserializeOpStartAssetBundleExp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32528,10 +32047,8 @@ func (m *awsRestjson1_deserializeOpStartAssetBundleImp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32576,7 +32093,6 @@ func (m *awsRestjson1_deserializeOpStartAssetBundleImp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32734,10 +32250,8 @@ func (m *awsRestjson1_deserializeOpStartDashboardSnaps
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32782,7 +32296,6 @@ func (m *awsRestjson1_deserializeOpStartDashboardSnaps
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32946,10 +32459,8 @@ func (m *awsRestjson1_deserializeOpStartDashboardSnaps
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32994,7 +32505,6 @@ func (m *awsRestjson1_deserializeOpStartDashboardSnaps
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33134,10 +32644,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33182,7 +32690,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33319,10 +32826,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33367,7 +32872,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33501,10 +33005,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountCustom
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33549,7 +33051,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountCustom
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33721,10 +33222,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33769,7 +33268,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33906,10 +33404,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysis) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33954,7 +33450,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysis) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34121,10 +33616,8 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysisPermi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34169,7 +33662,6 @@ func (m *awsRestjson1_deserializeOpUpdateAnalysisPermi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34332,10 +33824,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationWi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34380,7 +33870,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplicationWi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34520,10 +34009,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrand) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34563,7 +34050,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrand) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34701,10 +34187,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrandAssignme
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34744,7 +34228,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrandAssignme
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34881,10 +34364,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrandPublishe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34924,7 +34405,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrandPublishe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35061,10 +34541,8 @@ func (m *awsRestjson1_deserializeOpUpdateCustomPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35104,7 +34582,6 @@ func (m *awsRestjson1_deserializeOpUpdateCustomPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35260,10 +34737,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboard) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35303,7 +34778,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboard) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35483,10 +34957,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardLink
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35531,7 +35003,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardLink
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35685,10 +35156,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardPerm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35733,7 +35202,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardPerm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35901,10 +35369,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardPubl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35949,7 +35415,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardPubl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36104,10 +35569,8 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardsQAC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36152,7 +35615,6 @@ func (m *awsRestjson1_deserializeOpUpdateDashboardsQAC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36298,10 +35760,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36346,7 +35806,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36525,10 +35984,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSetPermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36573,7 +36030,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSetPermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36728,10 +36184,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36776,7 +36230,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36943,10 +36396,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSourcePer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36991,7 +36442,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSourcePer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37146,10 +36596,8 @@ func (m *awsRestjson1_deserializeOpUpdateDefaultQBusin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37194,7 +36642,6 @@ func (m *awsRestjson1_deserializeOpUpdateDefaultQBusin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37331,10 +36778,8 @@ func (m *awsRestjson1_deserializeOpUpdateFolder) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37379,7 +36824,6 @@ func (m *awsRestjson1_deserializeOpUpdateFolder) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37540,10 +36984,8 @@ func (m *awsRestjson1_deserializeOpUpdateFolderPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37583,7 +37025,6 @@ func (m *awsRestjson1_deserializeOpUpdateFolderPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37750,10 +37191,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37798,7 +37237,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37943,10 +37381,8 @@ func (m *awsRestjson1_deserializeOpUpdateIAMPolicyAssi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37991,7 +37427,6 @@ func (m *awsRestjson1_deserializeOpUpdateIAMPolicyAssi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38172,10 +37607,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdentityPropa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38220,7 +37653,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdentityPropa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38354,10 +37786,8 @@ func (m *awsRestjson1_deserializeOpUpdateIpRestriction
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38402,7 +37832,6 @@ func (m *awsRestjson1_deserializeOpUpdateIpRestriction
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38548,10 +37977,8 @@ func (m *awsRestjson1_deserializeOpUpdateKeyRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38591,7 +38018,6 @@ func (m *awsRestjson1_deserializeOpUpdateKeyRegistrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38723,10 +38149,8 @@ func (m *awsRestjson1_deserializeOpUpdatePublicSharing
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38771,7 +38195,6 @@ func (m *awsRestjson1_deserializeOpUpdatePublicSharing
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38908,10 +38331,8 @@ func (m *awsRestjson1_deserializeOpUpdateQPersonalizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38956,7 +38377,6 @@ func (m *awsRestjson1_deserializeOpUpdateQPersonalizat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39105,10 +38525,8 @@ func (m *awsRestjson1_deserializeOpUpdateQuickSightQSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39153,7 +38571,6 @@ func (m *awsRestjson1_deserializeOpUpdateQuickSightQSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39299,10 +38716,8 @@ func (m *awsRestjson1_deserializeOpUpdateRefreshSchedu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39347,7 +38762,6 @@ func (m *awsRestjson1_deserializeOpUpdateRefreshSchedu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39505,10 +38919,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoleCustomPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39548,7 +38960,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoleCustomPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39692,10 +39103,8 @@ func (m *awsRestjson1_deserializeOpUpdateSPICECapacity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39740,7 +39149,6 @@ func (m *awsRestjson1_deserializeOpUpdateSPICECapacity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39874,10 +39282,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39922,7 +39328,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40101,10 +39506,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateAlias
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40149,7 +39552,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplateAlias
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40288,10 +39690,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplatePermi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40336,7 +39736,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplatePermi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40499,10 +39898,8 @@ func (m *awsRestjson1_deserializeOpUpdateTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40547,7 +39944,6 @@ func (m *awsRestjson1_deserializeOpUpdateTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40726,10 +40122,8 @@ func (m *awsRestjson1_deserializeOpUpdateThemeAlias) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40774,7 +40168,6 @@ func (m *awsRestjson1_deserializeOpUpdateThemeAlias) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40919,10 +40312,8 @@ func (m *awsRestjson1_deserializeOpUpdateThemePermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40967,7 +40358,6 @@ func (m *awsRestjson1_deserializeOpUpdateThemePermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41130,10 +40520,8 @@ func (m *awsRestjson1_deserializeOpUpdateTopic) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41178,7 +40566,6 @@ func (m *awsRestjson1_deserializeOpUpdateTopic) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41348,10 +40735,8 @@ func (m *awsRestjson1_deserializeOpUpdateTopicPermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41396,7 +40781,6 @@ func (m *awsRestjson1_deserializeOpUpdateTopicPermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41562,10 +40946,8 @@ func (m *awsRestjson1_deserializeOpUpdateTopicRefreshS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41610,7 +40992,6 @@ func (m *awsRestjson1_deserializeOpUpdateTopicRefreshS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41780,10 +41161,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41828,7 +41207,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41973,10 +41351,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserCustomPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42021,7 +41397,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserCustomPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42164,10 +41539,8 @@ func (m *awsRestjson1_deserializeOpUpdateVPCConnection
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42212,7 +41585,6 @@ func (m *awsRestjson1_deserializeOpUpdateVPCConnection
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
