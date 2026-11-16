--- vendor/github.com/aws/aws-sdk-go-v2/service/appsync/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/appsync/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateApi) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpAssociateApi) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -205,10 +201,8 @@ func (m *awsRestjson1_deserializeOpAssociateMergedGrap
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -248,7 +242,6 @@ func (m *awsRestjson1_deserializeOpAssociateMergedGrap
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -372,10 +365,8 @@ func (m *awsRestjson1_deserializeOpAssociateSourceGrap
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -415,7 +406,6 @@ func (m *awsRestjson1_deserializeOpAssociateSourceGrap
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -539,10 +529,8 @@ func (m *awsRestjson1_deserializeOpCreateApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -582,7 +570,6 @@ func (m *awsRestjson1_deserializeOpCreateApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -703,10 +690,8 @@ func (m *awsRestjson1_deserializeOpCreateApiCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -746,7 +731,6 @@ func (m *awsRestjson1_deserializeOpCreateApiCache) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -867,10 +851,8 @@ func (m *awsRestjson1_deserializeOpCreateApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -910,7 +892,6 @@ func (m *awsRestjson1_deserializeOpCreateApiKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1037,10 +1018,8 @@ func (m *awsRestjson1_deserializeOpCreateChannelNamesp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1080,7 +1059,6 @@ func (m *awsRestjson1_deserializeOpCreateChannelNamesp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1207,10 +1185,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1250,7 +1226,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1371,10 +1346,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1414,7 +1387,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1529,10 +1501,8 @@ func (m *awsRestjson1_deserializeOpCreateFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1572,7 +1542,6 @@ func (m *awsRestjson1_deserializeOpCreateFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1693,10 +1662,8 @@ func (m *awsRestjson1_deserializeOpCreateGraphqlApi) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1736,7 +1703,6 @@ func (m *awsRestjson1_deserializeOpCreateGraphqlApi) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1860,10 +1826,8 @@ func (m *awsRestjson1_deserializeOpCreateResolver) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1903,7 +1867,6 @@ func (m *awsRestjson1_deserializeOpCreateResolver) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2024,10 +1987,8 @@ func (m *awsRestjson1_deserializeOpCreateType) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2067,7 +2028,6 @@ func (m *awsRestjson1_deserializeOpCreateType) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2188,10 +2148,8 @@ func (m *awsRestjson1_deserializeOpDeleteApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2203,7 +2161,6 @@ func (m *awsRestjson1_deserializeOpDeleteApi) HandleDe
 	output := &DeleteApiOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2291,10 +2248,8 @@ func (m *awsRestjson1_deserializeOpDeleteApiCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2306,7 +2261,6 @@ func (m *awsRestjson1_deserializeOpDeleteApiCache) Han
 	output := &DeleteApiCacheOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2391,10 +2345,8 @@ func (m *awsRestjson1_deserializeOpDeleteApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2406,7 +2358,6 @@ func (m *awsRestjson1_deserializeOpDeleteApiKey) Handl
 	output := &DeleteApiKeyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2488,10 +2439,8 @@ func (m *awsRestjson1_deserializeOpDeleteChannelNamesp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2503,7 +2452,6 @@ func (m *awsRestjson1_deserializeOpDeleteChannelNamesp
 	output := &DeleteChannelNamespaceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2591,10 +2539,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2606,7 +2552,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 	output := &DeleteDataSourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2691,10 +2636,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2706,7 +2649,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainName) H
 	output := &DeleteDomainNameOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2791,10 +2733,8 @@ func (m *awsRestjson1_deserializeOpDeleteFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2806,7 +2746,6 @@ func (m *awsRestjson1_deserializeOpDeleteFunction) Han
 	output := &DeleteFunctionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2891,10 +2830,8 @@ func (m *awsRestjson1_deserializeOpDeleteGraphqlApi) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2906,7 +2843,6 @@ func (m *awsRestjson1_deserializeOpDeleteGraphqlApi) H
 	output := &DeleteGraphqlApiOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2994,10 +2930,8 @@ func (m *awsRestjson1_deserializeOpDeleteResolver) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3009,7 +2943,6 @@ func (m *awsRestjson1_deserializeOpDeleteResolver) Han
 	output := &DeleteResolverOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3094,10 +3027,8 @@ func (m *awsRestjson1_deserializeOpDeleteType) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3109,7 +3040,6 @@ func (m *awsRestjson1_deserializeOpDeleteType) HandleD
 	output := &DeleteTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3194,10 +3124,8 @@ func (m *awsRestjson1_deserializeOpDisassociateApi) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3209,7 +3137,6 @@ func (m *awsRestjson1_deserializeOpDisassociateApi) Ha
 	output := &DisassociateApiOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3294,10 +3221,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMergedG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3337,7 +3262,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMergedG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3462,10 +3386,8 @@ func (m *awsRestjson1_deserializeOpDisassociateSourceG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3505,7 +3427,6 @@ func (m *awsRestjson1_deserializeOpDisassociateSourceG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3630,10 +3551,8 @@ func (m *awsRestjson1_deserializeOpEvaluateCode) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3673,7 +3592,6 @@ func (m *awsRestjson1_deserializeOpEvaluateCode) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3820,10 +3738,8 @@ func (m *awsRestjson1_deserializeOpEvaluateMappingTemp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3863,7 +3779,6 @@ func (m *awsRestjson1_deserializeOpEvaluateMappingTemp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4010,10 +3925,8 @@ func (m *awsRestjson1_deserializeOpFlushApiCache) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4025,7 +3938,6 @@ func (m *awsRestjson1_deserializeOpFlushApiCache) Hand
 	output := &FlushApiCacheOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4110,10 +4022,8 @@ func (m *awsRestjson1_deserializeOpGetApi) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4153,7 +4063,6 @@ func (m *awsRestjson1_deserializeOpGetApi) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4274,10 +4183,8 @@ func (m *awsRestjson1_deserializeOpGetApiAssociation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4317,7 +4224,6 @@ func (m *awsRestjson1_deserializeOpGetApiAssociation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4435,10 +4341,8 @@ func (m *awsRestjson1_deserializeOpGetApiCache) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4478,7 +4382,6 @@ func (m *awsRestjson1_deserializeOpGetApiCache) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4599,10 +4502,8 @@ func (m *awsRestjson1_deserializeOpGetChannelNamespace
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4642,7 +4543,6 @@ func (m *awsRestjson1_deserializeOpGetChannelNamespace
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4763,10 +4663,8 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4806,7 +4704,6 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4927,10 +4824,8 @@ func (m *awsRestjson1_deserializeOpGetDataSourceIntros
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4970,7 +4865,6 @@ func (m *awsRestjson1_deserializeOpGetDataSourceIntros
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5112,10 +5006,8 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5155,7 +5047,6 @@ func (m *awsRestjson1_deserializeOpGetDomainName) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5273,10 +5164,8 @@ func (m *awsRestjson1_deserializeOpGetFunction) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5316,7 +5205,6 @@ func (m *awsRestjson1_deserializeOpGetFunction) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5431,10 +5319,8 @@ func (m *awsRestjson1_deserializeOpGetGraphqlApi) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5474,7 +5360,6 @@ func (m *awsRestjson1_deserializeOpGetGraphqlApi) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5595,10 +5480,8 @@ func (m *awsRestjson1_deserializeOpGetGraphqlApiEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5638,7 +5521,6 @@ func (m *awsRestjson1_deserializeOpGetGraphqlApiEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5759,10 +5641,8 @@ func (m *awsRestjson1_deserializeOpGetIntrospectionSch
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5779,7 +5659,6 @@ func (m *awsRestjson1_deserializeOpGetIntrospectionSch
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5883,10 +5762,8 @@ func (m *awsRestjson1_deserializeOpGetResolver) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5926,7 +5803,6 @@ func (m *awsRestjson1_deserializeOpGetResolver) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6041,10 +5917,8 @@ func (m *awsRestjson1_deserializeOpGetSchemaCreationSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6084,7 +5958,6 @@ func (m *awsRestjson1_deserializeOpGetSchemaCreationSt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6215,10 +6088,8 @@ func (m *awsRestjson1_deserializeOpGetSourceApiAssocia
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6258,7 +6129,6 @@ func (m *awsRestjson1_deserializeOpGetSourceApiAssocia
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6376,10 +6246,8 @@ func (m *awsRestjson1_deserializeOpGetType) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6419,7 +6287,6 @@ func (m *awsRestjson1_deserializeOpGetType) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6540,10 +6407,8 @@ func (m *awsRestjson1_deserializeOpListApiKeys) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6583,7 +6448,6 @@ func (m *awsRestjson1_deserializeOpListApiKeys) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6710,10 +6574,8 @@ func (m *awsRestjson1_deserializeOpListApis) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6753,7 +6615,6 @@ func (m *awsRestjson1_deserializeOpListApis) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6877,10 +6738,8 @@ func (m *awsRestjson1_deserializeOpListChannelNamespac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6920,7 +6779,6 @@ func (m *awsRestjson1_deserializeOpListChannelNamespac
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7047,10 +6905,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7090,7 +6946,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7217,10 +7072,8 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7260,7 +7113,6 @@ func (m *awsRestjson1_deserializeOpListDomainNames) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7384,10 +7236,8 @@ func (m *awsRestjson1_deserializeOpListFunctions) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7427,7 +7277,6 @@ func (m *awsRestjson1_deserializeOpListFunctions) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7554,10 +7403,8 @@ func (m *awsRestjson1_deserializeOpListGraphqlApis) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7597,7 +7444,6 @@ func (m *awsRestjson1_deserializeOpListGraphqlApis) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7721,10 +7567,8 @@ func (m *awsRestjson1_deserializeOpListResolvers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7764,7 +7608,6 @@ func (m *awsRestjson1_deserializeOpListResolvers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7891,10 +7734,8 @@ func (m *awsRestjson1_deserializeOpListResolversByFunc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7934,7 +7775,6 @@ func (m *awsRestjson1_deserializeOpListResolversByFunc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8061,10 +7901,8 @@ func (m *awsRestjson1_deserializeOpListSourceApiAssoci
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8104,7 +7942,6 @@ func (m *awsRestjson1_deserializeOpListSourceApiAssoci
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8231,10 +8068,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8274,7 +8109,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8398,10 +8232,8 @@ func (m *awsRestjson1_deserializeOpListTypes) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8441,7 +8273,6 @@ func (m *awsRestjson1_deserializeOpListTypes) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8571,10 +8402,8 @@ func (m *awsRestjson1_deserializeOpListTypesByAssociat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8614,7 +8443,6 @@ func (m *awsRestjson1_deserializeOpListTypesByAssociat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8744,10 +8572,8 @@ func (m *awsRestjson1_deserializeOpPutGraphqlApiEnviro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8787,7 +8613,6 @@ func (m *awsRestjson1_deserializeOpPutGraphqlApiEnviro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8911,10 +8736,8 @@ func (m *awsRestjson1_deserializeOpStartDataSourceIntr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8954,7 +8777,6 @@ func (m *awsRestjson1_deserializeOpStartDataSourceIntr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9094,10 +8916,8 @@ func (m *awsRestjson1_deserializeOpStartSchemaCreation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9137,7 +8957,6 @@ func (m *awsRestjson1_deserializeOpStartSchemaCreation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9262,10 +9081,8 @@ func (m *awsRestjson1_deserializeOpStartSchemaMerge) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9305,7 +9122,6 @@ func (m *awsRestjson1_deserializeOpStartSchemaMerge) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9430,10 +9246,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9445,7 +9259,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9533,10 +9346,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9548,7 +9359,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9636,10 +9446,8 @@ func (m *awsRestjson1_deserializeOpUpdateApi) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9679,7 +9487,6 @@ func (m *awsRestjson1_deserializeOpUpdateApi) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9803,10 +9610,8 @@ func (m *awsRestjson1_deserializeOpUpdateApiCache) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9846,7 +9651,6 @@ func (m *awsRestjson1_deserializeOpUpdateApiCache) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9967,10 +9771,8 @@ func (m *awsRestjson1_deserializeOpUpdateApiKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10010,7 +9812,6 @@ func (m *awsRestjson1_deserializeOpUpdateApiKey) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10134,10 +9935,8 @@ func (m *awsRestjson1_deserializeOpUpdateChannelNamesp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10177,7 +9976,6 @@ func (m *awsRestjson1_deserializeOpUpdateChannelNamesp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10301,10 +10099,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10344,7 +10140,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10465,10 +10260,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10508,7 +10301,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainName) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10629,10 +10421,8 @@ func (m *awsRestjson1_deserializeOpUpdateFunction) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10672,7 +10462,6 @@ func (m *awsRestjson1_deserializeOpUpdateFunction) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10793,10 +10582,8 @@ func (m *awsRestjson1_deserializeOpUpdateGraphqlApi) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10836,7 +10623,6 @@ func (m *awsRestjson1_deserializeOpUpdateGraphqlApi) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10960,10 +10746,8 @@ func (m *awsRestjson1_deserializeOpUpdateResolver) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11003,7 +10787,6 @@ func (m *awsRestjson1_deserializeOpUpdateResolver) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11124,10 +10907,8 @@ func (m *awsRestjson1_deserializeOpUpdateSourceApiAsso
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11167,7 +10948,6 @@ func (m *awsRestjson1_deserializeOpUpdateSourceApiAsso
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11288,10 +11068,8 @@ func (m *awsRestjson1_deserializeOpUpdateType) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11331,7 +11109,6 @@ func (m *awsRestjson1_deserializeOpUpdateType) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
