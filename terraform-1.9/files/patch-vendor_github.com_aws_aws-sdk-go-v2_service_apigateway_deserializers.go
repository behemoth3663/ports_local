--- vendor/github.com/aws/aws-sdk-go-v2/service/apigateway/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/apigateway/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateApiKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -304,10 +300,8 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -347,7 +341,6 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -556,10 +549,8 @@ func (m *awsRestjson1_deserializeOpCreateBasePathMappi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -599,7 +590,6 @@ func (m *awsRestjson1_deserializeOpCreateBasePathMappi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -745,10 +735,8 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -788,7 +776,6 @@ func (m *awsRestjson1_deserializeOpCreateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -949,10 +936,8 @@ func (m *awsRestjson1_deserializeOpCreateDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -992,7 +977,6 @@ func (m *awsRestjson1_deserializeOpCreateDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1134,10 +1118,8 @@ func (m *awsRestjson1_deserializeOpCreateDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1177,7 +1159,6 @@ func (m *awsRestjson1_deserializeOpCreateDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1330,10 +1311,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1373,7 +1352,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1682,10 +1660,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainNameAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1725,7 +1701,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainNameAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1882,10 +1857,8 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1925,7 +1898,6 @@ func (m *awsRestjson1_deserializeOpCreateModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2089,10 +2061,8 @@ func (m *awsRestjson1_deserializeOpCreateRequestValida
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2132,7 +2102,6 @@ func (m *awsRestjson1_deserializeOpCreateRequestValida
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2287,10 +2256,8 @@ func (m *awsRestjson1_deserializeOpCreateResource) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2330,7 +2297,6 @@ func (m *awsRestjson1_deserializeOpCreateResource) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2490,10 +2456,8 @@ func (m *awsRestjson1_deserializeOpCreateRestApi) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2533,7 +2497,6 @@ func (m *awsRestjson1_deserializeOpCreateRestApi) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2770,10 +2733,8 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2813,7 +2774,6 @@ func (m *awsRestjson1_deserializeOpCreateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3079,10 +3039,8 @@ func (m *awsRestjson1_deserializeOpCreateUsagePlan) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3122,7 +3080,6 @@ func (m *awsRestjson1_deserializeOpCreateUsagePlan) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3297,10 +3254,8 @@ func (m *awsRestjson1_deserializeOpCreateUsagePlanKey)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3340,7 +3295,6 @@ func (m *awsRestjson1_deserializeOpCreateUsagePlanKey)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3495,10 +3449,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3538,7 +3490,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3709,10 +3660,8 @@ func (m *awsRestjson1_deserializeOpDeleteApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3730,7 +3679,6 @@ func (m *awsRestjson1_deserializeOpDeleteApiKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3815,10 +3763,8 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3836,7 +3782,6 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3921,10 +3866,8 @@ func (m *awsRestjson1_deserializeOpDeleteBasePathMappi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3942,7 +3885,6 @@ func (m *awsRestjson1_deserializeOpDeleteBasePathMappi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4027,10 +3969,8 @@ func (m *awsRestjson1_deserializeOpDeleteClientCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4048,7 +3988,6 @@ func (m *awsRestjson1_deserializeOpDeleteClientCertifi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4133,10 +4072,8 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4154,7 +4091,6 @@ func (m *awsRestjson1_deserializeOpDeleteDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4242,10 +4178,8 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4263,7 +4197,6 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4348,10 +4281,8 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4369,7 +4300,6 @@ func (m *awsRestjson1_deserializeOpDeleteDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4454,10 +4384,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4475,7 +4403,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4560,10 +4487,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainNameAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4581,7 +4506,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainNameAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4666,10 +4590,8 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayRespon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4687,7 +4609,6 @@ func (m *awsRestjson1_deserializeOpDeleteGatewayRespon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4772,10 +4693,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4793,7 +4712,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4878,10 +4796,8 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4899,7 +4815,6 @@ func (m *awsRestjson1_deserializeOpDeleteIntegrationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4984,10 +4899,8 @@ func (m *awsRestjson1_deserializeOpDeleteMethod) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5005,7 +4918,6 @@ func (m *awsRestjson1_deserializeOpDeleteMethod) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5087,10 +4999,8 @@ func (m *awsRestjson1_deserializeOpDeleteMethodRespons
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5108,7 +5018,6 @@ func (m *awsRestjson1_deserializeOpDeleteMethodRespons
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5193,10 +5102,8 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5214,7 +5121,6 @@ func (m *awsRestjson1_deserializeOpDeleteModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5299,10 +5205,8 @@ func (m *awsRestjson1_deserializeOpDeleteRequestValida
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5320,7 +5224,6 @@ func (m *awsRestjson1_deserializeOpDeleteRequestValida
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5405,10 +5308,8 @@ func (m *awsRestjson1_deserializeOpDeleteResource) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5426,7 +5327,6 @@ func (m *awsRestjson1_deserializeOpDeleteResource) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5511,10 +5411,8 @@ func (m *awsRestjson1_deserializeOpDeleteRestApi) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5532,7 +5430,6 @@ func (m *awsRestjson1_deserializeOpDeleteRestApi) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5617,10 +5514,8 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5638,7 +5533,6 @@ func (m *awsRestjson1_deserializeOpDeleteStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5726,10 +5620,8 @@ func (m *awsRestjson1_deserializeOpDeleteUsagePlan) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5747,7 +5639,6 @@ func (m *awsRestjson1_deserializeOpDeleteUsagePlan) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5832,10 +5723,8 @@ func (m *awsRestjson1_deserializeOpDeleteUsagePlanKey)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5853,7 +5742,6 @@ func (m *awsRestjson1_deserializeOpDeleteUsagePlanKey)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5938,10 +5826,8 @@ func (m *awsRestjson1_deserializeOpDeleteVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5959,7 +5845,6 @@ func (m *awsRestjson1_deserializeOpDeleteVpcLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6044,10 +5929,8 @@ func (m *awsRestjson1_deserializeOpFlushStageAuthorize
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6065,7 +5948,6 @@ func (m *awsRestjson1_deserializeOpFlushStageAuthorize
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6153,10 +6035,8 @@ func (m *awsRestjson1_deserializeOpFlushStageCache) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6174,7 +6054,6 @@ func (m *awsRestjson1_deserializeOpFlushStageCache) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6262,10 +6141,8 @@ func (m *awsRestjson1_deserializeOpGenerateClientCerti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6305,7 +6182,6 @@ func (m *awsRestjson1_deserializeOpGenerateClientCerti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6485,10 +6361,8 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6528,7 +6402,6 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6669,10 +6542,8 @@ func (m *awsRestjson1_deserializeOpGetApiKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6712,7 +6583,6 @@ func (m *awsRestjson1_deserializeOpGetApiKey) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6921,10 +6791,8 @@ func (m *awsRestjson1_deserializeOpGetApiKeys) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6964,7 +6832,6 @@ func (m *awsRestjson1_deserializeOpGetApiKeys) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7096,10 +6963,8 @@ func (m *awsRestjson1_deserializeOpGetAuthorizer) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7139,7 +7004,6 @@ func (m *awsRestjson1_deserializeOpGetAuthorizer) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7342,10 +7206,8 @@ func (m *awsRestjson1_deserializeOpGetAuthorizers) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7385,7 +7247,6 @@ func (m *awsRestjson1_deserializeOpGetAuthorizers) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7512,10 +7373,8 @@ func (m *awsRestjson1_deserializeOpGetBasePathMapping)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7555,7 +7414,6 @@ func (m *awsRestjson1_deserializeOpGetBasePathMapping)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7695,10 +7553,8 @@ func (m *awsRestjson1_deserializeOpGetBasePathMappings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7738,7 +7594,6 @@ func (m *awsRestjson1_deserializeOpGetBasePathMappings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7865,10 +7720,8 @@ func (m *awsRestjson1_deserializeOpGetClientCertificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7908,7 +7761,6 @@ func (m *awsRestjson1_deserializeOpGetClientCertificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8085,10 +7937,8 @@ func (m *awsRestjson1_deserializeOpGetClientCertificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8128,7 +7978,6 @@ func (m *awsRestjson1_deserializeOpGetClientCertificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8255,10 +8104,8 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8298,7 +8145,6 @@ func (m *awsRestjson1_deserializeOpGetDeployment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8453,10 +8299,8 @@ func (m *awsRestjson1_deserializeOpGetDeployments) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8496,7 +8340,6 @@ func (m *awsRestjson1_deserializeOpGetDeployments) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8626,10 +8469,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentationPar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8669,7 +8510,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentationPar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8805,10 +8645,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentationPar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8848,7 +8686,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentationPar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8975,10 +8812,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentationVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9018,7 +8853,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentationVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9162,10 +8996,8 @@ func (m *awsRestjson1_deserializeOpGetDocumentationVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9205,7 +9037,6 @@ func (m *awsRestjson1_deserializeOpGetDocumentationVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9332,10 +9163,8 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9375,7 +9204,6 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9681,10 +9509,8 @@ func (m *awsRestjson1_deserializeOpGetDomainNameAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9724,7 +9550,6 @@ func (m *awsRestjson1_deserializeOpGetDomainNameAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9851,10 +9676,8 @@ func (m *awsRestjson1_deserializeOpGetDomainNames) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9894,7 +9717,6 @@ func (m *awsRestjson1_deserializeOpGetDomainNames) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10021,10 +9843,8 @@ func (m *awsRestjson1_deserializeOpGetExport) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10046,7 +9866,6 @@ func (m *awsRestjson1_deserializeOpGetExport) HandleDe
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10173,10 +9992,8 @@ func (m *awsRestjson1_deserializeOpGetGatewayResponse)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10216,7 +10033,6 @@ func (m *awsRestjson1_deserializeOpGetGatewayResponse)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10366,10 +10182,8 @@ func (m *awsRestjson1_deserializeOpGetGatewayResponses
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10409,7 +10223,6 @@ func (m *awsRestjson1_deserializeOpGetGatewayResponses
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10536,10 +10349,8 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10579,7 +10390,6 @@ func (m *awsRestjson1_deserializeOpGetIntegration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10811,10 +10621,8 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10854,7 +10662,6 @@ func (m *awsRestjson1_deserializeOpGetIntegrationRespo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11004,10 +10811,8 @@ func (m *awsRestjson1_deserializeOpGetMethod) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11047,7 +10852,6 @@ func (m *awsRestjson1_deserializeOpGetMethod) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11236,10 +11040,8 @@ func (m *awsRestjson1_deserializeOpGetMethodResponse) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11279,7 +11081,6 @@ func (m *awsRestjson1_deserializeOpGetMethodResponse) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11408,10 +11209,8 @@ func (m *awsRestjson1_deserializeOpGetModel) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11451,7 +11250,6 @@ func (m *awsRestjson1_deserializeOpGetModel) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11609,10 +11407,8 @@ func (m *awsRestjson1_deserializeOpGetModels) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11652,7 +11448,6 @@ func (m *awsRestjson1_deserializeOpGetModels) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11779,10 +11574,8 @@ func (m *awsRestjson1_deserializeOpGetModelTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11822,7 +11615,6 @@ func (m *awsRestjson1_deserializeOpGetModelTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11944,10 +11736,8 @@ func (m *awsRestjson1_deserializeOpGetRequestValidator
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11987,7 +11777,6 @@ func (m *awsRestjson1_deserializeOpGetRequestValidator
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12136,10 +11925,8 @@ func (m *awsRestjson1_deserializeOpGetRequestValidator
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12179,7 +11966,6 @@ func (m *awsRestjson1_deserializeOpGetRequestValidator
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12306,10 +12092,8 @@ func (m *awsRestjson1_deserializeOpGetResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12349,7 +12133,6 @@ func (m *awsRestjson1_deserializeOpGetResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12500,10 +12283,8 @@ func (m *awsRestjson1_deserializeOpGetResources) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12543,7 +12324,6 @@ func (m *awsRestjson1_deserializeOpGetResources) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12670,10 +12450,8 @@ func (m *awsRestjson1_deserializeOpGetRestApi) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12713,7 +12491,6 @@ func (m *awsRestjson1_deserializeOpGetRestApi) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12947,10 +12724,8 @@ func (m *awsRestjson1_deserializeOpGetRestApis) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12990,7 +12765,6 @@ func (m *awsRestjson1_deserializeOpGetRestApis) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13117,10 +12891,8 @@ func (m *awsRestjson1_deserializeOpGetSdk) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13142,7 +12914,6 @@ func (m *awsRestjson1_deserializeOpGetSdk) HandleDeser
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13269,10 +13040,8 @@ func (m *awsRestjson1_deserializeOpGetSdkType) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13312,7 +13081,6 @@ func (m *awsRestjson1_deserializeOpGetSdkType) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13457,10 +13225,8 @@ func (m *awsRestjson1_deserializeOpGetSdkTypes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13500,7 +13266,6 @@ func (m *awsRestjson1_deserializeOpGetSdkTypes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13618,10 +13383,8 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13661,7 +13424,6 @@ func (m *awsRestjson1_deserializeOpGetStage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13927,10 +13689,8 @@ func (m *awsRestjson1_deserializeOpGetStages) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13970,7 +13730,6 @@ func (m *awsRestjson1_deserializeOpGetStages) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14094,10 +13853,8 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14137,7 +13894,6 @@ func (m *awsRestjson1_deserializeOpGetTags) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14255,10 +14011,8 @@ func (m *awsRestjson1_deserializeOpGetUsage) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14298,7 +14052,6 @@ func (m *awsRestjson1_deserializeOpGetUsage) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14452,10 +14205,8 @@ func (m *awsRestjson1_deserializeOpGetUsagePlan) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14495,7 +14246,6 @@ func (m *awsRestjson1_deserializeOpGetUsagePlan) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14664,10 +14414,8 @@ func (m *awsRestjson1_deserializeOpGetUsagePlanKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14707,7 +14455,6 @@ func (m *awsRestjson1_deserializeOpGetUsagePlanKey) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14856,10 +14603,8 @@ func (m *awsRestjson1_deserializeOpGetUsagePlanKeys) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14899,7 +14644,6 @@ func (m *awsRestjson1_deserializeOpGetUsagePlanKeys) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15026,10 +14770,8 @@ func (m *awsRestjson1_deserializeOpGetUsagePlans) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15069,7 +14811,6 @@ func (m *awsRestjson1_deserializeOpGetUsagePlans) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15196,10 +14937,8 @@ func (m *awsRestjson1_deserializeOpGetVpcLink) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15239,7 +14978,6 @@ func (m *awsRestjson1_deserializeOpGetVpcLink) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15407,10 +15145,8 @@ func (m *awsRestjson1_deserializeOpGetVpcLinks) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15450,7 +15186,6 @@ func (m *awsRestjson1_deserializeOpGetVpcLinks) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15577,10 +15312,8 @@ func (m *awsRestjson1_deserializeOpImportApiKeys) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15620,7 +15353,6 @@ func (m *awsRestjson1_deserializeOpImportApiKeys) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15749,10 +15481,8 @@ func (m *awsRestjson1_deserializeOpImportDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15792,7 +15522,6 @@ func (m *awsRestjson1_deserializeOpImportDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15921,10 +15650,8 @@ func (m *awsRestjson1_deserializeOpImportRestApi) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15964,7 +15691,6 @@ func (m *awsRestjson1_deserializeOpImportRestApi) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16204,10 +15930,8 @@ func (m *awsRestjson1_deserializeOpPutGatewayResponse)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16247,7 +15971,6 @@ func (m *awsRestjson1_deserializeOpPutGatewayResponse)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16403,10 +16126,8 @@ func (m *awsRestjson1_deserializeOpPutIntegration) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16446,7 +16167,6 @@ func (m *awsRestjson1_deserializeOpPutIntegration) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16684,10 +16404,8 @@ func (m *awsRestjson1_deserializeOpPutIntegrationRespo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16727,7 +16445,6 @@ func (m *awsRestjson1_deserializeOpPutIntegrationRespo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16883,10 +16600,8 @@ func (m *awsRestjson1_deserializeOpPutMethod) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16926,7 +16641,6 @@ func (m *awsRestjson1_deserializeOpPutMethod) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17124,10 +16838,8 @@ func (m *awsRestjson1_deserializeOpPutMethodResponse) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17167,7 +16879,6 @@ func (m *awsRestjson1_deserializeOpPutMethodResponse) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17305,10 +17016,8 @@ func (m *awsRestjson1_deserializeOpPutRestApi) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17348,7 +17057,6 @@ func (m *awsRestjson1_deserializeOpPutRestApi) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17588,10 +17296,8 @@ func (m *awsRestjson1_deserializeOpRejectDomainNameAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17609,7 +17315,6 @@ func (m *awsRestjson1_deserializeOpRejectDomainNameAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17694,10 +17399,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17715,7 +17418,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17803,10 +17505,8 @@ func (m *awsRestjson1_deserializeOpTestInvokeAuthorize
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17846,7 +17546,6 @@ func (m *awsRestjson1_deserializeOpTestInvokeAuthorize
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18022,10 +17721,8 @@ func (m *awsRestjson1_deserializeOpTestInvokeMethod) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18065,7 +17762,6 @@ func (m *awsRestjson1_deserializeOpTestInvokeMethod) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18232,10 +17928,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18253,7 +17947,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18341,10 +18034,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccount) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18384,7 +18075,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccount) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18531,10 +18221,8 @@ func (m *awsRestjson1_deserializeOpUpdateApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18574,7 +18262,6 @@ func (m *awsRestjson1_deserializeOpUpdateApiKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18789,10 +18476,8 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18832,7 +18517,6 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19041,10 +18725,8 @@ func (m *awsRestjson1_deserializeOpUpdateBasePathMappi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19084,7 +18766,6 @@ func (m *awsRestjson1_deserializeOpUpdateBasePathMappi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19230,10 +18911,8 @@ func (m *awsRestjson1_deserializeOpUpdateClientCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19273,7 +18952,6 @@ func (m *awsRestjson1_deserializeOpUpdateClientCertifi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19456,10 +19134,8 @@ func (m *awsRestjson1_deserializeOpUpdateDeployment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19499,7 +19175,6 @@ func (m *awsRestjson1_deserializeOpUpdateDeployment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19660,10 +19335,8 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19703,7 +19376,6 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19845,10 +19517,8 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19888,7 +19558,6 @@ func (m *awsRestjson1_deserializeOpUpdateDocumentation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20041,10 +19710,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20084,7 +19751,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20396,10 +20062,8 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayRespon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20439,7 +20103,6 @@ func (m *awsRestjson1_deserializeOpUpdateGatewayRespon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20595,10 +20258,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20638,7 +20299,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntegration) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20876,10 +20536,8 @@ func (m *awsRestjson1_deserializeOpUpdateIntegrationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20919,7 +20577,6 @@ func (m *awsRestjson1_deserializeOpUpdateIntegrationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21075,10 +20732,8 @@ func (m *awsRestjson1_deserializeOpUpdateMethod) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21118,7 +20773,6 @@ func (m *awsRestjson1_deserializeOpUpdateMethod) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21313,10 +20967,8 @@ func (m *awsRestjson1_deserializeOpUpdateMethodRespons
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21356,7 +21008,6 @@ func (m *awsRestjson1_deserializeOpUpdateMethodRespons
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21494,10 +21145,8 @@ func (m *awsRestjson1_deserializeOpUpdateModel) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21537,7 +21186,6 @@ func (m *awsRestjson1_deserializeOpUpdateModel) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21701,10 +21349,8 @@ func (m *awsRestjson1_deserializeOpUpdateRequestValida
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21744,7 +21390,6 @@ func (m *awsRestjson1_deserializeOpUpdateRequestValida
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21899,10 +21544,8 @@ func (m *awsRestjson1_deserializeOpUpdateResource) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21942,7 +21585,6 @@ func (m *awsRestjson1_deserializeOpUpdateResource) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22099,10 +21741,8 @@ func (m *awsRestjson1_deserializeOpUpdateRestApi) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22142,7 +21782,6 @@ func (m *awsRestjson1_deserializeOpUpdateRestApi) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22382,10 +22021,8 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22425,7 +22062,6 @@ func (m *awsRestjson1_deserializeOpUpdateStage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22691,10 +22327,8 @@ func (m *awsRestjson1_deserializeOpUpdateUsage) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22734,7 +22368,6 @@ func (m *awsRestjson1_deserializeOpUpdateUsage) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22894,10 +22527,8 @@ func (m *awsRestjson1_deserializeOpUpdateUsagePlan) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22937,7 +22568,6 @@ func (m *awsRestjson1_deserializeOpUpdateUsagePlan) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23112,10 +22742,8 @@ func (m *awsRestjson1_deserializeOpUpdateVpcLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23155,7 +22783,6 @@ func (m *awsRestjson1_deserializeOpUpdateVpcLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
