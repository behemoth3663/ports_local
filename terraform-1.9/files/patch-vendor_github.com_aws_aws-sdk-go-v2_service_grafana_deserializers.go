--- vendor/github.com/aws/aws-sdk-go-v2/service/grafana/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/grafana/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateLicense) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAssociateLicense) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -209,10 +205,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -252,7 +246,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -376,10 +369,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceApiK
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -419,7 +410,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceApiK
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -568,10 +558,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -611,7 +599,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -769,10 +756,8 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -812,7 +797,6 @@ func (m *awsRestjson1_deserializeOpCreateWorkspaceServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -957,10 +941,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1000,7 +982,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1124,10 +1105,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceApiK
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1167,7 +1146,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceApiK
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1304,10 +1282,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1347,7 +1323,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1484,10 +1459,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1527,7 +1500,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorkspaceServ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1673,10 +1645,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspace) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1716,7 +1686,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspace) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1837,10 +1806,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceAu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1880,7 +1847,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceAu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2004,10 +1970,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2047,7 +2011,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorkspaceCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2178,10 +2141,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLicense
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2221,7 +2182,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLicense
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2342,10 +2302,8 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2385,7 +2343,6 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2515,10 +2472,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2558,7 +2513,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2679,10 +2633,8 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2722,7 +2674,6 @@ func (m *awsRestjson1_deserializeOpListVersions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2852,10 +2803,8 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2895,7 +2844,6 @@ func (m *awsRestjson1_deserializeOpListWorkspaces) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3019,10 +2967,8 @@ func (m *awsRestjson1_deserializeOpListWorkspaceServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3062,7 +3008,6 @@ func (m *awsRestjson1_deserializeOpListWorkspaceServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3204,10 +3149,8 @@ func (m *awsRestjson1_deserializeOpListWorkspaceServic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3247,7 +3190,6 @@ func (m *awsRestjson1_deserializeOpListWorkspaceServic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3398,10 +3340,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3413,7 +3353,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3498,10 +3437,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3513,7 +3450,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3598,10 +3534,8 @@ func (m *awsRestjson1_deserializeOpUpdatePermissions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3641,7 +3575,6 @@ func (m *awsRestjson1_deserializeOpUpdatePermissions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3762,10 +3695,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspace) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3805,7 +3736,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspace) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3929,10 +3859,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceAuth
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3972,7 +3900,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceAuth
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4096,10 +4023,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4111,7 +4036,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorkspaceConf
 	output := &UpdateWorkspaceConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
