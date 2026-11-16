--- vendor/github.com/aws/aws-sdk-go-v2/service/account/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/account/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAcceptPrimaryEmailU
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAcceptPrimaryEmailU
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpDeleteAlternateCont
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -237,7 +231,6 @@ func (m *awsRestjson1_deserializeOpDeleteAlternateCont
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -322,10 +315,8 @@ func (m *awsRestjson1_deserializeOpDisableRegion) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -343,7 +334,6 @@ func (m *awsRestjson1_deserializeOpDisableRegion) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -428,10 +418,8 @@ func (m *awsRestjson1_deserializeOpEnableRegion) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -449,7 +437,6 @@ func (m *awsRestjson1_deserializeOpEnableRegion) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -534,10 +521,8 @@ func (m *awsRestjson1_deserializeOpGetAccountInformati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -577,7 +562,6 @@ func (m *awsRestjson1_deserializeOpGetAccountInformati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -721,10 +705,8 @@ func (m *awsRestjson1_deserializeOpGetAlternateContact
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -764,7 +746,6 @@ func (m *awsRestjson1_deserializeOpGetAlternateContact
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -885,10 +866,8 @@ func (m *awsRestjson1_deserializeOpGetContactInformati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -928,7 +907,6 @@ func (m *awsRestjson1_deserializeOpGetContactInformati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1049,10 +1027,8 @@ func (m *awsRestjson1_deserializeOpGetPrimaryEmail) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1092,7 +1068,6 @@ func (m *awsRestjson1_deserializeOpGetPrimaryEmail) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1217,10 +1192,8 @@ func (m *awsRestjson1_deserializeOpGetRegionOptStatus)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1260,7 +1233,6 @@ func (m *awsRestjson1_deserializeOpGetRegionOptStatus)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1391,10 +1363,8 @@ func (m *awsRestjson1_deserializeOpListRegions) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1434,7 +1404,6 @@ func (m *awsRestjson1_deserializeOpListRegions) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1561,10 +1530,8 @@ func (m *awsRestjson1_deserializeOpPutAccountName) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1582,7 +1549,6 @@ func (m *awsRestjson1_deserializeOpPutAccountName) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1664,10 +1630,8 @@ func (m *awsRestjson1_deserializeOpPutAlternateContact
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1685,7 +1649,6 @@ func (m *awsRestjson1_deserializeOpPutAlternateContact
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1767,10 +1730,8 @@ func (m *awsRestjson1_deserializeOpPutContactInformati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1788,7 +1749,6 @@ func (m *awsRestjson1_deserializeOpPutContactInformati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1870,10 +1830,8 @@ func (m *awsRestjson1_deserializeOpStartPrimaryEmailUp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1913,7 +1871,6 @@ func (m *awsRestjson1_deserializeOpStartPrimaryEmailUp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
