--- vendor/github.com/aws/aws-sdk-go-v2/service/taxsettings/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/taxsettings/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteTaxRegis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteTaxRegis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -202,10 +198,8 @@ func (m *awsRestjson1_deserializeOpBatchGetTaxExemptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -245,7 +239,6 @@ func (m *awsRestjson1_deserializeOpBatchGetTaxExemptio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -365,10 +358,8 @@ func (m *awsRestjson1_deserializeOpBatchPutTaxRegistra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -408,7 +399,6 @@ func (m *awsRestjson1_deserializeOpBatchPutTaxRegistra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -532,10 +522,8 @@ func (m *awsRestjson1_deserializeOpDeleteSupplementalT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -547,7 +535,6 @@ func (m *awsRestjson1_deserializeOpDeleteSupplementalT
 	output := &DeleteSupplementalTaxRegistrationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -629,10 +616,8 @@ func (m *awsRestjson1_deserializeOpDeleteTaxRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -644,7 +629,6 @@ func (m *awsRestjson1_deserializeOpDeleteTaxRegistrati
 	output := &DeleteTaxRegistrationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -726,10 +710,8 @@ func (m *awsRestjson1_deserializeOpGetTaxExemptionType
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -769,7 +751,6 @@ func (m *awsRestjson1_deserializeOpGetTaxExemptionType
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -884,10 +865,8 @@ func (m *awsRestjson1_deserializeOpGetTaxInheritance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -927,7 +906,6 @@ func (m *awsRestjson1_deserializeOpGetTaxInheritance) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1046,10 +1024,8 @@ func (m *awsRestjson1_deserializeOpGetTaxRegistration)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1089,7 +1065,6 @@ func (m *awsRestjson1_deserializeOpGetTaxRegistration)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1204,10 +1179,8 @@ func (m *awsRestjson1_deserializeOpGetTaxRegistrationD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1247,7 +1220,6 @@ func (m *awsRestjson1_deserializeOpGetTaxRegistrationD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1372,10 +1344,8 @@ func (m *awsRestjson1_deserializeOpListSupplementalTax
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1415,7 +1385,6 @@ func (m *awsRestjson1_deserializeOpListSupplementalTax
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1539,10 +1508,8 @@ func (m *awsRestjson1_deserializeOpListTaxExemptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1582,7 +1549,6 @@ func (m *awsRestjson1_deserializeOpListTaxExemptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1706,10 +1672,8 @@ func (m *awsRestjson1_deserializeOpListTaxRegistration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1749,7 +1713,6 @@ func (m *awsRestjson1_deserializeOpListTaxRegistration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1873,10 +1836,8 @@ func (m *awsRestjson1_deserializeOpPutSupplementalTaxR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1916,7 +1877,6 @@ func (m *awsRestjson1_deserializeOpPutSupplementalTaxR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2044,10 +2004,8 @@ func (m *awsRestjson1_deserializeOpPutTaxExemption) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2087,7 +2045,6 @@ func (m *awsRestjson1_deserializeOpPutTaxExemption) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2215,10 +2172,8 @@ func (m *awsRestjson1_deserializeOpPutTaxInheritance) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2230,7 +2185,6 @@ func (m *awsRestjson1_deserializeOpPutTaxInheritance) 
 	output := &PutTaxInheritanceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2312,10 +2266,8 @@ func (m *awsRestjson1_deserializeOpPutTaxRegistration)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2355,7 +2307,6 @@ func (m *awsRestjson1_deserializeOpPutTaxRegistration)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
