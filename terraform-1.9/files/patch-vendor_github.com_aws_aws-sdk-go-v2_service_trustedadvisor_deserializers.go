--- vendor/github.com/aws/aws-sdk-go-v2/service/trustedadvisor/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/trustedadvisor/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRecommen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateRecommen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -210,10 +206,8 @@ func (m *awsRestjson1_deserializeOpGetOrganizationReco
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -253,7 +247,6 @@ func (m *awsRestjson1_deserializeOpGetOrganizationReco
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -374,10 +367,8 @@ func (m *awsRestjson1_deserializeOpGetRecommendation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -417,7 +408,6 @@ func (m *awsRestjson1_deserializeOpGetRecommendation) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -538,10 +528,8 @@ func (m *awsRestjson1_deserializeOpListChecks) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -581,7 +569,6 @@ func (m *awsRestjson1_deserializeOpListChecks) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -708,10 +695,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -751,7 +736,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -881,10 +865,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -924,7 +906,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1054,10 +1035,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1097,7 +1076,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationRec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1224,10 +1202,8 @@ func (m *awsRestjson1_deserializeOpListRecommendationR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1267,7 +1243,6 @@ func (m *awsRestjson1_deserializeOpListRecommendationR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1397,10 +1372,8 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1440,7 +1413,6 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1567,10 +1539,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1588,7 +1558,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1676,10 +1645,8 @@ func (m *awsRestjson1_deserializeOpUpdateRecommendatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1697,7 +1664,6 @@ func (m *awsRestjson1_deserializeOpUpdateRecommendatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
