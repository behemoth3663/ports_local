--- vendor/github.com/aws/aws-sdk-go-v2/service/mpa/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mpa/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpCancelSession) Hand
 	output := &CancelSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -148,10 +144,8 @@ func (m *awsRestjson1_deserializeOpCreateApprovalTeam)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -191,7 +185,6 @@ func (m *awsRestjson1_deserializeOpCreateApprovalTeam)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -350,10 +343,8 @@ func (m *awsRestjson1_deserializeOpCreateIdentitySourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -393,7 +384,6 @@ func (m *awsRestjson1_deserializeOpCreateIdentitySourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -540,10 +530,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdentitySourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -561,7 +549,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdentitySourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -646,10 +633,8 @@ func (m *awsRestjson1_deserializeOpDeleteInactiveAppro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -661,7 +646,6 @@ func (m *awsRestjson1_deserializeOpDeleteInactiveAppro
 	output := &DeleteInactiveApprovalTeamVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -749,10 +733,8 @@ func (m *awsRestjson1_deserializeOpGetApprovalTeam) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -792,7 +774,6 @@ func (m *awsRestjson1_deserializeOpGetApprovalTeam) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1039,10 +1020,8 @@ func (m *awsRestjson1_deserializeOpGetIdentitySource) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1082,7 +1061,6 @@ func (m *awsRestjson1_deserializeOpGetIdentitySource) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1261,10 +1239,8 @@ func (m *awsRestjson1_deserializeOpGetPolicyVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1304,7 +1280,6 @@ func (m *awsRestjson1_deserializeOpGetPolicyVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1425,10 +1400,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1468,7 +1441,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1629,10 +1601,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1672,7 +1642,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1999,10 +1968,8 @@ func (m *awsRestjson1_deserializeOpListApprovalTeams) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2042,7 +2009,6 @@ func (m *awsRestjson1_deserializeOpListApprovalTeams) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2169,10 +2135,8 @@ func (m *awsRestjson1_deserializeOpListIdentitySources
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2212,7 +2176,6 @@ func (m *awsRestjson1_deserializeOpListIdentitySources
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2339,10 +2302,8 @@ func (m *awsRestjson1_deserializeOpListPolicies) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2382,7 +2343,6 @@ func (m *awsRestjson1_deserializeOpListPolicies) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2509,10 +2469,8 @@ func (m *awsRestjson1_deserializeOpListPolicyVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2552,7 +2510,6 @@ func (m *awsRestjson1_deserializeOpListPolicyVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2682,10 +2639,8 @@ func (m *awsRestjson1_deserializeOpListResourcePolicie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2725,7 +2680,6 @@ func (m *awsRestjson1_deserializeOpListResourcePolicie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2855,10 +2809,8 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2898,7 +2850,6 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3028,10 +2979,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3071,7 +3020,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3192,10 +3140,8 @@ func (m *awsRestjson1_deserializeOpStartActiveApproval
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3235,7 +3181,6 @@ func (m *awsRestjson1_deserializeOpStartActiveApproval
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3380,10 +3325,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3395,7 +3338,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3483,10 +3425,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3498,7 +3438,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3583,10 +3522,8 @@ func (m *awsRestjson1_deserializeOpUpdateApprovalTeam)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3626,7 +3563,6 @@ func (m *awsRestjson1_deserializeOpUpdateApprovalTeam)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
