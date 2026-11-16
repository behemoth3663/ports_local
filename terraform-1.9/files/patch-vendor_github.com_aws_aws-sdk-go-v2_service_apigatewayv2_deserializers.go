--- vendor/github.com/aws/aws-sdk-go-v2/service/apigatewayv2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apigatewayv2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -343,10 +339,8 @@ func (m *awsRestjson1_deserializeOpCreateApiMapping) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -386,7 +380,6 @@ func (m *awsRestjson1_deserializeOpCreateApiMapping) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -535,10 +528,8 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -578,7 +569,6 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -786,10 +776,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -829,7 +817,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1000,10 +987,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1043,7 +1028,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1210,10 +1194,8 @@ func (m *awsRestjson1_deserializeOpCreateIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1253,7 +1235,6 @@ func (m *awsRestjson1_deserializeOpCreateIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1534,10 +1515,8 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1577,7 +1556,6 @@ func (m *awsRestjson1_deserializeOpCreateIntegrationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1736,10 +1714,8 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1779,7 +1755,6 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1937,10 +1912,8 @@ func (m *awsRestjson1_deserializeOpCreateRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1980,7 +1953,6 @@ func (m *awsRestjson1_deserializeOpCreateRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2198,10 +2170,8 @@ func (m *awsRestjson1_deserializeOpCreateRouteResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2241,7 +2211,6 @@ func (m *awsRestjson1_deserializeOpCreateRouteResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2391,10 +2360,8 @@ func (m *awsRestjson1_deserializeOpCreateRoutingRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2434,7 +2401,6 @@ func (m *awsRestjson1_deserializeOpCreateRoutingRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2588,10 +2554,8 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2631,7 +2595,6 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2858,10 +2821,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2901,7 +2862,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3081,10 +3041,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccessLogSett
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3102,7 +3060,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccessLogSett
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3178,10 +3135,8 @@ func (m *awsRestjson1_deserializeOpDeleteApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3199,7 +3154,6 @@ func (m *awsRestjson1_deserializeOpDeleteApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3275,10 +3229,8 @@ func (m *awsRestjson1_deserializeOpDeleteApiMapping) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3296,7 +3248,6 @@ func (m *awsRestjson1_deserializeOpDeleteApiMapping) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3375,10 +3326,8 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3396,7 +3345,6 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3472,10 +3420,8 @@ func (m *awsRestjson1_deserializeOpDeleteCorsConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3493,7 +3439,6 @@ func (m *awsRestjson1_deserializeOpDeleteCorsConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3569,10 +3514,8 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3590,7 +3533,6 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3666,10 +3608,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3687,7 +3627,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3763,10 +3702,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3784,7 +3721,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3860,10 +3796,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3881,7 +3815,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3957,10 +3890,8 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3978,7 +3909,6 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3984,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4075,7 +4003,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4151,10 +4078,8 @@ func (m *awsRestjson1_deserializeOpDeleteRouteRequestP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4172,7 +4097,6 @@ func (m *awsRestjson1_deserializeOpDeleteRouteRequestP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4248,10 +4172,8 @@ func (m *awsRestjson1_deserializeOpDeleteRouteResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4269,7 +4191,6 @@ func (m *awsRestjson1_deserializeOpDeleteRouteResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4345,10 +4266,8 @@ func (m *awsRestjson1_deserializeOpDeleteRouteSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4366,7 +4285,6 @@ func (m *awsRestjson1_deserializeOpDeleteRouteSettings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4442,10 +4360,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingRule) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4463,7 +4379,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoutingRule) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4542,10 +4457,8 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4563,7 +4476,6 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4639,10 +4551,8 @@ func (m *awsRestjson1_deserializeOpDeleteVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4654,7 +4564,6 @@ func (m *awsRestjson1_deserializeOpDeleteVpcLink) Hand
 	output := &DeleteVpcLinkOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4730,10 +4639,8 @@ func (m *awsRestjson1_deserializeOpExportApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4750,7 +4657,6 @@ func (m *awsRestjson1_deserializeOpExportApi) HandleDe
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4851,10 +4757,8 @@ func (m *awsRestjson1_deserializeOpGetApi) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4894,7 +4798,6 @@ func (m *awsRestjson1_deserializeOpGetApi) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5142,10 +5045,8 @@ func (m *awsRestjson1_deserializeOpGetApiMapping) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5185,7 +5086,6 @@ func (m *awsRestjson1_deserializeOpGetApiMapping) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5331,10 +5231,8 @@ func (m *awsRestjson1_deserializeOpGetApiMappings) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5374,7 +5272,6 @@ func (m *awsRestjson1_deserializeOpGetApiMappings) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5498,10 +5395,8 @@ func (m *awsRestjson1_deserializeOpGetApis) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5541,7 +5436,6 @@ func (m *awsRestjson1_deserializeOpGetApis) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5665,10 +5559,8 @@ func (m *awsRestjson1_deserializeOpGetAuthorizer) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5708,7 +5600,6 @@ func (m *awsRestjson1_deserializeOpGetAuthorizer) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5910,10 +5801,8 @@ func (m *awsRestjson1_deserializeOpGetAuthorizers) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5953,7 +5842,6 @@ func (m *awsRestjson1_deserializeOpGetAuthorizers) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6077,10 +5965,8 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6120,7 +6006,6 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6285,10 +6170,8 @@ func (m *awsRestjson1_deserializeOpGetDeployments) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6328,7 +6211,6 @@ func (m *awsRestjson1_deserializeOpGetDeployments) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6452,10 +6334,8 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6495,7 +6375,6 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6653,10 +6532,8 @@ func (m *awsRestjson1_deserializeOpGetDomainNames) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6696,7 +6573,6 @@ func (m *awsRestjson1_deserializeOpGetDomainNames) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6820,10 +6696,8 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6863,7 +6737,6 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7138,10 +7011,8 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7181,7 +7052,6 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7334,10 +7204,8 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7377,7 +7245,6 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7501,10 +7368,8 @@ func (m *awsRestjson1_deserializeOpGetIntegrations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7544,7 +7409,6 @@ func (m *awsRestjson1_deserializeOpGetIntegrations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7668,10 +7532,8 @@ func (m *awsRestjson1_deserializeOpGetModel) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7711,7 +7573,6 @@ func (m *awsRestjson1_deserializeOpGetModel) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7863,10 +7724,8 @@ func (m *awsRestjson1_deserializeOpGetModels) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7906,7 +7765,6 @@ func (m *awsRestjson1_deserializeOpGetModels) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8030,10 +7888,8 @@ func (m *awsRestjson1_deserializeOpGetModelTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8073,7 +7929,6 @@ func (m *awsRestjson1_deserializeOpGetModelTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8189,10 +8044,8 @@ func (m *awsRestjson1_deserializeOpGetRoute) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8232,7 +8085,6 @@ func (m *awsRestjson1_deserializeOpGetRoute) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8444,10 +8296,8 @@ func (m *awsRestjson1_deserializeOpGetRouteResponse) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8487,7 +8337,6 @@ func (m *awsRestjson1_deserializeOpGetRouteResponse) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8631,10 +8480,8 @@ func (m *awsRestjson1_deserializeOpGetRouteResponses) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8674,7 +8521,6 @@ func (m *awsRestjson1_deserializeOpGetRouteResponses) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8798,10 +8644,8 @@ func (m *awsRestjson1_deserializeOpGetRoutes) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8841,7 +8685,6 @@ func (m *awsRestjson1_deserializeOpGetRoutes) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8965,10 +8808,8 @@ func (m *awsRestjson1_deserializeOpGetRoutingRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9008,7 +8849,6 @@ func (m *awsRestjson1_deserializeOpGetRoutingRule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9159,10 +8999,8 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9202,7 +9040,6 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9423,10 +9260,8 @@ func (m *awsRestjson1_deserializeOpGetStages) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9466,7 +9301,6 @@ func (m *awsRestjson1_deserializeOpGetStages) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9590,10 +9424,8 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9633,7 +9465,6 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9751,10 +9582,8 @@ func (m *awsRestjson1_deserializeOpGetVpcLink) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9794,7 +9623,6 @@ func (m *awsRestjson1_deserializeOpGetVpcLink) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9974,10 +9802,8 @@ func (m *awsRestjson1_deserializeOpGetVpcLinks) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10017,7 +9843,6 @@ func (m *awsRestjson1_deserializeOpGetVpcLinks) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10138,10 +9963,8 @@ func (m *awsRestjson1_deserializeOpImportApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10181,7 +10004,6 @@ func (m *awsRestjson1_deserializeOpImportApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10435,10 +10257,8 @@ func (m *awsRestjson1_deserializeOpListRoutingRules) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10478,7 +10298,6 @@ func (m *awsRestjson1_deserializeOpListRoutingRules) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10602,10 +10421,8 @@ func (m *awsRestjson1_deserializeOpPutRoutingRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10645,7 +10462,6 @@ func (m *awsRestjson1_deserializeOpPutRoutingRule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10799,10 +10615,8 @@ func (m *awsRestjson1_deserializeOpReimportApi) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10842,7 +10656,6 @@ func (m *awsRestjson1_deserializeOpReimportApi) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11096,10 +10909,8 @@ func (m *awsRestjson1_deserializeOpResetAuthorizersCac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11117,7 +10928,6 @@ func (m *awsRestjson1_deserializeOpResetAuthorizersCac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11193,10 +11003,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11208,7 +11016,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11290,10 +11097,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11311,7 +11116,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11393,10 +11197,8 @@ func (m *awsRestjson1_deserializeOpUpdateApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11436,7 +11238,6 @@ func (m *awsRestjson1_deserializeOpUpdateApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11690,10 +11491,8 @@ func (m *awsRestjson1_deserializeOpUpdateApiMapping) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11733,7 +11532,6 @@ func (m *awsRestjson1_deserializeOpUpdateApiMapping) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11882,10 +11680,8 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11925,7 +11721,6 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12133,10 +11928,8 @@ func (m *awsRestjson1_deserializeOpUpdateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12176,7 +11969,6 @@ func (m *awsRestjson1_deserializeOpUpdateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12347,10 +12139,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12390,7 +12180,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12554,10 +12343,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12597,7 +12384,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12878,10 +12664,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntegrationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12921,7 +12705,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntegrationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13080,10 +12863,8 @@ func (m *awsRestjson1_deserializeOpUpdateModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13123,7 +12904,6 @@ func (m *awsRestjson1_deserializeOpUpdateModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13281,10 +13061,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoute) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13324,7 +13102,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoute) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13542,10 +13319,8 @@ func (m *awsRestjson1_deserializeOpUpdateRouteResponse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13585,7 +13360,6 @@ func (m *awsRestjson1_deserializeOpUpdateRouteResponse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13735,10 +13509,8 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13778,7 +13550,6 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14005,10 +13776,8 @@ func (m *awsRestjson1_deserializeOpUpdateVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14048,7 +13817,6 @@ func (m *awsRestjson1_deserializeOpUpdateVpcLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
